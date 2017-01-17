package main

import (
	"goshop/src/server/config"
	"goshop/src/server/route"
)

func main() {
	//读取resources配置信息
	config.SetUpDefaultConfig()
	//连接数据库
	config.SetUpDBConfig()
	//连接redis
	config.SetUpRedis()
	//设置前端static
	config.SetUpStatic()
	//设置前端HtmlTemplates
	config.SetUpHtmlTemplates()
	//配置路由
	route.SetUpRoute()
}
