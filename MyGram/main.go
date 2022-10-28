package main

import (
	"mygram/database"
	"mygram/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")

}
