package helper

import (
	"fmt"
	"testing"
)

type TestData struct {
	Listen       string `json:"listen"`
	UserAPIURL   string `json:"userapisrv"`
	SecretKey    string `json:"secret_key"`
	AuthorityUrl string `json:"authorityapisrv"`
	ResourceKey  string `json:"resource_key"`
}

func TestGetStructFields(t *testing.T) {
	fl, err := GetStructFields(TestData{})
	if err != nil {
		panic(err)
	}
	for _, f := range fl {
		fmt.Println(f.Name, f.Title, f.JsonTitle)
	}
}
