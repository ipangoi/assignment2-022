package main

import (
	"assignment-2/database"
	routers "assignment-2/router"
)

//"fmt"

func main() {

	database.StartDB()

	ginSection()
}

func ginSection() {
	routers.SetOrderRouter().Run(":8080")

	return
}
