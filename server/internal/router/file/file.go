package file

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tech_platform/server/internal/pkg/response"
)

func uploadfile(c *gin.Context) {
	resp :=response.CreateBySuccess()
	var err error
	defer func() {
		if err!=nil{
			resp =response.CreateByErrorMessage(err)
		}
		c.JSON(http.StatusOK,resp)
	}()
	// single file
	file, err := c.FormFile("file")
	result,err :=srv.OSSHelper.Upload(file)
	if result!=""{
		resp = response.CreateBySuccessData(result)
	}

}
