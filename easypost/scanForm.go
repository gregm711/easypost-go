package easypost

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

//ScanForm is and EasyPost object and defines a form for a shipment
type ScanForm struct {
	ID        string    `json:"id"`
	Object    string    `json:"object"`
	Mode      string    `json:"mode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Status        string   `json:"status"`
	Message       string   `json:"message"`
	Address       Address  `json:"address"`
	TrackingCodes []string `json:"tracking_codes"`
	FormURL       string   `json:"form_url"`
	FormFileType  string   `json:"form_file_type"`
	BatchID       string   `json:"batch_id"`
}

//ScanFormList is the object for retrieving a list of scan forms
type ScanFormList struct {
	ScanForms     []ScanForm `json:"scan_forms"`
	StartDatetime string     `json:"start_datetime"`
	EndDatetime   string     `json:"end_datetime"`
	BeforeID      string     `json:"before_id"`
	AfterID       string     `json:"after_id"`
	PageSize      int        `json:"page_size"`

	HasMore bool `json:"has_more"`
}

func (sf *ScanForm) Get() error {
	obj, err := Request.do("GET", "scan_form", sf.ID, "")
	if err != nil {
		return errors.New("Failed to retrieve EasyPost scan form")
	}
	return json.Unmarshal(obj, &sf)
}

//Get retrieves a list of reports
func (sfl *ScanFormList) Get() error {
	var payload = ""
	if sfl.StartDatetime != "" {
		payload = fmt.Sprintf("%s&start_date=%s", payload, sfl.StartDatetime)
	}
	if sfl.EndDatetime != "" {
		payload = fmt.Sprintf("%s&end_date=%s", payload, sfl.EndDatetime)
	}
	if sfl.BeforeID != "" {
		payload = fmt.Sprintf("%s&before_id=%s", payload, sfl.BeforeID)
	}
	if sfl.AfterID != "" {
		payload = fmt.Sprintf("%s&after_id=%s", payload, sfl.AfterID)
	}
	if sfl.PageSize > 0 {
		payload = fmt.Sprintf("%s&page_size=%v", payload, sfl.PageSize)
	}

	obj, err := Request.do("GET", "report", "", payload)
	if err != nil {
		return errors.New("Failed to retrieve EasyPost report list")
	}
	return json.Unmarshal(obj, &sfl)
}
