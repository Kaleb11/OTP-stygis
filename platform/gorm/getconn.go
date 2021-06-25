package gorm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (grm *gormm) GetConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", grm.hostname, grm.port, grm.dbname, grm.username, grm.password)
	fmt.Println("constring", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dsn == "true" {
		print("Successfully connected to database")
	}

	if err != nil {
		panic("failed to connect database")
	}
	return db, err
}
func (grm *gormm) Migrate() error {
	db, err := grm.GetConnection()
	if err != nil {
		return err
	}
	db.AutoMigrate()
	fmt.Print("Database successfully migrated")

	return err
}
