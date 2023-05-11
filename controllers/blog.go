package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ishangoyal13/blogPost/models"
	logger "github.com/ishangoyal13/blogPost/pkg/log"
)

// get all blogs
func (base *BaseController) GetBlogs(c *gin.Context) {
	var (
		blogRepo = models.InitBlogRepo(base.DB)
	)

	ur, ok := c.Get("User")
	if !ok {
		logger.Info(ur)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
			"success": false,
		})
		return
	}

	user, ok := ur.(models.User)
	if !ok {
		logger.Info(user)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
			"success": false,
		})
		return
	}

	blogs, err := blogRepo.GetBlog(user.ID)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Something went wrong",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    blogs,
		"success": true,
	})
}

// post a new blog
func (base *BaseController) AddBlog(c *gin.Context) {
	var (
		request  AddBlogRequest
		blogRepo = models.InitBlogRepo(base.DB)
	)

	ur, ok := c.Get("User")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
			"success": false,
		})
		return
	}

	user, ok := ur.(models.User)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
			"success": false,
		})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to bind json",
			"success": false,
		})
		return
	}

	blog := models.Blog{
		Author:  request.Author,
		Content: request.Content,
		Title:   request.Title,
		UserID:  user.ID,
	}

	err := blogRepo.CreateBlog(&blog)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
		"success": true,
	})
}

// delete a blog
func (base *BaseController) DeleteBlog(c *gin.Context) {
	var (
		blogRepo = models.InitBlogRepo(base.DB)
	)

	ur, ok := c.Get("User")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
			"success": false,
		})
		return
	}

	user, ok := ur.(models.User)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
			"success": false,
		})
		return
	}

	blog_id := c.Param("id")
	if blog_id == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Blog id can't be empty",
			"success": false,
		})
		return
	}

	uintBlogId, err := strconv.ParseUint(blog_id, 10, 64)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Something went wrong",
			"success": false,
		})
		return
	}

	err = blogRepo.DeleteBlog(uintBlogId, user.ID)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Something went wrong",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully",
		"success": true,
	})
}
