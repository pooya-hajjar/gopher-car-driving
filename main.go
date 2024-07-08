package main

import (
	"github.com/pooya-hajjar/terminal-car-driving/config"
	"github.com/pooya-hajjar/terminal-car-driving/exception"
)

var (
	appConfig config.Config
)

func init() {
	conf, err := config.LoadConfig(".")
	exception.FatalIfError(err)
	appConfig = conf
}

func main() {

}
