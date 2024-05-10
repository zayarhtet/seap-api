package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/service"
)

type FamilyController interface {
	GetAllFamilies(*gin.Context)
	GetAllFamiliesWithMembers(*gin.Context)
	GetMemberByIdWithFamilies(*gin.Context)
	GetAllDutiesByFamilyId(*gin.Context)
	GetAllMembersByFamilyId(*gin.Context)
	SaveNewFamily(*gin.Context)
	AddNewMemberToFamily(*gin.Context)
	DeleteFamily(*gin.Context)
}

type familyControllerImpl struct {
	fs service.FamilyService
}

var familyControllerObj FamilyController

func initFamily() {
	if familyControllerObj != nil {
		return
	}
	fs := service.NewFamilyService()
	familyControllerObj = NewFamilyController(fs)
}

func (fc *familyControllerImpl) SetFamilyService(fs service.FamilyService) {
	fc.fs = fs
}

func NewFamilyController(fs service.FamilyService) FamilyController {
	return &familyControllerImpl{fs: fs}
}

func (fc *familyControllerImpl) GetAllFamilies(context *gin.Context) {
	getPaginatedResponseByCallBack(context, fc.fs.GetAllFamiliesResponse)
}

func (fc *familyControllerImpl) GetAllFamiliesWithMembers(context *gin.Context) {
	getPaginatedResponseByCallBack(context, fc.fs.GetAllFamiliesWithMembersResponse)
}

func (fc *familyControllerImpl) GetMemberByIdWithFamilies(context *gin.Context) {
	idRaw := context.Param("id")
	getOneResponseByCallBack(context, idRaw, fc.fs.GetMemberByIdWithFamiliesResponse)
}

func (fc *familyControllerImpl) GetAllDutiesByFamilyId(context *gin.Context) {
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

func (fc *familyControllerImpl) GetAllMembersByFamilyId(context *gin.Context) {
	idRaw := context.Param("famId")
	getOneResponseByCallBack(context, idRaw, fc.fs.GetFamilyByIdWithMembers)
}

func (fc *familyControllerImpl) SaveNewFamily(context *gin.Context) {
	idRaw := context.MustGet("username").(string)
	var familyName string = context.PostForm("familyName")
	var familyInfo string = context.PostForm("familyInfo")

	if len(familyName) == 0 {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, "Invalid Input")))
		return
	}
	var input dto.NewFamilyRequest = dto.NewFamilyRequest{FamilyName: familyName, FamilyInfo: familyInfo}
	file, _ := context.FormFile("familyIcon")

	newFamily, err := fc.fs.SaveNewFamily(idRaw, input, file)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, newFamily)
}

func (fc *familyControllerImpl) AddNewMemberToFamily(context *gin.Context) {
	famId := context.Param("famId")
	var input dto.MemberToFamilyRequest

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, "Invalid Input")))
		return
	}

	input.FamilyId = famId
	family, err := fc.fs.AddMemberToFamily(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, family)
}

func (fc *familyControllerImpl) DeleteFamily(context *gin.Context) {
	idRaw := context.Param("famId")
	getOneResponseByCallBack(context, idRaw, fc.fs.DeleteFamilyResponse)
}

func GetAllFamilies() gin.HandlerFunc {
	return familyControllerObj.GetAllFamilies
}

func GetAllFamiliesWithMembers() gin.HandlerFunc {
	return familyControllerObj.GetAllFamiliesWithMembers
}

func GetMemberByIdWithFamilies() gin.HandlerFunc {
	return familyControllerObj.GetMemberByIdWithFamilies
}

func SaveNewFamily() gin.HandlerFunc {
	return familyControllerObj.SaveNewFamily
}

func AddNewMemberToFamily() gin.HandlerFunc {
	return familyControllerObj.AddNewMemberToFamily
}

func GetAllDutiesByFamilyId() gin.HandlerFunc {
	return familyControllerObj.GetAllDutiesByFamilyId
}

func GetAllMembersByFamilyId() gin.HandlerFunc {
	return familyControllerObj.GetAllMembersByFamilyId
}

func DeleteFamily() gin.HandlerFunc {
	return familyControllerObj.DeleteFamily
}
