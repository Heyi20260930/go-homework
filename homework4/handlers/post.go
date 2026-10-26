package handlers

import (
	"homework4/database"
	"homework4/models"
	"homework4/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var in PostInput
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数校验失败: "+err.Error())
		return
	}
	userID := c.GetUint("userID")
	post := models.Post{Title: in.Title, Content: in.Content, UserID: userID}
	if err := database.DB.Create(&post).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "创建文章失败")
		return
	}
	utils.Success(c, http.StatusCreated, post)
}

func ListPosts(c *gin.Context) {
	var posts []models.Post
	if err := database.DB.Preload("User").Order("created_at DESC").Find(&posts).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取文章列表失败")
		return
	}
	utils.Success(c, http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if err := database.DB.Preload("User").First(&post, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "文章不存在")
		return
	}
	utils.Success(c, http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID")

	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "文章不存在")
		return
	}
	if post.UserID != userID {
		utils.Error(c, http.StatusForbidden, "无权修改他人的文章")
		return
	}
	var in PostInput
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数校验失败")
		return
	}
	post.Title, post.Content = in.Title, in.Content
	if err := database.DB.Save(&post).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "更新文章失败")
		return
	}
	// 更新成功后，返回完整的数据，重新预加载User
	database.DB.Preload("User").First(&post,post.ID)
	utils.Success(c, http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID")

	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "文章不存在")
		return
	}
	if post.UserID != userID {
		utils.Error(c, http.StatusForbidden, "无权删除他人的文章")
		return
	}
	if err := database.DB.Delete(&post).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "删除文章失败")
		return
	}
	utils.Success(c, http.StatusOK, gin.H{"deleted_id": id})
}