package routers

import (
	"github.com/ishangoyal13/blogPost/controllers"
	"github.com/ishangoyal13/blogPost/pkg/config"
	"github.com/ishangoyal13/blogPost/routers/middleware"
	"github.com/sirupsen/logrus"
)

func RegisterRoutes(app config.App) {

	ctrl := controllers.BaseController{
		DB:     app.DB,
		Config: app.Config,
		Log:    logrus.New(),
	}

	apiGroup := app.Router.Group("/api")
	blogGroup := apiGroup.Group("/blog")
	userGroup := apiGroup.Group("/user")

	// user routes
	userGroup.POST("/token", ctrl.GenerateToken)
	userGroup.POST("/register", ctrl.RegisterUser)

	// blog routes
	blogGroup.Use(middleware.Auth(app.DB))
	blogGroup.GET("", ctrl.GetBlogs)
	blogGroup.POST("", ctrl.AddBlog)
	blogGroup.DELETE("/:id", ctrl.DeleteBlog)
}
