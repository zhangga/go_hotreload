package patch_v1

import (
	"github.com/zhangga/go_hotreload/patch"
)

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
	//TODO implement me
	panic("implement me")
}

func (p *patchEntryImpl) Unpatch() error {
	p.Reset()
	return nil
}
