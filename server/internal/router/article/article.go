package article

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tech_platform/server/internal/model/article"
	"tech_platform/server/internal/pkg/response"
)

func getArticleById(c *gin.Context) {
	resp := response.CreateBySuccess()
	var err error
	defer func() {
		if err != nil {
			resp = response.CreateByErrorMessage(err)
		}
		c.JSON(http.StatusOK, resp)
	}()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	resp =srv.GetArticleById(c,id)
}

func list(c *gin.Context) {
	resp := response.CreateBySuccess()
	var err error
	defer func() {
		if err != nil {
			resp = response.CreateByErrorMessage(err)
		}
		c.JSON(http.StatusOK, resp)
	}()
	var req article.ListArticle
	err = c.Bind(&req)
	if err!=nil {
		return
	}

	resp = srv.List(c,req)
}
