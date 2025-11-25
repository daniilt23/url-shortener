package main

import (
	"url-shortener/app"
)

func main() {
	app := app.NewApp()
	app.Start()
}
