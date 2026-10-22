package models

import "gorm.io/gorm"

// 1.Post的钩子：创建文章时更新用户统计
// AfterCreate 钩子：在文章创建成功后触发
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 使用 tx (当前事务) 去更新对应用户的文章数量 +1
	// gorm.Expr 允许我们在更新时直接执行 SQL 运算
	err = tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + ?",1)).Error
		return
}

// 2.Post的钩子：在删除文章时更新用户的统计
// AfterDelete文章钩子：在文章删除成功后触发
func (p *Post) AfterDelete(tx *gorm.DB) (err error) {
	// 只有但post_count >0 时才做减法，防止出现负数
	err = tx.Model(&User{}).Where("id = ? AND post_count >0", p.UserID).
		Update("post_count", gorm.Expr("post_count -?", 1)).Error
		return
}

// 3.Comment的钩子：删除评论时检查状态
// 使用 AfterDelete（因为删除成功后，我们才能去统计剩下的数量）。
// 因为这里涉及到两步操作（查数量、更新状态），为了数据一致性，我们必须使用事务 tx.Transaction

// AfterDelete 钩子：在评论删除后触发
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 开启数据库事务，保证查数量和改状态要么全成功，要么全失败
	return tx.Transaction(func(tx *gorm.DB) error {
		// 1.统计该文章当前还剩下多少条评论
		var count int64
		if err := tx.Model(&Comment{}).Where("post_id =?", c.PostID).Count(&count).Error;
		err != nil {
			return err
		}
		
		// 2. 如果评论数量为 0，更新文章状态
		if count == 0 {
			if err := tx.Model(&Post{}).Where("id =?", c.PostID).
				Update("comment_status","无评论").Error; err != nil {
				return err 
			}
		}
		return nil
	})
}


// 4. Comment的钩子：新增评论时更新文章状态
// AfterCreate钩子：在评论创建成功后触发
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	// 只要有新评论产生，就把对应文章的状态更新为 "有评论"
	return tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_status", "有评论").Error
}