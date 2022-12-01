package ws

type Config struct {
	Port int
	Url  string
	// redis 配置
	Address  []string
	Password string
}
