package server

import (
	"fmt"
	"github.com/feiyangderizi/homeassistant-websocket-client-go/model"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Init(config *model.Config) {
	http.HandleFunc(config.HomeAssistant.Path, handleWebSocket)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%v", config.HomeAssistant.Port), nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
		log.Println("websocket服务端开始监听")
	}()

	select {}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		// 发送数据给客户端
		err := conn.WriteMessage(websocket.TextMessage, []byte("Hello, client!"))
		if err != nil {
			log.Println("Write error:", err)
			return
		}

		// 等待 2 秒钟
		time.Sleep(2 * time.Second)
	}
}
