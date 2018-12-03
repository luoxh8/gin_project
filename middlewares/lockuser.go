package middlewares

import (
	`gin_project/core`
	`gin_project/models`
	`github.com/astaxie/beego/orm`
	`github.com/gin-gonic/gin`
	`net/http`
)

type LockForm struct {
	Uid string `form:"uid" json:"uid" xml:"uid" binding:"required"`
}

var UserMiddleware core.Handler = core.Handler{
	"lock": func(context *gin.Context) {
		var (
			lock LockForm
			user models.User
			err  error
		)

		if err = context.ShouldBind(&lock); err != nil {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseParams)
			return
		}
		orm.NewOrm().QueryTable("user").Filter("uid", lock.Uid).One(&user)
		user.Lock = true
		orm.NewOrm().Update(&user)
	},
	"checkLock": func(context *gin.Context) {
		var (
			claims interface{}
			user   models.User
		)
		claims, _ = context.Get("claims")
		customClaims := claims.(*CustomClaims)
		orm.NewOrm().QueryTable("user").Filter("uid", customClaims.Uid).One(&user)

	},
}
