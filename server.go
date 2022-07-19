package main

import (
	"fmt"

	"github.com/ishangoyal13/blogPost/models"
	"github.com/ishangoyal13/blogPost/routers"
)

func main() {
	router := routers.RegisterRoutes()
	models.ConnectDatabase()

	fmt.Printf("\nSuccessfully connected to database :)\n\n")

	router.Run("localhost:8080")
}
