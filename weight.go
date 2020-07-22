package main

import (
	"fmt"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/golayui/combination"
	"reflect"
)

//通过weight生成对应的基础模板
func main() {
	t := models.Admin{}
	fieldList, err := getStructFields(t)
	if err != nil {
		panic(err)
	}
	combination.Build(fieldList, "views/tmp")
}

//从struct中取field信息
func getStructFields(data interface{}) (fieldList []*combination.Item, err error) {
	cType := reflect.TypeOf(data)
	if cType.Kind() != reflect.Struct {
		err = fmt.Errorf("no ptr")
		return
	}
	for i := 0; i < reflect.ValueOf(data).NumField(); i++ {
		sf := cType.Field(i)
		//排查没有json tag的
		if sf.Tag.Get("json") == "" {
			continue
		}
		fieldList = append(fieldList, &combination.Item{
			Title: sf.Name,
			Field: sf.Tag.Get("json"),
		})
	}
	return
}
