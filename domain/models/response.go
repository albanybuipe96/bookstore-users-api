package models

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"status"`
	Data    interface{} `json:"data"`
}
