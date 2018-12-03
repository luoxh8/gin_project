package middlewares

import (
	`gin_project/core`
	`gin_project/core/h_categories`
	`github.com/gin-gonic/gin`
	`github.com/go-redis/redis`
	`net/http`
	`time`
)

var LimiterMiddleware = core.Handler{
	"frequency": func(context *gin.Context) {
		var (
			client   *redis.Client
			ip       string
			count    int
			timeLess time.Duration
		)
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		ip = h_categories.GetIP(context.Request.RemoteAddr)
		count, _ = client.Get(ip).Int()
		timeLess = client.PTTL(ip).Val()
		if timeLess < 0 {
			client.Del(ip)
			client.Set(ip, 1, core.LimiterTime)
		} else {
			if count < core.LimiterCount {
				client.Set(ip, count+1, timeLess)
			} else {
				context.SecureJSON(http.StatusBadRequest, core.ErrBaseFrequentRequests)
				context.Abort()
				return
			}
		}
	},
}
