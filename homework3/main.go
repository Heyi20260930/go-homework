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
	
	
	// ================= 测试 Post 钩子与 AfterDelete 修复 =================
	fmt.Println("测试 Post 钩子与 AfterDelete---------------------------------")
	// 1.查找李四。如果没查到，才创建
	var lisi models.User
	errLisi := db.Where("email = ?", "lisi@test.com").First(&lisi).Error
	if errLisi != nil {
		// 如果报错说找不到，说明李四不在，这里才创建
		lisi = models.User{Name: "李四", Email: "lisi@test.com"}
		db.Create(&lisi) 
	}
	// 删光李四名下所有的旧文章（防止脏数据干扰）
	db.Where("user_id = ?", lisi.ID).Delete(&models.Post{})
	// 2. 强制把李四的统计数重置为 0
	db.Model(&lisi).Update("post_count", 0)
	
	// 2.不管怎样，每次测试前，强制给李四重建一篇文章
	lisiPost := models.Post{Title:"李四的测试文章", UserID: lisi.ID}
	db.Create(&lisiPost)
	// 3. 重新查一下李四现在的文章数（应该是1）
	var checkUser models.User
	db.First(&checkUser, lisi.ID)
	fmt.Printf("删除文章前，李四的文章数: %d\n", checkUser.PostCount)
	// 4. 删除刚刚建的文章（触发AfterDelete，post_count -1）
	db.Delete(&lisiPost)
	db.First(&checkUser, lisi.ID)
	fmt.Printf("删除文章后，李四的文章数: %d\n", checkUser.PostCount)
	
	
	// ================= 测试 Comment 的 AfterDelete 钩子 =================
	fmt.Println("测试评论的AfterDelete钩子---------------------------------")
	
	// 1.新建一篇专门用来测试评论的干净文章
	testPost := models.Post{Title: "测试评论的文章", UserID: lisi.ID}
	db.Create(&testPost) 

	// 【为了严谨测试，手动改成 "无评论"】
	db.Model(&testPost).Update("comment_status", "无评论")

	// 2. 给这篇文章插入一条评论（让他触发 Comment.AfterCreate 钩子）
	newComment := models.Comment{Content: "李四的沙发", PostID: testPost.ID}
	db.Create(&newComment)

	// 3. 查询文章状态（看看钩子有没有把 "无评论" 自动变回 "有评论"）
	var postForComment models.Post
	db.First(&postForComment, testPost.ID)
	fmt.Printf("添加评论后，文章状态: %s\n", postForComment.CommentStatus)

	// 4. 执行删除评论动作（让它触发 Comment.AfterDelete 钩子）
	db.Delete(&newComment)

	// 5. 重新加载文章，看看状态有没有被减回 "无评论"
	db.First(&postForComment, testPost.ID)
	fmt.Printf("删除评论后，文章状态: %s\n", postForComment.CommentStatus)

}
