package article

import "tech_platform/server/internal/model/article"

type Store interface {
	GetById(id int64)(article.Article,error)
	ListArticle(req article.ListArticle)([]article.ListArticleResponse,error)
}

func GetById(store Store,id int64)(article.Article,error){
	return store.GetById(id)
}

func ListArticle(store Store,req article.ListArticle)([]article.ListArticleResponse,error){
	return store.ListArticle(req)
}