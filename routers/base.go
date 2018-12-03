package routers

import (
	"gin_project/core"
	`gin_project/middlewares`
)

func init() {
	application := core.Instance()
	application.Router.GET("/", IMRouter["index"]) // websocket

	/**
		首页
	 */
	home := application.Router.Group("/homes")
	{
		home.GET("/index", HomeRouter["index"])
		home.GET("/list", HomeRouter["list"])
	}

	/**
		商品
	 */
	goods := application.Router.Group("/goods")
	{
		goods.GET("/list", GoodsRouter["list"])
		goods.GET("/detail", GoodsRouter["retrieve"])
		goods.POST("/bid", middlewares.JWTMiddleware["JWTAuth"], middlewares.PermissionMiddleware["user"], GoodsRouter["bid"])
	}

	/**
		最新成交
	 */
	latestNews := application.Router.Group("/latestNews")
	{
		latestNews.GET("/index", LatestNewsRouter["index"])
	}

	/**
		商品分类
	 */
	category := application.Router.Group("/categories")
	{
		category.GET("/index", CategoryRouter["index"])
	}

	/**
	   用户
	*/
	user := application.Router.Group("/users")
	{
		user.POST("/third/login", UsersRouter["thirdLogin"])
		user.POST("/login", UsersRouter["login"])
		user.POST("/register", UsersRouter["register"])
		user.POST("/send/code", UsersRouter["sendCode"])
		user.GET("/info", middlewares.JWTMiddleware["JWTAuth"], UsersRouter["info"])
	}

	/**
		其他
	 */
	other := application.Router.Group("/others",middlewares.LimiterMiddleware["frequency"])
	{
		other.GET("/config", OtherRouter["config"])
		other.GET("/time", OtherRouter["time"])
	}

}
