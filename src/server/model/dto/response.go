package dto

import (
	"github.com/zayarhtet/seap-api/src/util"
)

type Response interface{}

type DataResponse struct {
	RespondedAt string `json:"respondedAt"`
	TotalPage   int    `json:"totalPage"`
	Page        int    `json:"currentPage"`
	Total       int64  `json:"totalRow"`
	Size        uint   `json:"size"`
	StartAt     uint   `json:"startAt"`
	Data        any    `json:"data"`
}

type ErrorResponse struct {
	RespondedAt string          `json:"respondedAt"`
	Error       *map[string]any `json:"error"`
}

func NewDataResponse() *DataResponse {
	return &DataResponse{RespondedAt: util.CurrentTimeString()}
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{RespondedAt: util.CurrentTimeString()}
}
