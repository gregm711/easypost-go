package easypost

/**
Official API documentation available at:
https://www.easypost.com/docs/api.html#addresses
**/

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	DefaultNullPhone = "0000000000"
)

//Address is ah EasyPost object that defines a shipping address
type Address struct {
	ID        string    `json:"id"`
	Object    string    `json:"object"`
	Mode      string    `json:"mode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Street1         string        `json:"street1"`
	Street2         string        `json:"street2"`
	City            string        `json:"city"`
	State           string        `json:"state"`
	Zip             string        `json:"zip"`
	Country         string        `json:"country"`
	Residential     bool          `json:"residential"`
	CarrierFacility string        `json:"carrier_facility"`
	Name            string        `json:"name"`
	Company         string        `json:"company"`
	Phone           string        `json:"phone"`
	Email           string        `json:"email"`
	FederalTaxID    string        `json:"federal_tax_id"`
	StateTaxID      string        `json:"state_tax_id"`
	Verifications   Verifications `json:"verifications"`
	Verify          []string      `json:"verify"`
	VerifyStrict    []string      `json:"verify_strict"`
}

type Verifications struct {
	Zip4     Verification `json:"zip4"`
	Delivery Verification `json:"delivery"`
}

type Verification struct {
	Success bool                `json:"success"`
	Errors  []FieldError        `json:"errors"`
	Details VerificationDetails `json:"details"`
}

type VerificationDetails struct {
	Latitude  int `json:"latitude"`
	Longitude int `json:"longitude"`
}

//Create a new EasyPost address
func (a *Address) Create() error {
	obj, err := Request.do("POST", "address", "", a.getPayload("address"))
	if err != nil {
		return errors.New("Failed to request EasyPost Address creation")
	}
	return json.Unmarshal(obj, &a)
}

//Get Retrieves the address from EasyPost
func (a *Address) Get() error {
	obj, _ := Request.do("GET", "address", a.ID, "")
	return json.Unmarshal(obj, &a)
}

//getRequestPayload returns the payload to append to the EasyPost API request
func (a Address) getPayload(prefix string) string {
	payloadValues := reflect.ValueOf(a)
	bodyString := ""

	for i := 0; i < payloadValues.NumField(); i++ {
		fieldName := ""
		if payloadValues.Field(i).Interface() == nil || payloadValues.Field(i).Interface() == "" {
			continue
		}
		if payloadValues.Type().Field(i).Name == "Verify" && payloadValues.Field(i).Interface().(bool) {
			bodyString = fmt.Sprintf("%v&verify[]=%s", bodyString, payloadValues.Field(i).Interface())
			continue
		}
		if payloadValues.Type().Field(i).Name == "VerifyStrict" && payloadValues.Field(i).Interface().(bool) {
			bodyString = fmt.Sprintf("%v&verify_strict[]=%s", bodyString, payloadValues.Field(i).Interface())
			continue
		}
		if payloadValues.Type().Field(i).Name == "Residential" {
			continue
		}

		fieldName = strings.ToLower(payloadValues.Type().Field(i).Name)
		bodyString = fmt.Sprintf("%v&%v[%v]=%v", bodyString, prefix, fieldName, payloadValues.Field(i).Interface())
	}
	return bodyString
}
