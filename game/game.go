package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

// NewGame sets the window size and title then returns a pointer to Game struct
func NewGame() *Game {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("terminal car simulator")

	return &Game{}
}

// Run runs the game and returns a nillable error
func (g *Game) Run() error {
	return ebiten.RunGame(g)
}
