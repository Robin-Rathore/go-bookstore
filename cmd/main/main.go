package main

import (
	"log"
	"myapp/config" // Update with your actual module name
	"github.com/Robin-Rathore/go-bookstore/pkg/config"
)

func main() {
	config.Connect() // Initialize the database connection
	log.Println("🚀 Server is running...")
}
