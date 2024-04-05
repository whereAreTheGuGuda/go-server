package main

import (
	"fmt"
	"go-server/global"
	"go-server/initialize"
	"strconv"

	// 初始化数据库连接及日志文件
	_ "go-server/common"
	// 数据模型中init方法的执行
	_ "go-server/model"
	// 文档
	// "go get github.com/swaggo"
)

func main() {
	// 初始化自定义校验器
	initialize.InitValidate()
	//2.初始化路由
	router := initialize.Routers()
	// 获取端口号
	PORT := strconv.Itoa(global.ServerConfig.Port)
	fmt.Println(PORT + "当前端口")
	global.Logger.Sugar().Infof("服务已经启动:localhost:%s", PORT)

	// 启动服务
	if err := router.Run(fmt.Sprintf(":%s", PORT)); err != nil {
		global.Logger.Sugar().Panic("服务启动失败:%s", err.Error())
	}
}
