package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type CourseController struct{}
type CourseFunctionParams struct {
	Term             string `json:"term,omitempty"`
	Subject          string `json:"subject,omitempty"`
	Campus           string `json:"campus,omitempty"`
	EnrollmentStatus string `json:"enrollmentStatus,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	CourseNumber     string `json:"courseNumber,omitempty"`
}

func (c *CourseController) GetCourseInfo(context *gin.Context) {

	// get post BODY from request
	var params CourseFunctionParams
	if err := context.ShouldBindJSON(&params); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// send request to Pitt API
	jsonStr, err := GetCourse(params, context)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return

	}

	// return response
	context.Data(200, gin.MIMEJSON, []byte(jsonStr))
}

func (c *CourseController) GetCourse(context *gin.Context) {
	term := context.Param("term")
	courseNumber := context.Param("number")

	if term == "" {
		context.JSON(400, gin.H{"error": "term cannot be empty"})
		return
	}

	_, err := ValidateTerm(term)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if courseNumber == "" {
		context.JSON(400, gin.H{"error": "courseNumber cannot be empty"})
		return
	}

	params := CourseFunctionParams{
		Term:         term,
		CourseNumber: courseNumber,
	}

	jsonStr, err := GetCourse(params, context)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// inline string array

	finalJson, err := omitKeys(jsonStr, []string{"cfg", "additionalLinks"})

	if err != nil {
		context.Data(200, gin.MIMEJSON, []byte(jsonStr))
		// log a warning
		log.Fatal("error omitting cfg key")
	}

	context.Data(200, gin.MIMEJSON, []byte(finalJson))
}

func GetCourse(params CourseFunctionParams, ctx *gin.Context) (string, error) {

	_, err := ValidateParams(&params, ctx)
	if err != nil {
		return "", err
	}

	url, err := buildUrl(params, ctx)
	if err != nil {
		return "", err
	}

	client := resty.New()
	resp, err := client.R().EnableTrace().Get(url)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func ValidateParams(params *CourseFunctionParams, ctx *gin.Context) (bool, error) {
	termValid, err := ValidateTerm(params.Term)

	if err != nil {
		return false, err
	}

	if !termValid {
		return false, errors.New("Term [" + params.Term + "] is not valid")
	}

	if params.Subject == "" && ctx.Request.Method == "POST" {
		return false, errors.New("subject cannot be empty")
	}

	params.EnrollmentStatus = strings.ToUpper(params.EnrollmentStatus)

	// check to see if its o c or w or its not empty
	if params.EnrollmentStatus != "" && params.EnrollmentStatus != "O" && params.EnrollmentStatus != "C" && params.EnrollmentStatus != "W" {
		return false, errors.New("enrollmentStatus must be O, C, or W")
	}

	params.EnrollmentStatus = "O"

	if params.Campus == "" {
		params.Campus = "UPJ"
	}

	if params.PageNumber == 0 {
		params.PageNumber = 1
	}

	params.Subject = strings.ToUpper(params.Subject)
	params.Campus = strings.ToUpper(params.Campus)

	return true, nil
}

func buildUrl(params CourseFunctionParams, ctx *gin.Context) (string, error) {

	if ctx.Request.Method == "POST" {
		return "https://prd.ps.pitt.edu/psc/pitcsprd/EMPLOYEE/SA/s/WEBLIB_HCX_CM.H_CLASS_SEARCH.FieldFormula.IScript_ClassSearch?institution=UPITT&term=" + params.Term + "&date_from=&date_thru=&subject=" + params.Subject + "&subject_like=&catalog_nbr=&time_range=&days=&campus=" + params.Campus + "&location=&x_acad_career=&acad_group=&rqmnt_designtn=&instruction_mode=&keyword=&class_nbr=&acad_org=&enrl_stat=" + params.EnrollmentStatus + "&crse_attr=&crse_attr_value=&instructor_name=&session_code=&units=&page=" + strconv.Itoa(params.PageNumber), nil
	} else if ctx.Request.Method == "GET" {
		return "https://prd.ps.pitt.edu/psc/pitcsprd/EMPLOYEE/SA/s/WEBLIB_HCX_CM.H_CLASS_SEARCH.FieldFormula.IScript_ClassDetails?institution=UPITT&term=" + params.Term + "&class_nbr=" + params.CourseNumber, nil
	} else {
		return "", errors.New("invalid request method")
	}

}

func ValidateTerm(term string) (bool, error) {
	termRegex := "2\\d\\d[147]"
	// test the regex
	return regexp.Match(termRegex, []byte(term))
}

func omitKeys(jsonStr string, keys []string) (string, error) {
	var i interface{}
	if err := json.Unmarshal([]byte(jsonStr), &i); err != nil {
		return "", err
	}
	if m, ok := i.(map[string]interface{}); ok {
		for _, key := range keys {
			delete(m, key) // No problem if "foo" isn't in the map
		}
	}

	output, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
