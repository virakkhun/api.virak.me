package utils

import (
	"api.virak.me/shared/models"
)

func MergeMap(dest, src interface{}) models.Map {
	hashMap := make(models.Map)

	destData, _ := StructToMap(dest)
	srcData, _ := StructToMap(src)

	for k, v := range srcData {
		hashMap[k] = v
	}

	for k, v := range destData {
		hashMap[k] = v
	}

	return hashMap
}
