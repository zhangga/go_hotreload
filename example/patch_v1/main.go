package main

import (
	"fmt"
	"reflect"

	internal_player "github.com/zhangga/go_hotreload/example/internal/player"
	patch_v1_player "github.com/zhangga/go_hotreload/example/patch_v1/player"
	"github.com/zhangga/go_hotreload/patch"
)

// GlobalPatchEntry 热更新入口，名称固定为 GlobalPatchEntry
var GlobalPatchEntry patch.PatchEntry = &patchEntryImpl{
	BasePatchEntry: patch.NewBasePatchEntry(),
}

type patchEntryImpl struct {
	*patch.BasePatchEntry
}

func (p *patchEntryImpl) Tag() string {
	return "v1"
}

func (p *patchEntryImpl) Patch() error {
	// 检查结构体字段偏移量是否一致
	if err := p.CheckStructFieldOffset(internal_player.Player{}, patch_v1_player.Player{},
		"name", "name"); err != nil {
		return fmt.Errorf("check struct field offset failed: %w", err)
	}

	nameFunc, err := p.MakeValueByFunctionName((*internal_player.Player).Name,
		"github.com/zhangga/go_hotreload/example/internal/player.(*Player).Name")
	if err != nil {
		return fmt.Errorf("make value by function name failed: %w", err)
	}

	p.ApplyCore(nameFunc, reflect.ValueOf(patch_v1_player.Name_v1))
	return nil
}

func main() {
	fmt.Println("This is patch_v1 module")
}
