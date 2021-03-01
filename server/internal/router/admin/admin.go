package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tech_platform/server/internal/model/admin"
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
	resp = srv.AdminLogin(c, req)
}

func adminAdd(c *gin.Context) {
	resp := response.CreateBySuccess()
	var err error
	defer func() {
		if err != nil {
			resp = response.CreateByErrorMessage(err)
		}
		c.JSON(http.StatusOK, resp)
	}()
	is_admin, _ := c.Get("is_admin")
	admin1 := is_admin.(bool)
	if !admin1 {
		resp = response.CreateByErrorCodeMessage(response.ForbiddenCode)
	}
	var req admin.Admin
	err = c.Bind(&req)
	if err != nil {
		return
	}
	resp = srv.AdminAdd(c,req)
}
