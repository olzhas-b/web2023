package helper

import "github.com/gin-gonic/gin"

func GetUserID(c *gin.Context) uint64 {
	return c.GetUint64("userID")
}

func GetUsername(c *gin.Context) string {
	return c.GetString("username")
}
