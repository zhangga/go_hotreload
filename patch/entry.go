package patch

import "github.com/agiledragon/gomonkey/v2"

const GlobalPatchEntryVarName = "GlobalPatchEntry"

type PatchEntry interface {
	Tag() string
	Patch() error
	Unpatch() error
	GomonkeyPatches() *gomonkey.Patches
}
