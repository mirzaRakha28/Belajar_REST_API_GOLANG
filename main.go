package main

import (
	"ECHO-REST/db"
	"ECHO-REST/routes"
)

func main() {
	db.Init()

	e := routes.Init()
	e.Logger.Fatal(e.Start(":1234"))

}
