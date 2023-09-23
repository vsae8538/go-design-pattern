package memento

import "fmt"

type Game struct {
	hp, mp int
}

type gameMemento struct {
	hp, mp int
}

func (g *Game) Play(mpDelta, hpDelta int) {
	g.mp += mpDelta
	g.hp += hpDelta
}

func (g *Game) Save() *gameMemento {
	return &gameMemento{
		mp: g.mp,
		hp: g.hp,
	}
}

func (g *Game) Load(m *gameMemento) {
	g.mp = m.mp
	g.hp = m.hp
}

func (g *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", g.hp, g.mp)
}
