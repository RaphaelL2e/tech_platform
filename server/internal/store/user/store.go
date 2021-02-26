package user

import (
	"context"
	"log"
	"os"
	"tech_platform/server/internal/model/user"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Store interface {
	register(u user.User) (string, error)
	login(u user.User) (user.LoginResponse, error)
	updateUserinfo(ui user.Userinfo)(user.Userinfo,error)
	getUserinfo(userId string)(user.Userinfo,error)
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

func Register(c context.Context,u user.User)(string,error){
	return FromContext(c).register(u)
}


func Login(c context.Context,u user.User)(user.LoginResponse,error){
	return FromContext(c).login(u)
}

func UpdateUserinfo(c context.Context,ui user.Userinfo)(user.Userinfo,error){
	return FromContext(c).updateUserinfo(ui)
}

func GetUserinfo(c context.Context,userId string)(user.Userinfo,error){
	return FromContext(c).getUserinfo(userId)
}
