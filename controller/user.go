package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"errors"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// SignUpHandler handles user registration requests
func SignUpHandler(trans ut.Translator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p models.ParamSignUp
		err := c.ShouldBindJSON(&p)
		if err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				// not validation errors
				ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
				return
			}
			// params validation errors
			ResponseErrorWithMsg(c, CodeInvalidParams, errs.Translate(trans))
			return
		}
		// Call logic service
		err = logic.SignUp(p)
		if err != nil {
			// Handle specific errors
			if errors.Is(err, logic.ErrUserExists) {
				ResponseError(c, CodeUserExists)
				return
			}
			// Handle other errors
			ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
			return
		}

		// Success response
		ResponseSuccess(c, "注册成功")
	}
}

func LoginHandler(trans ut.Translator) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. get login params
		var p models.ParamLogin
		err := c.ShouldBindJSON(&p)
		if err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
				return
			}
			// params validation errors
			ResponseErrorWithMsg(c, CodeInvalidParams, errs.Translate(trans))
			return
		}
		// 2. logic
		token, err := logic.Login(p)
		if err != nil {
			if errors.Is(err, logic.ErrUserNotFound) {
				ResponseError(c, CodeInvalidPassword)
				return
			}

			ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
			return
		}

		// 3. return response
		// ResponseSuccess(c, "登录成功")
		ResponseSuccessWithMsg(c, "登录成功", token)
	}
}

func UserInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := GetId(c)
		if err != nil {
			ResponseErrorWithMsg(c, CodeInvalidToken, err.Error())
			return
		}
		ResponseSuccess(c, uid)
	}
}
