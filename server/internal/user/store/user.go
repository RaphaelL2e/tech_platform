package store

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"tech_platform/server/internal/user/model"
	"time"
)

func (d *DataHandler) register(u model.User) (string, error) {

	err := d.DB.Create(u).Error
	if err != nil {
		return "", err
	}

	//init userInfo
	ui := model.Userinfo{
		UserId:    u.Id,
		Name:      "polar_" + strconv.FormatInt(time.Now().Unix(), 10),
		Avatar:    model.DefaultAvatar,
		Introduce: model.DefaultIntroduce,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}

	err = d.DB.Create(ui).Error
	for err != nil {
		err = d.DB.Create(ui).Error
	}
	return u.Id, nil
}

func (d *DataHandler) login(u model.User) (model.LoginResponse, error) {
	err := d.DB.Where("username = ?", u.Username).Where("password = ?", u.Password).First(&u).Error
	if err != nil {
		return model.LoginResponse{}, err
	}
	return model.LoginResponse{
		Status: u.Status,
		UserId: u.Id,
	}, nil
}

func (d *DataHandler) updateUserinfo(ui model.Userinfo) (model.Userinfo, error) {
	ui.UpdateAt = time.Now()
	err := d.DB.Model(&ui).UpdateColumns(ui).Scan(&ui).Error
	if err != nil {
		return model.Userinfo{}, err
	}
	return ui, nil
}

func (d *DataHandler)getUserinfo(userId string)(model.Userinfo,error)  {
	ui := new(model.Userinfo)
	ui.UserId= userId
	err :=d.DB.First(&ui).Error
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return model.Userinfo{}, nil
		}
		return *ui, err
	}
	return *ui,nil
}
