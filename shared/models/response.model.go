package models

type BaseResponse struct {
	Message    interface{}
	Data       interface{}
	StatusCode int
}
