package main

import (
	"golang-final-project/Configs"
	"golang-final-project/Routes"
)

func main() {
	Configs.Connection()
	e := Routes.RouteVersion1()
	e.Start(":8080")
}
