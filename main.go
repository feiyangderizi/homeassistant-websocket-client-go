package main

import (
	"encoding/json"
	"fmt"
	"github.com/feiyangderizi/homeassistant-websocket-client-go/model"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"gopkg.in/yaml.v2"
)

//参考文档：//模式定义参考文档：https://developers.home-assistant.io/docs/api/websocket

func main() {
	// 读取配置文件
	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal("无法读取配置文件:", err)
	}

	var config model.Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal("无法解析配置文件:", err)
	}

	//server.Init(&config)
	//return

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("%s:%d", config.HomeAssistant.IP, config.HomeAssistant.Port),
		Path:   config.HomeAssistant.Path,
	}

	fmt.Printf("连接到: %s\n", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("连接错误:", err)
	}
	defer c.Close()

	// 构建授权信息
	authMessage := model.AuthMessage{
		Type:        "auth",
		AccessToken: config.HomeAssistant.Token, // 替换为实际的访问令牌
	}

	// 将授权信息转换为JSON格式
	authJSON, err := json.Marshal(authMessage)
	if err != nil {
		log.Fatal("JSON转换错误:", err)
	}

	//log.Printf("授权信息：", authJSON)
	// 发送授权信息
	err = c.WriteMessage(websocket.TextMessage, authJSON)
	if err != nil {
		log.Fatal("发送授权信息错误:", err)
	}

	subscribeEventMessage := model.SubscribeEventMessage{
		Id:        18,
		Type:      "subscribe_events",
		EventType: "state_changed",
	}
	// 将授权信息转换为JSON格式
	subscribeJSON, err := json.Marshal(subscribeEventMessage)
	if err != nil {
		log.Fatal("JSON转换错误:", err)
	}

	// 发送订阅信息
	err = c.WriteMessage(websocket.TextMessage, subscribeJSON)
	if err != nil {
		log.Fatal("发送订阅信息错误:", err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()

			if err != nil {
				log.Println("读取错误:", err)
				return
			}
			log.Printf("收到消息: %s", message)
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("接收到中断信号，关闭连接...")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("关闭连接错误:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
