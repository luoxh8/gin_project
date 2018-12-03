package routers

import (
	`gin_project/core`
	`github.com/gin-gonic/gin`
	`net/http`
)

type TestForm struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

var TestRouter = core.Handler{
	"test": func(context *gin.Context) {

		var json TestForm

		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"code":  400,
				"msg":   "",
				"error": err.Error(),
			})
			return
		}
		if json.User != "mamu" || json.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"status": "you are logged in",
		})
	},
}
