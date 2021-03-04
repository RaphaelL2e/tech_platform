package technology

import (
	"time"
)

type Technology struct {
	Id int `gorm:"id" gorm:"primaryKey" json:"id"`
	Name string `gorm:"name" json:"name"`
	Summary string `gorm:"summary" json:"summary"`
	Image string `gorm:"image" json:"image"`
	ImageType int `gorm:"image_type" json:"image_type"`
	Context string `gorm:"context" json:"context"`
	UserId string `gorm:"user_id" json:"user_id"`
	CreateAt time.Time `gorm:"create_at" json:"create_at"`
	UpdateAt time.Time `gorm:"update_at" json:"update_at"`
}

type AddTechnology struct {
	Technology
}

type DeleteTechnology struct {
	Id int `gorm:"id" gorm:"primaryKey" json:"id" binding:"required"`
}



type ListTechnology struct {
	Id int `gorm:"id" gorm:"primaryKey" json:"id"`
	Name string `gorm:"name" json:"name"`
	Summary string `gorm:"summary" json:"summary"`
	Image string `gorm:"image" json:"image"`
	ImageType int `gorm:"image_type" json:"image_type"`
	UserId string `gorm:"user_id" json:"user_id"`
	CreateAt time.Time `gorm:"create_at" json:"create_at"`
	UpdateAt time.Time `gorm:"update_at" json:"update_at"`
}