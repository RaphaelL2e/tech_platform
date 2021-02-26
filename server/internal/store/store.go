package store

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	userstore "tech_platform/server/internal/store/user"
	"time"
)
type Store interface {
	userstore.Store
}

type DataHandler struct {
	userstore.DataHandler
}

func New(resConfig string) Store {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 500 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:      logger.Info,            // Log level
			Colorful:      true,                  // 禁用彩色打印
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
		userstore.DataHandler{DB: db},
	}
}

