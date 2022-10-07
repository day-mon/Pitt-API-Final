package models

import "encoding/json"

type LaundryMachine struct {
	ApplianceDesc string  `json:"appliance_desc"`
	Type          string  `json:"type"`
	Percentage    float32 `json:"percentage"`
	TimeLeftLite  string  `json:"time_left_lite"`
	StatusToggle  int     `json:"status_toggle"`
}

type LaundryModel struct {
	Objects []LaundryMachine `json:"objects"`
}

type LaundryObject struct {
	Appliance     string `json:"appliance"`
	Type          string `json:"type"`
	IsWorking     bool   `json:"isWorking"`
	TimeRemaining string `json:"timeRemaining"`
	IsInUse       bool   `json:"isInUse"`
	Location      string `json:"location"`
}

func (l LaundryObject) AsLaundryObject(jsonStr string, location string) ([]LaundryObject, error) {
	var model LaundryModel
	err := json.Unmarshal([]byte(jsonStr), &model)

	if err != nil {
		return nil, err
	}

	var objects []LaundryObject
	for _, machine := range model.Objects {
		objects = append(objects, LaundryObject{
			Appliance:     machine.ApplianceDesc,
			Type:          machine.Type,
			IsWorking:     machine.Percentage <= 5,
			TimeRemaining: machine.TimeLeftLite,
			IsInUse:       machine.StatusToggle > 0 && machine.Percentage <= 5,
			Location:      location,
		})
	}

	return objects, nil
}
