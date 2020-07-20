package main

import (
	"fmt"
	"github.com/lvxin0315/gg/databases"
	model "github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/golayui/weight"
	"io/ioutil"
	"os"
	"reflect"
)

type field struct {
	Name string
	Type reflect.Type
}

func main() {
	databases.InitMysqlDB()
	m := model.Admin{}
	os.MkdirAll(m.TableName(), 0777)
	th, fh := makeHtml(m)
	ioutil.WriteFile(fmt.Sprintf("%s/table.html", m.TableName()), []byte(th), 0777)
	ioutil.WriteFile(fmt.Sprintf("%s/form.html", m.TableName()), []byte(fh), 0777)

}

//获取字段和类型
func getModelField(m interface{}) []*field {
	var fList []*field
	for _, f := range databases.NewDB().NewScope(m).Fields() {
		fList = append(fList, &field{
			Name: f.DBName,
			Type: f.StructField.Struct.Type,
		})
	}
	return fList
}

//根据字段类型，生成对应的表单html结构
func makeFormHTMLStruct(f *field) *weight.FormItemWeight {
	formItemWeight := new(weight.FormItemWeight)
	formItemWeight.Label = f.Name
	formItemWeight.Item = &weight.InputTextWeight{
		Attr: weight.Attr{
			Name: f.Name,
			Id:   f.Name,
		},
		Placeholder: f.Name,
	}
	return formItemWeight
}

//根据字段类型，生成对应的表格html结构
func makeTableHTMLStruct(f *field) *weight.TableTdWeight {
	tdWeight := new(weight.TableTdWeight)
	tdWeight.Content = f.Name
	return tdWeight
}

//遍历字段，生成对应结构
func makeHtml(m interface{}) (tableHtml, bodyHtml string) {
	//table
	tableWeight := new(weight.TableWeight)
	tableTrWeight := new(weight.TableTrWeight)
	tableDataTrWeight := new(weight.TableTrWeight)
	tableWeight.Thead = tableTrWeight
	tableWeight.Tbody = append(tableWeight.Tbody, tableDataTrWeight)
	//form
	formWeight := new(weight.FormWeight)

	fieldList := getModelField(m)
	for _, f := range fieldList {
		tableTrWeight.TdList = append(tableTrWeight.TdList, makeTableHTMLStruct(f))
		tableDataTrWeight.TdList = append(tableDataTrWeight.TdList, makeTableHTMLStruct(f))
		formWeight.Children = append(formWeight.Children, makeFormHTMLStruct(f))
	}
	tableHtml, _ = tableWeight.Output()
	bodyHtml, _ = formWeight.Output()
	return
}
