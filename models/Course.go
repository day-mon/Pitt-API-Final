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
	ClassEnrollInfo struct {
		IsRelated        bool `json:"is_related"`
		LastEnrlDtPassed bool `json:"last_enrl_dt_passed"`
	} `json:"class_enroll_info"`
	EnrollDates struct {
		AppDate     string `json:"app_date"`
		AppDateEnd  string `json:"app_date_end"`
		OpenDate    string `json:"open_date"`
		OpenDateEnd string `json:"open_date_end"`
	} `json:"enroll_dates"`
	Messages struct {
		NoMeetingInfo string `json:"noMeetingInfo"`
		ReserveInfo   string `json:"reserveInfo"`
		ShareLink     string `json:"shareLink"`
		ShareSocial   string `json:"shareSocial"`
	} `json:"messages"`
	SectionInfo struct {
		CatalogDescr struct {
			CrseCatalogDescription string `json:"crse_catalog_description"`
		} `json:"catalog_descr"`
		ClassAvailability struct {
			ClassCapacity       string `json:"class_capacity"`
			EnrollmentAvailable int    `json:"enrollment_available"`
			EnrollmentTotal     string `json:"enrollment_total"`
			WaitListCapacity    string `json:"wait_list_capacity"`
			WaitListTotal       string `json:"wait_list_total"`
		} `json:"class_availability"`
		ClassDetails struct {
			AcadCareer      string `json:"acad_career"`
			AcadCareerDescr string `json:"acad_career_descr"`
			Campus          string `json:"campus"`
			CampusCode      string `json:"campus_code"`
			CatalogNbr      string `json:"catalog_nbr"`
			ClassComponents string `json:"class_components"`
			ClassNumber     int    `json:"class_number"`
			ClassSection    string `json:"class_section"`
			Component       string `json:"component"`
			CourseID        string `json:"course_id"`
			CourseOfferNbr  int    `json:"course_offer_nbr"`
			CourseStatus    string `json:"course_status"`
			CourseTitle     string `json:"course_title"`
			GradingBasis    string `json:"grading_basis"`
			Institution     string `json:"institution"`
			InstructionMode string `json:"instruction_mode"`
			Location        string `json:"location"`
			SectionDescr    string `json:"section_descr"`
			Session         string `json:"session"`
			SessionCode     string `json:"session_code"`
			Status          string `json:"status"`
			Subject         string `json:"subject"`
			Topic           string `json:"topic"`
			Units           string `json:"units"`
		} `json:"class_details"`
		EnrollmentInformation struct {
			AddConsent         string `json:"add_consent"`
			ClassAttributes    string `json:"class_attributes"`
			DropConsent        string `json:"drop_consent"`
			EnrollRequirements string `json:"enroll_requirements"`
			RequirementDesig   string `json:"requirement_desig"`
		} `json:"enrollment_information"`
		IsCombined bool `json:"is_combined"`
		Materials  struct {
			TextbooksMessage       string `json:"textbooks_message"`
			TxbNone                string `json:"txb_none"`
			TxbSpecialInstructions string `json:"txb_special_instructions"`
			TxbStatus              string `json:"txb_status"`
		} `json:"materials"`
		Meetings []struct {
			BldgCd      string `json:"bldg_cd"`
			DateRange   string `json:"date_range"`
			EndDate     string `json:"end_date"`
			Instructors []struct {
				Email string `json:"email"`
				Name  string `json:"name"`
			} `json:"instructors"`
			MeetingTimeEnd   string `json:"meeting_time_end"`
			MeetingTimeStart string `json:"meeting_time_start"`
			MeetingTopic     string `json:"meeting_topic"`
			Meets            string `json:"meets"`
			ShowTopic        bool   `json:"show_topic"`
			StartDate        string `json:"start_date"`
			StndMtgPat       string `json:"stnd_mtg_pat"`
			Topic            string `json:"topic"`
		} `json:"meetings"`
		Notes struct {
			ClassNotes   string `json:"class_notes"`
			SubjectNotes string `json:"subject_notes"`
		} `json:"notes"`
		ReserveCaps   []interface{} `json:"reserve_caps"`
		ValidToEnroll string        `json:"valid_to_enroll"`
	} `json:"section_info"`
	ShowValidate   string `json:"show_validate"`
	ShowWaitlist   string `json:"show_waitlist"`
	SimilarClasses []struct {
		AcadCareer      string `json:"acad_career"`
		CatalogNbr      string `json:"catalog_nbr"`
		ClassNbr        int    `json:"class_nbr"`
		ClassSection    string `json:"class_section"`
		Descr           string `json:"descr"`
		MeetingPatterns []struct {
			ClassMtgNbr      int    `json:"class_mtg_nbr"`
			EndDt            string `json:"end_dt"`
			Fri              string `json:"fri"`
			MeetingTimeEnd   string `json:"meeting_time_end"`
			MeetingTimeStart string `json:"meeting_time_start"`
			Mon              string `json:"mon"`
			Sat              string `json:"sat"`
			StartDt          string `json:"start_dt"`
			StndMtgPat       string `json:"stnd_mtg_pat"`
			Sun              string `json:"sun"`
			Thurs            string `json:"thurs"`
			Tues             string `json:"tues"`
			Wed              string `json:"wed"`
		} `json:"meeting_patterns"`
		Subject string `json:"subject"`
		Topic   string `json:"topic"`
	} `json:"similar_classes"`
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

	convertedTime := convertDateToISO(strings.ReplaceAll(c.Meetings[0].StartDt, "/", "-"))

	time, err := time.Parse("2006-01-02", convertedTime)

	if err == nil {
		if isDaylightSavings(time) {
			course.Timezone = "EDT"
		} else {
			course.Timezone = "EST"
		}
	}
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

	startTimeObj, startErr := time.Parse("2006-01-02 15:04", startDate+" "+startTimeMilitary)
	endDateObj, endErr := time.Parse("2006-01-02 15:04", endDate+" "+endTimeMilitary)

	if isDaylightSavings(startTimeObj) {
		startTimeObj = startTimeObj.Add(time.Hour * 4)
		endDateObj = endDateObj.Add(time.Hour * 4)
		cm.Timezone = "EDT"
	} else {
		startTimeObj = startTimeObj.Add(time.Hour * 5)
		endDateObj = endDateObj.Add(time.Hour * 5)
		cm.Timezone = "EST"
	}

	if startErr != nil || endErr != nil {
		cm.StartDateAndStartTime = -1
		cm.EndDateAndEndTime = -1
		return
	}

	// convert Time object to Unix timestamp
	cm.StartDateAndStartTime = startTimeObj.Unix()
	cm.EndDateAndEndTime = endDateObj.Unix()

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
		startHour += 12
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
