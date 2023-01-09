package models

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var daysOFWeek = map[string]string{
	"Mo":  "Monday",
	"Tu":  "Tuesday",
	"We":  "Wednesday",
	"Th":  "Thursday",
	"Thr": "Thursday",
	"Fr":  "Friday",
	"Fri": "Friday",
}

type CourseResponse struct {
	EnrollDates     EnrollDates     `json:"enroll_dates"`
	ShowValidate    string          `json:"show_validate"`
	ShowWaitlist    string          `json:"show_waitlist"`
	SectionInfo     SectionInfo     `json:"section_info"`
	ClassEnrollInfo ClassEnrollInfo `json:"class_enroll_info"`
	AdditionalLinks []interface{}   `json:"additionalLinks"`
	Cfg             Cfg             `json:"cfg"`
	Messages        Messages        `json:"messages"`
}
type EnrollDates struct {
	AppDate     string `json:"app_date"`
	OpenDate    string `json:"open_date"`
	AppDateEnd  string `json:"app_date_end"`
	OpenDateEnd string `json:"open_date_end"`
}
type ClassDetails struct {
	Institution     string `json:"institution"`
	Subject         string `json:"subject"`
	CatalogNbr      string `json:"catalog_nbr"`
	Status          string `json:"status"`
	ClassNumber     int    `json:"class_number"`
	Component       string `json:"component"`
	CourseOfferNbr  int    `json:"course_offer_nbr"`
	Session         string `json:"session"`
	SessionCode     string `json:"session_code"`
	ClassSection    string `json:"class_section"`
	SectionDescr    string `json:"section_descr"`
	Units           string `json:"units"`
	AcadCareer      string `json:"acad_career"`
	AcadCareerDescr string `json:"acad_career_descr"`
	CourseID        string `json:"course_id"`
	CourseTitle     string `json:"course_title"`
	CourseStatus    string `json:"course_status"`
	InstructionMode string `json:"instruction_mode"`
	GradingBasis    string `json:"grading_basis"`
	Campus          string `json:"campus"`
	CampusCode      string `json:"campus_code"`
	Location        string `json:"location"`
	Topic           string `json:"topic"`
	ClassComponents string `json:"class_components"`
}
type Instructors struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Meetings struct {
	Meets            string        `json:"meets"`
	StndMtgPat       string        `json:"stnd_mtg_pat"`
	MeetingTimeStart string        `json:"meeting_time_start"`
	MeetingTimeEnd   string        `json:"meeting_time_end"`
	BldgCd           string        `json:"bldg_cd"`
	MeetingTopic     string        `json:"meeting_topic"`
	Instructors      []Instructors `json:"instructors"`
	StartDate        string        `json:"start_date"`
	EndDate          string        `json:"end_date"`
	Topic            string        `json:"topic"`
	ShowTopic        bool          `json:"show_topic"`
	DateRange        string        `json:"date_range"`
}
type EnrollmentInformation struct {
	AddConsent         string `json:"add_consent"`
	DropConsent        string `json:"drop_consent"`
	EnrollRequirements string `json:"enroll_requirements"`
	RequirementDesig   string `json:"requirement_desig"`
	ClassAttributes    string `json:"class_attributes"`
}
type ClassAvailability struct {
	ClassCapacity       string `json:"class_capacity"`
	EnrollmentTotal     string `json:"enrollment_total"`
	EnrollmentAvailable int    `json:"enrollment_available"`
	WaitListCapacity    string `json:"wait_list_capacity"`
	WaitListTotal       string `json:"wait_list_total"`
}
type Notes struct {
	ClassNotes   string `json:"class_notes"`
	SubjectNotes string `json:"subject_notes"`
}
type CatalogDescr struct {
	CrseCatalogDescription string `json:"crse_catalog_description"`
}
type Materials struct {
	TxbNone                string `json:"txb_none"`
	TxbStatus              string `json:"txb_status"`
	TxbSpecialInstructions string `json:"txb_special_instructions"`
	TextbooksMessage       string `json:"textbooks_message"`
}
type SectionInfo struct {
	ClassDetails          ClassDetails          `json:"class_details"`
	Meetings              []Meetings            `json:"meetings"`
	EnrollmentInformation EnrollmentInformation `json:"enrollment_information"`
	ClassAvailability     ClassAvailability     `json:"class_availability"`
	ReserveCaps           []interface{}         `json:"reserve_caps"`
	IsCombined            bool                  `json:"is_combined"`
	Notes                 Notes                 `json:"notes"`
	CatalogDescr          CatalogDescr          `json:"catalog_descr"`
	Materials             Materials             `json:"materials"`
	ValidToEnroll         string                `json:"valid_to_enroll"`
}
type ClassEnrollInfo struct {
	LastEnrlDtPassed bool `json:"last_enrl_dt_passed"`
	IsRelated        bool `json:"is_related"`
}
type Cfg struct {
	IsRelated             bool `json:"is_related"`
	ShowCrseID            bool `json:"show_crse_id"`
	ShowCrseOfferNbr      bool `json:"show_crse_offer_nbr"`
	ShowCampus            bool `json:"show_campus"`
	ShowLocation          bool `json:"show_location"`
	ShowConsentToAdd      bool `json:"show_consent_to_add"`
	ShowConsentToDrop     bool `json:"show_consent_to_drop"`
	ShowEnrollReq         bool `json:"show_enroll_req"`
	ShowReqDesig          bool `json:"show_req_desig"`
	ShowClassAttributes   bool `json:"show_class_attributes"`
	ShowClassAvailability bool `json:"show_class_availability"`
	ShowCombined          bool `json:"show_combined"`
	ShowClassNotes        bool `json:"show_class_notes"`
	ShowCatalogDescr      bool `json:"show_catalog_descr"`
	ShowTextbookInfo      bool `json:"show_textbook_info"`
	ShowCommonAttributes  bool `json:"show_common_attributes"`
	CanAddToPlanner       bool `json:"can_add_to_planner"`
	ShowEnroll            bool `json:"show_enroll"`
	CanAddToCart          bool `json:"can_add_to_cart"`
	CanEnrollClass        bool `json:"can_enroll_class"`
	CanValidateClass      bool `json:"can_validate_class"`
	CanEditClass          bool `json:"can_edit_class"`
	CanDeleteClass        bool `json:"can_delete_class"`
	ShowFriendSuggest     bool `json:"show_friend_suggest"`
	ShowBookstore         bool `json:"show_bookstore"`
	ShowShare             bool `json:"show_share"`
	ShowWaitList          bool `json:"show_wait_list"`
	ShowInstructionMode   bool `json:"show_instruction_mode"`
	ShowTopic             bool `json:"show_topic"`
	ShowAddToWishList     bool `json:"show_add_to_wish_list"`
	WishListEnabled       bool `json:"wish_list_enabled"`
	ShowActions           bool `json:"show_actions"`
}
type Messages struct {
	ShareLink     string `json:"shareLink"`
	ShareSocial   string `json:"shareSocial"`
	ReserveInfo   string `json:"reserveInfo"`
	NoMeetingInfo string `json:"noMeetingInfo"`
}
type CourseInfoResponse struct {
	Index                int    `json:"index"`
	CrseID               string `json:"crse_id"`
	CrseOfferNbr         int    `json:"crse_offer_nbr"`
	Strm                 string `json:"strm"`
	SessionCode          string `json:"session_code"`
	SessionDescr         string `json:"session_descr"`
	ClassSection         string `json:"class_section"`
	Location             string `json:"location"`
	LocationDescr        string `json:"location_descr"`
	StartDt              string `json:"start_dt"`
	EndDt                string `json:"end_dt"`
	ClassStat            string `json:"class_stat"`
	Campus               string `json:"campus"`
	CampusDescr          string `json:"campus_descr"`
	ClassNbr             int    `json:"class_nbr"`
	AcadCareer           string `json:"acad_career"`
	AcadCareerDescr      string `json:"acad_career_descr"`
	Component            string `json:"component"`
	Subject              string `json:"subject"`
	SubjectDescr         string `json:"subject_descr"`
	CatalogNbr           string `json:"catalog_nbr"`
	ClassType            string `json:"class_type"`
	SchedulePrint        string `json:"schedule_print"`
	AcadGroup            string `json:"acad_group"`
	InstructionMode      string `json:"instruction_mode"`
	InstructionModeDescr string `json:"instruction_mode_descr"`
	AcadOrg              string `json:"acad_org"`
	WaitTot              int    `json:"wait_tot"`
	WaitCap              int    `json:"wait_cap"`
	ClassCapacity        int    `json:"class_capacity"`
	EnrollmentTotal      int    `json:"enrollment_total"`
	EnrollmentAvailable  int    `json:"enrollment_available"`
	Descr                string `json:"descr"`
	RqmntDesigntn        string `json:"rqmnt_designtn"`
	Units                string `json:"units"`
	CombinedSection      string `json:"combined_section"`
	EnrlStat             string `json:"enrl_stat"`
	EnrlStatDescr        string `json:"enrl_stat_descr"`
	Topic                string `json:"topic"`
	Instructors          []struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"instructors"`
	SectionType string `json:"section_type"`
	Meetings    []struct {
		Days       string `json:"days"`
		StartTime  string `json:"start_time"`
		EndTime    string `json:"end_time"`
		StartDt    string `json:"start_dt"`
		EndDt      string `json:"end_dt"`
		Instructor string `json:"instructor"`
	} `json:"meetings"`
	CrseAttr      string        `json:"crse_attr"`
	CrseAttrValue string        `json:"crse_attr_value"`
	ReserveCaps   []interface{} `json:"reserve_caps"`
}

// convert CourseModel to go struct
type CourseModel struct {
	Name                   string   `json:"name"`
	Identifier             string   `json:"identifier"`
	Session                string   `json:"session"`
	ClassNumber            int      `json:"classNumber"`
	Career                 string   `json:"career"`
	StartDateAndStartTime  int64    `json:"startDateAndStartTime"`
	EndDateAndEndTime      int64    `json:"endDateAndEndTime"`
	StartDate              string   `json:"startDate"`
	EndDate                string   `json:"endDate"`
	StartTime              string   `json:"startTime"`
	EndTime                string   `json:"endTime"`
	Timezone               string   `json:"timezone"`
	Units                  string   `json:"units"`
	DropConsent            string   `json:"dropConsent"`
	Grading                string   `json:"grading"`
	Description            string   `json:"description"`
	EnrollmentRequirements string   `json:"enrollmentRequirements"`
	ClassAttributes        string   `json:"classAttributes"`
	Topic                  string   `json:"topic"`
	HasTopic               bool     `json:"hasTopic"`
	Instructors            []string `json:"instructors"`

	// class details
	MeetingDays []string `json:"meetingDays"`
	Campus      string   `json:"campus"`
	Building    string   `json:"building"`
	Location    string   `json:"location"`
	Components  string   `json:"components"`

	// class availability
	Status           string `json:"status"`
	WaitListTotal    int    `json:"waitListTotal"`
	WaitListCapacity int    `json:"waitListCapacity"`
	SeatsOpen        int    `json:"seatsOpen"`
	SeatsTaken       int    `json:"seatsTaken"`
	ClassCapacity    int    `json:"classCapacity"`
}

func ConvertToCourseModelFromCourseInfo(c CourseInfoResponse) CourseModel {
	var course CourseModel
	course.Location = c.LocationDescr
	course.StartTime = c.Meetings[0].StartTime
	course.EndTime = c.Meetings[0].EndTime
	course.StartDate = c.StartDt
	course.EndDate = c.EndDt
	course.Timezone = "America/New_York"
	course.Units = c.Units
	course.DropConsent = c.RqmntDesigntn
	course.Name = c.Descr
	course.Identifier = c.Subject + " " + c.CatalogNbr
	course.Session = c.SessionDescr
	course.ClassAttributes = c.CrseAttr
	course.Topic = c.Topic
	course.HasTopic = len(c.Topic) > 0
	for _, i := range c.Instructors {
		course.Instructors = append(course.Instructors, i.Name)
	}

	if len(c.Meetings) > 0 {
		re := regexp.MustCompile(`([A-Z][a-z]*)`)
		splitDays := re.FindAllString(c.Meetings[0].Days, -1)

		for _, day := range splitDays {
			course.MeetingDays = append(course.MeetingDays, daysOFWeek[day])
		}
	}

	course.Campus = c.CampusDescr
	course.Career = c.AcadCareerDescr
	course.Components = c.Component
	course.ClassNumber = c.ClassNbr
	course.Status = c.EnrlStatDescr
	course.WaitListTotal = c.WaitTot
	course.WaitListCapacity = c.WaitCap
	course.SeatsOpen = c.EnrollmentAvailable
	course.SeatsTaken = c.EnrollmentTotal
	course.ClassCapacity = c.ClassCapacity
	return course
}

func ConvertToCourseModelFromCourseResponse(course CourseResponse) CourseModel {
	// declare a new CourseModel
	var courseModel CourseModel

	// assign values to CourseModel
	courseModel.Name = course.SectionInfo.ClassDetails.CourseTitle
	courseModel.Identifier = course.SectionInfo.ClassDetails.SectionDescr
	courseModel.Session = course.SectionInfo.ClassDetails.Session
	courseModel.ClassNumber = course.SectionInfo.ClassDetails.ClassNumber
	courseModel.Career = course.SectionInfo.ClassDetails.AcadCareerDescr
	courseModel.Units = course.SectionInfo.ClassDetails.Units
	courseModel.Location = course.SectionInfo.ClassDetails.Location
	courseModel.Grading = course.SectionInfo.ClassDetails.GradingBasis
	courseModel.DropConsent = course.SectionInfo.EnrollmentInformation.DropConsent
	courseModel.Description = course.SectionInfo.CatalogDescr.CrseCatalogDescription
	courseModel.EnrollmentRequirements = course.SectionInfo.EnrollmentInformation.EnrollRequirements
	if len(course.SectionInfo.Meetings) != 0 && course.SectionInfo.Meetings[0].ShowTopic {
		courseModel.Topic = course.SectionInfo.Meetings[0].Topic
	}
	courseModel.Components = course.SectionInfo.ClassDetails.Component
	courseModel.Campus = course.SectionInfo.ClassDetails.Campus
	courseModel.Status = course.SectionInfo.ClassDetails.Status
	courseModel.Topic = course.SectionInfo.ClassDetails.Topic

	// class availability will fix later lol
	waitListTotal, err := strconv.Atoi(course.SectionInfo.ClassAvailability.WaitListTotal)
	if err != nil {
		waitListTotal = -1
	} else {
		courseModel.WaitListTotal = waitListTotal
	}
	waitListCapacity, err := strconv.Atoi(course.SectionInfo.ClassAvailability.WaitListCapacity)
	if err != nil {
		waitListCapacity = -1
	} else {
		courseModel.WaitListCapacity = waitListCapacity
	}
	courseModel.SeatsOpen = course.SectionInfo.ClassAvailability.EnrollmentAvailable
	classCapacity, err := strconv.Atoi(course.SectionInfo.ClassAvailability.ClassCapacity)
	if err != nil {
		classCapacity = -1
	} else {
		courseModel.ClassCapacity = classCapacity
	}
	seatsTaken, err := strconv.Atoi(course.SectionInfo.ClassAvailability.EnrollmentTotal)
	if err != nil {
		seatsTaken = -1
	} else {
		courseModel.SeatsTaken = seatsTaken
	}

	if len(course.SectionInfo.Meetings) != 0 {
		days := course.SectionInfo.Meetings[0].StndMtgPat

		re := regexp.MustCompile(`([A-Z][a-z]*)`)
		splitDays := re.FindAllString(days, -1)

		for _, day := range splitDays {
			courseModel.MeetingDays = append(courseModel.MeetingDays, daysOFWeek[day])
		}

		courseModel.Building = course.SectionInfo.Meetings[0].BldgCd

		for _, instructor := range course.SectionInfo.Meetings[0].Instructors {
			courseModel.Instructors = append(courseModel.Instructors, instructor.Name)
		}
		courseModel.HasTopic = course.SectionInfo.Meetings[0].ShowTopic
		courseModel.Topic = course.SectionInfo.Meetings[0].Topic

		courseModel.StartTime = course.SectionInfo.Meetings[0].MeetingTimeStart
		courseModel.EndTime = course.SectionInfo.Meetings[0].MeetingTimeEnd

		courseModel.StartDate = course.SectionInfo.Meetings[0].StartDate
		courseModel.EndDate = course.SectionInfo.Meetings[0].EndDate

		setCourseTimes(course, &courseModel)

		return courseModel
	}
	courseModel.HasTopic = false
	return courseModel
}

func setCourseTimes(c CourseResponse, cm *CourseModel) {

	// get start and end times
	startTime := c.SectionInfo.Meetings[0].MeetingTimeStart
	endTime := c.SectionInfo.Meetings[0].MeetingTimeEnd

	// get start and end dates
	startDate := strings.ReplaceAll(c.SectionInfo.Meetings[0].StartDate, "/", "-")
	endDate := strings.ReplaceAll(c.SectionInfo.Meetings[0].EndDate, "/", "-")

	// get start and end times in military time
	startTimeMilitary, startErr := convertTimeToMilitary(startTime)
	endTimeMilitary, endErr := convertTimeToMilitary(endTime)

	if startErr != nil || endErr != nil {
		cm.StartDateAndStartTime = -1
		cm.EndDateAndEndTime = -1
		return
	}

	startDate = convertDateToISO(startDate)
	endDate = convertDateToISO(endDate)

	startTimeObj, startErr := time.Parse("2006-01-02 15:04 MST", startDate+" "+startTimeMilitary+" EST")
	endDateObj, endErr := time.Parse("2006-01-02 15:04 MST", endDate+" "+endTimeMilitary+" EST")

	startHour := startTimeObj.Hour()
	endHour := endDateObj.Hour()

	startMinute := startTimeObj.Minute()
	endMinute := endDateObj.Minute()

	// go est time zone
	loc, _ := time.LoadLocation("America/New_York")

	// set start and end times in startTimeObj and endDateObj
	startTimeObj = time.Date(startTimeObj.Year(), startTimeObj.Month(), startTimeObj.Day(), startHour, startMinute, 0, 0, loc)
	endDateObj = time.Date(endDateObj.Year(), endDateObj.Month(), endDateObj.Day(), endHour, endMinute, 0, 0, loc)

	if startErr != nil || endErr != nil {
		cm.StartDateAndStartTime = -1
		cm.EndDateAndEndTime = -1
		return
	}

	// convert Time object to Unix timestamp
	cm.StartDateAndStartTime = startTimeObj.UnixMilli()
	cm.EndDateAndEndTime = endDateObj.UnixMilli()
	cm.Timezone = "America/New_York"

}

func isDaylightSavings(date time.Time) bool {
	month := date.Month()
	if month >= time.March && month <= time.November {
		return true
	}
	return false
}

func convertDateToISO(date string) string {
	splitDate := strings.Split(date, "-")
	year := splitDate[2]
	month := splitDate[0]
	day := splitDate[1]
	return year + "-" + month + "-" + day
}

func convertTimeToMilitary(time string) (string, error) {
	// convert time to military time
	if time == "TBA" {
		return "", errors.New("time is TBA")
	}

	if strings.Contains(time, "PM") {
		time = strings.Replace(time, "PM", "", -1)
		time = strings.TrimSpace(time)
		startTimeSplit := strings.Split(time, ":")
		startHour, _ := strconv.Atoi(startTimeSplit[0])
		if startHour != 12 {
			startHour += 12
		}
		time = strconv.Itoa(startHour) + ":" + startTimeSplit[1]
		return time, nil
	} else if strings.Contains(time, "AM") {
		time = strings.Replace(time, "AM", "", -1)
		time = strings.TrimSpace(time)
		startTimeSplit := strings.Split(time, ":")
		startHour, _ := strconv.Atoi(startTimeSplit[0])
		time = strconv.Itoa(startHour) + ":" + startTimeSplit[1]
		return time, nil
	}
	return "", errors.New("time is not in correct format")
}
