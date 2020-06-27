package etc

type WxConfig struct {
	AppID          string `yaml:"appID" env:"WxConfigAppID"`
	AppSecret      string `yaml:"appSecret" env:"WxConfigAppSecret"`
	Token          string `yaml:"token" env:"WxConfigToken"`
	EncodingAESKey string `yaml:"encodingAESKey" env:"WxConfigEncodingAESKey"`
}
