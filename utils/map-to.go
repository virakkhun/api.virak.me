package utils

// - MapTo
// - @arg values | type []T
// - @arg mapFunc | type func(v T) T
// - @return []T
func MapTo[T any](values []T, mapFunc func(v T) T) []T {
	var newValues []T
	for _, value := range values {
		newValue := mapFunc(value)
		newValues = append(newValues, newValue)
	}

	return newValues
}
