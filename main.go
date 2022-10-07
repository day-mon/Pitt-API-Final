package main

import (
	"awesomeProject/controllers"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	uuid "github.com/twinj/uuid"
	"net/http"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		uuid := uuid.NewV4()
		context.Writer.Header().Add("X-Request-ID", uuid.String())
		context.Next() // goes to next middleware
	}
}

func main() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(RequestIDMiddleware())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"healthy": "true",
		})
	})

	v1 := router.Group("/v1/api")
	{
		cc := new(controllers.CourseController)
		v1.POST("/courses/info", cc.GetCourseInfo)
		v1.GET("/course/:term/:number", cc.GetCourse)

		lc := new(controllers.LaundryController)
		v1.GET("/laundry/:dormitory", lc.GetByDormitory)

	}

	router.Run(":7777")

}
