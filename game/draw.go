package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"strings"
)

func (g *Game) Draw(screen *ebiten.Image) {
	var keyStrs []string
	var keyNames []string
	for _, k := range g.keys {
		keyStrs = append(keyStrs, k.String())
		if name := ebiten.KeyName(k); name != "" {
			keyNames = append(keyNames, name)
		}
	}

	// Use bitmapfont.Face instead of ebitenutil.DebugPrint, since some key names might not be printed with DebugPrint.
	textOp := &text.DrawOptions{}
	textOp.LineSpacing = fontFace.Metrics().HLineGap + fontFace.Metrics().HAscent + fontFace.Metrics().HDescent
	text.Draw(screen, strings.Join(keyStrs, ", ")+"\n"+strings.Join(keyNames, ", "), fontFace, textOp)
}
