package models

type Map map[string]interface{}

type AnyMap interface{}

type GenericMap[K comparable, V int | string | float32 | float64] map[K]V
