//go:build linux || darwin || freebsd

package internal

import "plugin"

var _ Loader = (*loaderPlugin)(nil)

type loaderPlugin struct {
	path string
	lib  *plugin.Plugin
}

func (l *loaderPlugin) Lookup(symName string) (any, error) {
	return l.lib.Lookup(symName)
}

func NewLoader(path string) (Loader, error) {
	lib, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	loader := &loaderPlugin{
		path: path,
		lib:  lib,
	}
	return loader, nil
}
