package controller

import (
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
	submitDuty(*gin.Context)
	deleteDuty(*gin.Context)
	getMyGrading(*gin.Context)
	triggerPluginExecution(*gin.Context)
	getStaticReport(*gin.Context)
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
	dutyId := context.Param("dutyId")
	getOneResponseByCallBack(context, dutyId, dc.ds.GetDutyByIdResponse)
}

func (dc *dutyControllerImpl) submitDuty(context *gin.Context) {
	gradingId := context.Param("gradingId")
	dutyId := context.Param("dutyId")
	err := dc.ds.SubmitDutyResponse(gradingId, dutyId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (dc *dutyControllerImpl) deleteDuty(context *gin.Context) {
	dutyId := context.Param("dutyId")
	getOneResponseByCallBack(context, dutyId, dc.ds.DeleteDutyResponse)
}

func (dc *dutyControllerImpl) getMyGrading(context *gin.Context) {
	dutyId := context.Param("dutyId")
	username := context.MustGet("username").(string)
	resp, err := dc.ds.GetMyGradingResponse(dutyId, username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, resp)
}

func (dc *dutyControllerImpl) triggerPluginExecution(context *gin.Context) {
	dutyId := context.Param("dutyId")
	resp, err := dc.ds.ExecutePlugin(dutyId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, resp)
}
func (dc *dutyControllerImpl) getStaticReport(context *gin.Context) {
	gradingId := context.Param("gradingId")
	content, err := dc.ds.GetReportContent(gradingId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.String(http.StatusOK, content)
}

func TriggerExecution() gin.HandlerFunc {
	return dutyControllerObj.triggerPluginExecution
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

func SubmitDutyByTutee() gin.HandlerFunc {
	return dutyControllerObj.submitDuty
}

func DeleteDuty() gin.HandlerFunc {
	return dutyControllerObj.deleteDuty
}

func GetMyGradingDetail() gin.HandlerFunc {
	return dutyControllerObj.getMyGrading
}

func GetDutyReport() gin.HandlerFunc {
	return dutyControllerObj.getStaticReport
}
