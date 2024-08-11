package keyboard

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pooya-hajjar/terminal-car-driving/exception"
	"github.com/pooya-hajjar/terminal-car-driving/models"
)

const (
	TotalGameKeys = 14
)

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
func GetKeyRegistry(config models.Config) models.KeyRegistry {
	keyRegistry := make(map[ebiten.Key]models.Key, TotalGameKeys)

	keyRegistry[getKeyById(config.GasKey)] = models.Key{OnPress: pressGas, OnRelease: releaseGas}
	keyRegistry[getKeyById(config.BrakeKey)] = models.Key{OnPress: pressBrake, OnRelease: releaseBrake}
	keyRegistry[getKeyById(config.ClutchKey)] = models.Key{OnPress: pressClutch, OnRelease: releaseClutch}
	keyRegistry[getKeyById(config.HandBrakeKey)] = models.Key{OnPress: pressHandBrake, OnRelease: releaseHandBrake}
	keyRegistry[getKeyById(config.PowerKey)] = models.Key{OnPress: pressPower, OnRelease: releasePower}
	keyRegistry[getKeyById(config.IGearKey)] = models.Key{OnPress: pressIGear, OnRelease: releaseIGear}
	keyRegistry[getKeyById(config.DGearKey)] = models.Key{OnPress: pressDGear, OnRelease: releaseDGear}

	keyRegistry[getKeyById(config.Gear1Key)] = models.Key{OnPress: pressGear1, OnRelease: releaseGear1}
	keyRegistry[getKeyById(config.Gear2Key)] = models.Key{OnPress: pressGear2, OnRelease: releaseGear2}
	keyRegistry[getKeyById(config.Gear3Key)] = models.Key{OnPress: pressGear3, OnRelease: releaseGear3}
	keyRegistry[getKeyById(config.Gear4Key)] = models.Key{OnPress: pressGear4, OnRelease: releaseGear4}
	keyRegistry[getKeyById(config.Gear5Key)] = models.Key{OnPress: pressGear5, OnRelease: releaseGear5}
	keyRegistry[getKeyById(config.GearNKey)] = models.Key{OnPress: pressGearN, OnRelease: releaseGearN}
	keyRegistry[getKeyById(config.GearRKey)] = models.Key{OnPress: pressGearR, OnRelease: releaseGearR}

	return keyRegistry
}
