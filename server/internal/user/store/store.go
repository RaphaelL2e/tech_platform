package store

import (
	"context"
	"log"
	"os"
	"tech_platform/server/internal/user/model"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Store interface {
	register(u model.User) (string, error)
	login(u model.User) (model.LoginResponse, error)
	updateUserinfo(ui model.Userinfo)(model.Userinfo,error)
	getUserinfo(userId string)(model.Userinfo,error)
}

type DataHandler struct {
	DB *gorm.DB
}

func New(resConfig string) *DataHandler {
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
		DB: db,
	}
}

func Register(c context.Context,u model.User)(string,error){
	return FromContext(c).register(u)
}


func Login(c context.Context,u model.User)(model.LoginResponse,error){
	return FromContext(c).login(u)
}

func UpdateUserinfo(c context.Context,ui model.Userinfo)(model.Userinfo,error){
	return FromContext(c).updateUserinfo(ui)
}

func GetUserinfo(c context.Context,userId string)(model.Userinfo,error){
	return FromContext(c).getUserinfo(userId)
}
