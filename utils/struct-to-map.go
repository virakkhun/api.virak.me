package utils

import (
	"errors"
	"reflect"

	"api.virak.me/shared/models"
)

// @ref to https://github.com/Klathmon/StructToMap/blob/3d0229e2dce7fa9c1d30993603145e4932abf608/structtomap.go#L31
func StructToMap(str interface{}) (models.Map, error) {

	hashMap := make(models.Map)

	sType := StructType(str)

	if sType.Kind() != reflect.Struct {
		return hashMap, errors.New("[str] is not a struct or a pointer to a struct")
	}

	for i := 0; i < sType.NumField(); i++ {
		fieldName := sType.Field(i).Name
		value := reflect.ValueOf(str)
		hashMap[fieldName] = value.FieldByName(fieldName).Interface()
	}

	return hashMap, nil
}

func StructType(str interface{}) reflect.Type {
	sType := reflect.TypeOf(str)

	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
	}

	return sType
}
