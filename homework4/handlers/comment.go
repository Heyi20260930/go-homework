package handlers

import (
	"homework4/database"
	"homework4/models"
	"homework4/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentInput struct {
	Content string `json:"content" binding:"required"`
}

func CreateComment(c *gin.Context) {
	postID, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if err := database.DB.First(&post, postID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "文章不存在")
		return
	}
	var in CommentInput
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数校验失败")
		return
	}
	comment := models.Comment{
		Content: in.Content,
		UserID:  c.GetUint("userID"),
		PostID:  uint(postID),
	}
	// 1.先写入数据库
	if err := database.DB.Create(&comment).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "创建评论失败")
		return
	}
	// 2.把刚刚创建好的comment关联的User查询出来
	database.DB.Preload("User").First(&comment, comment.ID)
	
	// 3.返回完整的数据
	utils.Success(c, http.StatusCreated, comment)
}

func ListComments(c *gin.Context) {
	postID, _ := strconv.Atoi(c.Param("id"))
	var comments []models.Comment
	if err := database.DB.Preload("User").Where("post_id = ?", postID).
		Order("created_at DESC").Find(&comments).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取评论失败")
		return
	}
	utils.Success(c, http.StatusOK, comments)
}