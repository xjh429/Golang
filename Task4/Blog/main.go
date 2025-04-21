package main

import (
	// "blog/config"
	// "blog/models"
	"blog/routes"
)

func main() {
	// db := config.GetDB()
	// // 自动迁移表
	// _ = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	router := routes.SetupRouter()
	_ = router.Run(":8080")

}
