package config

type App struct {
	Name  string `mapstructure:"name"`
	Port  int    `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}

type Default struct {
	Driver          string `mapstructure:"driver"`
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Name            string `mapstructure:"name"`
	Prefix          string `mapstructure:"prefix"`
	MaxIdleConn     int    `mapstructure:"maxIdleConn"`
	MaxOpenConn     int    `mapstructure:"maxOpenConn"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

type Database struct {
	Default Default `mapstructure:"default"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type Logger struct {
	Path       string `mapstructure:"path"`
	File       string `mapstructure:"file"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Compress   bool   `mapstructure:"compress"`
	Level      string `mapstructure:"level"`
}

type Config struct {
	App      App      `mapstructure:"app"`
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
	Logger   Logger   `mapstructure:"logger"`
}
