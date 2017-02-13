package easypost

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	//ReportTypeShipment is the repoprt type for shipments
	ReportTypeShipment = "shipment"
	//ReportTypePaymentLog is the repoprt type for payment logs
	ReportTypePaymentLog = "payment_log"
	//ReportTypeTracker is the repoprt type for trackers
	ReportTypeTracker = "tracker"
)

//Report is a report for EasyPost shipments, payment logs and trackers
type Report struct {
	ID              string `json:"id"`
	Object          string `json:"object"`
	Mode            string `json:"mode"`
	Status          string `json:"status"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	IncludeChildren bool   `json:"include_children"`
	URL             string `json:"url"`
	URLExpiresAt    string `json:"url_expires_at"`
	SendEmail       bool   `json:"send_email"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Type            string `json:"type"`
}

//ReportList is the object for retrieving a list of reports
type ReportList struct {
	Type      string `json:"type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	BeforeID  string `json:"before_id"`
	AfterID   string `json:"after_id"`
	PageSize  int    `json:"page_size"`

	Reports []Report `json:"reports"`
	HasMore bool     `json:"has_more"`
}

//Create creates a new report
func (r *Report) Create() error {
	var payload = ""
	//TODO
	obj, err := Request.do("POST", "report", r.Type, payload)
	if err != nil {
		return errors.New("Failed to request EasyPost report creation")
	}
	return json.Unmarshal(obj, &r)
}

//Get retrieves a list of reports
func (rl *ReportList) Get() error {
	var payload = ""
	if rl.StartDate != "" {
		payload = fmt.Sprintf("%s&start_date=%s", payload, rl.StartDate)
	}
	if rl.EndDate != "" {
		payload = fmt.Sprintf("%s&end_date=%s", payload, rl.EndDate)
	}
	if rl.BeforeID != "" {
		payload = fmt.Sprintf("%s&before_id=%s", payload, rl.BeforeID)
	}
	if rl.AfterID != "" {
		payload = fmt.Sprintf("%s&after_id=%s", payload, rl.AfterID)
	}
	if rl.PageSize > 0 {
		payload = fmt.Sprintf("%s&page_size=%v", payload, rl.PageSize)
	}

	obj, err := Request.do("GET", "report", rl.Type, payload)
	if err != nil {
		return errors.New("Failed to retrieve EasyPost report list")
	}
	return json.Unmarshal(obj, &rl)
}

//Get retrieves a report
func (r *Report) Get() error {
	obj, err := Request.do("GET", "report", fmt.Sprintf("%s/%s", r.Type, r.ID), "")
	if err != nil {
		return errors.New("Failed to retrieve EasyPost report")
	}
	return json.Unmarshal(obj, &r)
}
