package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/pkg/snowflake"
	"bluebell/routes"
	"bluebell/settings"
	"fmt"
)

func main() {
	// 初始化 配置器
	if err := settings.InitConfig(); err != nil {
		fmt.Printf("Init Config error: %v\n", err)
		return
	}

	// 初始化 mysql
	if err := mysql.Init(); err != nil {
		fmt.Printf("Init mysql error: %v\n", err)
		return
	}
	defer mysql.GetDB().Close()

	// 初始化 UUID
	if err := snowflake.Init("2020-11-01", 1); err != nil {
		fmt.Printf("Init snowflake error: %v\n", err)
		return
	}

	// 初始化全局翻译器
	trans, err := controller.InitTrans("zh")
	if err != nil {
		fmt.Printf("Init Translator error: %v\n", err)
		return
	}

	// 初始化 Gin
	r := routes.Init(trans)

	// 启动 Gin 服务器
	r.Run(":9090")
}
