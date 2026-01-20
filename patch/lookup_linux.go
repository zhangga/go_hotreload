//go:build linux

package patch

/*
   #cgo linux LDFLAGS: -ldl
   #include <dlfcn.h>
   #include <limits.h>
   #include <stdlib.h>
   #include <stdint.h>

   #include <stdio.h>
   #ifndef RTLD_MAIN_ONLY
   #define RTLD_MAIN_ONLY 0
   #endif

   static void* localLookup(const char* name, char** err) {
   	void* r = dlsym(RTLD_MAIN_ONLY, name);
   	if (r == NULL) {
   		*err = (char*)dlerror();
   	}
   	return r;
   }

   static void* localLookup4So(const char* soPath, const char* name, char** err) {
	dlerror();
	void* h = dlopen(soPath, 0x102); //flags: RTLD_NOW|RTLD_GLOBAL
	if (h == NULL) {
   		*err = (char*)dlerror();
		return NULL;
	}
   	void* r = dlsym(h, name);
   	if (r == NULL) {
   		*err = (char*)dlerror();
   	}
	dlclose(h);
   	return r;
   }
*/
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

func (b *BasePatchEntry) MakeValueByFunctionName(target any, name string) (reflect.Value, error) {
	src := reflect.ValueOf(target)
	if src.Kind() != reflect.Func {
		return src, fmt.Errorf("%s is not function", src.String())
	}
	ptr, err := findFuncWithName(name)
	if err != nil {
		return src, err
	}
	val := (*[2]uintptr)(unsafe.Pointer(&src))
	(*val)[1] = uintptr(makePtr(ptr))
	return src, nil
}

func (b *BasePatchEntry) MakeValueByFunctionName4So(target any, soPath, name string) (reflect.Value, error) {
	src := reflect.ValueOf(target)
	if src.Kind() != reflect.Func {
		return src, fmt.Errorf("%s is not function", src.String())
	}
	ptr, err := findFuncWithName4So(soPath, name)
	if err != nil {
		return src, err
	}
	fmt.Printf("============ MakeValueByFunctionName4So(%s, %s), ptr=0x%08x\n", soPath, name, ptr)
	val := (*[2]uintptr)(unsafe.Pointer(&src))
	(*val)[1] = uintptr(makePtr(ptr))
	return src, nil
}

func findFuncWithName(name string) (uintptr, error) {
	var cErr *C.char
	nameStr := make([]byte, len(name)+1)
	copy(nameStr, name)

	handle := C.localLookup((*C.char)(unsafe.Pointer(&nameStr[0])), &cErr)
	if handle == nil {
		return 0, errors.New(C.GoString(cErr))
	}
	return uintptr(handle), nil
}

func findFuncWithName4So(soPath, name string) (uintptr, error) {
	var cErr *C.char
	nameStr := make([]byte, len(name)+1)
	copy(nameStr, name)

	cPath := make([]byte, C.PATH_MAX+1)
	cRelName := make([]byte, len(soPath)+1)
	copy(cRelName, soPath)
	if C.realpath(
		(*C.char)(unsafe.Pointer(&cRelName[0])),
		(*C.char)(unsafe.Pointer(&cPath[0]))) == nil {
		return 0, fmt.Errorf(`findFuncWithName4So("%s"): realpath failed`, soPath)
	}

	handle := C.localLookup4So((*C.char)(unsafe.Pointer(&cPath[0])), (*C.char)(unsafe.Pointer(&nameStr[0])), &cErr)
	if handle == nil {
		return 0, errors.New(C.GoString(cErr))
	}
	return uintptr(handle), nil
}

// makePtr 防止逃逸检测失败
//
//go:noinline
func makePtr(p uintptr) unsafe.Pointer {
	return unsafe.Pointer(&p)
}
