package utils

import (
	"strings"

	"api.virak.me/shared/models"
)

// - MapTo
// - @arg values | type []T
// - @arg mapFunc | type func(v T) T
// - @return []T
func LowerCaseKeyMap(values models.Map) models.Map {
	newValues := make(models.Map)
	for k, v := range values {
		newValues[strings.ToLower(k)] = v
	}
	return newValues
}
