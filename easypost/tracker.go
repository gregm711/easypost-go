package easypost

/*
Official API documentation available at:
https://www.easypost.com/docs/api.html#trackers
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

//these are the possible tracker statuses
const (
	TrackerStatusUnknown            = "unknown"
	TrackerStatusPreTransit         = "pre_transit"
	TrackerStatusInTransit          = "in_transit"
	TrackerStatusOutForDelivery     = "out_for_delivery"
	TrackerStatusDelivered          = "delivered"
	TrackerStatusAvailableForPickup = "available_for_pickup"
	TrackerStatusReturnToSender     = "return_to_sender"
	TrackerStatusFailure            = "failure"
	TrackerStatusCancelled          = "cancelled"
	TrackerStatusError              = "error"
)

//CarrierUPSTrackingURL is the default tracking base URL for UPS
const CarrierUPSTrackingURL = "https://wwwapps.ups.com/tracking/tracking.cgi?tracknum="

//CarrierUSPSTrackingURL is the default tracking base URL for USPS
const CarrierUSPSTrackingURL = "https://tools.usps.com/go/TrackConfirmAction?qtc_tLabels1="

//Tracker is an EasyPost object that defines a shipping tracker
type Tracker struct {
	ID        string    `json:"id"`
	Object    string    `json:"object"`
	Mode      string    `json:"mode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	TrackingCode    string            `json:"tracking_code"`
	Status          string            `json:"status"`
	NumShipments    int               `json:"num_shipments"`
	Reference       string            `json:"reference"`
	ScanForm        string            `json:"scan_form"`
	Shipments       []Shipment        `json:"shipments"`
	Pickup          string            `json:"pickup"`
	LabelURL        string            `json:"label_url"`
	SignedBy        string            `json:"signed_by"`
	Weight          float64           `json:"weight"`
	EstDeliveryDate *time.Time        `json:"est_delivery_date"`
	ShipmentID      string            `json:"shipment_id"`
	Carrier         string            `json:"carrier"`
	TrackingDetails []TrackingDetails `json:"tracking_details"`
	CarrierDetail   Carrier           `json:"carrier_detail"`
	Fees            []Fee             `json:"fees"`
}

//TrackingDetails is an EasyPost object that defines the details for a shipping tracker
type TrackingDetails struct {
	Object           string    `json:"object"`
	Message          string    `json:"message"`
	Status           string    `json:"status"`
	Datetime         time.Time `json:"datetime"`
	Source           string    `json:"source"`
	TrackingLocation Location  `json:"tracking_location"`
}

//Fee is an EasyPost object that defines the fee details for a shipping tracker
type Fee struct {
	Object   string `json:"object"`
	Type     string `json:"type"`
	Amount   string `json:"amount"`
	Charged  bool   `json:"charged"`
	Refunded bool   `json:"refunded"`
}

//Carrier is a carrier
type Carrier struct {
	Object               string     `json:"object"`
	Service              string     `json:"service"`
	ContainerType        string     `json:"container_type"`
	EstDeliveryDateLocal *time.Time `json:"est_delivery_date_local"`
	EstDeliveryTimeLocal *time.Time `json:"est_delivery_time_local"`
}

//Location is a tracker location
type Location struct {
	Object  string `json:"object"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
}

//NewTracker returns a new instance of Tracker
func NewTracker(id string, createdAt, updatedAt time.Time) Tracker {
	return Tracker{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Object:    "Tracker",
	}
}

//TrackerList is the object required to request a list of trackers
type TrackerList struct {
	BeforeID      string `json:"before_id"`
	AfterID       string `json:"after_id"`
	StartDatetime string `json:"start_datetime"`
	EndDatetime   string `json:"end_datetime"`
	PageSize      int    `json:"page_size"`
	TrackingCode  string `json:"tracking_code"`
	Carrier       string `json:"carrier"`

	Trackers []Tracker `json:"trackers"`
	HasMore  bool      `json:"has_more"`
}

//Create creates a new tracker on Easypost
func (t *Tracker) Create() error {
	obj, err := Request.do("POST", "tracker", "", t.getCreatePayload("tracker"))
	if err != nil {
		return errors.New("Failed to request EasyPost tracker creation")
	}
	return json.Unmarshal(obj, &t)
}

//Get retrieves a tracker
func (t *Tracker) Get() error {
	obj, _ := Request.do("GET", "tracker", t.ID, "")
	return json.Unmarshal(obj, &t)
}

//Get retrieve the list of trackers based on the given query
func (tl *TrackerList) Get() error {
	var payload = ""
	if tl.BeforeID != "" {
		payload = fmt.Sprintf("%s&before_id=%s", payload, tl.BeforeID)
	}
	if tl.AfterID != "" {
		payload = fmt.Sprintf("%s&after_id=%s", payload, tl.AfterID)
	}
	if tl.StartDatetime != "" {
		payload = fmt.Sprintf("%s&start_datetime=%s", payload, tl.StartDatetime)
	}
	if tl.PageSize > 0 {
		payload = fmt.Sprintf("%s&page_size=%v", payload, tl.PageSize)
	}
	if tl.TrackingCode != "" {
		payload = fmt.Sprintf("%s&tracking_code=%s", payload, tl.TrackingCode)
	}
	if tl.Carrier != "" {
		payload = fmt.Sprintf("%s&carrier=%s", payload, tl.Carrier)
	}
	obj, _ := Request.do("GET", "tracker", "", payload)
	return json.Unmarshal(obj, &tl)
}

func (t Tracker) getCreatePayload(prefix string) string {
	var bodyString = ""
	bodyString = fmt.Sprintf("%s&%s[tracking_code]=%s", bodyString, prefix, t.TrackingCode)
	bodyString = fmt.Sprintf("%s&%s[carrier]=%s", bodyString, prefix, t.Carrier)

	return bodyString
}
