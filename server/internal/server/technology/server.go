package technology

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tech_platform/server/internal/model"
	"tech_platform/server/internal/model/technology"
	"tech_platform/server/internal/pkg/response"
	"tech_platform/server/internal/store"
	technologystore "tech_platform/server/internal/store/technology"
)

type Handler struct {
}

func (h Handler) AddTechnology(c context.Context,req technology.Technology) response.ServerResponse {

	tech_store := store.FromContext(c)
	at,err :=technologystore.Add(tech_store,req)
	if err!=nil{
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccessData(at)
}

func (h Handler) GetTechnology(c *gin.Context, id int64) response.ServerResponse {
	tech_store := store.FromContext(c)
	oneById,err :=technologystore.GetOneById(tech_store,id)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return response.CreateByErrorCodeMessage(response.NotFoundCode)
		}
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccessData(oneById)
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) UpdateTechnology(c context.Context,req technology.Technology) response.ServerResponse {
	tech_store := store.FromContext(c)
	ut,err :=technologystore.Update(tech_store,req)
	if err!=nil{
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccessData(ut)
}

func (h Handler) DeleteTechnology(c *gin.Context, req technology.DeleteTechnology) response.ServerResponse {
	tech_store := store.FromContext(c)
	result,err :=technologystore.Delete(tech_store,req)
	if err!=nil {
		return response.CreateByErrorMessage(err)
	}
	if result{
		return response.CreateBySuccess()
	}else {
		return response.CreateByError()
	}

}

func (h Handler) ListTechnology(c *gin.Context, req model.ListModel) response.ServerResponse {
	tech_store := store.FromContext(c)
	list,err :=technologystore.List(tech_store,req)
	if err!=nil {
		return response.CreateByErrorMessage(err)
	}
	count,err :=technologystore.Count(tech_store)
	if err!=nil{
		count = 10
	}
	m :=make(map[string]interface{})
	m["list"] = list
	m["count"] = count
	return response.CreateBySuccessData(m)
}

func (h Handler) AddATT(c *gin.Context, req technology.ATT) response.ServerResponse {
	tech_store :=store.FromContext(c)
	err :=technologystore.AddATT(tech_store,req)
	if err!=nil {
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccess()

}

func (h Handler) DelATT(c *gin.Context, req technology.ATT) response.ServerResponse {
	tech_store :=store.FromContext(c)
	err :=technologystore.DelATT(tech_store,req)
	if err!=nil {
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccess()

}

func (h Handler) ListArticles(c *gin.Context, req technology.ListArticle) response.ServerResponse {
	tech_store :=store.FromContext(c)
	list,err :=technologystore.ListArticles(tech_store,req)
	if err!=nil {
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccessData(list)
}
