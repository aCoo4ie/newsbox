package controller

import (
	"bluebell/logic"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommunityHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取社区名字以及ID
		data, err := logic.GetCommunityList()
		if err != nil {
			ResponseError(c, CodeServerBusy)
			return
		}
		ResponseSuccess(c, data)
	}
}

func CommunityDetailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从路径参数获取 id
		idStr := c.Param("id")
		// 参数校验
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			ResponseError(c, CodeInvalidParams)
			return
		}
		data, err := logic.GetCommunityDetailByID(id)
		fmt.Println(data, err)
		if err != nil {
			ResponseError(c, CodeServerBusy)
			return
		}
		ResponseSuccessWithMsg(c, CodeSuccess, data)
	}
}
