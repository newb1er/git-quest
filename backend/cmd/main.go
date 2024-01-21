package main

import (
	"fmt"
	"git-quest-be/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func SocketHandler(c *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},

		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
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

func main() {
	router := api.NewRouter()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
