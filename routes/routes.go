package routes

import (
	"bluebell/controller"
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
	registerSignUpRoute(r, trans)

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

// registerSignUpRoute registers the signup endpoint
func registerSignUpRoute(r *gin.Engine, trans ut.Translator) {
	r.POST("/signup", controller.SignUpHandler(trans))
	r.POST("/login", controller.LoginHandler(trans))
}

// func SignUpRoute() *gin.Engine {
// 	// r := gin.New()
// 	// r.Use(logger, Recovery)
// 	// 使用默认的Gin
// 	r := gin.Default()

// 	r.POST("/signup", controller.SignUpHandler)

// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "OK")
// 	})

// 	r.NoRoute(func(c *gin.Context) {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"msg": "Not Found",
// 		})
// 	})
// }
