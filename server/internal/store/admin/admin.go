package admin

import (
	"gorm.io/gorm"
	"tech_platform/server/internal/model/admin"
	"time"
)

type AdminDataHandler struct {
	DB *gorm.DB
}

func (d *AdminDataHandler) AdminLogin(u string) (string, error) {
	user_id := ""
	err := d.DB.Table("admin").Select("user_id").Where("user_id = ?", u).Pluck("user_id", &user_id).Error
	if err != nil {
		return "", err
	}
	return user_id, nil
}

func (d *AdminDataHandler) AdminAdd(a admin.Admin) (bool, error) {
	a.CreateAt =time.Now()
	a.UpdateAt =time.Now()
	err :=d.DB.FirstOrCreate(&a,a).Error
	if err!=nil{
		return false, err
	}
	return true,nil
}
