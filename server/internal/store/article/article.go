package article

import (
	"gorm.io/gorm"
	"tech_platform/server/internal/model/article"
	"time"
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

func (d *ArticleHandler) AddArticle(a article.Article) (article.Article, error) {
	a.CreateAt = time.Now()
	a.UpdateAt = a.CreateAt
	a.Status = 0 // 待审核
	err := d.DB.Create(&a).Error
	if err != nil {
		return article.Article{}, err
	}
	return a, err
}

func (d *ArticleHandler) UpdateArticle(a article.Article) (article.Article, error) {
	a.UpdateAt = time.Now()
	a.Status = 0 // 待审核
	err := d.DB.Model(&a).Updates(a).Scan(&a).Error
	if err != nil {
		return article.Article{}, err
	}
	return a, err
}

func (d *ArticleHandler) CheckAuthority(uid string, aid int) bool {
	var suid string
	err := d.DB.Model(&article.Article{}).Select("user_id").Where("id = ?", aid).Pluck("user_id", &suid).Error
	if err != nil {
		return false
	}
	return suid == uid
}
