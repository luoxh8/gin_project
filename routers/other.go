package routers

import (
	`fmt`
	"gin_project/core"
	"github.com/gin-gonic/gin"
	`net/http`
	`time`
)

var Timer *time.Timer

var OtherRouter = core.Handler{
	/**
	    配置信息
	*/
	"config": func(c *gin.Context) {
		// goods, err := models.GetGoodsDetail(1)
		// fmt.Println(goods, err)
		c.SecureJSON(http.StatusOK, gin.H{
			"code": core.SuccessCode,
			"data": gin.H{
				"ios_test":     false,
				"force_update": false,
				"check_update": gin.H{
					"appVersion":  100,
					"changeLog":   "更新内容",
					"downloadUrl": "https://www.baidu.com",
				},
			},
		})
	},
	/**
	    时间
	*/
	"time": func(context *gin.Context) {
		var (
			currentTime time.Time
		)
		Timer = time.NewTimer(3 * time.Second)
		go func() {
			currentTime = <-Timer.C
			fmt.Println("time", <-Timer.C)
		}()
		context.JSON(200, gin.H{
			"now": currentTime,
		})
	},
}
