package rabbitmq

type Config struct {
	Username  string
	Password  string
	Host      string
	Port      int
	queueName string
}
