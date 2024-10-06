package model

type Response struct {
	Status string      `json:"status"`
	Detail string      `json:"detail"`
	Data   interface{} `json:"data"`
}
