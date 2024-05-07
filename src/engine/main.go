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

func ABSOLUTE_INPUT_FILE_PATH() string { return util.ABSOLUTE_STORAGE_PATH() + "input-files/" }

func Init() {
	if err := DiscoverAndRegisterPlugins(RELATIVE_PLUGINS_PATH); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func ExecuteDuty(pluginName, dutyDir string) error {
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
		go Worker(inputFileCh, pluginName, &wg, &mu)
	}

	wg.Wait()

	fmt.Println("All executions completed")
	return nil
}

func Worker(inputFileCh <-chan string, pluginName string, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for inputFile := range inputFileCh {
		newPlugin, _ := GetNewPlugin(pluginName)
		newPlugin.Initialize(inputFile)
		newPlugin.Execute()
		mu.Lock()
		fmt.Printf("Finished executing %s with input file: %s\n", pluginName, inputFile)
		mu.Unlock()
	}
}
