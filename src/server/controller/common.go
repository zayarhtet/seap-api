package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zayarhtet/seap-api/src/server/config/constant"
	"strconv"
)

func Init() {
	initRole()
}

func paginated(context *gin.Context) (int, int) {
	size, err := strconv.Atoi(context.Query("size"))
	if err != nil || size <= 0 {
		size = constant.SIZE
	}
	page, err := strconv.Atoi(context.Query("page"))
	if err != nil || page <= 0 {
		page = constant.PAGE
	}

	return size, page
}
