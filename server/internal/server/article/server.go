package article

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tech_platform/server/internal/model/article"
	"tech_platform/server/internal/pkg/response"
	"tech_platform/server/internal/store"
	articlestore "tech_platform/server/internal/store/article"
	userstore "tech_platform/server/internal/store/user"
)

type Handler struct {
}

func (h Handler) GetArticleById(c *gin.Context, id int64) response.ServerResponse {
	s := store.FromContext(c)
	a, err := articlestore.GetById(s, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.CreateByErrorCodeMessage(response.NotFoundCode)
		}
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccessData(a)
}

func (h Handler) List(c *gin.Context, req article.ListArticle) response.ServerResponse {
	s := store.FromContext(c)
	list, err := articlestore.ListArticle(s, req)

	if err != nil {
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccessData(list)
}

func (h Handler) AddArticle(c *gin.Context, req article.Article) response.ServerResponse {
	s := store.FromContext(c)
	user, err := userstore.GetUserinfo(s, req.UserId)
	req.Author = user.Name
	a1, err := articlestore.AddArticle(s, req)


	if err != nil {
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccessData(a1)

}

func (h Handler) UpdateArticle(c *gin.Context, req article.Article) response.ServerResponse {
	s := store.FromContext(c)
	//判断是否有权限更新
	result :=articlestore.CheckAuthority(s,req.UserId,req.Id)
	if !result{
		return response.CreateByErrorCodeMessage(response.ForbiddenCode)
	}
	user, err := userstore.GetUserinfo(s, req.UserId)
	req.Author = user.Name
	a1, err := articlestore.UpdateArticle(s, req)
	if err != nil {
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccessData(a1)
}

func (h Handler) DeleteArticle(c *gin.Context, req article.Article) response.ServerResponse {
	s:=store.FromContext(c)
	admin,_ := c.Get("is_admin")
	isAdmin := admin.(bool)
	if !isAdmin{
		result :=articlestore.CheckAuthority(s,req.UserId,req.Id)
		if !result{
			return response.CreateByErrorCodeMessage(response.ForbiddenCode)
		}
	}
	aid :=req.Id
	err :=articlestore.DeleteArticle(s,aid)
	if err!=nil{
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccess()
}

func NewHandler() *Handler {
	return &Handler{}
}
