package lib

import (
	"encoding/json"
	"os"
	"sync"
)

type SeePluginStandardLibrary struct {
	memoryFile string
	dir        string
}

func (pc *SeePluginStandardLibrary) AddMemoryFile(message string) {
	pc.memoryFile = message
}

func (pc *SeePluginStandardLibrary) SetDir(dir string) {
	pc.dir = dir
}

func (pc *SeePluginStandardLibrary) GetDir() string {
	return pc.dir
}

var fileReadMutexes = make(map[string]*sync.Mutex)
var fileReadMutexesMutex sync.Mutex

// ReadFileConcurrentlySafe reads the content of the file safely for concurrent access
func (pc *SeePluginStandardLibrary) ReadFileConcurrentlySafe(filePath string) ([]byte, error) {
	fileReadMutexesMutex.Lock()
	fileMutex, ok := fileReadMutexes[filePath]
	if !ok {
		fileMutex = &sync.Mutex{}
		fileReadMutexes[filePath] = fileMutex
	}
	fileReadMutexesMutex.Unlock()

	fileMutex.Lock()
	defer fileMutex.Unlock()

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

// ReadJSONFile reads the content of a JSON file and parses it into a map
func (pc *SeePluginStandardLibrary) ReadJSONFile(filePath string) (map[string]interface{}, error) {
	fileContent, err := pc.ReadFileConcurrentlySafe(filePath)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(fileContent, &data); err != nil {
		return nil, err
	}

	return data, nil
}
