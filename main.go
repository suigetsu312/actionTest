package main

import (
	"github.com/gin-gonic/gin"
	"example/gintest/db"
	"example/gintest/models"
	"example/gintest/routes"
)

func main() {
	// 初始化數據庫
	db.InitDatabase()
	// 自動遷移模型
	db.DB.AutoMigrate(&models.Article{})
	// test commit
	// 初始化 Gin 路由
	r := gin.Default()

	// 註冊文章相關路由
	routes.RegisterArticleRoutes(r)

	// 啟動服務
	r.Run(":8080")
}
