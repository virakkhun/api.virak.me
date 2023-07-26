package utils

import "api.virak.me/src/shared/models"

func DeleteMultiKeys(value *models.Map, keys []string) {
	for _, k := range keys {
		delete(*value, k)
	}
}
