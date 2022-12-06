package database

import (
	"BackEnd/models"
	"BackEnd/pkg/connection"
	"fmt"
)

// automatic migration database if running app
func RunMigration() {
	err := connection.DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Successful")
}
