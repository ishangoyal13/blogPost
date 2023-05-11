package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ishangoyal13/blogPost/auth"
	"github.com/ishangoyal13/blogPost/models"
	logger "github.com/ishangoyal13/blogPost/pkg/log"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			user models.User
		)

		bearerToken := context.GetHeader("Authorization")
		if bearerToken == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		tokenString := strings.Split(bearerToken, "Bearer ")

		user_phoneNumber, user_id, err := auth.ParseClaims(tokenString[1])
		if err != nil {
			logger.Error(err)
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = db.Model(&models.User{}).Where(&models.User{PhoneNumber: user_phoneNumber, ID: user_id}).Find(&user).Error
		if err != nil {
			logger.Error(err)
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authentication invalid",
				"success": false,
			})
			return
		}

		context.Set("User", user)
		context.Next()
	}
}
