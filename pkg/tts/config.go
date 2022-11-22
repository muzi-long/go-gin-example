package tts

type Config struct {
	AccessKeyId     string `yaml:"access_key_id" mapstructure:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret" mapstructure:"access_key_secret"`
	AppKey          string `yaml:"app_key" mapstructure:"app_key"`
}
