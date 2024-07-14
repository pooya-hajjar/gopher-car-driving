package game

import (
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var fontFace = text.NewGoXFace(bitmapfont.Face)

type Game struct {
	keys []ebiten.Key
}

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
