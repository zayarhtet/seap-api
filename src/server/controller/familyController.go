package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/service"
)

type FamilyController interface {
	getAllFamilies(*gin.Context)
	getAllFamiliesWithMembers(*gin.Context)
	getMemberByIdWithFamilies(*gin.Context)
	saveNewFamily(*gin.Context)
	addNewMemberToFamily(*gin.Context)
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
	getPaginatedResponseByCallBack(context, fc.fs.GetAllFamiliesResponse)
}

func (fc *familyControllerImpl) getAllFamiliesWithMembers(context *gin.Context) {
	getPaginatedResponseByCallBack(context, fc.fs.GetAllFamiliesWithMembersResponse)
}

func (fc *familyControllerImpl) getMemberByIdWithFamilies(context *gin.Context) {
	idRaw := context.Param("id")
	getOneResponseByCallBack(context, idRaw, fc.fs.GetMemberByIdWithFamiliesResponse)
}

func (fc *familyControllerImpl) saveNewFamily(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	var input dto.NewFamilyRequest

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, "Invalid Input")))
		return
	}
	newFamily, err := fc.fs.SaveNewFamily(idRaw, input)
	if err != nil {
		context.JSON(http.StatusBadRequest, newFamily)
		return
	}
	context.JSON(http.StatusOK, newFamily)
}

func (fc *familyControllerImpl) addNewMemberToFamily(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	var input dto.MemberToFamilyRequest

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, "Invalid Input")))
		return
	}
	isTutor, err := fc.fs.IsTutorInFamily(idRaw, input.FamilyId)
	if err != nil || !isTutor {
		context.JSON(http.StatusUnauthorized, service.BeforeErrorResponse(service.PrepareErrorMap(401, "You are not a tutor of this family.")))
		return
	}
	family, err := fc.fs.AddMemberToFamily(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, family)
		return
	}
	context.JSON(http.StatusOK, family)
}

func GetAllFamilies() func(*gin.Context) {
	return familyControllerObj.getAllFamilies
}

func GetAllFamiliesWithMembers() func(*gin.Context) {
	return familyControllerObj.getAllFamiliesWithMembers
}

func GetMemberByIdWithFamilies() func(*gin.Context) {
	return familyControllerObj.getMemberByIdWithFamilies
}

func SaveNewFamily() func(*gin.Context) {
	return familyControllerObj.saveNewFamily
}

func AddNewMemberToFamily() func(*gin.Context) {
	return familyControllerObj.addNewMemberToFamily
}
