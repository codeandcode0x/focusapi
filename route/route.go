package route

import (
	"focusapi/controller"
	"focusapi/middleware"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	TIME_DURATION   = 10
	INTERNAL_SERVER = "http://127.0.0.1:22520/"
)

func DefinitionRoute(router *gin.Engine) {
	// set run mode
	gin.SetMode(gin.DebugMode)
	// middleware
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(middleware.Tracing())
	router.Use(middleware.UseCookieSession())
	router.Use(middleware.TimeoutHandler(time.Second * TIME_DURATION))
	// home
	var userController *controller.UserController
	// var instanceController *controller.InstanceController
	var gatewayController *controller.GatewayController
	router.Static("/web/assets", "./web/assets")
	router.StaticFS("/web/upload", http.Dir("/web/upload"))
	router.LoadHTMLGlob("web/*.tmpl")
	// login
	router.GET("/login", userController.Login)
	router.POST("/dologin", userController.DoLogin)

	auth := router.Group("/")
	auth.Use(middleware.AuthMiddle())
	{
		auth.GET("/", userController.UserHome)
		auth.GET("/logout", userController.Logout)
		// user
		auth.GET("/users", userController.GetAllUsers)
		auth.GET("/user/add", userController.AddUser) //web ui
		auth.GET("/user/search", userController.SearchUsersByKeys)
		auth.POST("/user/create", userController.CreateUser)
		auth.POST("/user/update", userController.UpdateUser)
		auth.POST("/user/delete", userController.DeleteUser)

		// auth.GET("/apis", instanceController.GetAllInstances)
		// auth.GET("/api/add", instanceController.AddInstance) //web ui
		// auth.GET("/api/search", instanceController.SearchInstancesByKeys)
		// auth.POST("/api/create", instanceController.CreateInstance)
		// auth.POST("/api/update", instanceController.UpdateInstance)
		// auth.POST("/api/delete", instanceController.DeleteInstance)

		auth.GET("/:module", func(c *gin.Context) {
			method, _ := c.GetQuery("module")
			log.Println("okk", method)
			c.Redirect(http.StatusMovedPermanently, INTERNAL_SERVER+method)
		})
		auth.GET("/:module/add", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, INTERNAL_SERVER)
		})
		auth.GET("/:module/search", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, INTERNAL_SERVER)
		})
		auth.POST("/:module/create", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, INTERNAL_SERVER)
		})
		auth.POST("/:module/update", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, INTERNAL_SERVER)
		})
		auth.POST("/:module/delete", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, INTERNAL_SERVER)
		})
	}

	// no route
	router.NoRoute(gatewayController.Gateway(router))
	// api doc
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "USE_SWAGGER"))

}

// no route
func NoRouteResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":  404,
		"error": "oops, page not exists!",
	})
}
