# GoPOP

A simple Go POP3 server library

```go

import (
    "github.com/Jinnrry/gopop"
    "pmail/config"
)

type action struct {
}

func (a action) User(ctx *gopop.Data, username string) error {
    //TODO implement me
    panic("implement me")
}

func (a action) Pass(ctx *gopop.Data, password string) error {
    //TODO implement me
    panic("implement me")
}

func (a action) Apop(ctx *gopop.Data, username, digest string) error {
    //TODO implement me
    panic("implement me")
}

func (a action) Stat(ctx *gopop.Data) (msgNum, msgSize int64, err error) {
    //TODO implement me
    panic("implement me")
}

func (a action) Uidl(ctx *gopop.Data, id int64) (string, error) {
    //TODO implement me
    panic("implement me")
}

func (a action) List(ctx *gopop.Data, msg string) ([]gopop.MailInfo, error) {
    //TODO implement me
    panic("implement me")
}

func (a action) Retr(ctx *gopop.Data, id int64) (string, int64, error) {
    //TODO implement me
    panic("implement me")
}

func (a action) Delete(ctx *gopop.Data, id int64) error {
    //TODO implement me
    panic("implement me")
}

func (a action) Rest(ctx *gopop.Data) error {
    //TODO implement me
    panic("implement me")
}

func (a action) Top(ctx *gopop.Data, id int64, n int) (string, error) {
    //TODO implement me
    panic("implement me")
}

func (a action) Noop(ctx *gopop.Data) error {
    //TODO implement me
    panic("implement me")
}

func (a action) Quit(ctx *gopop.Data) error {
    //TODO implement me
    panic("implement me")
}

func Start() {
    pop3instance := gopop.NewPop3Server(110, "domain.com", false, action{})
    pop3instance.Start()

}

```