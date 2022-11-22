package config

import (
	"path"
	"strings"

	"github.com/spf13/viper"
)

// New 加载配置文件
func New(file string, config interface{}) error {
	vp := viper.New()
	filePath, fileName := path.Split(file)
	fileType := strings.TrimLeft(path.Ext(file), ".")
	vp.SetConfigName(fileName) //设置文件名时不要带后缀
	viper.SetConfigType(fileType)
	viper.AddConfigPath(filePath) //搜索路径可以设置多个，viper 会根据设置顺序依次查找
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	err := viper.Unmarshal(config)
	if err != nil {
		return err
	}
	return nil
}
