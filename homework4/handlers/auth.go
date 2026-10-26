package handlers

import (
	"homework4/database"
	"homework4/models"
	"homework4/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var in RegisterInput
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数校验失败: "+err.Error())
		return
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "密码加密失败")
		return
	}
	user := models.User{Username: in.Username, Password: string(hashed), Email: in.Email}
	if err := database.DB.Create(&user).Error; err != nil {
		utils.Error(c, http.StatusConflict, "用户名或邮箱已存在")
		return
	}
	utils.Success(c, http.StatusCreated, gin.H{"id": user.ID, "username": user.Username})
}

func Login(c *gin.Context) {
	var in LoginInput
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数校验失败: "+err.Error())
		return
	}
	var user models.User
	if err := database.DB.Where("username = ?", in.Username).First(&user).Error; err != nil {
		utils.Error(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		utils.Error(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "生成令牌失败")
		return
	}
	utils.Success(c, http.StatusOK, gin.H{"token": token})
}