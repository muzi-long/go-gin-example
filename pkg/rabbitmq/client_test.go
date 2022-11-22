package rabbitmq

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

var cfg = &Config{
	Username:  "guest",
	Password:  "guest",
	Host:      "127.0.0.1",
	Port:      5672,
	queueName: "aaa",
}

func TestClient_Send(t *testing.T) {
	client, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	tt := time.NewTicker(time.Second)
	for {
		select {
		case <-tt.C:
			err = client.Send(uuid.NewString())
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	t.Log("success")
}

func TestClient_Read(t *testing.T) {
	client, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	for msg := range client.Read() {
		t.Log(string(msg))
	}
}
