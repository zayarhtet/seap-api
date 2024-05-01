package util

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CurrentTimeString() string {
	return time.Now().Format(YYYY_MM_DDTHH_MM_SS)
}

func Encrypt(password *string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	*password = string(passwordHash)
	return nil
}

func ValidatePassword(password, existedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(existedPassword), []byte(password))
}

func NewUUID() string {
	return uuid.New().String()
}

func VerifyImageFile(file *multipart.FileHeader) error {
	// Open the uploaded file
	uploadedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer uploadedFile.Close()

	// Read the first 512 bytes to determine the file type
	buffer := make([]byte, 512)
	_, err = uploadedFile.Read(buffer)
	if err != nil {
		return err
	}

	// Reset file pointer
	_, err = uploadedFile.Seek(0, 0)
	if err != nil {
		return err
	}

	// Get the file extension
	extension := filepath.Ext(file.Filename)

	// Verify the file type based on the magic number and file extension
	if !isImage(buffer) || !isSupportedImageExtension(strings.ToLower(extension)) {
		return errors.New("uploaded file is not an image")
	}

	return nil
}

func isImage(buffer []byte) bool {
	// Check for common image file signatures
	if len(buffer) >= 3 && buffer[0] == 0xff && buffer[1] == 0xd8 && buffer[2] == 0xff {
		return true // JPEG/JFIF
	}
	if len(buffer) >= 4 && buffer[0] == 0x89 && buffer[1] == 0x50 && buffer[2] == 0x4e && buffer[3] == 0x47 {
		return true // PNG
	}
	if len(buffer) >= 2 && buffer[0] == 0x47 && buffer[1] == 0x49 {
		return true // GIF
	}
	if len(buffer) >= 4 && buffer[0] == 0x49 && buffer[1] == 0x49 && buffer[2] == 0x2a && buffer[3] == 0x00 {
		return true // TIFF
	}
	if len(buffer) >= 4 && buffer[0] == 0x42 && buffer[1] == 0x4d {
		return true // BMP
	}
	return false
}
func isSupportedImageExtension(extension string) bool {
	supportedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
	return supportedExtensions[extension]
}

const ABSOLUTE_STORAGE_PATH = "/home/zayar/seap_storage/"
const ABSOLUTE_ICONS_PATH = ABSOLUTE_STORAGE_PATH + "family-icons/"
const ABSOLUTE_GIVEN_STORAGE_PATH = ABSOLUTE_STORAGE_PATH + "given-files/"
const ABSOLUTE_SUBMITTED_STORAGE_PATH = ABSOLUTE_STORAGE_PATH + "submitted-files/"
const ABSOLUTE_REPORT_STORAGE_PATH = ABSOLUTE_STORAGE_PATH + "report-files/"

func SaveIcons(fileHeader *multipart.FileHeader, id string) error {
	return SaveFile(fileHeader, filepath.Join(ABSOLUTE_ICONS_PATH, id))
}

func createDirectoryIfNotExist(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0777)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func SaveGivenFiles(fileHeaders []*multipart.FileHeader, id string) (map[string]string, error) {
	if len(fileHeaders) == 0 {
		return map[string]string{}, nil
	}
	filePath := filepath.Join(ABSOLUTE_GIVEN_STORAGE_PATH, id)
	createDirectoryIfNotExist(filePath)
	return SaveFiles(fileHeaders, filePath)
}

func SaveSubmittedFiles(fileHeaders []*multipart.FileHeader, dutyId, username string) (map[string]string, error) {
	if len(fileHeaders) == 0 {
		return map[string]string{}, nil
	}
	filePath := filepath.Join(ABSOLUTE_SUBMITTED_STORAGE_PATH, dutyId, username)
	createDirectoryIfNotExist(filePath)
	return SaveFiles(fileHeaders, filePath)
}

func SaveFiles(fileHeaders []*multipart.FileHeader, filePath string) (map[string]string, error) {
	result := map[string]string{}
	createDirectoryIfNotExist(filePath)
	errorMessage := ""
	for _, fh := range fileHeaders {
		id := NewUUID()
		fullFilePath := filepath.Join(filePath, fh.Filename)
		if FileExists(fullFilePath) {
			continue
		}
		err := SaveFile(fh, fullFilePath)
		if err != nil {
			errorMessage = errorMessage + "error in saving " + fh.Filename + "\n"
			continue
		}
		result[id] = fh.Filename
	}

	if len(errorMessage) != 0 {
		return result, errors.New(errorMessage)
	}

	return result, nil
}

func SaveFile(fileHeader *multipart.FileHeader, filePath string) error {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new file on the server
	destination, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer destination.Close()

	// Save the uploaded file directly to the destination
	_, err = io.Copy(destination, file)
	if err != nil {
		return err
	}

	return nil
}

func GetFamilyIconAbsolutePath(fileName string) string {
	absFilePath := filepath.Join(ABSOLUTE_ICONS_PATH, fileName)
	if FileExists(absFilePath) {
		return absFilePath
	}
	return ""
}

func GetGivenFileAbsolutePath(fileName string, dutyId string) string {
	absFilePath := filepath.Join(ABSOLUTE_GIVEN_STORAGE_PATH, dutyId, fileName)
	if FileExists(absFilePath) {
		return absFilePath
	}
	return ""
}

func GetSubmittedFileAbsolutePath(dutyId, username, fileName string) string {
	absFilePath := filepath.Join(ABSOLUTE_SUBMITTED_STORAGE_PATH, dutyId, username, fileName)
	if FileExists(absFilePath) {
		return absFilePath
	}
	return ""
}

func GetIndividualDutyReport(fileName string, username, dutyId string) string {
	absFilePath := filepath.Join(ABSOLUTE_REPORT_STORAGE_PATH, dutyId, username, fileName)
	if FileExists(absFilePath) {
		return absFilePath
	}
	return ""
}

func DeleteFile(filePath string) error {
	if FileExists(filePath) {
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteDirectory(absDirPath string) error {
	err := os.RemoveAll(absDirPath)
	if err != nil {
		return err
	}
	return nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func RemoveElementsInPlace[T any](slice *[]T, condition func(T) bool) {
	newIndex := 0
	for _, item := range *slice {
		if condition(item) {
			(*slice)[newIndex] = item
			newIndex++
		}
	}
	*slice = (*slice)[:newIndex]
}
