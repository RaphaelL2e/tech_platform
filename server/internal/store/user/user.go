package user

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"tech_platform/server/internal/model/user"
	"time"
)

func (d *DataHandler) Register(u user.User) (string, error) {

	err := d.DB.Create(u).Error
	if err != nil {
		return "", err
	}

	//init userInfo
	ui := user.Userinfo{
		UserId:    u.Id,
		Name:      "polar_" + strconv.FormatInt(time.Now().Unix(), 10),
		Avatar:    user.DefaultAvatar,
		Introduce: user.DefaultIntroduce,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}

	err = d.DB.Create(ui).Error
	for err != nil {
		err = d.DB.Create(ui).Error
	}
	return u.Id, nil
}

func (d *DataHandler) Login(u user.User) (user.LoginResponse, error) {
	err := d.DB.Where("username = ?", u.Username).Where("password = ?", u.Password).First(&u).Error
	if err != nil {
		return user.LoginResponse{}, err
	}
	return user.LoginResponse{
		Status: u.Status,
		UserId: u.Id,
	}, nil
}

func (d *DataHandler) UpdateUserinfo(ui user.Userinfo) (user.Userinfo, error) {
	ui.UpdateAt = time.Now()
	err := d.DB.Model(&ui).UpdateColumns(ui).Scan(&ui).Error
	if err != nil {
		return user.Userinfo{}, err
	}
	return ui, nil
}

func (d *DataHandler) GetUserinfo(userId string)(user.Userinfo,error)  {
	ui := new(user.Userinfo)
	ui.UserId= userId
	err :=d.DB.First(&ui).Error
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return user.Userinfo{}, nil
		}
		return *ui, err
	}
	return *ui,nil
}
