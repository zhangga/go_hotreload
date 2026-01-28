//go:build windows

package internal

var _ Loader = (*loaderDll)(nil)

type loaderDll struct {
	path string
}

func (l *loaderDll) Lookup(symName string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func NewLoader(path string) (Loader, error) {
	loader := &loaderDll{
		path: path,
	}
	return loader, nil
}
