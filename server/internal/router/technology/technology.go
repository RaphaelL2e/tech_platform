package technology

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tech_platform/server/internal/model/technology"
	"tech_platform/server/internal/pkg/response"
)

func addTechnology(c *gin.Context) {
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
		return
	}
	var req technology.Technology
	err = c.Bind(&req)
	if err != nil {
		return
	}
	user_id, _ := c.Get("user_id")
	req.UserId = user_id.(string)

	resp = srv.AddTechnology(c, req)
}
func getTechnology(c *gin.Context){
	resp := response.CreateBySuccess()
	var err error
	defer func() {
		if err != nil {
			resp = response.CreateByErrorMessage(err)
		}
		if resp.Code == response.NotFoundCode.Code{
			c.JSON(http.StatusNotFound, resp)
		}else {
			c.JSON(http.StatusOK, resp)
		}
	}()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	resp = srv.GetTechnology(c,id)
}

func updateTechnology(c *gin.Context){
	resp := response.CreateBySuccess()
	var err error
	defer func() {
		if err!=nil{
			resp = response.CreateByErrorMessage(err)
		}
		c.JSON(http.StatusOK,resp)
	}()
	is_admin, _ := c.Get("is_admin")
	admin1 := is_admin.(bool)
	if !admin1 {
		resp = response.CreateByErrorCodeMessage(response.ForbiddenCode)
		return
	}
	var req technology.Technology
	err = c.Bind(&req)
	if err != nil {
		return
	}
	user_id, _ := c.Get("user_id")
	req.UserId = user_id.(string)

	resp = srv.UpdateTechnology(c,req)
}

func deleteTechnology(c *gin.Context){
	resp := response.CreateBySuccess()
	var err error
	defer func() {
		if err!=nil{
			resp = response.CreateByErrorMessage(err)
		}
		c.JSON(http.StatusOK,resp)
	}()
	is_admin, _ := c.Get("is_admin")
	admin1 := is_admin.(bool)
	if !admin1 {
		resp = response.CreateByErrorCodeMessage(response.ForbiddenCode)
		return
	}
	var req technology.DeleteTechnology
	err = c.Bind(&req)
	if err != nil {
		return
	}

	resp = srv.DeleteTechnology(c,req)
}


func listTechnology(c *gin.Context){
	resp := response.CreateBySuccess()
	var err error
	defer func() {
		if err != nil {
			resp = response.CreateByErrorMessage(err)
		}
		if resp.Code == response.NotFoundCode.Code{
			c.JSON(http.StatusNotFound, resp)
		}else {
			c.JSON(http.StatusOK, resp)
		}
	}()

	var req technology.ListModel
	err = c.Bind(&req)
	if err != nil {
		return
	}
	resp = srv.ListTechnology(c,req)
}