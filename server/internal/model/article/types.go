package article

import "time"

type Article struct {
	Id int `gorm:"id" json:"id" gorm:"primaryKey"`
	Title string `gorm:"title" json:"title"`
	Summary string `gorm:"Summary" json:"summary"`
	Image string `gorm:"image" json:"image"`
	UserId string `gorm:"user_id" json:"user_id"`
	Author string `gorm:"author" json:"author" `
	Context string `gorm:"context" json:"context"`
	Status int `gorm:"status" json:"status"`
	CreateAt time.Time `gorm:"create_at" json:"create_at"`
	UpdateAt time.Time `gorm:"update_at" json:"update_at"`
}
