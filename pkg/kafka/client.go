package kafka

import (
	"strings"

	"github.com/Shopify/sarama"
)

type (
	Producer struct {
		Client *sarama.AsyncProducer
		Config *Config
	}
	Consumer struct {
		Client *sarama.Consumer
		Config *Config
		ch     chan string
	}
)

// NewProducer 生产者
func NewProducer(cfg *Config) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	producer, err := sarama.NewAsyncProducer(strings.Split(cfg.Host, ","), config)
	if err != nil {
		return nil, err
	}
	return &Producer{
		Client: &producer,
		Config: cfg,
	}, nil
}

// Send 生产者生产消费
func (p *Producer) Send(content string, key string) error {
	msg := &sarama.ProducerMessage{
		Topic: p.Config.Topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(content),
	}
	(*p.Client).Input() <- msg
	select {
	case <-(*p.Client).Successes():
		return nil
	case fail := <-(*p.Client).Errors():
		return fail.Err
	}
}

// NewConsumer 消费者
func NewConsumer(cfg *Config) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Version = sarama.V0_11_0_2
	// consumer
	consumer, err := sarama.NewConsumer(strings.Split(cfg.Host, ","), config)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		Client: &consumer,
		Config: cfg,
		ch:     make(chan string),
	}, nil
}

// 消费者消费消息
func (c *Consumer) Read() chan string {
	consumer := *c.Client
	partitions, err := consumer.Partitions(c.Config.Topic)
	if err != nil {
		panic(err)
	}
	for _, partition := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(c.Config.Topic, partition, sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		go func(pc *sarama.PartitionConsumer) {
			defer (*pc).Close()
			for {
				select {
				case msg := <-(*pc).Messages():
					c.ch <- string(msg.Value)
				case <-(*pc).Errors():
					return
				}
			}
		}(&partitionConsumer)
	}
	return c.ch
}
