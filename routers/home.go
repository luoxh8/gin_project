package routers

import (
	"gin_project/core"
	`gin_project/models`
	`github.com/astaxie/beego/orm`
	`github.com/gin-gonic/gin`
	`net/http`
)

var HomeRouter = core.Handler{
	/**
		静态接口：banner，smallIcon
	 */
	"index": func(context *gin.Context) {
		var (
			o          orm.Ormer
			banners    []*models.Banner
			smallIcons []*models.SmallIcon
		)
		o = orm.NewOrm()
		o.QueryTable("banner").All(&banners)
		o.QueryTable("small_icon").All(&smallIcons)
		context.SecureJSON(http.StatusOK, gin.H{
			"code": core.SuccessCode,
			"data": gin.H{
				"banner":    banners,
				"smallIcon": smallIcons,
			},
		})
	},
	/**
		动态接口：所有的商品列表
	 */
	"list": func(context *gin.Context) {
		var (
			goodType  string
			homeGoods []*models.HomeGoods
		)
		goodType = context.DefaultQuery("good_type", "0")
		if goodType != "0" && goodType != "1" && goodType != "2" {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseParams)
		}

		if goodType == "0" {
			orm.NewOrm().QueryTable("home_goods").Filter("like", true).All(&homeGoods)
			context.SecureJSON(http.StatusOK, gin.H{
				"code": core.SuccessCode,
				"data": homeGoods,
			})
			return
		}
		if goodType == "1" {

		}
		if goodType == "2" {

		}
	},
}
