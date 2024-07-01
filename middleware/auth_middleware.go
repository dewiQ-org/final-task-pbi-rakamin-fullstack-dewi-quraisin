package middleware

import (
	"project-api-golang/errorhandler"
	"project-api-golang/helper"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			errorhandler.HandleError(c, &errorhandler.UnauthorizedError{Message: "Unauthorized"})
			c.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenString)
		if err != nil {
			errorhandler.HandleError(c, &errorhandler.UnauthorizedError{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", *userId)
		c.Next()
	}
}
