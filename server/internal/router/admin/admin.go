package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tech_platform/server/internal/model/user"
	"tech_platform/server/internal/pkg/response"
)

func adminLogin(c *gin.Context) {
	resp := response.CreateBySuccess()
	var err error
	defer func() {
		if err != nil {
			resp = response.CreateByErrorMessage(err)
		}
		c.JSON(http.StatusOK, resp)
	}()

	var req user.LoginRequest
	err = c.Bind(&req)
	if err != nil {
		return
	}
	resp.Data, err = srv.AdminLogin(c, req)
}
