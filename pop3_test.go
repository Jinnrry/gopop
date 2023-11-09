package gopop

import (
	"errors"
	"fmt"
	"testing"
)

type action struct {
}

func (a action) User(ctx *Data, username string) error {
	ctx.User = username
	return nil
}

func (a action) Pass(ctx *Data, password string) error {
	ctx.Status = TRANSACTION
	return nil
}

func (a action) Apop(ctx *Data, username, digest string) error {
	ctx.User = username
	ctx.Status = TRANSACTION

	return nil
}

func (a action) Stat(ctx *Data) (msgNum, msgSize int64, err error) {
	return 0, 0, err
}

func (a action) Uidl(ctx *Data, id int64) (string, error) {
	return fmt.Sprintf("%d", id), nil
}

func (a action) List(ctx *Data, msg string) ([]MailInfo, error) {
	return nil, nil
}

func (a action) Retr(ctx *Data, id int64) (string, int64, error) {
	return "", 0, nil
}

func (a action) Delete(ctx *Data, id int64) error {
	return nil
}

func (a action) Rest(ctx *Data) error {
	return nil
}

func (a action) Top(ctx *Data, id int64, n int) (string, error) {
	return "", errors.New("not supported")
}

func (a action) Noop(ctx *Data) error {
	return nil
}

func (a action) Quit(ctx *Data) error {
	return nil
}

func TestServer_Start(t *testing.T) {
	s := NewPop3Server(110, "domain.com", false, action{})
	s.Start()
}
