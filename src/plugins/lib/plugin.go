package lib

type Plugin interface {
	Initialize(string) error
	Execute(string) error
	Close()
	Name() string
}
