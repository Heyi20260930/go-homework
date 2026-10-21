package main

import (
	"fmt"
	"homework3/models" // 引入models包
	"log"

	"github.com/glebarez/sqlite" // 引入纯Go的SQLite驱动
	"gorm.io/gorm"               // <--- 加上这一行！！！
)

func main() {
	// 1. 连接 SQLite 数据库
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	fmt.Println("数据库连接成功!")

	// 2. 自动迁移，创建表结构和外键
	// 注意顺序：先建 User，再建 Post，最后建 Comment
	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatalf("建表失败： %v", err)
	}

	fmt.Println("博客系统表结构（User, Post, Comment）创建成功！")

	// 3. 插入测试数据 (加在这里)
	// 利用 GORM 的关联创建，直接造一个嵌套结构体
	/* user := models.User{
		Name:  "张三",
		Email: "zhangsan@test.com",
		Posts: []models.Post{
			{
				Title:   "Go语言初体验",
				Content: "GORM 真的是太好用了！",
				Comments: []models.Comment{
					{Content: "沙发！"},
					{Content: "顶一个！"},
				},
			},
		},
	}

	// 只要执行这一句，GORM 会把 User、Post、Comment 全部按顺序插入，并自动填好外键
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("插入测试数据失败： %v", err)
	}
	fmt.Println("测试数据插入成功！") */

	// 查询 ID 为 1 的文章，并把它的评论一起查出来（预加载）
	/* var post models.Post
	db.Preload("Comments").First(&post,1)
	fmt.Printf("文章标题： %s\n", post.Title)
	fmt.Printf("文章数量： %d\n", len(post.Comments))
	for _, c := range post.Comments{
		fmt.Printf(" - %s\n", c.Content)
	}	 */

	// 从 User 开始查
	var user models.User
	// 从 User 开始查，预加载 Posts，再预加载 Posts 里的 Comments
	db.Preload("Posts.Comments").First(&user, 1)
	// 打印结果
	fmt.Printf("用户: %s\n", user.Name)

	for _, p := range user.Posts {
		fmt.Printf("文章标题： %v\n", p.Title)
		for _, c := range p.Comments {
			fmt.Printf("  - 评论: %s\n", c.Content)
		}
	}

}
