package main

import (
	"golang-final-project/Configs"
	"golang-final-project/Routes"
)

func main() {
	Configs.Connection()
	e := Routes.NewRoute()
	e.Start(":8000")
}
