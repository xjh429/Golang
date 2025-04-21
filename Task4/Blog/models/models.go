package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"unique;not null" json:"email"`
	Posts    []Post `gorm:"foreignKey:UserID" json:"posts"`
}

type Post struct {
	gorm.Model
	Title    string    `gorm:"not null" json:"title"`
	Content  string    `gorm:"not null;type:text" json:"content"`
	UserID   uint      `json:"user_id"`
	User     User      `gorm:"references:ID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null;type:text" json:"content"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
	Post    Post   `gorm:"references:ID"`
	User    User   `gorm:"references:ID"`
}
