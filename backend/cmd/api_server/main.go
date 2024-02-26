package main

import (
	"fmt"
	"git-quest-be/internal/api"
)

func main() {
	// Start a new http server with Gin and listen on port 8080
	router := api.NewRouter()

	fmt.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
