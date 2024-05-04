package lib

type Plugin interface {
	Initialize(string) (string, error)
	Execute() error
	Name() string
}
