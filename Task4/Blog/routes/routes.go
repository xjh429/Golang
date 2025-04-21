package routes

import (
	"blog/controller"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 用户注册路由
	router.POST("/register", controller.RegisterUser)
	// 用户登录路由
	router.POST("/login", controller.LoginUser)

	// 受保护的路由组，需要认证
	protected := router.Group("/protected", middleware.AuthMiddleware())
	{
		// 文章相关路由
		protected.POST("/create-article", controller.CreateArticle)
		protected.GET("/articles", controller.GetArticles)
		protected.GET("/articles/:id", controller.GetArticle)
		protected.PUT("/articles/:id", controller.UpdateArticle)
		protected.DELETE("/articles/:id", controller.DeleteArticle)

		// 评论相关路由
		protected.POST("/user/:post_id/comments", controller.CreateComment)
		protected.GET("/user/:post_id/comments", controller.GetComments)
	}

	return router
}
