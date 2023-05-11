package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishangoyal13/blogPost/auth"
	"github.com/ishangoyal13/blogPost/models"
	logger "github.com/ishangoyal13/blogPost/pkg/log"
)

func (base *BaseController) RegisterUser(ctx *gin.Context) {
	var (
		userRepo = models.InitUserRepo(base.DB)
		request  RegisterUser
	)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
			"success": false,
		})
		return
	}

	newUser := &models.User{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
	}

	err = userRepo.CreateUser(newUser)
	if err != nil {
		logger.Error(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"success": false,
		})
		return
	}

	jwtToken, err := auth.GenerateJWT(newUser.PhoneNumber, newUser.ID)
	if err != nil {
		logger.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"phone_number": newUser.PhoneNumber,
		"token":        jwtToken,
		"success":      true,
	})
}
