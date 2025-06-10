package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 从数据库中获取所有社区名称
	return mysql.GetCommunityList()
}
