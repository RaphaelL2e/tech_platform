package user

import (
	"gorm.io/gorm"
	"tech_platform/server/internal/model"
	"tech_platform/server/internal/model/user"
)

type Store interface {
	Register(u user.User) (string, error)
	Login(u user.User) (user.LoginResponse, error)
	UpdateUserinfo(ui user.Userinfo)(user.Userinfo,error)
	GetUserinfo(userId string)(user.Userinfo,error)
	ListUser(lm model.ListModel)([]user.ListUser,int64,error)
}

type UserDataHandler struct {
	DB *gorm.DB
}


func Register(store Store,u user.User)(string,error){
	return store.Register(u)
}


func Login(store Store,u user.User)(user.LoginResponse,error){
	return store.Login(u)
}

func UpdateUserinfo(store Store,ui user.Userinfo)(user.Userinfo,error){
	return store.UpdateUserinfo(ui)
}

func GetUserinfo(store Store,userId string)(user.Userinfo,error){
	return store.GetUserinfo(userId)
}

func ListUser(store Store,lm model.ListModel)([]user.ListUser,int64,error){
	return store.ListUser(lm)
}