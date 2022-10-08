package main

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/twinj/uuid"
	"net/http"
	"pittapi/controllers"
	_ "pittapi/docs"
	"time"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		uuid := uuid.NewV4()
		context.Writer.Header().Add("X-Request-ID", uuid.String())
		context.Next() // goes to next middleware
	}
}

// @title University of Pittsburgh API
// @description This is a sample server for Pitt API.
// @version 1.0.0
// @BasePath /v1/api
// @host 0.0.0.0:8080
// @schemes http https
// @name University of Pittsburgh API
// @description Github: https://github.com/day-mon/Pitt-API-Final
func main() {
	router := gin.New()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(RequestIDMiddleware())
	router.Use(gin.Logger())

	// run gin in release mode if we are in production environment if not run in debug mode

	// @Summary Simple health check
	// @Description Simple health check
	// @Tags Health
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} map[string]string
	// @Router /health [get]
	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"healthy": "true",
		})
	})

	store := persistence.NewInMemoryStore(time.Second)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.NoRoute(func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	v1 := router.Group("/v1/api")
	{
		cc := new(controllers.CourseController)
		v1.POST("/courses/info", cc.GetCourseInfo)
		v1.GET("/course/:term/:number", cache.CachePage(store, time.Minute, cc.GetCourse))

		lc := new(controllers.LaundryController)
		v1.GET("/laundry/:dormitory", cache.CachePage(store, time.Minute, lc.GetByDormitory))

	}

	gin.SetMode(gin.TestMode)
	router.Run(":8080")

}
