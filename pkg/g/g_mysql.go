package g

import (
	"github.com/github-muzilong/go-toolkit/mysql"
	"gorm.io/gorm"

	"github.com/github-muzilong/go-gin-example/config"
)

var gDB *gorm.DB

func InitMysql(c *config.Config) {
	cfg := c.Database.Default
	mCfg := &mysql.Config{
		Host:            cfg.Host,
		Port:            cfg.Port,
		Username:        cfg.Username,
		Password:        cfg.Password,
		Dbname:          cfg.Name,
		MaxIdleConn:     cfg.MaxIdleConn,
		MaxOpenConn:     cfg.MaxOpenConn,
		ConnMaxLifeTime: cfg.ConnMaxLifeTime,
	}
	db, err := mysql.New(mCfg)
	if err != nil {
		panic(err)
	}
	gDB = db
}
