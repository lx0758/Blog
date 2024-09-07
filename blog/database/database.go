package database

import (
	"blog/bean/po"
	"blog/config"
	bloglogger "blog/logger"
	"blog/util"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var gDb *gorm.DB

func GetDB() *gorm.DB {
	if gDb == nil {
		conf := config.Config().DataSource
		var driver gorm.Dialector
		switch conf.Type {
		case "sqlite":
			driver = sqlite.Open(conf.Dsn)
			break
		case "postgres":
			driver = postgres.Open(conf.Dsn)
			break
		default:
			bloglogger.Panic("Unsupported database types\n")
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
					LogLevel:                  util.If(conf.Debug, gormlogger.Info, gormlogger.Warn),
					IgnoreRecordNotFoundError: false,
					ParameterizedQueries:      false,
					Colorful:                  true,
				},
			),
		})
		if err != nil {
			bloglogger.Panicf("Failed to connect to database:%s\n", err)
		}

		if conf.Debug {
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
			bloglogger.Panicf("Failed auto migrate database:%s\n", err)
		}

		gDb = db
	}
	return gDb
}

func Pagination[PO interface{}](db *gorm.DB, pageNum int, pageSize int) po.Pagination[PO] {
	var total int64
	var pos = make([]PO, 0)
	db.Count(&total).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&pos)
	lastCount := util.If(int(total)%pageSize > 0, 1, 0)
	return po.Pagination[PO]{
		PageNum:  pageNum,
		PageSize: pageSize,
		Size:     int(total)/pageSize + lastCount,
		Total:    int(total),
		List:     pos,
	}
}
