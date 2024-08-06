package main

import (
	"log"

	"github.com/Francisco-frc/APIsRest/config"
	"github.com/Francisco-frc/APIsRest/schemas"
)

var (
	logger *config.Logger
)

/*
func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Inicialize()
}
*/

func main() {
	db, err := config.InitializeSQLite()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Test inserting a record
	opening := schemas.Opening{
		Role:     "Developer",
		Company:  "Tech Corp",
		Location: "Remote",
		Remote:   true,
		Link:     "https://example.com",
		Salary:   100000,
	}

	result := db.Create(&opening)
	if result.Error != nil {
		log.Fatalf("Failed to create record: %v", result.Error)
	}

	log.Println("Record inserted successfully!")
}
