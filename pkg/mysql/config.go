package mysql

type Config struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Dbname          string
	MaxIdleConn     int //空闲连接，建议 10
	MaxOpenConn     int //打开的连接， 建议 100
	ConnMaxLifeTime int // 单位：小时
}
