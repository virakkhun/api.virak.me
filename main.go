package main

import (
	"api.virak.me/src/app"
	"api.virak.me/src/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	app.Bootstrap()
}
