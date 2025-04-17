package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
type User struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar(20);not null"`
	Email      string `gorm:"type:varchar(50);uniqueIndex"`
	Posts      []Post `gorm:"foreignKey:UserID"` // 一对多关系配置
	PostsCount int    `gorm:"default:0"`         // 用户文章数量统计
}

type Post struct {
	ID             uint      `gorm:"primaryKey"`
	Title          string    `gorm:"type:varchar(100);not null"`
	Content        string    `gorm:"type:text"`
	UserID         uint      `gorm:"not null"`
	User           User      `gorm:"foreignKey:UserID"`              // 多对一关系配置
	Comments       []Comment `gorm:"foreignKey:PostID"`              // 一对多关系配置
	CommentsStatus string    `gorm:"type:varchar(20);default:'无评论'"` // 文章评论状态，显示评论数量或'无评论'
}

// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 更新用户的文章数量统计
	err := tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("posts_count", gorm.Expr("posts_count + 1")).Error
	if err != nil {
		return err
	}
	// 初始化文章评论状态
	return tx.Model(p).Update("comments_status", "无评论").Error
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text"`
	PostID  uint   `gorm:"not null"`
	Post    Post   `gorm:"foreignKey:PostID"`
}

func (c *Comment) AfterCreate(tx *gorm.DB) error {
	// 获取文章的评论数量
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	// 更新文章的评论状态
	return tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comments_status", fmt.Sprintf("%d条评论", count)).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 检查文章的评论数量
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	// 更新文章的评论状态
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comments_status", "无评论").Error
	} else {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comments_status", fmt.Sprintf("%d条评论", count)).Error
	}
	return nil
}

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/db_shard_0?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// 自动迁移数据库表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		fmt.Println("数据库迁移失败:", err)
		return
	}
	fmt.Println("数据库迁移成功")

	// 插入测试数据
	users := []User{
		{Name: "张三", Email: "zhangsan@example.com"},
		{Name: "李四", Email: "lisi@example.com"},
		{Name: "王五", Email: "wangwu@example.com"},
		{Name: "赵六", Email: "zhaoliu@example.com"},
		{Name: "钱七", Email: "qianqi@example.com"},
	}
	db.Create(&users)

	posts := []Post{
		{Title: "第一篇文章", Content: "这是张三的第一篇文章内容", UserID: users[0].ID},
		{Title: "第二篇文章", Content: "这是李四的第一篇文章内容", UserID: users[1].ID},
		{Title: "第三篇文章", Content: "这是王五的第一篇文章内容", UserID: users[2].ID},
		{Title: "第四篇文章", Content: "这是赵六的第一篇文章内容", UserID: users[3].ID},
		{Title: "第五篇文章", Content: "这是钱七的第一篇文章内容", UserID: users[4].ID},
	}
	db.Create(&posts)

	comments := []Comment{
		{Content: "张三对第一篇文章的评论", PostID: posts[0].ID},
		{Content: "李四对第一篇文章的评论", PostID: posts[0].ID},
		{Content: "王五对第二篇文章的评论", PostID: posts[1].ID},
		{Content: "赵六对第三篇文章的评论", PostID: posts[2].ID},
		{Content: "钱七对第四篇文章的评论", PostID: posts[3].ID},
	}
	db.Create(&comments)

	// 题目2：关联查询
	// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	userID := users[0].ID
	var userWithPosts User
	err = db.Preload("Posts.Comments").First(&userWithPosts, userID).Error
	if err != nil {
		fmt.Println("查询用户文章失败:", err)
		return
	}
	fmt.Printf("用户 %s 的文章及评论:\n", userWithPosts.Name)
	for _, post := range userWithPosts.Posts {
		fmt.Printf("文章ID: %d, 标题: %s\n", post.ID, post.Title)
		for _, comment := range post.Comments {
			fmt.Printf("\t评论ID: %d, 内容: %s\n", comment.ID, comment.Content)
		}
	}

	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	var mostCommentedPost Post
	err = db.Model(&Post{}).
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		First(&mostCommentedPost).Error
	if err != nil {
		fmt.Println("查询最多评论文章失败:", err)
		return
	}
	// 获取该文章的评论数量
	var commentCount int64
	db.Model(&Comment{}).Where("post_id = ?", mostCommentedPost.ID).Count(&commentCount)
	fmt.Printf("评论最多的文章: %s (评论数: %d)\n", mostCommentedPost.Title, commentCount)

}
