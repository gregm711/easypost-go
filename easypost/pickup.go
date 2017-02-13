package easypost

/*
Official API documentation available at:
https://www.easypost.com/docs/api.html#pickups
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

//Pickup is an easypost pickup
type Pickup struct {
	ID        string    `json:"id"`
	Object    string    `json:"object"`
	Mode      string    `json:"mode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Reference        string           `json:"reference"`
	Status           string           `json:"status"`
	MinDatetime      string           `json:"min_datetime"`
	MaxDatetime      string           `json:"max_datetime"`
	IsAccountAddress bool             `json:"is_account_address"`
	Instructions     string           `json:"instructions"`
	Messages         Message          `json:"messages"`
	Confirmation     string           `json:"confirmation"`
	Shipment         *Shipment        `json:"shipment"`
	Address          Address          `json:"address"`
	Batch            *Batch           `json:"batch"`
	CarrierAccounts  []CarrierAccount `json:"carrier_accounts"`
	Carrier          string           `json:"carrier"`
	Service          string           `json:"service"`
	PickupRates      []PickupRate     `json:"pickup_rates"`
}

//PickupRate is an easypost pickup rate
type PickupRate struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Mode      string `json:"mode"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	Service  string `json:"service"`
	Carrier  string `json:"carrier"`
	Rate     string `json:"rate"`
	Currency string `json:"currency"`
	PickupID string `json:"pickup_id"`
}

//Buy purchases a pickup
func (p *Pickup) Buy() error {
	if p.Carrier == "" {
		return errors.New("Carrier is missing")
	}
	if p.Service == "" {
		return errors.New("Service rate is missing")
	}

	obj, err := Request.do("POST", "pickup", fmt.Sprintf("%v/buy", p.ID), fmt.Sprintf("carrier=%v&service=%v", p.Carrier, p.Service))
	if err != nil {
		return errors.New("Failed to request EasyPost pickup purchase")
	}
	return json.Unmarshal(obj, &p)
}

//Cancel an EasyPost pickup
func (p *Pickup) Cancel() error {
	obj, err := Request.do("POST", "pickup", fmt.Sprintf("%s/cancel", p.ID), "")
	if err != nil {
		return errors.New("Failed to request EasyPost pickup cancellation")
	}
	return json.Unmarshal(obj, &p)
}

//Create an EasyPost pickup
func (p *Pickup) Create() error {
	obj, err := Request.do("POST", "pickup", "", p.getCreatePayload())
	if err != nil {
		return errors.New("Failed to request EasyPost pickup creation")
	}
	return json.Unmarshal(obj, &p)
}

//Get retrieves a pickup
func (p *Pickup) Get() error {
	obj, err := Request.do("GET", "pickup", p.ID, "")
	if err != nil {
		return errors.New("Failed to retrieve EasyPost pickup")
	}
	return json.Unmarshal(obj, &p)
}

//GetByReference retrieves a pickup by reference value
func (p *Pickup) GetByReference() error {
	obj, err := Request.do("GET", "pickup", p.Reference, "")
	if err != nil {
		return errors.New("Failed to retrieve EasyPost pickup by reference")
	}
	return json.Unmarshal(obj, &p)
}

func (p Pickup) getCreatePayload() string {
	var bodyString = ""
	bodyString = fmt.Sprintf("%s&%s", bodyString, p.Address.getPayload("pickup[address]"))
	if p.Shipment != nil {
		bodyString = fmt.Sprintf("%s&%s", bodyString, p.Shipment.getCreatePayload("pickup[shipment]"))
	}
	if p.Batch != nil {
		bodyString = fmt.Sprintf("%s&%s", bodyString, p.Batch.getCreatePayload("pickup[batch]"))
	}
	bodyString = fmt.Sprintf("%s&pickup[instructions]=%s", bodyString, p.Instructions)
	bodyString = fmt.Sprintf("%s&pickup[reference]=%s", bodyString, p.Reference)
	bodyString = fmt.Sprintf("%s&pickup[is_account_address]=%v", bodyString, p.IsAccountAddress)
	bodyString = fmt.Sprintf("%s&pickup[min_datetime]=%s", bodyString, p.MinDatetime)
	bodyString = fmt.Sprintf("%s&pickup[max_datetime]=%s", bodyString, p.MaxDatetime)

	return bodyString
}
