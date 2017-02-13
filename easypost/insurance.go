package easypost

/*
Official API documentation available at:
https://www.easypost.com/docs/api.html#insurance
*/

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Insurance struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Mode      string `json:"mode"`
	Reference string `json:"reference"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	Amount       string   `json:"amount"`
	Carrier      string   `json:"carrier"`
	Provider     string   `json:"provider"`
	ProviderID   string   `json:"provider_id"`
	ShipmentID   string   `json:"shipment_id"`
	TrackingCode string   `json:"tracking_code"`
	Status       string   `json:"status"`
	Tracker      Tracker  `json:"tracker"`
	ToAddress    Address  `json:"to_address"`
	FromAddress  Address  `json:"from_address"`
	Fee          Fee      `json:"fee"`
	Messages     []string `json:"messages"`
}

//Create an EasyPost insurance
func (i *Insurance) Create() error {
	obj, err := Request.do("POST", "insurance", "", i.getCreatePayload("insurance"))
	if err != nil {
		return errors.New("Failed to request EasyPost insurance creation")
	}
	return json.Unmarshal(obj, &i)
}

//Get Retrieves the insurance from EasyPost
func (i *Insurance) Get() error {
	obj, _ := Request.do("GET", "insurance", i.ID, "")
	return json.Unmarshal(obj, &i)
}

//GetAll Retrieves the insurance from EasyPost
func (i *Insurance) GetAll() ([]Insurance, error) {
	var insurances []Insurance
	obj, _ := Request.do("GET", "insurance", "", "")
	var err = json.Unmarshal(obj, &insurances)
	return insurances, err
}

func (i Insurance) getCreatePayload(prefix string) string {
	var bodyString = ""

	bodyString = fmt.Sprintf("%s&%s[reference]=%s", bodyString, prefix, i.Reference)
	bodyString = fmt.Sprintf("%s&%s[amount]=%s", bodyString, prefix, i.Amount)
	bodyString = fmt.Sprintf("%s&%s[carrier]=%s", bodyString, prefix, i.Carrier)
	bodyString = fmt.Sprintf("%s&%s[provider]=%s", bodyString, prefix, i.Provider)
	bodyString = fmt.Sprintf("%s&%s[provider_id]=%s", bodyString, prefix, i.ProviderID)
	bodyString = fmt.Sprintf("%s&%s[shipment_id]=%s", bodyString, prefix, i.ShipmentID)
	bodyString = fmt.Sprintf("%s&%s[tracking_code]=%s", bodyString, prefix, i.TrackingCode)
	bodyString = fmt.Sprintf("%s&%s", bodyString, i.ToAddress.getPayload(fmt.Sprintf("%s[to_address]", prefix)))
	bodyString = fmt.Sprintf("%s&%s[]=%s", bodyString, prefix, i.FromAddress.getPayload(fmt.Sprintf("%s[from_address]", prefix)))

	return bodyString
}
