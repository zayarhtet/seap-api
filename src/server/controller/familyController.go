package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/service"
)

type FamilyController interface {
	getAllFamilies(*gin.Context)
	getAllFamiliesWithMembers(*gin.Context)
}

type familyControllerImpl struct {
	fs service.FamilyService
}

var familyControllerObj FamilyController

func initFamily() {
	if familyControllerObj != nil {
		return
	}
	familyControllerObj = &familyControllerImpl{fs: service.NewFamilyService()}
}

func (fc *familyControllerImpl) getAllFamilies(context *gin.Context) {
	size, page := paginated(context)
	response, err := fc.fs.GetAllFamiliesResponse(size, page)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, response)
}

func (fc *familyControllerImpl) getAllFamiliesWithMembers(context *gin.Context) {
	size, page := paginated(context)
	response, err := fc.fs.GetAllFamiliesWithMembersResponse(size, page)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, response)
}

func GetAllFamilies() func(*gin.Context) {
	return familyControllerObj.getAllFamilies
}

func GetAllFamiliesWithMembers() func(*gin.Context) {
	return familyControllerObj.getAllFamiliesWithMembers
}
