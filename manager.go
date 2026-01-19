package go_hotreload

import (
	"context"
	"errors"
	"os"
	"sync"

	"github.com/zhangga/go_hotreload/internal"
)

var manager = Manager{
	patcherLoaded: map[string]*internal.Patcher{},
}

// Manager 热更新管理器
type Manager struct {
	loadLock      sync.Mutex
	patcherLoaded map[string]*internal.Patcher
}

func LoadPatch(ctx context.Context, patchPath string) error {
	// 1.判断patch是否存在
	_, err := os.Stat(patchPath)
	if err != nil {
		return errors.New("patch_v1 file not found: " + patchPath)
	}

	// 2.判断是否已经有补丁加载
	p, err := addLoadedPatcher(patchPath)
	if err != nil {
		return err
	}

	// 3.加载补丁
	return p.Patch(ctx)
}

func RevertPatch(ctx context.Context, patchPath string) error {
	manager.loadLock.Lock()
	defer manager.loadLock.Unlock()
	return nil
}

func addLoadedPatcher(patchPath string) (*internal.Patcher, error) {
	manager.loadLock.Lock()
	defer manager.loadLock.Unlock()
	if _, ok := manager.patcherLoaded[patchPath]; ok {
		return nil, errors.New("patch_v1 file already loaded: " + patchPath)
	}
	p := internal.NewPatcher(patchPath)
	manager.patcherLoaded[patchPath] = p
	return p, nil
}
