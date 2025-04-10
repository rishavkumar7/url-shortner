package main

import (
	"github.com/joho/godotenv"
	"github.com/rishavkumar7/url-shortner/database"
	"github.com/rishavkumar7/url-shortner/routes"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	database.ConnectDb()
}

func main() {
	routes.ClientRoutes()
}
