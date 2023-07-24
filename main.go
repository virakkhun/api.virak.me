package main

import (
	"api.virak.me/app"
	"api.virak.me/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	app.Bootstrap()
}
