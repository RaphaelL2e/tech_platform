package technology

import (
	"gorm.io/gorm"
	"tech_platform/server/internal/model/technology"
	"time"
)

type TechnologyDataHandler struct {
	DB *gorm.DB
}

func (d *TechnologyDataHandler) Add(t technology.Technology) (technology.AddTechnology,error) {
	t.CreateAt = time.Now()
	t.UpdateAt = t.CreateAt
	nt :=technology.AddTechnology{}
	err :=d.DB.Create(&t).Scan(&nt).Error
	if err!=nil{
		return technology.AddTechnology{},err
	}
	return nt,nil
}
