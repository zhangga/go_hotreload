//go:build linux
// +build linux

package patch

func (b *BasePatchEntry) MakeValueByFunctionName(target interface{}, name string) (reflect.Value, error) {
	src := reflect.ValueOf(target)
	if src.Kind() != reflect.Func {
		return src, fmt.Errorf("%s is not function", src.String())
	}
	ptr, err := FindFuncWithName(name)
	if err != nil {
		return src, err
	}
	val := (*[2]uintptr)(unsafe.Pointer(&src))
	(*val)[1] = uintptr(makePtr(ptr))
	return src, nil
}
