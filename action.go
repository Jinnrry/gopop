package gopop

type MailInfo struct {
	Id   int64
	Size int64
}

type Action interface {
	User(ctx *Data, username string) error
	Pass(ctx *Data, password string) error
	Apop(ctx *Data, username, digest string) error
	Stat(ctx *Data) (msgNum, msgSize int64, err error)
	Uidl(ctx *Data, id int64) (string, error)
	List(ctx *Data, msg string) ([]MailInfo, error)
	Retr(ctx *Data, id int64) (string, int64, error)
	Delete(ctx *Data, id int64) error
	Rest(ctx *Data) error
	Top(ctx *Data, id int64, n int) (string, error)
	Noop(ctx *Data) error
	Quit(ctx *Data) error
}
