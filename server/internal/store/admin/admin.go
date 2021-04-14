package admin

import (
	"gorm.io/gorm"
	"tech_platform/server/internal/model"
	"tech_platform/server/internal/model/admin"
	"time"
)

type AdminDataHandler struct {
	DB *gorm.DB
}

func (d *AdminDataHandler) AdminLogin(u string) (string, error) {
	user_id := ""
	err := d.DB.Table("admins").Select("user_id").Where("user_id = ?", u).Pluck("user_id", &user_id).Error
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

func (d *AdminDataHandler)AdminList(req model.ListModel) ([]admin.AdminModel,int64,error){
	list := []admin.AdminModel{}
	count :=int64(0)
	err :=d.DB.Table("admins").Select("admins.id","admins.user_id","userinfos.name","userinfos.avatar","admins.create_at","admins.update_at").Joins("JOIN userinfos ON admins.user_id = userinfos.user_id").Offset(((req.PageNum - 1) * req.PageSize)).Order("create_at").Scan(&list).Count(&count).Error
	if err != nil {
		return nil,0, err
	}
	return list,count, nil
}
