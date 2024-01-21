package api_routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsRoutes = []Routes{
	{
		Path:    "",
		Method:  http.MethodGet,
		Handler: websocketHandler,
	},
}

func InitWsRoutes(r *gin.RouterGroup) {
	for _, route := range wsRoutes {
		r.Handle(route.Method, route.Path, route.Handler)
	}
}

func upgradeWebsocket(c *gin.Context) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},

		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return nil, fmt.Errorf("error upgrading websocket: %v", err)
	}

	return ws, nil
}

func websocketHandler(c *gin.Context) {
	ws, err := upgradeWebsocket(c)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := ws.Close()
		if err != nil {
			panic(err)
		}
	}()

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			if msgType == -1 {
				break
			}

			panic(err)
		}

		fmt.Printf("Message Type: %d, Message: %s\n", msgType, string(msg))

		err = ws.WriteJSON(struct {
			Reply string `json:"reply"`
		}{
			Reply: string(msg),
		})
		if err != nil {
			panic(err)
		}
	}
}
