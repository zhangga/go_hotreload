//go:build (!linux && !freebsd && !darwin) || !cgo

package patch

import "reflect"

func (b *BasePatchEntry) MakeValueByFunctionName(target any, name string) (reflect.Value, error) {
	panic("unsupported platform")
}

func (b *BasePatchEntry) MakeValueByFunctionName4So(target any, soPath, name string) (reflect.Value, error) {
	panic("unsupported platform")
}
