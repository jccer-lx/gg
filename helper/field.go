package helper

import (
	"fmt"
	"github.com/lvxin0315/gg/impl"
	"reflect"
)

//从struct中取field信息
func GetStructFields(data interface{}) (fieldList []*impl.DataField, err error) {
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
		fieldList = append(fieldList, &impl.DataField{
			Name:      sf.Name,
			Title:     sf.Name,
			JsonTitle: sf.Tag.Get("json"),
		})
	}
	return
}
