package middleware

import (
	"fmt"
	"go-server/utils"

	"github.com/gin-gonic/gin"
)

func RecoverMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.Fail(c, fmt.Sprint(err))
				c.Abort()
				return
			}
		}()
	}
}
