package engine

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"runtime"
	"sync"

	"github.com/zayarhtet/seap-api/src/plugins/lib"
	"github.com/zayarhtet/seap-api/src/util"
)

var (
	pluginsMu sync.RWMutex
	plugins   = make(map[string]func() lib.Plugin)
)

// DiscoverAndRegisterPlugins scans the specified directory for plugin implementations and registers them with the framework.
func DiscoverAndRegisterPlugins(directory string) error {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.Name() == "lib" {
			continue
		}

		pluginGo := filepath.Join(directory, entry.Name(), entry.Name()+".go")
		if !util.FileExists(pluginGo) {
			continue
		}

		os := runtime.GOOS
		var ext string
		if os == "windows" {
			ext = ".dll"
		} else {
			ext = ".so"
		}

		pluginBinary := filepath.Join(directory, entry.Name(), entry.Name()+ext)

		cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", pluginBinary, pluginGo)

		err := cmd.Run()
		if err != nil {
			fmt.Println("Error compiling .so file:", err)
			return err
		}

		p, err := plugin.Open(pluginBinary)
		if err != nil {
			return err
		}

		symbol, err := p.Lookup("NewPlugin")
		if err != nil {
			return err
		}

		newPluginFunc, ok := symbol.(func() lib.Plugin)
		if !ok {
			return errors.New("unexpected type from plugin symbol")
		}

		RegisterPlugin(entry.Name(), newPluginFunc)
	}

	return nil
}

// RegisterPlugin registers a plugin with the framework.
func RegisterPlugin(name string, pluginInstance func() lib.Plugin) {
	pluginsMu.Lock()
	defer pluginsMu.Unlock()
	if pluginInstance == nil {
		panic("seengine: RegisterPlugin plugin is nil")
	}
	if _, dup := plugins[name]; dup {
		panic("seengine: RegisterPlugin called twice for plugin " + name)
	}
	plugins[name] = pluginInstance
}

// GetNewPlugin returns a registered plugin by name.
func GetNewPlugin(name string) (lib.Plugin, error) {
	pluginsMu.RLock()
	defer pluginsMu.RUnlock()
	plugin, ok := plugins[name]
	if !ok {
		return nil, errors.New("seengine: plugin not found")
	}
	return plugin(), nil
}

func GetPluginList() *[]string {
	var keys []string

	for k := range plugins {
		keys = append(keys, k)
	}
	return &keys
}
