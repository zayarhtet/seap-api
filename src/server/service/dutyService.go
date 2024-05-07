package service

import (
	"errors"
	"mime/multipart"
	"os"
	"path"
	"time"

	"github.com/zayarhtet/seap-api/src/server/model/dao"
	"github.com/zayarhtet/seap-api/src/server/model/dto"
	"github.com/zayarhtet/seap-api/src/server/repository"
	"github.com/zayarhtet/seap-api/src/util"
)

type DutyService interface {
	GetAllDutiesResponse(int, int) (dto.Response, error)
	GetAllDutiesByMemberResponse(string) (dto.Response, error)
	AddNewGradeResponse(dto.NewGradeRequest) (dto.Response, error)
	SaveNewDutyResponse(dao.Duty) (dto.Response, error)
	GetGradingByDutyIdResponse(string, int, int) (dto.Response, error)
	GetDutyByIdResponse(string) (dto.Response, error)
	CreateGivenFiles([]*multipart.FileHeader, string) error
	GetGivenFilePath(string, string) (string, error)
	UploadSubmittedFiles([]*multipart.FileHeader, string, string, string) (dto.Response, error)
	GetSubmittedFilePath(string, string, string, string) (string, error)
	SubmitDutyResponse(string, string) error

	DeleteSubmittedFileResponse(string, string, string) error
	DeleteDutyResponse(string) (dto.Response, error)

	GetMyGradingResponse(string, string) (dto.Response, error)

	GetReportContent(string) (string, error)

	SaveInputFiles([]*multipart.FileHeader, string) error
}

type dutyServiceImpl struct {
	dr repository.DutyRepository
	fr repository.FamilyRepository
}

func (ds dutyServiceImpl) GetRowCount() *int64 {
	return ds.dr.GetRowCount()
}

func (ds dutyServiceImpl) GetAllDutiesResponse(size, page int) (dto.Response, error) {
	var newResp dto.Response

	total, offset := calculateOffset(ds, size, page)

	var data *[]dao.Duty
	if offset == -1 {
		data = &[]dao.Duty{}
	} else {
		data = ds.dr.GetAllDuties(offset, size)
	}

	newResp = BeforeDataResponse[dao.Duty](data, *total, size, page)

	return newResp, nil
}

func (ds dutyServiceImpl) GetAllDutiesByMemberResponse(username string) (dto.Response, error) {
	var data *dao.MyDuty = &dao.MyDuty{Username: username}

	duties := ds.dr.GetDutiesByUsername(data)

	util.RemoveElementsInPlace[dao.MyDuty](duties, func(duty dao.MyDuty) bool {
		return time.Now().After(duty.Duty_.PublishingDate)
	})

	newResp := BeforeDataResponse[dao.MyDuty](duties, 1)
	return newResp, nil
}

func (ds dutyServiceImpl) AddNewGradeResponse(request dto.NewGradeRequest) (dto.Response, error) {
	var updatedGradingMap map[string]any = map[string]any{"points": request.Points, "has_graded": true}
	if len(request.GradeComment) != 0 {
		updatedGradingMap["grade_comment"] = request.GradeComment
	}
	var grading *dao.Grading = &dao.Grading{GradingId: request.GradingId}
	err := ds.dr.UpdateGrading(updatedGradingMap, grading)

	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(400, err.Error())), err
	}
	return "HELLO SUCCESS", nil
}

func (ds dutyServiceImpl) SaveNewDutyResponse(duty dao.Duty) (dto.Response, error) {
	duty.DutyId = util.NewUUID()
	err := ds.dr.SaveDuty(&duty)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(400, err.Error())), err
	}

	var memberList *dao.FamilyWithMembers = &dao.FamilyWithMembers{
		FamilyId: duty.FamilyId,
	}
	err = ds.fr.GetFamilyById(memberList)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(400, err.Error())), err
	}

	var newGrades []*dao.Grading

	for _, member := range memberList.Members {
		if member.MemberRole.Name != "tutee" {
			continue
		}
		newGrades = append(newGrades, &dao.Grading{
			GradingId: util.NewUUID(),
			DutyId:    duty.DutyId,
			Username:  member.Username,
			FamilyId:  duty.FamilyId,
		})
	}
	err = ds.dr.CreateGrades(newGrades)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(400, err.Error())), err
	}
	return duty, err
}

func (ds dutyServiceImpl) CreateGivenFiles(files []*multipart.FileHeader, dutyId string) error {
	var err error
	var result map[string]string
	result, err = util.SaveGivenFiles(files, dutyId)

	var givenFileList []*dao.GivenFile

	for id, fileName := range result {
		var eachFile *dao.GivenFile = &dao.GivenFile{
			FileId:   id,
			DutyId:   dutyId,
			FilePath: fileName,
		}
		givenFileList = append(givenFileList, eachFile)
	}

	err = ds.dr.InsertGivenFilesMetadata(givenFileList)
	if err != nil {
		return err
	}

	return nil
}

func (ds dutyServiceImpl) GetGradingByDutyIdResponse(dutyId string, size, page int) (dto.Response, error) {
	var newResp dto.Response

	total, offset := calculateOffset(ds, size, page)

	var data *[]dao.Grading
	if offset == -1 {
		data = &[]dao.Grading{}
	} else {
		condition := &dao.Grading{DutyId: dutyId}
		data = ds.dr.GetAllGradingByDutyId(condition, offset, size)
	}

	newResp = BeforeDataResponse[dao.Grading](data, *total, size, page)

	return newResp, nil
}

func (ds dutyServiceImpl) GetDutyByIdResponse(dutyId string) (dto.Response, error) {
	var duty *dao.Duty = &dao.Duty{
		DutyId: dutyId,
	}
	err := ds.dr.GetDutyById(duty)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "Username does not exist.")), err
	}
	newResp := BeforeDataResponse[dao.Duty](&[]dao.Duty{*duty}, 1)
	return newResp, nil
}

func (ds dutyServiceImpl) GetGivenFilePath(dutyId string, fileId string) (string, error) {
	var givenFile *dao.GivenFile = &dao.GivenFile{
		FileId: fileId,
	}

	err := ds.dr.GetGivenFileById(givenFile)
	if err != nil {
		return "", err
	}

	return util.GetGivenFileAbsolutePath(givenFile.FilePath, dutyId), nil
}

func (ds dutyServiceImpl) UploadSubmittedFiles(files []*multipart.FileHeader, dutyId, famId, username string) (dto.Response, error) {
	var duty *dao.Duty = &dao.Duty{
		DutyId: dutyId,
	}
	err := ds.dr.GetDutyById(duty)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "record not found.")), err
	}

	closedTime, _ := time.Parse(util.YYYY_MM_DDTHH_MM_SS, duty.ClosingDate.String())

	if time.Now().After(closedTime) {
		return BeforeErrorResponse(PrepareErrorMap(400, "due date expired.")), err
	}

	var grading *dao.Grading = &dao.Grading{
		FamilyId: famId,
		DutyId:   dutyId,
		Username: username,
	}
	err = ds.dr.GetGradingByStructCondition(grading, grading)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "record not found.")), err
	}

	if grading.Submitted {
		var subFile *dao.SubmittedFile = &dao.SubmittedFile{
			GradingId: grading.GradingId,
		}

		allFiles := ds.dr.GetAllSubmittedFilesMetadata(subFile)

		return BeforeDataResponse[dao.SubmittedFile](allFiles, 1), nil
	}

	result, err := util.SaveSubmittedFiles(files, dutyId, username)
	var subFileList []*dao.SubmittedFile

	for id, fileName := range result {
		var eachFile *dao.SubmittedFile = &dao.SubmittedFile{
			FileId:      id,
			GradingId:   grading.GradingId,
			FilePath:    fileName,
			SubmittedAt: util.NewWrapTime(time.Now()),
		}
		subFileList = append(subFileList, eachFile)
	}

	err = ds.dr.InsertSubmittedFilesMetadata(subFileList)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(404, "record not found.")), err
	}

	var subFile *dao.SubmittedFile = &dao.SubmittedFile{
		GradingId: grading.GradingId,
	}

	allFiles := ds.dr.GetAllSubmittedFilesMetadata(subFile)

	return BeforeDataResponse[dao.SubmittedFile](allFiles, 1), nil
}

func (ds dutyServiceImpl) GetSubmittedFilePath(dutyId string, fileId string, username string, familyRole string) (string, error) {
	var subFile *dao.SubmittedFile = &dao.SubmittedFile{
		FileId: fileId,
	}

	err := ds.dr.GetSubmittedFileById(subFile)
	if err != nil {
		return "", err
	}

	var grading *dao.Grading = &dao.Grading{
		DutyId:    dutyId,
		GradingId: subFile.GradingId,
	}
	err = ds.dr.GetGradingByStructCondition(grading, grading)
	if err != nil {
		return "", err
	}

	if familyRole == "tutee" && username != grading.Username {
		return "", errors.New("unauthorized file access")
	} else {
		username = grading.Username
	}

	return util.GetSubmittedFileAbsolutePath(dutyId, username, subFile.FilePath), nil
}

func (ds dutyServiceImpl) DeleteSubmittedFileResponse(fileId, dutyId, username string) error {
	var duty *dao.Duty = &dao.Duty{
		DutyId: dutyId,
	}
	err := ds.dr.GetDutyById(duty)
	if err != nil {
		return err
	}

	closedTime, _ := time.Parse(util.YYYY_MM_DDTHH_MM_SS, duty.ClosingDate.String())

	if time.Now().After(closedTime) {
		return err
	}

	var subFile *dao.SubmittedFile = &dao.SubmittedFile{
		FileId: fileId,
	}
	err = ds.dr.GetSubmittedFileById(subFile)
	if err != nil {
		return err
	}

	var grading *dao.Grading = &dao.Grading{
		DutyId:    dutyId,
		GradingId: subFile.GradingId,
	}
	err = ds.dr.GetGradingByStructCondition(grading, grading)
	if err != nil {
		return err
	}

	if grading.Username != username || grading.Submitted {
		return errors.New("unauthorized")
	}
	filePath := subFile.FilePath
	err = ds.dr.DeleteSubmittedFileById(subFile)
	if err != nil {
		return errors.New("unauthorized")
	}

	err = util.DeleteFile(util.GetSubmittedFileAbsolutePath(dutyId, username, filePath))
	if err != nil {
		return err
	}

	return nil
}

func (ds dutyServiceImpl) SubmitDutyResponse(gradingId string, dutyId string) error {
	var duty *dao.Duty = &dao.Duty{
		DutyId: dutyId,
	}
	err := ds.dr.GetDutyById(duty)
	if err != nil {
		return err
	}

	closedTime, _ := time.Parse(util.YYYY_MM_DDTHH_MM_SS, duty.ClosingDate.String())

	if time.Now().After(closedTime) {
		return err
	}

	var updatedGradingMap map[string]any = map[string]any{"submitted": true}

	var grading *dao.Grading = &dao.Grading{GradingId: gradingId}
	err = ds.dr.UpdateGrading(updatedGradingMap, grading)

	if err != nil {
		return err
	}
	return nil
}

func (ds dutyServiceImpl) DeleteDutyResponse(dutyId string) (dto.Response, error) {
	var duty *dao.Duty = &dao.Duty{DutyId: dutyId}
	err := ds.dr.DeleteDutyById(duty)
	if err != nil {
		return "", err
	}
	err = util.DeleteDirectory(path.Join(util.ABSOLUTE_GIVEN_STORAGE_PATH(), dutyId))
	if err != nil {
		return nil, err
	}
	err = util.DeleteDirectory(path.Join(util.ABSOLUTE_SUBMITTED_STORAGE_PATH(), dutyId))
	if err != nil {
		return nil, err
	}
	return "success", nil
}

func (ds dutyServiceImpl) GetMyGradingResponse(dutyId string, username string) (dto.Response, error) {
	var grading *dao.Grading = &dao.Grading{
		DutyId:   dutyId,
		Username: username,
	}
	err := ds.dr.GetGradingByStructCondition(grading, grading)
	if err != nil {
		return BeforeErrorResponse(PrepareErrorMap(400, err.Error())), err
	}
	return BeforeDataResponse[dao.Grading](&[]dao.Grading{*grading}, 1), nil
}

func (ds dutyServiceImpl) GetReportContent(gradingId string) (string, error) {
	var grading *dao.Grading = &dao.Grading{
		GradingId: gradingId,
	}
	err := ds.dr.GetGradingByStructCondition(grading, grading)
	if err != nil {
		return "", err
	}

	reportPath := util.GetIndividualDutyReport("log.html", grading.Username, grading.DutyId)

	if len(reportPath) == 0 {
		return "hello", nil
	}

	htmlContent, err := os.ReadFile(reportPath)
	if err != nil {
		return "", err
	}
	return string(htmlContent), nil
}

func (ds dutyServiceImpl) SaveInputFiles(headers []*multipart.FileHeader, dutyId string) error {
	_, err := util.SaveInputFiles(headers, dutyId)
	return err
}

func NewDutyService() DutyService {
	return &dutyServiceImpl{dr: repository.DutyRepositoryImpl{}, fr: repository.FamilyRepositoryImpl{}}
}
