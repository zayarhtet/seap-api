//go:build plugin

package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/zayarhtet/seap-api/src/plugins/lib"
	"github.com/zayarhtet/seap-api/src/util"
)

type FpClean struct {
	lib.SeePluginStandardLibrary

	testCaseList []TestCase
}

type TestCase struct {
	MethodName string              `json:"method_name"`
	Arguments  []map[string]string `json:"test_cases"`
}

func NewPlugin() lib.Plugin {
	return &FpClean{}
}

func (p *FpClean) Initialize(inputDir string) error {
	p.InitializeLibrary(inputDir)
	err := p.ReadJSONFileAsStruct(filepath.Join(inputDir, "testcase.json"), &p.testCaseList)

	if err != nil {
		return err
	}

	return nil
}

func (p *FpClean) Execute(targetDir string) error {
	p.SetUsername(targetDir)

	entries := p.ReadDirectory(targetDir)
	if len(entries) == 0 {
		return nil
	}

	targetFilePath := filepath.Join(targetDir, entries[0].Name())
	tempFileName := "temp"

	linesWithoutComment, err := p.LoadFileContent(targetFilePath, tempFileName)
	if err != nil {
		return err
	}

	for _, tc := range p.testCaseList {
		for _, args := range tc.Arguments {

			mainFunctionLine := fmt.Sprintf("Start = %s %s", tc.MethodName, args["args"])
			p.ReportAddMiniHeaderAndParagraph("method name", tc.MethodName)
			p.ReportAddMiniHeaderAndParagraph("arguments", args["args"])
			p.ReportAddMiniHeaderAndParagraph("expected", args["expected"])

			tempFileContent := append(linesWithoutComment, mainFunctionLine)

			tempDirPath := p.CreateAndWriteFileInTemp(tempFileContent, tempFileName+".icl")

			tempFilePath := filepath.Join(tempDirPath, tempFileName)

			output, errOutput, err := p.ExecuteCommandWithTimeout("clm", "-I", tempDirPath, tempFileName, "-o", tempFilePath)

			if err != nil {
				p.ReportAddMiniHeaderAndParagraph("actual", output)
				p.ReportAddMiniHeaderAndParagraph("error", err.Error()+"\n"+errOutput)
				p.ReportAddHorizontalBar()
				util.DeleteDirectory(tempDirPath)
				continue
			}

			output, errOutput, err = p.ExecuteCommandWithTimeout(tempFilePath)
			p.ReportAddMiniHeaderAndParagraph("actual", output)
			if err != nil {
				p.ReportAddMiniHeaderAndParagraph("error", err.Error()+"\n"+errOutput)
			} else {
				p.ReportAddMiniHeaderAndParagraph("error", errOutput)
			}
			p.ReportAddHorizontalBar()
			util.DeleteDirectory(tempDirPath)
		}
	}
	return nil
}

func (p *FpClean) LoadFileContent(targetFilePath string, tempFileName string) ([]string, error) {
	linesWithoutComment, err := p.ReadProgrammingFileWithoutComment(targetFilePath, "//", "/*", "*/")
	if err != nil {
		return nil, err
	}

	util.RemoveElementsInPlace[string](&linesWithoutComment, func(s string) bool {
		if strings.HasPrefix(strings.TrimSpace(s), "module ") {
			return false
		}
		if strings.TrimSpace(s) == "import StdEnv" {
			return false
		}
		return true
	})

	linesWithoutComment = append([]string{
		"module " + tempFileName,
		"import StdEnv",
	}, linesWithoutComment...)

	return linesWithoutComment, nil
}

func (p *FpClean) Close() {
	p.CloseLibrary()
}

func (p *FpClean) Name() string {
	return "FPClean"
}
