package model

type BaseResponse struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
}
