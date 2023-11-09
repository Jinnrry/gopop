package gopop

import (
	"testing"
	"time"
)

type action struct {
}

func (a action) User(ctx *Data, username string) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Pass(ctx *Data, password string) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Apop(ctx *Data, username, digest string) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Stat(ctx *Data) (msgNum, msgSize int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (a action) Uidl(ctx *Data, id int64) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a action) List(ctx *Data, msg string) ([]MailInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (a action) Retr(ctx *Data, id int64) (string, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a action) Delete(ctx *Data, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Rest(ctx *Data) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Top(ctx *Data, id int64, n int) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a action) Noop(ctx *Data) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Quit(ctx *Data) error {
	//TODO implement me
	panic("implement me")
}

func TestServer_Start(t *testing.T) {
	s := NewPop3Server(110, "", false, action{})
	go s.Start()

	time.Sleep(2 * time.Second)

	s.Stop()

	time.Sleep(10 * time.Second)

}
