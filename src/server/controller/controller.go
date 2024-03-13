package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/service"
)

func getPaginatedResponseByCallBack(context *gin.Context, callback func(int, int) (dto.Response, error)) {
	size, page := paginated(context)
	response, err := callback(size, page)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, response)
}

func getOneResponseByCallBack(context *gin.Context, id string, callback func(string) (dto.Response, error)) {
	resp, err := callback(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, err.Error())))
		return
	}
	context.JSON(http.StatusOK, resp)
}
