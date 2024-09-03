package db

import (
	"blog/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var gDb *gorm.DB

func DB() *gorm.DB {
	if gDb == nil {
		dataSource := config.Config().DataSource
		driver := (gorm.Dialector)(nil)
		switch dataSource.Type {
		case "sqlite":
			driver = sqlite.Open(dataSource.Dsn)
			break
		case "postgres":
			driver = postgres.Open(dataSource.Dsn)
			break
		default:
			panic("Unsupported database types")
		}
		db, err := gorm.Open(driver, &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database")
		}
		gDb = db
	}
	return gDb
}
