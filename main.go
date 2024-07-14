package main

import (
	"github.com/pooya-hajjar/terminal-car-driving/config"
	"github.com/pooya-hajjar/terminal-car-driving/exception"
	"github.com/pooya-hajjar/terminal-car-driving/game"
)

var (
	Config config.Config
	Game   *game.Game
)

func init() {
	conf, err := config.LoadConfig(".")
	exception.FatalIfError(err)
	Config = conf

	Game = game.NewGame()
}

func main() {
	runGameErr := Game.Run()
	exception.FatalIfError(runGameErr)

}
