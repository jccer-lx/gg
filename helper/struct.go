package helper

import "reflect"

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
