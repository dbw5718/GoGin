package routes

import (
	"github.com/gin-gonic/gin"
	"jassue-gin/app/common/request"
	"net/http"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/user/resigter", func(c *gin.Context) {
		var form request.Resigter
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": request.GetErrorMsg(form, err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
}
