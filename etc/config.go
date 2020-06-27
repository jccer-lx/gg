package etc

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	APPName string       `default:"app name"`
	Port    string       `default:"8088"`
	DB      *mysqlConfig `default:"db"`
	MemDB   *memDBConfig `yaml:"memDB"`
	Wx      *WxConfig    `yaml:"wx"`
}{}

func init() {
	configor.Load(&Config, "config.yml")
}
