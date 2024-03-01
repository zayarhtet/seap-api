package service

import (
	"github.com/zayarhtet/seap-api/src/server/model/dto"
)

func BeforeDataResponse[T any](data *[]T, total int64, args ...int) dto.Response {
	newDataResp := dto.NewDataResponse("mg wade")

	newDataResp.Data = data
	newDataResp.Size = uint(len(*data))
	newDataResp.Total = total

	if len(args) >= 2 {
		page := args[1]
		limit := args[0]
		newDataResp.Page = page

		totalPage := total / int64(limit)
		if total%int64(limit) > 0 {
			totalPage++
		}
		newDataResp.TotalPage = int(totalPage)
	} else {
		newDataResp.Page = 1
		newDataResp.TotalPage = 1
	}
	newDataResp.StartAt = 0
	return &newDataResp
}

func BeforeErrorResponse(err *map[string]any) dto.Response {
	newErrorResp := dto.NewErrorResponse("Mg wade")
	newErrorResp.Error = err

	return &newErrorResp
}

func PrepareErrorMap(code uint, msg string) *map[string]any {
	return &map[string]any{
		"code":    code,
		"message": msg,
	}
}

func calculatePagination(total int64, size, page int64) int {
	totalPage := total / size
	if total%size > 0 {
		totalPage++
	}
	start := size * (page - 1)

	if page > totalPage || start > total {
		return -1
	}

	return int(start)
}
