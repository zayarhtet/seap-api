//go:build plugin

package main

import (
	"github.com/zayarhtet/seap-api/src/plugins/lib"
)

type PluginTemplate struct {
	lib.SeePluginStandardLibrary
	// member variables goes here.
}

func NewPlugin() lib.Plugin {
	return &PluginTemplate{}
}

func (p *PluginTemplate) Initialize(inputDir string) error {
	p.InitializeLibrary(inputDir)

	/*
		Member variable initialization goes here.
	*/

	return nil
}

func (p *PluginTemplate) Execute(targetDir string) error {
	p.SetUsername(targetDir)

	/*
		Business logic goes here.
	*/

	return nil
}

func (p *PluginTemplate) Close() {
	/*
		Close everything here
	*/
	p.CloseLibrary()
}

func (p *PluginTemplate) Name() string {
	return "PluginTemplate"
}
