package services

import (
	"fmt"
	"github.com/lvxin0315/gg/databases"
	"testing"
)

func TestSaveOpenid(t *testing.T) {
	databases.InitMysqlDB()
	userModel, err := SaveOpenid("12345678901234567890")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(userModel.Openid)
}
