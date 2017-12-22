package easypost

/*
Official API documentation available at:
https://www.easypost.com/docs/api.html#events
*/

import "time"

const (
	//EventObjectTracker is the type 'Tracker' for an event
	EventObjectTracker = "Tracker"
	//EventObjectScanForm is the type 'ScanForm' for an event
	EventObjectScanForm = "ScanForm"
	//EventObjectBatch is the type 'Batch' for an event
	EventObjectBatch = "Batch"
)

//Event is created by changes in objects created via the API
type Event struct {
	ID        string    `json:"id"`
	Object    string    `json:"object"`
	Mode      string    `json:"mode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Description        string     `json:"description"`
	PreviousAttributes Attributes `json:"previous_attributes"`
	PendingUrls        []string   `json:"pending_urls"`
	CompletedUrls      []string   `json:"completed_urls"`
}

//GenericEvent is a generic webhook event
type GenericEvent struct {
	Event
	Result EasypostObject `json:"result"`
}

//BatchEvent is the event received from a batch webhook
type BatchEvent struct {
	Event
	Result Batch `json:"result"`
}

//ScanFormEvent is the event received from a scan form webhook
type ScanFormEvent struct {
	Event
	Result ScanForm `json:"result"`
}

//TrackerEvent is the event received from a tracker webhook
type TrackerEvent struct {
	Event
	Result Tracker `json:"result"`
}

//Attributes are attributes
type Attributes struct {
	Status string `json:"status"`
}

//NewGenericEvent returns a new istance of Event
func NewGenericEvent(id string, createdAt, updatedAt time.Time) GenericEvent {
	return GenericEvent{
		Event: Event{
			ID:        id,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
			Object:    "Event"},
	}
}

//NewScanFormEvent returns a new istance of Event
func NewScanFormEvent(id string, createdAt, updatedAt time.Time) ScanFormEvent {
	return ScanFormEvent{
		Event: Event{
			ID:        id,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
			Object:    "Event"},
	}
}

//NewTrackerEvent returns a new istance of Event
func NewTrackerEvent(id string, createdAt, updatedAt time.Time) TrackerEvent {
	return TrackerEvent{
		Event: Event{
			ID:        id,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
			Object:    "Event"},
	}
}
