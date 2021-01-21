package main

import (
	"galaxy-golang/server/global"
	"galaxy-golang/server/initialize"
	"galaxy-golang/server/routes"
)

func main() {
	// Our server will live in the routes package
	global.GVA_DB = initialize.Gorm()     // gorm连接数据库
	routes.Run()
}

