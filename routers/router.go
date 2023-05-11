package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishangoyal13/blogPost/pkg/config"
	"github.com/ishangoyal13/blogPost/routers/middleware"
)

func SetupAndRunServer(app *config.App) {
	environment := app.Config.Server.Debug
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	// if app.Config.Server.VPCProxyCIDR != "" {
	// 	router.SetTrustedProxies([]string{app.Config.Server.VPCProxyCIDR})
	// }
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	app.Router = router

	// Register routes
	registerRoutes(*app)
	// Run Server after InitRoutes
	runServer(*app)
}

// registerRoutes add all routing list here automatically get main router
func registerRoutes(app config.App) {
	app.Router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	app.Router.GET("/health", func(ctx *gin.Context) {
		// Send a ping to make sure the database connection is alive.
		// db, err := app.DB.DB()
		// if err != nil {
		// 	ctx.JSON(http.StatusServiceUnavailable, gin.H{"live": "not ok"})
		// 	return
		// }
		// err = db.PingContext(ctx)
		// if err != nil {
		// 	ctx.JSON(http.StatusServiceUnavailable, gin.H{"live": "not ok"})
		// 	return
		// }
		ctx.JSON(http.StatusOK, gin.H{"live": "ok"})
	})

	// Register All routes
	RegisterRoutes(app)
}
