package database

import (
	"blog/bean/po"
	"blog/config"
	bloglogger "blog/logger"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var gDb *gorm.DB

func Init() {
	GetDB()
}

func GetDB() *gorm.DB {
	if gDb == nil {
		dataSource := config.Config().DataSource
		var driver gorm.Dialector
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
		db, err := gorm.Open(driver, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "t_",
				SingularTable: true,
				NoLowerCase:   false,
				NameReplacer:  nil,
			},
			Logger: gormlogger.New(
				bloglogger.Logger(),
				gormlogger.Config{
					SlowThreshold:             time.Second,
					LogLevel:                  gormlogger.Info,
					IgnoreRecordNotFoundError: false,
					ParameterizedQueries:      false,
					Colorful:                  true,
				},
			),
		})
		if err != nil {
			panic("Failed to connect to database")
		}

		if config.Config().DataSource.Debug {
			db.Debug()
		}

		err = db.AutoMigrate(
			&po.User{},
			&po.Config{},
			&po.Feature{},
			&po.Link{},

			&po.Url{},
			&po.File{},

			&po.Fragment{},
			&po.Article{},
			&po.ArticleUrl{},
			&po.Category{},
			&po.Tag{},

			&po.Comment{},
		)
		if err != nil {
			panic("Failed auto migrate database")
		}

		gDb = db
	}
	return gDb
}

func Pagination(db *gorm.DB, pageNum int, pageSize int) po.Pagination {
	var total int64
	db.Count(&total)
	return po.Pagination{
		PageNum:  pageNum,
		PageSize: pageSize,
		Size:     int(total) / pageSize,
		Total:    int(total),
	}
}
