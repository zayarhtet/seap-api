package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/service"
)

type DutyController interface {
	getAllDuties(*gin.Context)
	addNewGrade(*gin.Context)
	saveNewDuty(*gin.Context)
	getGradingByDutyId(*gin.Context)
	getDutyById(*gin.Context)
}

type dutyControllerImpl struct {
	ds service.DutyService
}

var dutyControllerObj DutyController

func initDuty() {
	if dutyControllerObj != nil {
		return
	}
	dutyControllerObj = &dutyControllerImpl{ds: service.NewDutyService()}
}

func (dc *dutyControllerImpl) getAllDuties(context *gin.Context) {
	getPaginatedResponseByCallBack(context, dc.ds.GetAllDutiesResponse)
}

func (dc *dutyControllerImpl) addNewGrade(context *gin.Context) {
	var input dto.NewGradeRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, "Invalid Input")))
		return
	}
	savedGrade, err := dc.ds.AddNewGradeResponse(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, savedGrade)
		return
	}
	context.JSON(http.StatusOK, savedGrade)
}

func (dc *dutyControllerImpl) saveNewDuty(context *gin.Context) {
	var input dao.Duty
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, "Invalid Input")))
		return
	}
	savedGrade, err := dc.ds.SaveNewDutyResponse(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, savedGrade)
		return
	}
	context.JSON(http.StatusOK, savedGrade)
}

func (dc *dutyControllerImpl) getGradingByDutyId(context *gin.Context) {
	dutyId := context.Param("dutyId")
	getPaginatedResponseWithIdByCallBack(context, dutyId, dc.ds.GetGradingByDutyIdResponse)
}

func (dc *dutyControllerImpl) getDutyById(context *gin.Context) {
	// check if he is a member of dutyId and family
	dutyId := context.Param("dutyId")
	getOneResponseByCallBack(context, dutyId, dc.ds.GetDutyByIdResponse)
}

func GetAllDuties() gin.HandlerFunc {
	return dutyControllerObj.getAllDuties
}

func AddNewGrade() gin.HandlerFunc {
	return dutyControllerObj.addNewGrade
}

func SaveNewDuty() gin.HandlerFunc {
	return dutyControllerObj.saveNewDuty
}

func GetGradingByDutyId() gin.HandlerFunc {
	return dutyControllerObj.getGradingByDutyId
}

func GetDutyById() gin.HandlerFunc {
	return dutyControllerObj.getDutyById
}
