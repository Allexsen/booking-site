package main

import (
	"github.com/Allexsen/booking-site/api/routes"
	database "github.com/Allexsen/booking-site/db"
)

func main() {
	db := database.Connect()
	defer db.Close()
	routes.Init()
}
