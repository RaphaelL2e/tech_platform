package store

import (
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Store interface {
}

type DataHandler struct {
	DB *gorm.DB
}

func New(resConfig string) *DataHandler {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 500 * time.Microsecond, // 慢 SQL 阈值
			LogLevel:      logger.Warn,            // Log level
			Colorful:      false,                  // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(resConfig), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database connection failed")
	}

	return &DataHandler{
		DB: db,
	}
}
