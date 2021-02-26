package admin

import "tech_platform/server/internal/model/user"

type Store interface {
	AdminLogin(u user.User)
}


