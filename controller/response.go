package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RespData struct {
	Code RespCode    `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code RespCode) {
	d := &RespData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, d)
}

func ResponseErrorWithMsg(c *gin.Context, code RespCode, msg interface{}) {
	d := &RespData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, d)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	d := &RespData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, d)
}

func ResponseSuccessWithMsg(c *gin.Context, msg interface{}, data interface{}) {
	d := &RespData{
		Code: CodeSuccess,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, d)
}
