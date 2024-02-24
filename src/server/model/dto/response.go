package dto

import (
	"time"
)

type Response[T any] struct {
	RespondedAt time.Time `json:"respondedAt"`
	Username    string    `json:"username"`
	Total       uint      `json:"total"`
	MaxResults  uint      `json:"maxResults"`
	StartAt     uint      `json:"startAt"`
	Data        T         `json:"data"`
}

func NewResponse[T any]() *Response[T] {
	return &Response[T]{RespondedAt: time.Now()}
}
