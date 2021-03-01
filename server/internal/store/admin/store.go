package admin

import "tech_platform/server/internal/model/user"

type Store interface {
	AdminLogin(u user.User) (string,error)
}

func AdminLogin(store Store, u user.User) (string,error) {
	return store.AdminLogin(u)
}
