package main

import (
	"awesomeProject/controllers"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	uuid "github.com/twinj/uuid"
	"net/http"
)

// client := resty.New()
// // get start time
// r := gin.Default()
//
//	r.GET("/ping", func(context *gin.Context) {
//		startTime := time.Now()
//		context.JSON(http.StatusOK, gin.H{
//			"message":       "pong",
//			"response_time": time.Since(startTime),
//		})
//	})
//
//	r.GET("/course/:term/:subject", func(context *gin.Context) {
//		term := context.Param("term")
//		subject := context.Param("subject")
//		var url = "https://prd.ps.pitt.edu/psc/pitcsprd/EMPLOYEE/SA/s/WEBLIB_HCX_CM.H_CLASS_SEARCH.FieldFormula.IScript_ClassSearch?institution=UPITT&term=" + term + "&date_from=&date_thru=&subject=" + subject + "&subject_like=&catalog_nbr=&time_range=&days=&campus=UPJ&location=&x_acad_career=&acad_group=&rqmnt_designtn=&instruction_mode=&keyword=&class_nbr=&acad_org=&enrl_stat=O&crse_attr=&crse_attr_value=&instructor_name=&session_code=&units=&page=1"
//		resp, err := client.R().EnableTrace().Get(url)
//		if err != nil {
//			context.JSON(http.StatusInternalServerError, gin.H{
//				"message": err,
//			})
//		}
//		context.String(http.StatusOK, resp.String())
//	})
//
// err := r.Run()
//
//	if err != nil {
//		return
//	}

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

	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"healthy": "true",
		})
	})

	v1 := router.Group("/v1")
	{
		cc := new(controllers.CourseController)
		v1.POST("/courses/info", cc.GetByTermAndSubject)
	}

	router.NoRoute(func(context *gin.Context) {
		context.HTML(404, "html/404.html", gin.H{})
	})

	router.Run()

}
