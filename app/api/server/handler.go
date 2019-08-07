package server

import (
	"github.com/PhantomX7/go-cleanarch/app/api/middleware"
	"github.com/PhantomX7/go-cleanarch/util/validators"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/go-playground/validator.v8"
	"log"
	"net/http"
	"os"
)

// All handler that need to be registered MUST implement this interface
type Handler interface {
	Register(r *gin.Engine, m *middleware.Middleware)
}

func BuildHandler(middleware *middleware.Middleware, handlers ...Handler) http.Handler {

	if (os.Getenv("APP_ENV") == "production") {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// register all validator here
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("unique", validators.CustomValidator.Unique() )
		if err != nil {
			log.Println("error when applying validator")
		}
	}

	// apply global middleware here
	router.Use(cors.Default(), middleware.ErrorHandle())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/healthz", healthz)
	//router.GET("/test",test)

	// start registering routes from all handlers
	for _, reg := range handlers {
		reg.Register(router, middleware)
	}

	// 404 not found function
	router.NoRoute(notFound)

	return router
}

func healthz(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

//func test(c *gin.Context) {
//	err := errors.New("testtt")
//	if err != nil {
//		// put the Error to gin.Context.Errors
//		_ = c.Error(err)
//		return
//	}
//
//	c.String(http.StatusOK, "ok")
//}

func notFound(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}

