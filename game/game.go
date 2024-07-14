package game

import (
	"fmt"
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/pooya-hajjar/terminal-car-driving/config"
	"github.com/pooya-hajjar/terminal-car-driving/keyboard"
)

var fontFace = text.NewGoXFace(bitmapfont.Face)

type Game struct {
	keyRegistry keyboard.KeyRegistry
	heldKeys    []ebiten.Key
}

// NewGame sets the window size and title then returns a pointer to Game struct
func NewGame(keyRegistry keyboard.KeyRegistry, config config.Config) *Game {
	title := fmt.Sprintf("%s %s", config.CarBrand, config.CarModel)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle(title)

	return &Game{
		keyRegistry: keyRegistry,
		heldKeys:    make([]ebiten.Key, 0, keyboard.TotalGameKeys),
	}
}

// Run runs the game and returns a nillable error
func (g *Game) Run() error {
	return ebiten.RunGame(g)
}
