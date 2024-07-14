package keyboard

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pooya-hajjar/terminal-car-driving/config"
	"github.com/pooya-hajjar/terminal-car-driving/exception"
)

type onKeyPress = func()
type onKeyRelease = func()

const (
	TotalGameKeys = 14
)

type Key struct {
	OnRelease onKeyRelease
	OnPress   onKeyPress

	PressDelay   int64 // ms
	ReleaseDelay int64 // ms
}

type KeyRegistry = map[ebiten.Key]Key

func getKeyById(key string) ebiten.Key {
	k, ok := keyMap[key]
	if !ok {
		exception.FatalIfError(errors.New(fmt.Sprintf("the %s key is not supported", key)))
	}
	return k
}

// GetKeyRegistry initializes and returns a KeyRegistry based on the provided configuration.
// It maps configured key identifiers from the given config.Config to their corresponding
// ebiten.Key values, associating each key with its respective Key structure that defines
// OnPress and OnRelease handlers. The returned KeyRegistry allows easy access and management
// of key bindings for various game actions such as gas, brake, clutch, gear shifting, and more.
func GetKeyRegistry(config config.Config) KeyRegistry {
	keyRegistry := make(map[ebiten.Key]Key, TotalGameKeys)

	keyRegistry[getKeyById(config.GasKey)] = Key{OnPress: pressGas, OnRelease: releaseGas}
	keyRegistry[getKeyById(config.BrakeKey)] = Key{OnPress: pressBrake, OnRelease: releaseBrake}
	keyRegistry[getKeyById(config.ClutchKey)] = Key{OnPress: pressClutch, OnRelease: releaseClutch}
	keyRegistry[getKeyById(config.HandBrakeKey)] = Key{OnPress: pressHandBrake, OnRelease: releaseHandBrake}
	keyRegistry[getKeyById(config.PowerKey)] = Key{OnPress: pressPower, OnRelease: releasePower}
	keyRegistry[getKeyById(config.IGearKey)] = Key{OnPress: pressIGear, OnRelease: releaseIGear}
	keyRegistry[getKeyById(config.DGearKey)] = Key{OnPress: pressDGear, OnRelease: releaseDGear}

	keyRegistry[getKeyById(config.Gear1Key)] = Key{OnPress: pressGear1, OnRelease: releaseGear1}
	keyRegistry[getKeyById(config.Gear2Key)] = Key{OnPress: pressGear2, OnRelease: releaseGear2}
	keyRegistry[getKeyById(config.Gear3Key)] = Key{OnPress: pressGear3, OnRelease: releaseGear3}
	keyRegistry[getKeyById(config.Gear4Key)] = Key{OnPress: pressGear4, OnRelease: releaseGear4}
	keyRegistry[getKeyById(config.Gear5Key)] = Key{OnPress: pressGear5, OnRelease: releaseGear5}
	keyRegistry[getKeyById(config.GearNKey)] = Key{OnPress: pressGearN, OnRelease: releaseGearN}
	keyRegistry[getKeyById(config.GearRKey)] = Key{OnPress: pressGearR, OnRelease: releaseGearR}

	return keyRegistry
}
