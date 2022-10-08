package main

import (
	"awesomeProject/controllers"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	uuid "github.com/twinj/uuid"
	"log"
	"net/http"
	"time"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		uuid := uuid.NewV4()
		context.Writer.Header().Add("X-Request-ID", uuid.String())
		context.Next() // goes to next middleware
	}
}

func ExitMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("Example")
	}
}

func main() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(RequestIDMiddleware())

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.StaticFile("/", "./html/404.html")

	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"healthy": "true",
		})
	})

	store := persistence.NewInMemoryStore(time.Second)

	v1 := router.Group("/v1/api")
	{
		cc := new(controllers.CourseController)
		v1.POST("/courses/info", cc.GetCourseInfo)
		v1.GET("/course/:term/:number", cache.CachePage(store, time.Minute, cc.GetCourse))

		lc := new(controllers.LaundryController)
		v1.GET("/laundry/:dormitory", cache.CachePage(store, time.Minute, lc.GetByDormitory))

	}

	router.Run(":7777")

}
