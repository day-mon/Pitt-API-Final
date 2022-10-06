package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
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
}

func (c *CourseController) GetCourseInfo(context *gin.Context) {

	// get post BODY from request
	var params CourseFunctionParams
	if err := context.ShouldBindJSON(&params); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// send request to Pitt API
	json, err := GetCourse(params)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return

	}

	// return response
	context.Data(200, gin.MIMEJSON, []byte(json))
}

func (c *CourseController) GetCourse(context *gin.Context) {
	// get url / params
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

	// send request to Pitt API
	url := "https://prd.ps.pitt.edu/psc/pitcsprd/EMPLOYEE/SA/s/WEBLIB_HCX_CM.H_CLASS_SEARCH.FieldFormula.IScript_ClassDetails?institution=UPITT&term=" + term + "&class_nbr=" + courseNumber

	client := resty.New()
	resp, err := client.R().EnableTrace().Get(url)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
	}

	// return response
	context.Data(200, gin.MIMEJSON, []byte(resp.String()))
}

func GetCourse(params CourseFunctionParams) (string, error) {

	// validate the params
	_, err := ValidateParams(&params)
	if err != nil {
		return "", err
	}

	baseUrl := "https://prd.ps.pitt.edu/psc/pitcsprd/EMPLOYEE/SA/s/WEBLIB_HCX_CM.H_CLASS_SEARCH.FieldFormula.IScript_ClassSearch?institution=UPITT&term=" + params.Term + "&date_from=&date_thru=&subject=" + params.Subject + "&subject_like=&catalog_nbr=&time_range=&days=&campus=" + params.Campus + "&location=&x_acad_career=&acad_group=&rqmnt_designtn=&instruction_mode=&keyword=&class_nbr=&acad_org=&enrl_stat=" + params.EnrollmentStatus + "&crse_attr=&crse_attr_value=&instructor_name=&session_code=&units=&page=" + strconv.Itoa(params.PageNumber)

	client := resty.New()
	resp, err := client.R().EnableTrace().Get(baseUrl)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func ValidateParams(params *CourseFunctionParams) (bool, error) {
	termValid, err := ValidateTerm(params.Term)

	if err != nil {
		return false, err
	}

	if !termValid {
		// format string to include the term
		return false, errors.New("Term [" + params.Term + "] is not valid")
	}

	if params.Subject == "" {
		return false, errors.New("subject cannot be empty")
	}

	params.EnrollmentStatus = strings.ToUpper(params.EnrollmentStatus)

	if params.EnrollmentStatus != "O" && params.EnrollmentStatus != "C" && params.EnrollmentStatus != "W" {
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

func ValidateTerm(term string) (bool, error) {
	termRegex := "2\\d\\d[147]"
	// test the regex
	return regexp.Match(termRegex, []byte(term))
}
