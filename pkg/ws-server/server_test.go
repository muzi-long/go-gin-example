package ws_server

import (
	"fmt"
	"log"
	"testing"

	"github.com/gorilla/websocket"
)

func TestServer(t *testing.T) {
	s := NewServer(8080, "/wss")
	s.On("open", func(conn *websocket.Conn, data []byte) {
		log.Printf("客户端：%p 连接上了, data: %s\n", conn, data)
		fmt.Println(s.cons.data)
	})
	s.On("message", func(conn *websocket.Conn, data []byte) {
		log.Printf("收到客户端:%p 消息：%s", conn, data)
		if string(data) == "exit" {
			s.Close(conn)
		}
		if string(data) == "all" {
			if client, ok := s.cons.data[conn]; ok {
				s.SendToClient(client, []byte("向所有conn发送消息"))
			}
		}
	})
	s.On("close", func(conn *websocket.Conn, data []byte) {
		log.Printf("客户端：%p 关闭了", conn)
		fmt.Println(s.cons.data)
	})
	s.Run()
}
