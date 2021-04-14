package admin

import (
	"tech_platform/server/internal/model"
	"tech_platform/server/internal/model/admin"
)

type Store interface {
	AdminLogin(u string) (string,error)
	AdminAdd(a admin.Admin) (bool,error)
	AdminList(req model.ListModel) ([]admin.AdminModel,int64,error)
}

func AdminLogin(store Store, userId string) (string,error) {
	return store.AdminLogin(userId)
}


func AdminAdd(store Store,a admin.Admin)(bool,error){
	return store.AdminAdd(a)
}

func AdminList(store Store,req model.ListModel)([]admin.AdminModel,int64,error){
	return store.AdminList(req)
}