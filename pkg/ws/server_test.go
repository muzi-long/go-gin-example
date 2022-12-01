package ws

import (
	"testing"
)

func TestSubscribeAndPublish(t *testing.T) {

	s := NewServer(&Config{
		Port:     8080,
		Url:      "/ws",
		Address:  []string{"127.0.0.1:6379"},
		Password: "root",
	})

	s.Run()
}
