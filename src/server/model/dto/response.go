package dto

import (
	"time"
)

type Response struct {
	RespondedAt time.Time 	`json:"respondedAt"`
	Username    string    	`json:"username"`
	Total       uint      	`json:"total"`
	Size  		uint      	`json:"size"`
	StartAt     uint      	`json:"startAt"`
	Data        any      	`json:"data"`
}

func NewResponse() *Response {
	return &Response{RespondedAt: time.Now()}
}
