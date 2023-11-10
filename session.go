package gopop

import "context"

type Status int8

const (
	UNAUTHORIZED Status = 1
	TRANSACTION  Status = 2
	UPDATE       Status = 3
)

type Session struct {
	Status    Status
	User      string
	DeleteIds []int64
	Ctx       context.Context
}
