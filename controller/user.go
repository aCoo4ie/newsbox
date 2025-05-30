package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"net/http"

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
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": err.Error(),
				})
				return
			}
			// params validation errors
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": errs.Translate(trans),
			})
			return
		}
		// Call logic service
		err = logic.SignUp(p)
		if err != nil {
			// Handle specific errors
			if errors.Is(err, logic.ErrUserExists) {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "用户已存在，无法注册",
				})
				return
			}
			// Handle other errors
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}

		// Success response
		c.JSON(http.StatusOK, gin.H{
			"msg": "SignUp Success",
		})
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
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": err.Error(),
				})
				return
			}
			// params validation errors
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": errs.Translate(trans),
			})
			return
		}
		// 2. logic
		err = logic.Login(p)
		if err != nil {
			if errors.Is(err, logic.ErrUserNotFound) {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "用户名或密码错误",
				})
				return
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}

		// 3. return response
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	}
}
