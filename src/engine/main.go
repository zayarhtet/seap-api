package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/zayarhtet/seap-api/src/util"
)

const RELATIVE_PLUGINS_PATH = "src/plugins"

func Init() {
	if err := DiscoverAndRegisterPlugins(RELATIVE_PLUGINS_PATH); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func ExecuteDuty(pluginName, dutyId string) error {
	dutyDir := filepath.Join(util.ABSOLUTE_SUBMITTED_STORAGE_PATH(), dutyId)
	pluginInputDir := filepath.Join(util.ABSOLUTE_INPUT_FILE_PATH(), dutyId)
	entries, err := os.ReadDir(dutyDir)
	if err != nil {
		return err
	}

	inputFileCh := make(chan string, len(entries))

	for _, entry := range entries {
		eachUserDir := filepath.Join(dutyDir, entry.Name())
		inputFileCh <- eachUserDir
	}
	close(inputFileCh)

	maxThreads := runtime.NumCPU() - 2
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < maxThreads; i++ {
		wg.Add(1)
		go Worker(inputFileCh, pluginName, pluginInputDir, &wg, &mu)
	}

	wg.Wait()

	fmt.Println("All executions completed")
	return nil
}

func Worker(inputFileCh <-chan string, pluginName, pluginInputDir string, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for submittedDir := range inputFileCh {
		newPlugin, _ := GetNewPlugin(pluginName)
		err := newPlugin.Initialize(pluginInputDir)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = newPlugin.Execute(submittedDir)
		newPlugin.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
		mu.Lock()
		fmt.Printf("Finished executing %s with input file: %s\n", pluginName, submittedDir)
		mu.Unlock()
	}
}
