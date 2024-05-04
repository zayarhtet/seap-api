package lib

type PluginCommonLibrary struct {
	memoryFile string
	dir        string
}

func (pc *PluginCommonLibrary) AddMemoryFile(message string) {
	pc.memoryFile = message
}

func (pc *PluginCommonLibrary) SetDir(dir string) {
	pc.dir = dir
}

func (pc *PluginCommonLibrary) GetDir() string {
	return pc.dir
}
