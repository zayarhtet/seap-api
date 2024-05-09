package lib

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/zayarhtet/seap-api/src/util"
)

type SeePluginStandardLibrary struct {
	dutyId   string
	username string
	htmlBody HTMLElement
}

var (
	fileReadMutexes      = make(map[string]*sync.Mutex)
	fileReadMutexesMutex sync.Mutex
)

func (pc *SeePluginStandardLibrary) InitializeLibrary(inputDir string) {
	pc.dutyId = filepath.Base(inputDir)
	pc.htmlBody = HTMLElement{nodeName: "div", classes: "flex flex-col gap-y-6 gap-x-4"}
}

// ReadFileConcurrentlySafe reads the content of the file safely for concurrent access
func (pc *SeePluginStandardLibrary) ReadFileConcurrentlySafe(filePath string) ([]byte, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []byte(""), err
	}
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
func (pc *SeePluginStandardLibrary) ReadJSONFileAsStruct(filePath string, dest any) error {
	fileContent, err := pc.ReadFileConcurrentlySafe(filePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileContent, dest); err != nil {
		return err
	}

	return nil
}

func (pc *SeePluginStandardLibrary) ReadDirectory(path string) []os.DirEntry {
	if !util.FileExists(path) {
		return nil
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return []os.DirEntry{}
	}
	return entries
}

// ReadProgrammingFileWithoutComment reads lines from a file and filters out comments
func (pc *SeePluginStandardLibrary) ReadProgrammingFileWithoutComment(filePath, singleLineComment, multiLineStart, multiLineEnd string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	insideMultiLineComment := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, multiLineStart) {
			insideMultiLineComment = true
		}

		if insideMultiLineComment {
			if strings.Contains(line, multiLineEnd) {
				insideMultiLineComment = false
			}
			continue
		}

		if index := strings.Index(line, singleLineComment); index != -1 {
			line = line[:index]
		}
		//line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// WriteLinesToFile writes lines to a new file specified by filePath
func (pc *SeePluginStandardLibrary) WriteLinesToFile(lines []string, filePath string) error {
	err := util.CreateDirectoryIfNotExist(filepath.Dir(filePath))
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (pc *SeePluginStandardLibrary) CreateAndWriteFileInTemp(lines []string, fileName string) string {
	tempPath := pc.GetNewTemporaryDirectory()
	tempFile := filepath.Join(tempPath, fileName)
	err := pc.WriteLinesToFile(lines, tempFile)
	if err != nil {
		return ""
	}
	return tempPath
}

func (pc *SeePluginStandardLibrary) GetNewTemporaryDirectory() string {
	tempDirName := fmt.Sprintf("temp_%s", util.NewUUID())

	newTempDirPath := filepath.Join(util.ABSOLUTE_TEMP_PATH(), tempDirName)

	err := os.Mkdir(newTempDirPath, 0755)
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
		return ""
	}
	return newTempDirPath
}

// ExecuteCommandWithTimeout executes the OS command and wait the output.
func (pc *SeePluginStandardLibrary) ExecuteCommandWithTimeout(command ...string) (string, string, error) {
	cmd := exec.Command(command[0], command[1:]...)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return "", "", fmt.Errorf("error starting command: %v", err)
	}

	select {
	case <-ctx.Done():
		return stdout.String(), stderr.String(), fmt.Errorf("command execution timed out")
	case err := <-func() <-chan error {
		errCh := make(chan error, 1)
		go func() {
			errCh <- cmd.Wait()
		}()
		return errCh
	}():
		if err != nil {
			return stdout.String(), stderr.String(), fmt.Errorf("command execution failed: %v", err)
		}
		return stdout.String(), stderr.String(), nil
	}
}

func (pc *SeePluginStandardLibrary) SetUsername(targetDir string) {
	pc.username = filepath.Base(targetDir)
}

func (pc *SeePluginStandardLibrary) ReportAddHTMLTable(tableHeadAndData *[][]string) {}

func (pc *SeePluginStandardLibrary) ReportAddMiniHeader(header string) {
	pc.htmlBody.AddChild(pc.getHeaderElement(header))
}

func (pc *SeePluginStandardLibrary) ReportAddParagraph(para string) {
	pc.htmlBody.AddChild(pc.getParagraphElement(para))
}

func (pc *SeePluginStandardLibrary) ReportAddHorizontalBar() {
	element := HTMLElement{
		nodeName: "div",
		classes:  "my-6 border-2 border-sky-500 border-dashed",
	}
	pc.htmlBody.AddChild(element)
}

func (pc *SeePluginStandardLibrary) ReportAddMiniHeaderAndParagraph(header, param string) {
	/*
	   <div class="flex flex-col gap-y-4">
	       <h1
	           class="text-sm md:text-md font-semibold tracking-wide"
	       >
	           method name
	       </h1>
	       <p class="text-sm tracking-wide">oneForAll</p>
	   </div>
	*/
	mainBody := HTMLElement{
		nodeName: "div",
		classes:  "flex flex-col gap-y-4",
	}
	mainBody.AddChild(pc.getHeaderElement(header))
	if len(param) == 0 {
		param = "--"
	}
	mainBody.AddChild(pc.getParagraphElement(param))
	pc.htmlBody.AddChild(mainBody)
}

func (pc *SeePluginStandardLibrary) getHeaderElement(header string) HTMLElement {
	return HTMLElement{
		nodeName: "h1",
		classes:  "text-sm md:text-md font-semibold tracking-wide",
		value:    header,
	}
}

func (pc *SeePluginStandardLibrary) getParagraphElement(p string) HTMLElement {
	return HTMLElement{
		nodeName: "p",
		classes:  "text-sm tracking-wide",
		value:    p,
	}
}

func (pc *SeePluginStandardLibrary) CloseLibrary() {
	reportPath := util.GetAbsoluteReportPath("log.html", pc.username, pc.dutyId)
	pc.WriteLinesToFile([]string{pc.htmlBody.GetHTML()}, reportPath)
}
