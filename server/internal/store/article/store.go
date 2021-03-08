package article

import "tech_platform/server/internal/model/article"

type Store interface {
	GetById(id int64)(article.Article,error)
	ListArticle(req article.ListArticle)([]article.ListArticleResponse,error)
	AddArticle(article article.Article)(article.Article,error)
	UpdateArticle(article article.Article)(article.Article,error)
	CheckAuthority(uid string,aid int)(bool)
	DeteleArticle(aid int)(error)
	UpdateArticleStatus(a article.Article)(article.Article,error)
}

func GetById(store Store,id int64)(article.Article,error){
	return store.GetById(id)
}

func ListArticle(store Store,req article.ListArticle)([]article.ListArticleResponse,error){
	return store.ListArticle(req)
}

func AddArticle(store Store,article article.Article)(article.Article,error){
	return store.AddArticle(article)
}

func UpdateArticle(store Store,article article.Article)(article.Article,error){
	return store.UpdateArticle(article)
}

func CheckAuthority(store Store,uid string,aid int)(bool){
	return store.CheckAuthority(uid,aid)
}

func DeleteArticle(store Store,aid int)error{
	return store.DeteleArticle(aid)
}

func ReviewArticle(store Store,a article.Article)(article.Article,error){
	return store.UpdateArticleStatus(a)
}