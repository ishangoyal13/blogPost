package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"

	"github.com/ishangoyal13/blogPost/models"
	"github.com/ishangoyal13/blogPost/routers"
)

func main() {
	router := routers.RegisterRoutes()
	models.ConnectDatabase()

	router.Use(static.Serve("/", static.LocalFile("./views/frontend/public", true)))
	fmt.Printf("\nSuccessfully connected to database :)\n\n")

	router.Run("localhost:8080")
}
