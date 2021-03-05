package article

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tech_platform/server/internal/pkg/response"
	"tech_platform/server/internal/store"
	articlestore "tech_platform/server/internal/store/article"
)

type Handler struct {

}

func (h Handler) GetArticleById(c *gin.Context, id int64) response.ServerResponse {
	s :=store.FromContext(c)
	article,err :=articlestore.GetById(s,id)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return response.CreateByErrorCodeMessage(response.NotFoundCode)
		}
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccessData(article)
}

func NewHandler()*Handler{
	return &Handler{}
}