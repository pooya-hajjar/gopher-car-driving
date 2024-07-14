package keyboard

import "github.com/hajimehoshi/ebiten/v2"

var keyMap = map[string]ebiten.Key{
	// Letters
	"a": ebiten.KeyA,
	"b": ebiten.KeyB,
	"c": ebiten.KeyC,
	"d": ebiten.KeyD,
	"e": ebiten.KeyE,
	"f": ebiten.KeyF,
	"g": ebiten.KeyG,
	"h": ebiten.KeyH,
	"i": ebiten.KeyI,
	"j": ebiten.KeyJ,
	"k": ebiten.KeyK,
	"l": ebiten.KeyL,
	"m": ebiten.KeyM,
	"n": ebiten.KeyN,
	"o": ebiten.KeyO,
	"p": ebiten.KeyP,
	"q": ebiten.KeyQ,
	"r": ebiten.KeyR,
	"s": ebiten.KeyS,
	"t": ebiten.KeyT,
	"u": ebiten.KeyU,
	"v": ebiten.KeyV,
	"w": ebiten.KeyW,
	"x": ebiten.KeyX,
	"y": ebiten.KeyY,
	"z": ebiten.KeyZ,

	// Numbers
	"0": ebiten.Key0,
	"1": ebiten.Key1,
	"2": ebiten.Key2,
	"3": ebiten.Key3,
	"4": ebiten.Key4,
	"5": ebiten.Key5,
	"6": ebiten.Key6,
	"7": ebiten.Key7,
	"8": ebiten.Key8,
	"9": ebiten.Key9,

	// Arrow keys
	"up":    ebiten.KeyUp,
	"down":  ebiten.KeyDown,
	"left":  ebiten.KeyLeft,
	"right": ebiten.KeyRight,

	// Special keys
	"enter":     ebiten.KeyEnter,
	"escape":    ebiten.KeyEscape,
	"space":     ebiten.KeySpace,
	"tab":       ebiten.KeyTab,
	"backspace": ebiten.KeyBackspace,
	"shift":     ebiten.KeyShift,
	"ctrl":      ebiten.KeyControl,
	"alt":       ebiten.KeyAlt,
}
