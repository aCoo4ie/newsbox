package controller

import (
	"bluebell/logic"

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
