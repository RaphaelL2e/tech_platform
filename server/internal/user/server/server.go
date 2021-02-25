package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"tech_platform/server/internal/user/model"
	"tech_platform/server/pkg/jwtutil"
	"time"

	mymd5 "github.com/leeyf888/go-tools/md5"

	"tech_platform/server/internal/user/store"
	"tech_platform/server/pkg/snowflake"
)

type Handler struct {
	jwtHelper jwtutil.JWTHelper
}

func NewHandler(helper jwtutil.JWTHelper) *Handler {
	return &Handler{jwtHelper: helper}
}

func (h *Handler) Register(c context.Context, req model.RegisterRequest) (string, error) {
	u := model.User{
		Id:       snowflake.GetUniqueId(),
		Username: req.Username,
		Password: mymd5.Encryption(req.Password),
		Status:   model.ACTIVE,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	userId, err := store.Register(c, u)
	if err != nil {
		return "", errors.New("register error")
	}

	return userId, nil
}

func (h *Handler) Login(c context.Context, req model.LoginRequest) (model.LoginResponse, error) {
	u := model.User{
		Username: req.Username,
		Password: mymd5.Encryption(req.Password),
	}
	us, err := store.Login(c, u)
	if err != nil {
		return model.LoginResponse{}, fmt.Errorf("login error")
	}
	if us.Status == model.Forbidden {
		return model.LoginResponse{}, fmt.Errorf("this account was forbidden")
	}

	token, err := h.jwtHelper.GenToken(us.UserId)
	if err != nil {
		return model.LoginResponse{}, err
	}
	us.Token = token
	return us, nil
}

func (h *Handler) UpdateUserinfo(c context.Context, ui model.Userinfo) (model.Userinfo,error) {
	ui,err :=store.UpdateUserinfo(c,ui)
	if err!=nil{
		return model.Userinfo{},fmt.Errorf("update userinfo error")
	}
	return ui,nil
}

func (h *Handler) GetUserinfo(c *gin.Context, userId string) (model.Userinfo, error) {
	ui,err := store.GetUserinfo(c,userId)
	if err!=nil{
		return model.Userinfo{},fmt.Errorf("get userinfo error")
	}
	return ui,nil
}
