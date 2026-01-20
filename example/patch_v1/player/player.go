package player

import (
	"unsafe"

	"github.com/zhangga/go_hotreload/example/internal/player"
)

type Player struct {
	name  string
	level int32
}

func Name_v1(p *player.Player) string {
	player := (*Player)(unsafe.Pointer(p))
	return "new: " + player.name
}

//go:noinline
func (p *Player) Level_v1() int32 {
	return -p.level
}
