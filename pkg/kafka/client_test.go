package kafka

import (
	"fmt"
	"testing"
)

func TestProducer_Send(t *testing.T) {
	cfg := &Config{
		Host:  "127.0.0.1:9092",
		Topic: "cdr",
	}

	p, err := NewProducer(cfg)
	if err != nil {
		t.Fatal(err)
	}
	err = p.Send("ddd", "a")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(err)
}

func TestConsumer_Read(t *testing.T) {
	cfg := &Config{
		Host:  "127.0.0.1:9092",
		Topic: "cdr",
	}
	c, err := NewConsumer(cfg)
	if err != nil {
		t.Fatal(err)
	}
	for i := range c.Read() {
		fmt.Println(i)
	}
}
