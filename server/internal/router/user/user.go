package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tech_platform/server/internal/model"
	"tech_platform/server/internal/model/user"
	"tech_platform/server/internal/pkg/response"
)

func register(c *gin.Context) {
	resp := response.CreateBySuccess()
	var err error
	defer func() {
		if err != nil {
			resp = response.CreateByErrorMessage(err)
		}
		c.JSON(http.StatusOK, resp)
	}()

	var req user.RegisterRequest
	err = c.Bind(&req)
	if err != nil {
		return
	}
	m := make(map[string]string)
	m["userId"], err = srv.Register(c, req)
	resp.Data = m
}

func login(c *gin.Context) {
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

	resp.Data, err = srv.Login(c, req)
}

func updateUserInfo(c *gin.Context) {
	resp := response.CreateBySuccess()
	var err error
	defer func() {
		if err != nil {
			resp = response.CreateByErrorMessage(err)
		}
		c.JSON(http.StatusOK, resp)
	}()

	var req user.Userinfo
	err = c.Bind(&req)
	if err != nil {
		return
	}
	userId, _ := c.Get("user_id")
	req.UserId = userId.(string)

	resp.Data, err = srv.UpdateUserinfo(c, req)
}

func getUserinfo(c *gin.Context) {

	resp := response.CreateBySuccess()
	defer func() {
		c.JSON(http.StatusOK, resp)
	}()

	userId := c.Param("userId")

	ui, err := srv.GetUserinfo(c, userId)
	if err != nil {
		resp = response.CreateByErrorMessage(err)
	}
	fmt.Printf("%v", ui)
	if "" == ui.UserId {
		resp = response.CreateBySuccessMessage("not found userinfo")
	} else {
		resp = response.CreateBySuccessData(ui)
	}

}

func listUser(c *gin.Context) {
	resp := response.CreateBySuccess()
	var err error
	defer func() {
		if err != nil {
			resp = response.CreateByErrorMessage(err)
		}
		if resp.Code == response.NotFoundCode.Code {
			c.JSON(http.StatusNotFound, resp)
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}()

	is_admin, _ := c.Get("is_admin")
	admin1 := is_admin.(bool)
	if !admin1 {
		resp = response.CreateByErrorCodeMessage(response.ForbiddenCode)
		return
	}
	var req model.ListModel
	err = c.Bind(&req)
	if err != nil {
		return
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.PageNum <= 0 {
		req.PageNum = 1
	}

	resp = srv.ListUser(c, req)

}
