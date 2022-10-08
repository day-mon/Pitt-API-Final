package controllers

import (
	models "awesomeProject/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"strings"
)

type LaundryController struct{}

var laundryApiCalls = map[string]string{
	"HICKORY":  "5813396",
	"BRIAR":    "581339005",
	"BUCKHORN": "5813393",
	"LLC":      "58133912",
	"OAK":      "5813391",
	"HAWTHORN": "5813397",
	"HEATHER":  "5813398",
	"HEMLOCK":  "5813392",
	"MAPLE":    "5813399",
	"WILLOW":   "58133912",
	"LARKSPUR": "58133911",
	"LAUREL":   "5813394",
	"CPAS":     "581339013",
}

func (c *LaundryController) GetByDormitory(context *gin.Context) {
	// get url / params
	dormitory := strings.ToUpper(context.Param("dormitory"))

	if laundryApiCalls[dormitory] == "" {
		context.JSON(400, gin.H{"error": "invalid dormitory"})
		return
	}

	// send request to Pitt API
	json, err := getLaundry(dormitory)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// convert json response to Laundry struct
	laundry, err := models.LaundryObject{}.AsLaundryObject(json, dormitory)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return response
	context.JSON(200, laundry)

}

func getLaundry(dormitory string) (string, error) {

	const baseUrl = "https://www.laundryview.com/api/currentRoomData?school_desc_key=4590&location="
	var dormitoryUrl = baseUrl + laundryApiCalls[dormitory]

	// send request to Pitt API
	client := resty.New()
	resp, err := client.R().Get(dormitoryUrl)

	if resp.Status() != "200 OK" {
		return "", errors.New("error getting laundry data")
	}

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
