package patch_v1

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/zhangga/go_hotreload/patch"
)

var GlobalPatchEntry patch.PatchEntry = &patchEntryImpl{
	Patches: gomonkey.NewPatches(),
}

type patchEntryImpl struct {
	*gomonkey.Patches
}

func (p *patchEntryImpl) Tag() string {
	return "v1"
}

func (p *patchEntryImpl) Patch() error {
	//TODO implement me
	panic("implement me")
}

func (p *patchEntryImpl) Unpatch() error {
	p.Reset()
	return nil
}

func (p *patchEntryImpl) GomonkeyPatches() *gomonkey.Patches {
	return p.Patches
}
