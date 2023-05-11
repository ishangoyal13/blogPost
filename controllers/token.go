package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishangoyal13/blogPost/auth"
	"github.com/ishangoyal13/blogPost/models"
	logger "github.com/ishangoyal13/blogPost/pkg/log"
)

func (base *BaseController) GenerateToken(ctx *gin.Context) {
	var (
		request  TokenRequest
		userRepo = models.InitUserRepo(base.DB)
	)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check if email exists and password is correct
	user, err := userRepo.GetUser(request.PhoneNumber)
	if user.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Phone Number",
			"success": false,
		})
		return
	}
	if err != nil {
		logger.Error(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
			"success": false,
		})
		return
	}

	credentialError := userRepo.CheckPassword(request.Password, user)
	if credentialError != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Invalid Password",
			"success": false,
		})
		return
	}

	tokenString, err := auth.GenerateJWT(user.PhoneNumber, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":   tokenString,
		"success": true,
		"name":    user.Name,
	})
}
