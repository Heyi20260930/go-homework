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
	user := models.User{
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
	//if err := db.Create(&user).Error; err != nil {
	//	log.Fatalf("插入测试数据失败： %v", err)
	//}
	// 先按邮箱查找张三，如果找不到，才执行创建
	if err := db.Where(models.User{Email: "zhangsan@test.com"}).FirstOrCreate(&user).Error; err != nil {
    log.Fatalf("插入测试数据失败： %v", err)
}
	fmt.Println("测试数据插入成功！")

	// 查询 ID 为 1 的文章，并把它的评论一起查出来（预加载）
	/* var post models.Post
	db.Preload("Comments").First(&post,1)
	fmt.Printf("文章标题： %s\n", post.Title)
	fmt.Printf("文章数量： %d\n", len(post.Comments))
	for _, c := range post.Comments{
		fmt.Printf(" - %s\n", c.Content)
	}	 */

	// 从 User 开始查
	//var user models.User
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

	// 编写Go代码，使用Gorm查询评论数量最多的文章信息
	// 定义一个结构体来接收统计结果
	type PostStat struct {
		ID uint
		Title string
		CommentCount int64
	}
	
	var topPost PostStat
	
	errCount := db.Model(&models.Post{}).
		Select("posts.id, posts.title, COUNT(comments.id) AS comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&topPost).Error
		
	if errCount != nil {
		log.Fatalf("统计失败: %v", errCount)
	}
	
	fmt.Printf("评论最多的文章：%s (ID: %d)，评论数：%d\n", topPost.Title,topPost.ID, topPost.CommentCount)
	
	// 1. 测试 Post 钩子
	// var user models.User
	db.First(&user, 1) // 假设刚才插入了张三（ID=1）
	fmt.Printf("添加文章前，张三的文章数: %d\n", user.PostCount)
	
	
	// 测试钩子：创建一个全新的用户（李四），插入新文章
	lisi := models.User{
		Name:  "李四",
		Email: "lisi@test.com", // 全新的邮箱
		Posts: []models.Post{
			{Title: "李四的文章"},
		},
	}
	db.Create(&lisi) // 这会百分百触发 AfterCreate 钩子

	// 重新查李四
	var checkUser models.User
	db.Where("email = ?", "lisi@test.com").First(&checkUser)
	fmt.Printf("李四的文章数统计: %d\n", checkUser.PostCount) 
	
	
	
	// 测试 Comment 的 AfterDelete 钩子
	fmt.Printf("测试 Comment 的 AfterDelete 钩子------------------------------------------------------------\n") 
	// 刚才给李四插入了一篇文章，但没插评论。我们先给他插一条评论
	newComment := models.Comment{Content: "李四的沙发", PostID:2}
	db.Create(&newComment)

	// 此时文章状态应该是默认的 "有评论"
	var lisiPost models.Post
	db.First(&lisiPost, 2)
	fmt.Printf("删除评论前，文章状态: %s\n", lisiPost.CommentStatus)

	// 执行删除评论动作，这会触发你的 AfterDelete 钩子
	db.Delete(&newComment)

	// 重新加载文章，看看状态有没有变成 "无评论"
	db.First(&lisiPost,2)
	fmt.Printf("删除评论后，文章状态: %s\n", lisiPost.CommentStatus)

}
