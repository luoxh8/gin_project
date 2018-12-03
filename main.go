package main

import (
	"gin_project/core"
	_ "gin_project/models"
	_ "gin_project/routers"
	`github.com/gin-gonic/gin`
)

func main() {
	application := core.Instance()
	gin.SetMode(core.RunMode)
	application.Router.Run(":5566")
}
