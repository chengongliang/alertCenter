package middleware

import (
	"alertCenter/config"
	"alertCenter/controller/common"

	"github.com/gin-gonic/gin"
)

// TokenRequired header 必须包含 token
func TokenRequired(c *gin.Context) {
	SendJSON := common.SendJSON
	tokenString := c.GetHeader("token")
	if tokenString != config.Server.Token {
		SendJSON(c, make(map[string]string, 0), 9001, "token 错误!")
		return
	}
	c.Next()
}
