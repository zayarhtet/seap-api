package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/service"
)

type DutyController interface {
	GetAllDuties(*gin.Context)
	AddNewGrade(*gin.Context)
	SaveNewDuty(*gin.Context)
	GetGradingByDutyId(*gin.Context)
	GetDutyById(*gin.Context)
	SubmitDuty(*gin.Context)
	DeleteDuty(*gin.Context)
	GetMyGrading(*gin.Context)
	TriggerPluginExecution(*gin.Context)
	GetStaticReport(*gin.Context)
	GetPluginList(*gin.Context)
}

type dutyControllerImpl struct {
	ds service.DutyService
	es service.EngineService
}

var dutyControllerObj DutyController

func initDuty() {
	if dutyControllerObj != nil {
		return
	}
	ds := service.NewDutyService()
	es := service.NewEngineService()
	dutyControllerObj = NewDutyController(ds, es)
}

func (dc *dutyControllerImpl) SetDutyService(ds service.DutyService, es service.EngineService) {
	dc.ds = ds
	dc.es = es
}

func NewDutyController(ds service.DutyService, es service.EngineService) DutyController {
	return &dutyControllerImpl{ds: ds, es: es}
}

func (dc *dutyControllerImpl) GetAllDuties(context *gin.Context) {
	getPaginatedResponseByCallBack(context, dc.ds.GetAllDutiesResponse)
}

func (dc *dutyControllerImpl) AddNewGrade(context *gin.Context) {
	var input dto.NewGradeRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, "Invalid Input")))
		return
	}
	savedGrade, err := dc.ds.AddNewGradeResponse(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, savedGrade)
}

func (dc *dutyControllerImpl) SaveNewDuty(context *gin.Context) {
	var input dao.Duty
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, "Invalid Input")))
		return
	}
	savedGrade, err := dc.ds.SaveNewDutyResponse(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.JSON(http.StatusOK, savedGrade)
}

func (dc *dutyControllerImpl) GetGradingByDutyId(context *gin.Context) {
	dutyId := context.Param("dutyId")
	getPaginatedResponseWithIdByCallBack(context, dutyId, dc.ds.GetGradingByDutyIdResponse)
}

func (dc *dutyControllerImpl) GetDutyById(context *gin.Context) {
	dutyId := context.Param("dutyId")
	getOneResponseByCallBack(context, dutyId, dc.ds.GetDutyByIdResponse)
}

func (dc *dutyControllerImpl) SubmitDuty(context *gin.Context) {
	gradingId := context.Param("gradingId")
	dutyId := context.Param("dutyId")
	err := dc.ds.SubmitDutyResponse(gradingId, dutyId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (dc *dutyControllerImpl) DeleteDuty(context *gin.Context) {
	dutyId := context.Param("dutyId")
	getOneResponseByCallBack(context, dutyId, dc.ds.DeleteDutyResponse)
}

func (dc *dutyControllerImpl) GetMyGrading(context *gin.Context) {
	dutyId := context.Param("dutyId")
	username := context.MustGet("username").(string)
	resp, err := dc.ds.GetMyGradingResponse(dutyId, username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, resp)
}

func (dc *dutyControllerImpl) TriggerPluginExecution(context *gin.Context) {
	dutyId := context.Param("dutyId")
	resp, err := dc.es.ExecuteSubmittedFile(dutyId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, resp)
}
func (dc *dutyControllerImpl) GetStaticReport(context *gin.Context) {
	gradingId := context.Param("gradingId")
	content, err := dc.ds.GetReportContent(gradingId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.String(http.StatusOK, content)
}

func (dc *dutyControllerImpl) GetPluginList(context *gin.Context) {
	content, err := dc.es.GetPluginListResponse()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, content)
}

func TriggerExecution() gin.HandlerFunc {
	return dutyControllerObj.TriggerPluginExecution
}

func GetAllDuties() gin.HandlerFunc {
	return dutyControllerObj.GetAllDuties
}

func AddNewGrade() gin.HandlerFunc {
	return dutyControllerObj.AddNewGrade
}

func SaveNewDuty() gin.HandlerFunc {
	return dutyControllerObj.SaveNewDuty
}

func GetGradingByDutyId() gin.HandlerFunc {
	return dutyControllerObj.GetGradingByDutyId
}

func GetDutyById() gin.HandlerFunc {
	return dutyControllerObj.GetDutyById
}

func SubmitDutyByTutee() gin.HandlerFunc {
	return dutyControllerObj.SubmitDuty
}

func DeleteDuty() gin.HandlerFunc {
	return dutyControllerObj.DeleteDuty
}

func GetMyGradingDetail() gin.HandlerFunc {
	return dutyControllerObj.GetMyGrading
}

func GetDutyReport() gin.HandlerFunc {
	return dutyControllerObj.GetStaticReport
}

func GetPluginList() gin.HandlerFunc {
	return dutyControllerObj.GetPluginList
}
