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

	//EventTypeTrackerCreated is the event for a new created tracker
	EventTypeTrackerCreated = "tracker.created"
	//EventTypeTrackerUpdated is the event for an updated tracker
	EventTypeTrackerUpdated = "tracker.updated"
	//EventTypeBatchCreated is the event for a new created batch
	EventTypeBatchCreated = "batch.created"
	//EventTypeBatchUpdated is the event for an updated batch
	EventTypeBatchUpdated = "batch.updated"
	//EventTypeScanFormCreated is the event for a new created scan form
	EventTypeScanFormCreated = "scan_form.created"
	//EventTypeScanFormUpdated is the event for an updated scan form
	EventTypeScanFormUpdated = "scan_form.updated"
	//EventTypeInsurancePurchased is the event for a purchased insurance
	EventTypeInsurancePurchased = "insurance.purchased"
	//EventTypeInsuranceCancelled is the event for a cancelled insurance
	EventTypeInsuranceCancelled = "insurance.cancelled"
	//EventTypeRefundSuccessful is the event for a successful refund
	EventTypeRefundSuccessful = "refund.successful"
	//EventTypePaymentCreated is the event for a new created payment
	EventTypePaymentCreated = "payment.created"
	//EventTypePaymentCompleted is the event for a completed payment
	EventTypePaymentCompleted = "payment.completed"
	//EventTypePaymentFailed is the event for a failed payment
	EventTypePaymentFailed = "payment.failed"
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
