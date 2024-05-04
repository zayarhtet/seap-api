package lib

type SeePluginCommonLibrary struct {
	memoryFile string
	dir        string
}

func (pc *SeePluginCommonLibrary) AddMemoryFile(message string) {
	pc.memoryFile = message
}

func (pc *SeePluginCommonLibrary) SetDir(dir string) {
	pc.dir = dir
}

func (pc *SeePluginCommonLibrary) GetDir() string {
	return pc.dir
}
