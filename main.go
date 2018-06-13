package main

import "fmt"

const decksize = 40

type Player struct {
	Name           string
	PicksRemaining int
}

func (p *Player) takepick() {
	p.PicksRemaining--
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:           name,
		PicksRemaining: decksize,
	}
}

func main() {
	gid := NewPlayer("gideon")
	fmt.Println(gid.PicksRemaining)
	gid.takepick()
	fmt.Println(gid.PicksRemaining)

}
