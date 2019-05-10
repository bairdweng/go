package models

import "reflect"

type Relations struct {
	Id         int
	OtherId    int
	Relation   string
	RelationCn string
	CreateTime string
	RecordId   int
	UserInfo   Users
}

// 通过反射的方式设置user的值
func SetValueToStruct(key string, value interface{}, relation Relations) Relations {
	v := reflect.ValueOf(&relation).Elem().FieldByName(key)
	if v.IsValid() {
		v.Set(reflect.ValueOf(value))
	}
	return relation
}
