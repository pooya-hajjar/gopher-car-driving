package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"time"
)

func (g *Game) Update() error {
	g.heldKeys = inpututil.AppendPressedKeys(make([]ebiten.Key, 0))

outer:
	for key, val := range g.keyRegistry {
		for _, k := range g.heldKeys {
			if key == k {
				if val.PressDelay > 0 {
					time.Sleep(time.Millisecond * time.Duration(val.PressDelay))
				}
				val.OnPress()
				continue outer
			}
		}
		if val.ReleaseDelay > 0 {
			time.Sleep(time.Millisecond * time.Duration(val.ReleaseDelay))
		}
		val.OnRelease()
	}

	return nil
}
