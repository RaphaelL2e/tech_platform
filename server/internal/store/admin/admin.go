package admin

import (
	"gorm.io/gorm"
	"tech_platform/server/internal/model/user"
)

type AdminDataHandler struct {
	DB *gorm.DB
}

func (d *AdminDataHandler) AdminLogin(u user.User) (string,error) {
	user_id := ""
	err :=d.DB.Table("admin").Select("user_id").Where("user_id = ?",u.Id).Pluck("user_id",&user_id).Error
	if err!=nil {
		return "",err
	}
	return user_id, nil
}
