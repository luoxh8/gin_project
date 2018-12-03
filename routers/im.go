package routers

import (
	`fmt`
	`gin_project/core`
	`gin_project/core/h_categories`
	`gin_project/socket`
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	`time`
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

var IMRouter = core.Handler{
	"index": func(context *gin.Context) {

		var (
			wsConn     *websocket.Conn
			err        error
			data       []byte
			connection *socket.Connection
		)

		if wsConn, err = upgrader.Upgrade(context.Writer, context.Request, nil); err != nil {
			context.JSON(500, gin.H{
				"code": 500,
				"msg":  "header里面没有upgrade字段",
			})
			goto ERR
		}

		if connection, err = socket.InitConnection(wsConn); err != nil {
			goto ERR
		}

		go func() {
			var (
				err error
			)
			for {
				if err = connection.WriteMessage([]byte("pong")); err != nil {
					goto ERR
				}
				time.Sleep(5 * time.Second)
			}
		ERR:
			connection.Close()
		}()

		for {

			var (
				_data *h_categories.Dict
			)

			if data, err = connection.ReadMessage(); err != nil {
				goto ERR
			} else {
				if _data, err = h_categories.ByteToDict(data); err != nil {
					fmt.Println(_data)
					var (
						err error
					)
					if err = connection.WriteMessage(data); err != nil {
						goto ERR
					}
				}
			}

			if err = connection.WriteMessage(data); err != nil {
				goto ERR
			}
		}
	ERR:
		connection.Close()
	},
}
