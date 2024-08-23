package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseWithoutData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
