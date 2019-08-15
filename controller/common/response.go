package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendJSON 发送JSON
func SendJSON(c *gin.Context, data interface{}, args ...interface{}) {
	errID := 0
	reason := ""
	if len(args) == 2 {
		theErrID, ok := args[0].(int)
		if !ok {
			panic("errID 不正确")
		}
		errID = theErrID
		theCtx, ok := args[1].(string)
		if !ok {
			panic("缺少 Reason")
		}
		reason = theCtx
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"error": gin.H{
			"id":     errID,
			"reason": reason,
		},
	})
	// 终止请求链
	if errID != 0 {
		c.Abort()
	}
}
