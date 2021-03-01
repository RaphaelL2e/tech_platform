package admin

type Store interface {
	AdminLogin(u string) (string,error)
}

func AdminLogin(store Store, userId string) (string,error) {
	return store.AdminLogin(userId)
}
