package app

import (
	"github.com/gin-gonic/gin"
	"jassue-gin/app/common/request"
	"jassue-gin/app/common/response"
	"jassue-gin/app/services"
)

func Register(c *gin.Context) {
	var form request.Resigter
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
