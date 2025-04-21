package controller

import (
	"blog/config"
	"blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateArticle 处理文章创建请求
func CreateArticle(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "上下文中未找到用户ID"})
		return
	}
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserID = uint(userID.(uint))
	// 数据库连接
	db := config.GetDB()
	if db == nil {
		return
	}
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article created successfully", "article": post})
}

// GetArticles 处理获取所有文章请求
func GetArticles(c *gin.Context) {
	// 数据库连接
	db := config.GetDB()
	var posts []models.Post
	//Preload 预加载User
	if err := db.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"articles": posts})
}

// GetArticle 处理获取单个文章请求
func GetArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}
	// 数据库连接
	db := config.GetDB()
	var post models.Post
	//不展示User set User = nil
	if err := db.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章未找到"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"article": post})
}

// UpdateArticle 处理文章更新请求
func UpdateArticle(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "上下文中未找到用户ID"})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}
	// 数据库连接
	db := config.GetDB()
	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章未找到"})
		return
	}
	if post.UserID != uint(userID.(uint)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "您不是该文章的作者"})
		return
	}
	var updatedPost models.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.Title = updatedPost.Title
	post.Content = updatedPost.Content
	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article updated successfully", "article": post})
}

// DeleteArticle 处理文章删除请求
func DeleteArticle(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "上下文中未找到用户ID"})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}
	// 数据库连接
	db := config.GetDB()
	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章未找到"})
		return
	}
	if post.UserID != uint(userID.(uint)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "您不是该文章的作者"})
		return
	}
	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
