package main

import (
	"Hacktiv10JWT/database"
	"Hacktiv10JWT/router"
)

func main() {
	database.StartDB()
	defer database.CloseDB()
	r := router.StartApp()
	r.Run(":8080")
}
