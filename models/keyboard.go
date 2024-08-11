package models

import "github.com/hajimehoshi/ebiten/v2"

type onKeyPress = func()
type onKeyRelease = func()

type Key struct {
	OnRelease onKeyRelease
	OnPress   onKeyPress

	PressDelay   int64 // ms
	ReleaseDelay int64 // ms
}

type KeyRegistry = map[ebiten.Key]Key
