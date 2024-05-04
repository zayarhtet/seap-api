//go:build plugin

package main

import (
	"fmt"

	"github.com/zayarhtet/seap-api/src/plugins/lib"
)

type FpClean struct {
	dir string
}

func NewPlugin() lib.Plugin {
	return &FpClean{}
}

func (p *FpClean) Initialize(userDir string) (string, error) {
	// Plugin initialization
	p.dir = userDir
	return "", nil
}

func (p *FpClean) Execute() error {
	// Plugin execution
	fmt.Println("Executing plugin..." + p.dir)
	return nil
}

func (p *FpClean) Name() string {
	// Plugin initialization
	return "FPClean"
}
