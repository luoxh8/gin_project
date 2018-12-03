package middlewares

import (
	`gin_project/core`
	`gin_project/models`
	`github.com/astaxie/beego/orm`
	`github.com/gin-gonic/gin`
	`net/http`
)

var PermissionMiddleware = core.Handler{
	"admin": func(context *gin.Context) {
		var (
			claims interface{}
			user   models.User
		)
		claims, _ = context.Get("claims")
		customClaims := claims.(*CustomClaims)
		orm.NewOrm().QueryTable("user").Filter("uid", customClaims.Uid).One(&user)
		if user.AccountType != 0 {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseIllegalPermission)
			context.Abort()
			return
		}
	},
	"user": func(context *gin.Context) {
		var (
			claims interface{}
			user   models.User
		)
		claims, _ = context.Get("claims")
		customClaims := claims.(*CustomClaims)
		orm.NewOrm().QueryTable("user").Filter("uid", customClaims.Uid).One(&user)
		if user.AccountType != 1 && user.AccountType != 0 {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseIllegalPermission)
			context.Abort()
			return
		}
	},
}
