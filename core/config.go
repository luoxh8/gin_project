package core

import (
	"github.com/gin-gonic/gin"
	"time"
)

var Debug = true
var OrmForce = true
var OrmVerbose = false
var RunMode = gin.DebugMode
var AppSecretKeys = map[string]interface{}{
	"android": "128ijwoklms",
}
var LimiterTime = time.Duration(60)
var LimiterCount = 100
var SessionName = "gin_session"
var SecretKey = "02joepwkmldsa02jpoqkwmdlas"
var Mysql = "root:@/gin_project?charset=utf8"
