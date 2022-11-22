package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Client struct {
	Config *Config
	Ch     *amqp.Channel
	Queue  *amqp.Queue

	//消费者读取消息
	MsgCh chan []byte
}

func New(cfg *Config) (*Client, error) {
	res := new(Client)
	res.Config = cfg
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.Username, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	queue, err := ch.QueueDeclare(
		cfg.queueName, // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return nil, err
	}
	return &Client{
		Config: cfg,
		Ch:     ch,
		Queue:  &queue,
		MsgCh:  make(chan []byte),
	}, nil

}

// Send 发送消息
func (c *Client) Send(content string) error {
	err := c.Ch.Publish("", c.Queue.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(content)})
	if err != nil {
		return err
	}
	return nil
}

// Read 读取消息
func (c *Client) Read() chan []byte {
	msgs, err := c.Ch.Consume(
		c.Config.queueName, // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	if err != nil {
		panic(err)
	}
	go func() {
		for d := range msgs {
			c.MsgCh <- d.Body
		}
	}()
	return c.MsgCh
}
