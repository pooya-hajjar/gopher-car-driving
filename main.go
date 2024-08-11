package main

import (
	"github.com/pooya-hajjar/terminal-car-driving/config"
	"github.com/pooya-hajjar/terminal-car-driving/exception"
	"github.com/pooya-hajjar/terminal-car-driving/game"
	"github.com/pooya-hajjar/terminal-car-driving/keyboard"
	"github.com/pooya-hajjar/terminal-car-driving/models"
)

var (
	Game        *game.Game
	Config      models.Config
	KeyRegistry models.KeyRegistry
)

func init() {
	conf, err := config.LoadConfig(".")
	exception.FatalIfError(err)
	Config = conf

	KeyRegistry = keyboard.GetKeyRegistry(Config)

	Game = game.NewGame(KeyRegistry, Config)
}

func main() {
	runGameErr := Game.Run()
	exception.FatalIfError(runGameErr)
}
