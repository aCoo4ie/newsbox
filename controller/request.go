package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const CtxUserId = "user_id"

var (
	ErrUserNotLogin  = errors.New("用户没有登录")
	ErrInvalidUserID = errors.New("用户ID格式有误")
)

func GetId(c *gin.Context) (int64, error) {
	uid, ok := c.Get(CtxUserId)
	if !ok {
		return 0, ErrUserNotLogin
	}
	id, ok := uid.(int64)
	if !ok {
		return 0, ErrInvalidUserID
	}
	return id, nil
}
