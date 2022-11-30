package ws

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Server struct {
	Port    int    // 运行端口
	Url     string // 运行的连接地址
	upgrade websocket.Upgrader
	events  struct {
		mu   sync.Mutex
		data map[string]func(conn *websocket.Conn, data []byte)
	}
	cons struct {
		mu   sync.Mutex
		data map[*websocket.Conn]string
	}
}

func NewServer(port int, url string) *Server {
	return &Server{
		Port: port,
		Url:  url,
		upgrade: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		events: struct {
			mu   sync.Mutex
			data map[string]func(conn *websocket.Conn, data []byte)
		}{mu: sync.Mutex{}, data: make(map[string]func(conn *websocket.Conn, data []byte))},
		cons: struct {
			mu   sync.Mutex
			data map[*websocket.Conn]string
		}{mu: sync.Mutex{}, data: make(map[*websocket.Conn]string)},
	}
}

func (s *Server) Run() {
	url := "/"
	pushUrl := "/push"
	if s.Url != "" {
		url = s.Url
		pushUrl = strings.TrimRight(url, "/") + pushUrl
	}
	http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		conn, err := s.upgrade.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		data := r.URL.Query().Get("data")
		go s.handleConnOpen(conn, []byte(data))
		go s.handleConnMessage(conn)
		go s.handleConnClose(conn)
	})
	http.HandleFunc(pushUrl, func(w http.ResponseWriter, r *http.Request) {
		to := r.URL.Query().Get("to")
		content := r.URL.Query().Get("content")
		s.SendToClient(to, []byte(content))
		w.WriteHeader(http.StatusOK)
	})
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.Port), nil)
	if err != nil {
		panic(err)
	}
}

func (s *Server) On(event string, f func(conn *websocket.Conn, data []byte)) {
	s.events.mu.Lock()
	defer s.events.mu.Unlock()
	if _, ok := s.events.data[event]; !ok {
		s.events.data[event] = f
	}
}

// 收到客户端连接建立时
func (s *Server) handleConnOpen(conn *websocket.Conn, data []byte) {
	s.events.mu.Lock()
	if f, ok := s.events.data["open"]; ok {
		if len(data) > 0 {
			s.cons.mu.Lock()
			s.cons.data[conn] = string(data)
			s.cons.mu.Unlock()
		}
		f(conn, data)
	}
	s.events.mu.Unlock()
}

// 收到客户端消息时
func (s *Server) handleConnMessage(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("handleConnMessage error", err)
			return
		}
		s.events.mu.Lock()
		if f, ok := s.events.data["message"]; ok {
			f(conn, message)
		}
		s.events.mu.Unlock()
	}
}

// 收到客户端关闭时
func (s *Server) handleConnClose(conn *websocket.Conn) {
	for {
		time.Sleep(time.Second)
		err := conn.WriteMessage(websocket.TextMessage, []byte("keep heart"))
		if err != nil {
			log.Println("handleConnClose error", err)

			s.cons.mu.Lock()
			delete(s.cons.data, conn)
			s.cons.mu.Unlock()

			s.events.mu.Lock()
			if f, ok := s.events.data["close"]; ok {
				f(conn, nil)
			}
			s.events.mu.Unlock()

			return
		}
	}
}

// Send 服务端主动向客户端发送消息
func (s *Server) Send(conn *websocket.Conn, data []byte) error {
	err := conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}
	return nil
}

// SendToClient 向客户端所有的conn发送消息
func (s *Server) SendToClient(client string, content []byte) {
	for conn, data := range s.cons.data {
		if client == data {
			err := s.Send(conn, content)
			if err != nil {
				continue
			}
		}
	}
}

// Close 服务端主动关闭某个客户端
func (s *Server) Close(conn *websocket.Conn) {
	_ = conn.Close()
}
