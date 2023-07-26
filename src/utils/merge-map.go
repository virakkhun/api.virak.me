package utils

import (
	"api.virak.me/src/shared/models"
)

func MergeMap(dest, src models.Map) models.Map {
	hashMap := make(models.Map)

	for k, v := range dest {
		hashMap[k] = v
	}

	for k, v := range src {
		hashMap[k] = v
	}

	return hashMap
}
