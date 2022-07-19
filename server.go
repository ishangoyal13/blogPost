package main

import (
	"fmt"

	"github.com/ishangoyal/blogPost/routers"
)

func main() {
	router := routers.RegisterRoutes()

	fmt.Printf("\nSuccessfully connected to database :)\n\n")

	router.Run("localhost:8080")
}
