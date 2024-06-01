package main

import (
	"admin/core"
	"admin/db"
	"admin/global"
	"fmt"
)

func main() {
	// 初始化yaml配置
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = db.InitGorm()
	fmt.Println(global.DB)
}
