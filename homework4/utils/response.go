package utils

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{"code":0, "message":"success", "data":data})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"code": code, "messge": msg})
}