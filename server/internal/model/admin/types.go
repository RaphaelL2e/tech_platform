package admin

import "time"

type Admin struct {
	id int `gorm:"primaryKey"`
	UserId string `gorm:"user_id" json:"user_id"`
	CreateAt time.Time `gorm:"create_at"`
	UpdateAt time.Time `gorm:"update_at"`
}

type AdminModel struct {
	id int `gorm:"primaryKey"`
	UserId string `gorm:"user_id" json:"user_id"`
	Name      string    `gorm:"name" json:"name"`
	Avatar    string    `gorm:"avatar" json:"avatar"`
	CreateAt time.Time `gorm:"create_at"`
	UpdateAt time.Time `gorm:"update_at"`

}
