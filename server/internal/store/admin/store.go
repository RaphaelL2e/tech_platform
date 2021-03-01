package admin

import (
	"tech_platform/server/internal/model/admin"
)

type Store interface {
	AdminLogin(u string) (string,error)
	AdminAdd(a admin.Admin) (bool,error)
}

func AdminLogin(store Store, userId string) (string,error) {
	return store.AdminLogin(userId)
}


func AdminAdd(store Store,a admin.Admin)(bool,error){
	return store.AdminAdd(a)
}