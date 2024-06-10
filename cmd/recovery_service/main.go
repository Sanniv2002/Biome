package main

import (
	"time"

	service "github.com/sanniv2002/nether/internal/recovery/service"
)

func main() {
	// Start the recovery service with a specified interval
	go service.InitializeAndRun(10 * time.Second) // Adjust the interval as needed

	// Keep the main function running
	select {}
}
