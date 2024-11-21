package main

import (
	"log"
	"mygosimplerestfulapi/api"
	_ "mygosimplerestfulapi/docs"
)

// @title User Service API
// @version 1.0
// @description 這是一個用戶管理服務的API文檔
// @host localhost:8080
// @BasePath /
func main() {
	server := api.NewServer()
	if err := server.Run(":8080"); err != nil {
		log.Fatal("服務器啟動失敗:", err)
	}
}
