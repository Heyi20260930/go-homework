package models

import "time"

// Comment 评论表 (必须定义在最前面，因为 Post 要用)
type Comment struct {
	ID uint `gorm:"primaryKey"`
	Content string `gorm:"type:text;not null"`
	PostID uint `gorm:"index"`  // 外键，加index提升性能
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Post 文章表 (定义在 Comment 之后，User 之前)
type Post struct {
	ID uint `gorm:"primaryKey"`
	Title string `gorm:"size:128;not null"`
	Content string `gorm:"type:text"`
	UserID uint `gorm:"index"`  // 外键，加index提升性能
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments []Comment // // 一对多：自动寻找Comment里的PostID
}

// User 用户表
type User struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"size:64;not null"`
	Email string `gorm:"size:128;uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Posts []Post  // 一对多：自动寻找 Post 里的 UserID
}