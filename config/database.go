package initializer

import (
	"log"
	"packages/helper/db"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	newLogger := &db.CustomLogger{
		Config: logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	}

	DB, err = gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/test"), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
