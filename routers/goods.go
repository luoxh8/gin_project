package routers

import (
	`gin_project/core`
	`gin_project/core/h_categories`
	`gin_project/middlewares`
	`gin_project/models`
	`github.com/astaxie/beego/orm`
	`github.com/gin-gonic/gin`
	`net/http`
)

type RetrieveForm struct {
	GoodsId string `form:"goods_id" json:"goods_id" xml:"goods_id" binding:"required"`
}

var bidArray h_categories.Array

var GoodsRouter = core.Handler{
	/**
	    获取所有的商品
	*/
	"list": func(context *gin.Context) {
		var (
			goodsList []*models.Goods
		)
		orm.NewOrm().QueryTable("goods").All(&goodsList)
		context.SecureJSON(http.StatusOK, gin.H{
			"code": core.SuccessCode,
			"data": goodsList,
		})
	},
	/**
	    单个商品接口
	*/
	"retrieve": func(context *gin.Context) {
		var (
			retrieve       RetrieveForm
			goods          models.Goods
			goodsPics      []*models.GoodsBannerPics
			goodsRoomCount models.GoodsRoomCount
			err            error
		)

		/**
			1、参数校验
		 */
		if err = context.ShouldBind(&retrieve); err != nil {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseParams)
			return
		}

		if err = orm.NewOrm().QueryTable("goods").Filter("goods_id", retrieve.GoodsId).One(&goods); err != nil {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseDataNotFound)
			return
		}

		orm.NewOrm().QueryTable("goods_banner_pics").Filter("goods_id", retrieve.GoodsId).All(&goodsPics)
		orm.NewOrm().QueryTable("goods_room_count").Filter("goods_id", retrieve.GoodsId).One(&goodsRoomCount)

		goods.GoodsBannerPics = goodsPics
		goods.Onlookers = goodsRoomCount.Onlookers
		goods.Bidders = goodsRoomCount.Bidders

		/**
			2、返回数据
		 */
		context.SecureJSON(http.StatusOK, gin.H{
			"code": core.SuccessCode,
			"data": goods,
		})
	},
	/**
	    出价
	*/
	"bid": func(context *gin.Context) {
		var (
			form   RetrieveForm
			goods  models.Goods
			ub     models.UserBalance
			claims interface{}
			err    error
		)

		if err = context.ShouldBind(&form); err != nil {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseParams)
			return
		}
		if !goods.FindByGoodsId(form.GoodsId) {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseDataNotFound)
			return
		}
		claims, _ = context.Get("claims")
		customClaims := claims.(*middlewares.CustomClaims)
		if !ub.FindByUserId(customClaims.Uid) {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseDataNotFound)
			return
		}

		/**
		    出价数组判断
		*/
		if bidArray.In(customClaims.Uid) {
			context.SecureJSON(http.StatusBadRequest, core.ErrGoodsBid)
			return
		} else {
			bidArray.Append(customClaims.Uid)
		}

		/**
			参数校验成功，拿到了User，开始出价过程
		 */
		g := ub.Balance
		if g >= int64(goods.CostMoney) {
			less := g - int64(goods.CostMoney)
			ub.Balance = less
			goods.CurrentPrice += 0.1
		} else {
			context.SecureJSON(http.StatusBadRequest, core.ErrUserBalance)
			bidArray.Pop(customClaims.Uid)
			return
		}

		/**
		    数据更新，重要关系到金额
		*/
		o := orm.NewOrm()
		o.Begin()
		if _, err = o.Update(&ub); err != nil {
			o.Rollback()
			context.SecureJSON(http.StatusBadRequest, core.ErrUserUnusual)
			return
		}
		if _, err = o.Update(&goods); err != nil {
			o.Rollback()
			context.SecureJSON(http.StatusBadRequest, core.ErrGoodsUnusual)
			return
		}
		o.Commit()
		bidArray.Pop(customClaims.Uid)
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"code": core.SuccessCode,
			"msg":  "竞拍成功",
		})
	},
}
