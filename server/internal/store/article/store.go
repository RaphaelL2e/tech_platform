package article

import "tech_platform/server/internal/model/article"

type Store interface {
	GetById(id int64)(article.Article,error)
}

func GetById(store Store,id int64)(article.Article,error){
	return store.GetById(id)
}