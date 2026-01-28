//go:build linux || darwin || freebsd

package internal

import "plugin"

var _ Loader = (*loaderPlugin)(nil)

type loaderPlugin struct {
	path   string
	plugin *Plugin
}

func (l *loaderPlugin) Lookup(symName string) (any, error) {
	return l.plugin.Lookup(symName)
}

func NewLoader(path string) (Loader, error) {
	lib, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	loader := &loaderPlugin{
		path:   path,
		plugin: lib,
	}
	return loader, nil
}
