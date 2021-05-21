package main

import (
	"url-shortener/app"
)

func main() {
	app := app.NewServer("localhost:6796", "", 0)
	app.Run(":8080")
}
