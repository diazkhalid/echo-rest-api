package main

import (
	"rest-api-echo/db"
	"rest-api-echo/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))
}
