package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"tech_platform/server/internal/model/user"
	"tech_platform/server/internal/store"
	userstore "tech_platform/server/internal/store/user"
	"tech_platform/server/pkg/jwtutil"
	"time"

	mymd5 "github.com/leeyf888/go-tools/md5"

	"tech_platform/server/pkg/snowflake"
)

type Handler struct {
	jwtHelper jwtutil.JWTHelper
}

func NewHandler(helper jwtutil.JWTHelper) *Handler {
	return &Handler{jwtHelper: helper}
}

func (h *Handler) Register(c context.Context, req user.RegisterRequest) (string, error) {
	u := user.User{
		Id:       snowflake.GetUniqueId(),
		Username: req.Username,
		Password: mymd5.Encryption(req.Password),
		Status:   user.ACTIVE,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	userId, err := userstore.Register(store.FromContext(c),u)
	if err != nil {
		return "", errors.New("register error")
	}

	return userId, nil
}

func (h *Handler) Login(c context.Context, req user.LoginRequest) (user.LoginResponse, error) {
	u := user.User{
		Username: req.Username,
		Password: mymd5.Encryption(req.Password),
	}
	us, err := userstore.Login(store.FromContext(c), u)
	if err != nil {
		return user.LoginResponse{}, fmt.Errorf("login error")
	}
	if us.Status == user.Forbidden {
		return user.LoginResponse{}, fmt.Errorf("this account was forbidden")
	}

	token, err := h.jwtHelper.GenToken(us.UserId)
	if err != nil {
		return user.LoginResponse{}, err
	}
	us.Token = token
	return us, nil
}

func (h *Handler) UpdateUserinfo(c context.Context, ui user.Userinfo) (user.Userinfo,error) {
	ui,err := userstore.UpdateUserinfo(store.FromContext(c),ui)
	if err!=nil{
		return user.Userinfo{},fmt.Errorf("update userinfo error")
	}
	return ui,nil
}

func (h *Handler) GetUserinfo(c *gin.Context, userId string) (user.Userinfo, error) {
	ui,err := userstore.GetUserinfo(store.FromContext(c),userId)
	if err!=nil{
		return user.Userinfo{},fmt.Errorf("get userinfo error")
	}
	return ui,nil
}
