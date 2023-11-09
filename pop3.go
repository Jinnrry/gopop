package gopop

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

// server Pop3服务实例
type server struct {
	Domain     string // 域名
	Port       int    // 端口
	TlsEnabled bool   //是否启用Tls
	Action     Action
	stop       chan bool
	close      bool
	lck        sync.Mutex
}

// NewPop3Server 新建一个服务实例
func NewPop3Server(port int, domain string, tlsEnabled bool, action Action) *server {
	return &server{
		Domain:     domain,
		Port:       port,
		TlsEnabled: tlsEnabled,
		Action:     action,
		stop:       make(chan bool, 1),
	}
}

// Start 启动服务
func (s *server) Start() error {
	if s.lck.TryLock() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
		if err != nil {
			return err
		}
		s.close = false
		defer func() {
			listener.Close()
		}()

		go func() {
			for {
				conn, err := listener.Accept()
				if err != nil {
					if s.close {
						break
					} else {
						continue
					}
				}
				_ = conn
			}
		}()
		<-s.stop
	} else {
		return errors.New("server Is Running")
	}

	return nil
}

// Stop 停止服务
func (s *server) Stop() {
	s.close = true
	s.stop <- true
}

func (s *server) handleClient(conn net.Conn) {
	defer conn.Close()

	data := &Data{}
	var (
		eol = "\r\n"
	)
	reader := bufio.NewReader(conn)

	fmt.Fprintf(conn, "+OK %s POP3 Server powered by gopop"+eol, s.Domain)

	for {
		rawLine, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			return
		}

		cmd, args := getCommand(rawLine)

		var cmdError error

		/*
			user<SP>username<CRLF>	      user 命令是POP3客户端程序与POP3邮件服务器建立连接后通常发送的第一条命令，参数 username 表示收件人的帐户名称。
			pass<SP>password<CRLF>	      pass 命令是在user命令成功通过后，POP3客户端程序接着发送的命令，它用于传递帐户的密码，参数 password 表示帐户的密码。
			apop<SP>name,digest<CRLF>	      apop 命令用于替代user和pass命令，它以MD5 数字摘要的形式向POP3邮件服务器提交帐户密码。
			stat<CRLF>	     stat 命令用于查询邮箱中的统计信息，例如：邮箱中的邮件数量和邮件占用的字节大小等。
			uidl<SP>msg#<CRLF>	     uidl 命令用于查询某封邮件的唯一标志符，参数msg#表示邮件的序号，是一个从1开始编号的数字。
			list<SP>[MSG#]<CRLF>	     list 命令用于列出邮箱中的邮件信息，参数 msg#是一个可选参数，表示邮件的序号。当不指定参数时，POP3服务器列出邮箱中所有的邮件信息；当指定参数msg#时，POP3服务器只返回序号对应的邮件信息。
			retr<SP>msg#<CRLF>	    retr 命令用于获取某封邮件的内容，参数 msg#表示邮件的序号。
			dele<SP>msg#<CRLF>	    dele 命令用于在某封邮件上设置删除标记，参数msg#表示邮件的序号。POP3服务器执行dele命令时，只是为邮件设置了删除标记，并没有真正把邮件删除掉，只有POP3客户端发出quit命令后，POP3服务器才会真正删除所有设置了删除标记的邮件。
			rest<CRLF>	    rest 命令用于清除所有邮件的删除标记。
			top<SP>msg#<SP>n<CRLF>	    top 命令用于获取某封邮件的邮件头和邮件体中的前n行内容，参数msg#表示邮件的序号，参数n表示要返回邮件的前几行内容。使用这条命令以提高 Web Mail系统（通过Web站点上收发邮件）中的邮件列表显示的处理效率，因为这种情况下不需要获取每封邮件的完整内容，而是仅仅需要获取每封邮件的邮件头信息。
			noop<CRLF>	    noop 命令用于检测POP3客户端与POP3服务器的连接情况。
			quit<CRLF>	    quit 命令表示要结束邮件接收过程，POP3服务器接收到此命令后，将删除所有设置了删除标记的邮件，并关闭与POP3客户端程序的网络连接。
		*/

		switch cmd {
		case "USER":
			userName := getSafeArg(args, 0)
			cmdError = s.Action.User(data, userName)
			if cmdError != nil {
				fmt.Fprintf(conn, "-ERR The user %s doesn't belong here!"+eol, userName)
			} else {
				fmt.Fprintf(conn, "+OK"+eol)
			}
		case "PASS":
			password := getSafeArg(args, 0)
			cmdError = s.Action.Pass(data, password)
			if cmdError != nil {
				fmt.Fprintf(conn, "-ERR Password incorrect!"+eol)
			} else {
				fmt.Fprintf(conn, "+OK User signed in"+eol)
			}
		case "STAT":
			if data.Status == TRANSACTION {
				num, size, err := s.Action.Stat(data)
				if err != nil {
					fmt.Fprintf(conn, "-ERR %s%s", err.Error(), eol)
				} else {
					fmt.Fprintf(conn, "+OK %d %d %s", num, size, eol)
				}
			}
		case "LIST":
			if data.Status == TRANSACTION {
				infos, err := s.Action.List(data, getSafeArg(args, 0))
				if err != nil {
					fmt.Fprintf(conn, "-ERR %s %s", err.Error(), eol)
				} else {
					fmt.Fprintf(conn, "+OK"+eol)
					for _, info := range infos {
						fmt.Fprintf(conn, "%d %d%s", info.Id, info.Size, eol)
					}
					fmt.Fprintf(conn, "."+eol)
				}
			}
		case "UIDL":
			if data.Status == TRANSACTION {
				id, err := strconv.ParseInt(getSafeArg(args, 0), 10, 64)
				if err != nil {
					fmt.Fprintf(conn, "-ERR %s %s", "params error", eol)
				} else {
					unionId, err := s.Action.Uidl(data, id)
					if err != nil {
						fmt.Fprintf(conn, "-ERR %s %s", err.Error(), eol)
					} else {
						fmt.Fprintf(conn, "+OK %d %s%s", id, unionId, eol)
					}
				}

			}
		case "TOP":
			if data.Status == TRANSACTION {
				id, err1 := strconv.ParseInt(getSafeArg(args, 0), 10, 64)
				line, err2 := strconv.Atoi(getSafeArg(args, 1))
				if err1 != nil || err2 != nil {
					fmt.Fprintf(conn, "-ERR %s %s", "params error", eol)
				} else {
					res, err := s.Action.Top(data, id, line)
					if err != nil {
						fmt.Fprintf(conn, "-ERR %s %s", err.Error(), eol)
					} else {
						fmt.Fprintf(conn, "+OK"+eol)
						fmt.Fprintf(conn, "%s%s", res, eol)
						fmt.Fprintf(conn, "."+eol)
					}
				}
			}
		case "RETR":
			if data.Status == TRANSACTION {
				id, err := strconv.ParseInt(getSafeArg(args, 0), 10, 64)
				if err != nil {
					fmt.Fprintf(conn, "-ERR %s %s", "params error", eol)
				} else {
					res, size, err := s.Action.Retr(data, id)
					if err != nil {
						fmt.Fprintf(conn, "-ERR %s %s", err.Error(), eol)
					} else {
						fmt.Fprintf(conn, "+OK %d%s", size, eol)
						fmt.Fprintf(conn, "%s%s", res, eol)
						fmt.Fprintf(conn, "."+eol)
					}
				}
			}
		case "DELE":
			if data.Status == TRANSACTION {
				id, err := strconv.ParseInt(getSafeArg(args, 0), 10, 64)
				if err != nil {
					fmt.Fprintf(conn, "-ERR %s %s", "params error", eol)
				} else {
					err := s.Action.Delete(data, id)
					if err != nil {
						fmt.Fprintf(conn, "-ERR %s %s", err.Error(), eol)
					} else {
						fmt.Fprintf(conn, "+OK %s", eol)
					}
				}
			}
		case "QUIT":
			if data.Status == TRANSACTION {
				s.Action.Quit(data)
				fmt.Fprintf(conn, "+OK Bye %s", eol)
				conn.Close()
			}
		case "NOOP":
			fmt.Fprintf(conn, "+OK %s", eol)
		}

	}
}

// cuts the line into command and arguments
func getCommand(line string) (string, []string) {
	line = strings.Trim(line, "\r \n")
	cmd := strings.Split(line, " ")
	return cmd[0], cmd[1:]
}

func getSafeArg(args []string, nr int) string {
	if nr < len(args) {
		return args[nr]
	}
	return ""
}
