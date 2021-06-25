package gorm

import (
	"fmt"

	"gorm.io/gorm"
)

type gormm struct {
	hostname string
	port     int
	dbname   string
	username string
	password string
}
type Gorm interface {
	Migrate() error
	Seed() error
	GetConnection() (*gorm.DB, error)
}

func Initialize(hostname string, port int, dbname string, username string, password string) Gorm {
	return &gormm{
		hostname,
		port,
		dbname,
		username,
		password,
	}
}

func Migrate(gr gormm) error {
	init := Initialize(gr.hostname, gr.port, gr.dbname, gr.username, gr.password)
	db, err := init.GetConnection()
	//db := Initialize()
	if err != nil {
		return err
	}
	// db.AutoMigrate(model.User{})
	db.AutoMigrate()
	fmt.Print("Database successfully migrated")

	return err
}

func (g *gormm) Seed() error {

	return nil
}
