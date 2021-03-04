package user

import (
	"time"
)

const (
	ACTIVE           = 1
	Forbidden        = 0
	DefaultAvatar    = "https://lh3.googleusercontent.com/-kWpOOLkyhXQ/AAAAAAAAAAI/AAAAAAAAAAA/AMZuucmdhtDbsqtO7w6WSVZRUeTxSHQIPw/photo.jpg"
	DefaultIntroduce = "Stay hungry, Stay foolish."
)

type User struct {
	Id       string    `gorm:"id" gorm:"primaryKey"`
	Username string    `gorm:"username"`
	Password string    `gorm:"password"`
	Status   int       `gorm:"status"`
	CreateAt time.Time `gorm:"create_at"`
	UpdateAt time.Time `gorm:"update_at"`
}

type Userinfo struct {
	UserId    string    `gorm:"primaryKey" json:"user_id" `
	Name      string    `gorm:"name" json:"name"`
	Avatar    string    `gorm:"avatar" json:"avatar"`
	Introduce string    `gorm:"introduce" json:"introduce"`
	CreateAt  time.Time `gorm:"create_at" json:"create_at"`
	UpdateAt  time.Time `gorm:"update_at" json:"update_at"`
}

type RegisterRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginResponse struct {
	UserId string `form:"userId" json:"userId"`
	Token  string `json:"token"`
	Status int    `json:"status"`
}

type UpdateUserinfoRequest struct {
	Userinfo
}

type ListUser struct {
	UserId    string    `gorm:"primaryKey" json:"user_id" `
	Username  string    `gorm:"username" json:"username"`
	Status    int       `gorm:"status" json:"status"`
	Name      string    `gorm:"name" json:"name"`
	Avatar    string    `gorm:"avatar" json:"avatar"`
	Introduce string    `gorm:"introduce" json:"introduce"`
	CreateAt  time.Time `gorm:"create_at" json:"create_at"`
	UpdateAt  time.Time `gorm:"update_at" json:"update_at"`
}
