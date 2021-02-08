package model

import (
	"time"
)

type User struct {
	Id       string
	Account  string
	Password string
	Status   int
	CreateAt time.Time
	UpdateAt time.Time
}

type Userinfo struct {
	Id        int
	UserId    string
	Avatar    string
	Introduce string
	CreateAt  time.Time
	UpdateAt  time.Time
}
