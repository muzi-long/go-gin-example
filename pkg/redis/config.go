package redis

type Config struct {
	Address  []string
	Password string
	Db       int
}
