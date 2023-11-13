package gopop

import (
	"testing"
)

type action struct {
}

func (a action) User(session *Session, username string) error {
	session.User = username
	return nil
}

func (a action) Pass(session *Session, password string) error {
	session.Status = TRANSACTION
	return nil
}

func (a action) Apop(session *Session, username, digest string) error {
	session.User = username
	session.Status = TRANSACTION
	return nil
}

func (a action) Stat(session *Session) (msgNum, msgSize int64, err error) {
	return 0, 0, err
}

func (a action) Uidl(session *Session, msg string) ([]UidlItem, error) {
	return nil, nil
}

func (a action) List(session *Session, msg string) ([]MailInfo, error) {
	return nil, nil
}

func (a action) Retr(session *Session, id int64) (string, int64, error) {
	return "", 0, nil
}

func (a action) Delete(session *Session, id int64) error {
	return nil
}

func (a action) Rest(session *Session) error {
	return nil
}

func (a action) Top(session *Session, id int64, n int) (string, error) {
	return "", nil
}

func (a action) Noop(session *Session) error {
	return nil
}

func (a action) Quit(session *Session) error {
	return nil
}

func (a action) Capa(session *Session) ([]string, error) {
	return nil, nil
}

func (a action) Custom(session *Session, cmd string, args []string) ([]string, error) {
	return nil, nil
}

func TestServer_Start(t *testing.T) {
	s := NewPop3Server(110, "domain.com", false, action{})
	s.Start()
}
