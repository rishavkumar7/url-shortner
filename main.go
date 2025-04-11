package main

import (
	"github.com/rishavkumar7/url-shortner/database"
	"github.com/rishavkumar7/url-shortner/routes"
)

func init() {
	// err := godotenv.Load(".env")  // only used for local development
	// if err != nil {
	// 	panic(err)
	// }

	database.ConnectDb()
}

func main() {
	routes.ClientRoutes()
}
