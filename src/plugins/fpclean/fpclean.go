//go:build plugin

package main

import (
	"fmt"

	"github.com/zayarhtet/seap-api/src/plugins/lib"
)

type FpClean struct {
	lib.SeePluginStandardLibrary
}

func NewPlugin() lib.Plugin {
	return &FpClean{}
}

func (p *FpClean) Initialize(inputDir string) error {
	// Plugin initialization
	p.SetDir(inputDir)
	p.AddMemoryFile("HELLO")
	fmt.Println(inputDir)

	return nil
}

func (p *FpClean) Execute(targetDir string) error {
	// Plugin execution
	fmt.Println("Executing plugin..." + p.GetDir())
	return nil
}

func (p *FpClean) Name() string {
	// Plugin initialization
	return "FPClean"
}
