package main

import (
	"fmt"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/golayui/combination"
	"github.com/lvxin0315/golayui/weight"
	"reflect"
	"strings"
)

//通过weight生成对应的基础模板
func main() {
	t := models.JdItemModel{}
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
		//类型处理
		weightType := getTypeByJsonTag(sf.Tag.Get("json"))
		fieldList = append(fieldList, &combination.Item{
			Title: sf.Name,
			Field: sf.Tag.Get("json"),
			Type:  weightType,
		})
	}
	return
}

//根据json的tag名称判断类型
func getTypeByJsonTag(jsonTag string) string {
	//包含_id -> select
	if strings.Index(jsonTag, "_id") > 0 {
		return weight.Select
	}
	//包含is_ -> switch
	if strings.Index(jsonTag, "is_") >= 0 {
		return weight.Switch
	}
	//包含_radio -> radio
	if strings.Index(jsonTag, "_radio") > 0 {
		return weight.Radio
	}
	//包含_image_json -> 多图上传
	if strings.Index(jsonTag, "_image_json") > 0 {
		return weight.Upload
	}
	//包含_info -> textarea
	if strings.Index(jsonTag, "_info") > 0 {
		return weight.Textarea
	}
	//包含image -> upload
	if strings.Index(jsonTag, "_image") > 0 {
		return weight.Upload
	}
	//包含_json -> checkbox
	if strings.Index(jsonTag, "_json") > 0 {
		return weight.CheckBox
	}
	return weight.InputText
}
