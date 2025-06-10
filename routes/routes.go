package routes

import (
	"bluebell/controller"
	"bluebell/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

// Init initializes all routes and returns the Gin engine
func Init(trans ut.Translator) *gin.Engine {
	r := gin.Default()

	// register all routes, starting with lower case can decouple
	registerDefaultRoute(r)
	registerNotFoundRoute(r)
	registerUserRoute(r, trans)
	registerPostRoute(r)

	return r
}

func registerDefaultRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
}

func registerNotFoundRoute(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})
}

// registerUserRoute registers the user endpoint
func registerUserRoute(r *gin.Engine, trans ut.Translator) {
	r.POST("/signup", controller.SignUpHandler(trans))
	r.POST("/login", controller.LoginHandler(trans))
	r.GET("/userinfo", middlewares.JWTAuthMiddleware(), controller.UserInfoHandler())
}

// registerPostRoute registers the post endpoint
func registerPostRoute(r *gin.Engine) {
	r.GET("/community", middlewares.JWTAuthMiddleware(), controller.CommunityHandler())
}
