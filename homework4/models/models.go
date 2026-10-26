
package models

import "gorm.io/gorm"


type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Email    string `gorm:"unique;not null" json:"email"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"user_id"`
	User    User    `json:"user,omitempty"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"Content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user,omitempty"`
	PostID  uint   `json:"post_id"`
	Post    Post   `json:"-"`
}