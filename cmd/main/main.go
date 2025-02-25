package main

import (
	"fmt"

	"github.com/Robin-Rathore/go-bookstore/pkg/config"
)

func main() {
	fmt.Println("Starting the application...")

	// Initialize the database connection
	config.Connect()
}
