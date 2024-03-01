package dto

import (
	"github.com/zayarhtet/seap-api/src/server/util"
)

type Response interface{}

type DataResponse struct {
	RespondedAt string `json:"respondedAt"`
	Username    string `json:"username"`
	TotalPage   int    `json:"totalPage"`
	Page        int    `json:"currentPage"`
	Total       int64  `json:"totalRow"`
	Size        uint   `json:"size"`
	StartAt     uint   `json:"startAt"`
	Data        any    `json:"data"`
}

type ErrorResponse struct {
	RespondedAt string          `json:"respondedAt"`
	Username    string          `json:"username"`
	Error       *map[string]any `json:"error"`
}

func NewDataResponse(uname string) *DataResponse {
	return &DataResponse{RespondedAt: util.CurrentTimeString(), Username: uname}
}

func NewErrorResponse(uname string) *ErrorResponse {
	return &ErrorResponse{RespondedAt: util.CurrentTimeString(), Username: uname}
}
