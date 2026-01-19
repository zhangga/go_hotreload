package internal

import (
	"context"
	"errors"
	"fmt"
	"log"
	"plugin"
	"reflect"
	"sync"

	"github.com/zhangga/go_hotreload/patch"
)

type Patcher struct {
	path       string
	lock       sync.Mutex
	patchEntry patch.PatchEntry
	isPatched  bool
}

func NewPatcher(path string) *Patcher {
	return &Patcher{
		path:      path,
		isPatched: false,
	}
}

func (p *Patcher) Patch(ctx context.Context) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	lib, err := plugin.Open(p.path)
	if err != nil {
		return err
	}

	if sym, err := lib.Lookup(patch.GlobalPatchEntryVarName); err != nil {
		return err
	} else if entryPtr, ok := sym.(*patch.PatchEntry); !ok {
		return fmt.Errorf("type error:%+v", reflect.TypeOf(sym).String())
	} else {
		entry := *entryPtr
		p.patchEntry = entry
		if err = p.patchEntry.Patch(); err != nil {
			log.Printf("patch error, path=%s, unpatching: %v", p.path, err)
			if uperr := p.patchEntry.Unpatch(); uperr != nil {
				log.Printf("unpatch path=%s, error: %v", p.path, uperr)
			}
			return err
		}
		p.isPatched = true
		log.Printf("patch success, path=%s", p.path)
		return nil
	}
}

func (p *Patcher) Unpatch(ctx context.Context) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if !p.isPatched {
		return errors.New("not patched: " + p.path)
	}
	if p.patchEntry == nil {
		return errors.New("patch entry is nil: " + p.path)
	}
	if err := p.patchEntry.Unpatch(); err != nil {
		return err
	}
	p.isPatched = false
	p.patchEntry = nil
	return nil
}
