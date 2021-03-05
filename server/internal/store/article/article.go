package article

import (
	"gorm.io/gorm"
	"tech_platform/server/internal/model/article"
)

type ArticleHandler struct {
	DB *gorm.DB
}

func (d *ArticleHandler) GetById(id int64) (article.Article, error) {
	a := new(article.Article)
	a.Id = int(id)
	err := d.DB.First(a).Error
	if err != nil {
		return article.Article{}, nil
	}
	return *a, nil
}

func (d *ArticleHandler) ListArticle(req article.ListArticle) ([]article.ListArticleResponse, error) {
	list := []article.ListArticleResponse{}
	a := new(article.Article)
	a.Status = req.Status
	a.UserId = req.UserId
	err := d.DB.Model(&article.Article{}).Where(a).Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
