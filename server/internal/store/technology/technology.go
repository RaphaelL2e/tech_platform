package technology

import (
	"gorm.io/gorm"
	"tech_platform/server/internal/model"
	"tech_platform/server/internal/model/article"
	"tech_platform/server/internal/model/technology"
	"time"
)

type TechnologyDataHandler struct {
	DB *gorm.DB
}

func (d *TechnologyDataHandler) Add(t technology.Technology) (technology.AddTechnology, error) {
	t.CreateAt = time.Now()
	t.UpdateAt = t.CreateAt
	nt := technology.AddTechnology{}
	err := d.DB.Create(&t).Scan(&nt).Error
	if err != nil {
		return technology.AddTechnology{}, err
	}
	return nt, nil
}

func (d *TechnologyDataHandler) GetOneById(id int64) (technology.Technology, error) {

	t := new(technology.Technology)
	t.Id = int(id)
	err :=d.DB.First(t).Error
	if err!=nil{
		return *t, err
	}
	return *t,nil
}

func (d *TechnologyDataHandler) Update(t technology.Technology) (technology.Technology, error) {
	t.UpdateAt = time.Now()
	err := d.DB.Model(&t).Updates(t).Scan(&t).Error
	if err != nil {
		return technology.Technology{}, err
	}
	return t, nil
}

func (d *TechnologyDataHandler) Delete(t technology.DeleteTechnology) (bool, error) {

	err := d.DB.Delete(&technology.Technology{},"id = ?",t.Id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (d *TechnologyDataHandler) List(lm model.ListModel) ([]technology.ListTechnology, error) {
	var list []technology.ListTechnology
	err :=d.DB.Model(&technology.Technology{}).Limit(lm.PageSize).Offset((lm.PageNum-1)*lm.PageSize).Scan(&list).Error
	if err!=nil{
		return nil, err
	}
	return list, nil
}

func (d *TechnologyDataHandler) AddATT(att technology.ATT)(error){
	err :=d.DB.Table("article_technology").Create(&att).Error
	return err
}

func (d *TechnologyDataHandler) DelATT(att technology.ATT)(error){
	err :=d.DB.Table("article_technology").Delete(&att,att).Error
	return err
}

func (d *TechnologyDataHandler)ListArticles(la technology.ListArticle)(list []article.ListArticleResponse,err error){
	err =d.DB.Table("article_technology").Joins("join articles on article_technology.article_id = articles.id").Where("technology_id = ?",la.TechnologyId).Scan(&list).Error
	if err!=nil{
		return nil, err
	}
	return list,nil
}

func (d *TechnologyDataHandler)Count()(int64,error){
	count :=int64(0)
	err :=d.DB.Model(&technology.Technology{}).Count(&count).Error
	if err!=nil{
		return 0,err
	}
	return count,nil
}