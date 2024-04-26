package controller

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/zayarhtet/seap-api/src/server/service"
)

type CDNController interface {
	getProfileImage(*gin.Context)
	saveGivenFiles(*gin.Context)
	downloadGivenFile(*gin.Context)
	uploadSubmittedFiles(*gin.Context)
	downloadSubmittedFile(*gin.Context)
	deleteSubmittedFile(*gin.Context)
}

type CDNControllerImpl struct {
	fs service.FamilyService
	ds service.DutyService
}

var cdnControllerObj CDNController

func initCDN() {
	if cdnControllerObj != nil {
		return
	}
	cdnControllerObj = &CDNControllerImpl{fs: service.NewFamilyService(), ds: service.NewDutyService()}
}

func (cc *CDNControllerImpl) getProfileImage(context *gin.Context) {
	idRaw := context.Param("famId")
	//username := context.MustGet("username").(string)

	path, err := cc.fs.GetFamilyProfileImagePath(idRaw)
	if err != nil {
		context.JSON(http.StatusBadRequest, service.BeforeErrorResponse(service.PrepareErrorMap(400, err.Error())))
		return
	}
	context.File(path)
}

func (cc *CDNControllerImpl) saveGivenFiles(context *gin.Context) {
	dutyId := context.Param("dutyId")
	err := context.Request.ParseMultipartForm(10 << 20) // 10 MB maximum size
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form data"})
		return
	}

	files := context.Request.MultipartForm.File["files"] // Assuming files are uploaded with key "files"

	err = cc.ds.CreateGivenFiles(files, dutyId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create assignment and upload files"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "uploaded."})
}

func (cc *CDNControllerImpl) uploadSubmittedFiles(context *gin.Context) {
	dutyId := context.Param("dutyId")
	famId := context.Param("famId")
	username := context.MustGet("username").(string)
	err := context.Request.ParseMultipartForm(10 << 20) // 10 MB maximum size
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form data"})
		return
	}
	files := context.Request.MultipartForm.File["files"]

	resp, err := cc.ds.UploadSubmittedFiles(files, dutyId, famId, username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create assignment and upload files"})
		return
	}

	context.JSON(http.StatusOK, resp)
}

func (cc *CDNControllerImpl) downloadGivenFile(context *gin.Context) {
	dutyId := context.Param("dutyId")
	fileId := context.Param("fileId")

	filePath, err := cc.ds.GetGivenFilePath(dutyId, fileId)
	if err != nil || len(filePath) == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve file path"})
		return
	}
	context.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath)) // You may want to adjust the filename as needed
	//context.Header("Content-Disposition", "attachment; filename="+fileId) // You may want to adjust the filename as needed

	context.File(filePath)
}

func (cc *CDNControllerImpl) downloadSubmittedFile(context *gin.Context) {
	dutyId := context.Param("dutyId")
	fileId := context.Param("fileId")
	familyRole := context.MustGet("familyRole").(string)
	username := context.MustGet("username").(string)

	filePath, err := cc.ds.GetSubmittedFilePath(dutyId, fileId, username, familyRole)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if len(filePath) == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve file path"})
		return
	}
	context.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath)) // You may want to adjust the filename as needed

	context.File(filePath)
}

func (cc *CDNControllerImpl) deleteSubmittedFile(context *gin.Context) {
	dutyId := context.Param("dutyId")
	fileId := context.Param("fileId")
	username := context.MustGet("username").(string)

	err := cc.ds.DeleteSubmittedFileResponse(fileId, dutyId, username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func CDNProfileImage() gin.HandlerFunc {
	return cdnControllerObj.getProfileImage
}

func SaveGivenFiles() gin.HandlerFunc {
	return cdnControllerObj.saveGivenFiles
}

func DownloadGivenFile() gin.HandlerFunc {
	return cdnControllerObj.downloadGivenFile
}

func UploadSubmittedFiles() gin.HandlerFunc {
	return cdnControllerObj.uploadSubmittedFiles
}

func DownloadSubmittedFile() gin.HandlerFunc {
	return cdnControllerObj.downloadSubmittedFile
}

func DeleteSubmittedFile() gin.HandlerFunc {
	return cdnControllerObj.deleteSubmittedFile
}
