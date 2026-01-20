package player

var APlayer = &Player{
	name:  "Hero",
	level: 18,
}

type Player struct {
	name  string
	level int32
}

//go:noinline
func (p *Player) Name() string {
	return p.name
}

//go:noinline
func (p *Player) Level() int32 {
	return p.level
}
