package patch

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/agiledragon/gomonkey/v2"
)

const GlobalPatchEntryVarName = "GlobalPatchEntry"

type PatchEntry interface {
	Tag() string
	Patch() error
	Unpatch()
}

type BasePatchEntry struct {
	*gomonkey.Patches
}

func NewBasePatchEntry() *BasePatchEntry {
	return &BasePatchEntry{
		Patches: gomonkey.NewPatches(),
	}
}

// CheckStructFieldOffset 检查两个结构体字段偏移是否一致
func (b *BasePatchEntry) CheckStructFieldOffset(target, double any, tarFieldName, doubleFieldName string) error {
	tarVal := reflect.ValueOf(target)
	doubleVal := reflect.ValueOf(double)
	if tarVal.Kind() != reflect.Struct || doubleVal.Kind() != reflect.Struct {
		return errors.New("both target and double must be struct")
	}

	if len(tarFieldName) == 0 || len(doubleFieldName) == 0 {
		return errors.New("field names cannot be empty")
	}

	tarType := tarVal.Type()
	doubleType := doubleVal.Type()

	var tarOff, doubleOff uintptr
	found := false

	for i := 0; i < tarType.NumField(); i++ {
		if tarType.Field(i).Name == tarFieldName {
			tarOff = tarType.Field(i).Offset
			found = true
			break
		}
	}
	if !found {
		return errors.New("target field not found: " + tarFieldName)
	}

	found = false
	for i := 0; i < doubleType.NumField(); i++ {
		if doubleType.Field(i).Name == doubleFieldName {
			doubleOff = doubleType.Field(i).Offset
			found = true
			break
		}
	}
	if !found {
		return errors.New("double field not found: " + doubleFieldName)
	}

	if tarOff != doubleOff {
		return fmt.Errorf("field offset mismatch: target.%s offset %d != double.%s offset %d",
			tarFieldName, tarOff, doubleFieldName, doubleOff)
	}
	return nil
}

func (b *BasePatchEntry) Unpatch() {
	b.Patches.Reset()
}
