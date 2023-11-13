package gopop

import (
	"testing"
)

type action struct {
}

func (a action) User(session *Session, username string) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Pass(session *Session, password string) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Apop(session *Session, username, digest string) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Stat(session *Session) (msgNum, msgSize int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (a action) Uidl(session *Session, msg string) ([]UidlItem, error) {
	//TODO implement me
	panic("implement me")
}

func (a action) List(session *Session, msg string) ([]MailInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (a action) Retr(session *Session, id int64) (string, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a action) Delete(session *Session, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Rest(session *Session) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Top(session *Session, id int64, n int) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a action) Noop(session *Session) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Quit(session *Session) error {
	//TODO implement me
	panic("implement me")
}

func (a action) Capa(session *Session) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (a action) Custom(session *Session, cmd string, args []string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func TestServer_Start(t *testing.T) {
	s := NewPop3Server(110, "domain.com", false, action{})
	s.Start()
}
