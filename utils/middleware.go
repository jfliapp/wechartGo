package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// jwt token 中间件校验
func MiddleTokenParse(c *gin.Context) {
	token := c.GetHeader("Authorization")
	fmt.Println(token)
	// claim, err := ParseToken(token)
	tokenSplit := strings.Split(token, " ")
	token = strings.Join(tokenSplit[1:], "")
	_, err := ParseToken(token)
	if err != nil {
		fmt.Println("Authrization not right", err)
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "token 不对",
		})
		c.Abort()
		return
	}

	c.Next()
}
