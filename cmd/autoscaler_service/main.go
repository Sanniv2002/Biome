package main

import (
	"time"

	service "github.com/sanniv2002/nether/internal/autoscaler/service"
)

func main() {
	// Start the monitoring service with a specified interval
	go service.InitializeAndRun(5 * time.Second) // Adjust the interval as needed

	// Keep the main function running
	select {}
}
