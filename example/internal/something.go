package internal

import (
	"log"

	"github.com/zhangga/go_hotreload/example/internal/player"
)

func DoSomething() error {
	log.Printf("player.Name() = %s\n", player.APlayer.Name())
	log.Printf("player.Level() = %d\n", player.APlayer.Level())
	return nil
}
