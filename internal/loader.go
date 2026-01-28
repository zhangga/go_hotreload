package internal

type Loader interface {
	Lookup(symName string) (any, error)
}
