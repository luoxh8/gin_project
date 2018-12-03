package core

import (
	`github.com/gin-contrib/sessions`
	`github.com/gin-contrib/sessions/cookie`
	"github.com/gin-gonic/gin"
	`net/http`
	"sync"
)

type Handler map[string]gin.HandlerFunc

type application struct {
	Router *gin.Engine
}

var (
	instance *application
	once     sync.Once
)

func Instance() *application {
	once.Do(func() {
		gin.SetMode(RunMode)
		router := gin.Default()
		router.Use(sessions.Sessions(SessionName, cookie.NewStore([]byte(SecretKey))))
		/**
			page not found
		 */
		router.NoRoute(func(context *gin.Context) { context.JSON(http.StatusNotFound, ErrSysPageNotFound) })
		instance = &application{router}
	})
	return instance
}
