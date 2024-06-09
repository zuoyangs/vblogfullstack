package common

import (
	"time"
)

func NewScoppe() *Scope {
	return &Scope{
		CreatedAt: time.Now().Unix(),
	}
}

type Scope struct {
	Uname string `json:"Username"`
}
