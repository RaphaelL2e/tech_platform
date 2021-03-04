package technology

import (
	"gorm.io/gorm"
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

	technology := new(technology.Technology)
	technology.Id = int(id)
	err :=d.DB.First(technology,technology).Error
	if err!=nil{
		return *technology, err
	}
	return *technology,nil
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