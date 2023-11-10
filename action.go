package gopop

type MailInfo struct {
	Id   int64
	Size int64
}

type Action interface {
	User(ctx *Session, username string) error
	Pass(ctx *Session, password string) error
	Apop(ctx *Session, username, digest string) error
	Stat(ctx *Session) (msgNum, msgSize int64, err error)
	Uidl(ctx *Session, id int64) (string, error)
	List(ctx *Session, msg string) ([]MailInfo, error)
	Retr(ctx *Session, id int64) (string, int64, error)
	Delete(ctx *Session, id int64) error
	Rest(ctx *Session) error
	Top(ctx *Session, id int64, n int) (string, error)
	Noop(ctx *Session) error
	Quit(ctx *Session) error
	Capa(ctx *Session) ([]string, error)
}
