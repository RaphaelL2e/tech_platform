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
