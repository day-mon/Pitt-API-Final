package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type CourseController struct{}
type CourseFunctionParams struct {
	term             string
	subject          string
	campus           string
	enrollmentStatus string
	pageNum          int
}

func (c *CourseController) GetByTermAndSubject(context *gin.Context) {
	term := context.Param("term")

	subject := strings.ToUpper(context.Param("subject"))

	params := CourseFunctionParams{
		term:    term,
		subject: subject,
	}

	course, err := GetCourse(params)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": course,
	})

}

// generate a function that sends a request to the Pitt API and returns the response in JSON format
func GetCourse(params CourseFunctionParams) (string, error) {

	// validate the params
	_, err := ValidateParams(params)
	if err != nil {
		return "", err
	}

	baseUrl := "https://prd.ps.pitt.edu/psc/pitcsprd/EMPLOYEE/SA/s/WEBLIB_HCX_CM.H_CLASS_SEARCH.FieldFormula.IScript_ClassSearch?institution=UPITT&term=" + params.term + "&date_from=&date_thru=&subject=" + params.subject + "&subject_like=&catalog_nbr=&time_range=&days=&campus=" + params.campus + "&location=&x_acad_career=&acad_group=&rqmnt_designtn=&instruction_mode=&keyword=&class_nbr=&acad_org=&enrl_stat=" + params.enrollmentStatus + "&crse_attr=&crse_attr_value=&instructor_name=&session_code=&units=&page=" + strconv.Itoa(params.pageNum)

	client := resty.New()
	resp, err := client.R().EnableTrace().Get(baseUrl)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func ValidateParams(params CourseFunctionParams) (bool, error) {
	termValid, err := ValidateTerm(params.term)

	if err != nil {
		return false, err
	}

	if !termValid {
		return false, errors.New("term is not valid")
	}

	params.enrollmentStatus = strings.ToUpper(params.enrollmentStatus)

	if params.enrollmentStatus != "O" && params.enrollmentStatus != "C" && params.enrollmentStatus != "W" {
		return false, errors.New("enrollmentStatus must be O, C, or W")
	}

	params.campus = strings.ToUpper(params.campus)

	return true, nil
}

func ValidateTerm(term string) (bool, error) {
	termRegex := "2\\d\\d[147]"
	// test the regex
	return regexp.Match(termRegex, []byte(term))
}
