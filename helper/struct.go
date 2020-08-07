package helper

import (
	"encoding/json"
	"reflect"
)

/**
 * 2结构体通过反射互转
 * @Author: cs_shuai
 * @Date: 2020-06-20
 */
func ReflectiveStructToStruct(newStruct interface{}, srcStruct interface{}) {
	v := reflect.ValueOf(newStruct).Elem()
	sv := reflect.ValueOf(srcStruct).Elem()
	st := reflect.TypeOf(srcStruct).Elem()

	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		value := sv.FieldByName(fieldInfo.Name)
		_, ok := st.FieldByName(fieldInfo.Name)
		if ok && value.Type() == v.FieldByName(fieldInfo.Name).Type() {
			v.FieldByName(fieldInfo.Name).Set(value)
		}
	}
}

/**
 * 2结构体通过反射互转(判断json类型)
 * @Author: lvxin
 * @Date: 2020-08-06
 */
func ReflectiveStructToStructWithJson(newStruct interface{}, srcStruct interface{}) {
	v := reflect.ValueOf(newStruct).Elem()
	sv := reflect.ValueOf(srcStruct).Elem()
	st := reflect.TypeOf(srcStruct).Elem()

	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		value := sv.FieldByName(fieldInfo.Name)
		_, ok := st.FieldByName(fieldInfo.Name)
		if ok {
			//类型相同
			if value.Type() == v.FieldByName(fieldInfo.Name).Type() {
				v.FieldByName(fieldInfo.Name).Set(value)
			} else {
				if value.Type().String() == "json.Number" {
					//源结构体是json类型
					//根据新结构体类型转换
					switch v.FieldByName(fieldInfo.Name).Type().Kind() {
					case reflect.Int64:
						v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(JsonNumber2Int64(value.Interface().(json.Number))))
					case reflect.Int:
						v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(JsonNumber2Int(value.Interface().(json.Number))))
					case reflect.Uint:
						v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(JsonNumber2Uint(value.Interface().(json.Number))))
					case reflect.Float64:
						v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(JsonNumber2Float64(value.Interface().(json.Number))))
					case reflect.String:
						v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value.Interface().(json.Number).String()))
					}
				} else {
					switch value.Type().Kind() {
					case reflect.String: //源string
						switch v.FieldByName(fieldInfo.Name).Type().Kind() {
						case reflect.Int64:
							v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(String2Int64(value.Interface().(string))))
						case reflect.Int:
							v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(String2Int(value.Interface().(string))))
						case reflect.Uint:
							v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(String2Uint(value.Interface().(string))))
						case reflect.Float64:
							v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(String2Float64(value.Interface().(string))))
						}
					case reflect.Int64:
					}
				}
			}
		}
	}
}
