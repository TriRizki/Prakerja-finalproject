package main

import (
	"Prakerja-finalproject/db"
	"Prakerja-finalproject/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}