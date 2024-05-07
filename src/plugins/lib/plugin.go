package lib

type Plugin interface {
	Initialize(string) error
	Execute() error
	Name() string
}
