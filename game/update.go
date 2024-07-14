package game

import (
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	return nil
}
