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
	getAllDutiesByFamilyId(*gin.Context)
	getAllMembersByFamilyId(*gin.Context)
	saveNewFamily(*gin.Context)
	addNewMemberToFamily(*gin.Context)
	deleteFamily(*gin.Context)
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

func (fc *familyControllerImpl) getAllDutiesByFamilyId(context *gin.Context) {
	idRaw := context.Param("famId")
	username := context.MustGet("username").(string)

	//getOneResponseByCallBack(context, idRaw, fc.fs.GetFamilyByIdWithDuties)
	resp, err := fc.fs.GetFamilyByIdWithDuties(idRaw, username)
	if err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, err.Error())))
		return
	}
	context.JSON(http.StatusOK, resp)
}

func (fc *familyControllerImpl) getAllMembersByFamilyId(context *gin.Context) {
	idRaw := context.Param("famId")
	getOneResponseByCallBack(context, idRaw, fc.fs.GetFamilyByIdWithMembers)
}

func (fc *familyControllerImpl) saveNewFamily(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	var familyName string = context.PostForm("familyName")
	var familyInfo string = context.PostForm("familyInfo")

	if len(familyName) == 0 || len(familyInfo) == 0 {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, "Invalid Input")))
		return
	}
	var input dto.NewFamilyRequest = dto.NewFamilyRequest{FamilyName: familyName, FamilyInfo: familyInfo}
	file, _ := context.FormFile("familyIcon")

	newFamily, err := fc.fs.SaveNewFamily(idRaw, input, file)
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

func (fc *familyControllerImpl) deleteFamily(context *gin.Context) {
	idRaw := context.Param("famId")
	getOneResponseByCallBack(context, idRaw, fc.fs.DeleteFamilyResponse)
}

func GetAllFamilies() gin.HandlerFunc {
	return familyControllerObj.getAllFamilies
}

func GetAllFamiliesWithMembers() gin.HandlerFunc {
	return familyControllerObj.getAllFamiliesWithMembers
}

func GetMemberByIdWithFamilies() gin.HandlerFunc {
	return familyControllerObj.getMemberByIdWithFamilies
}

func SaveNewFamily() gin.HandlerFunc {
	return familyControllerObj.saveNewFamily
}

func AddNewMemberToFamily() gin.HandlerFunc {
	return familyControllerObj.addNewMemberToFamily
}

func GetAllDutiesByFamilyId() gin.HandlerFunc {
	return familyControllerObj.getAllDutiesByFamilyId
}

func GetAllMembersByFamilyId() gin.HandlerFunc {
	return familyControllerObj.getAllMembersByFamilyId
}

func DeleteFamily() gin.HandlerFunc {
	return familyControllerObj.deleteFamily
}
