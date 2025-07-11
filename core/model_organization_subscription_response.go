/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// checks if the OrganizationSubscriptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationSubscriptionResponse{}

// OrganizationSubscriptionResponse struct for OrganizationSubscriptionResponse
type OrganizationSubscriptionResponse struct {
	Active bool `json:"active"`
	CancelAtPeriodEnd bool `json:"cancelAtPeriodEnd"`
	CreatedAt string `json:"createdAt"`
	Currency string `json:"currency"`
	CurrentPeriodEnd string `json:"currentPeriodEnd"`
	Id string `json:"id"`
	Name string `json:"name"`
	OrganizationId string `json:"organizationId"`
	ProductId string `json:"productId"`
	ProductType ProductType `json:"productType"`
	Seats int32 `json:"seats"`
	SeatsTaken int32 `json:"seatsTaken"`
	Status string `json:"status"`
	// total amount in cents
	TotalAmount int32 `json:"totalAmount"`
	Trial bool `json:"trial"`
	TrialEnd string `json:"trialEnd"`
	UpdatedAt string `json:"updatedAt"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationSubscriptionResponse OrganizationSubscriptionResponse

// NewOrganizationSubscriptionResponse instantiates a new OrganizationSubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSubscriptionResponse(active bool, cancelAtPeriodEnd bool, createdAt string, currency string, currentPeriodEnd string, id string, name string, organizationId string, productId string, productType ProductType, seats int32, seatsTaken int32, status string, totalAmount int32, trial bool, trialEnd string, updatedAt string) *OrganizationSubscriptionResponse {
	this := OrganizationSubscriptionResponse{}
	this.Active = active
	this.CancelAtPeriodEnd = cancelAtPeriodEnd
	this.CreatedAt = createdAt
	this.Currency = currency
	this.CurrentPeriodEnd = currentPeriodEnd
	this.Id = id
	this.Name = name
	this.OrganizationId = organizationId
	this.ProductId = productId
	this.ProductType = productType
	this.Seats = seats
	this.SeatsTaken = seatsTaken
	this.Status = status
	this.TotalAmount = totalAmount
	this.Trial = trial
	this.TrialEnd = trialEnd
	this.UpdatedAt = updatedAt
	return &this
}

// NewOrganizationSubscriptionResponseWithDefaults instantiates a new OrganizationSubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSubscriptionResponseWithDefaults() *OrganizationSubscriptionResponse {
	this := OrganizationSubscriptionResponse{}
	return &this
}

// GetActive returns the Active field value
func (o *OrganizationSubscriptionResponse) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *OrganizationSubscriptionResponse) SetActive(v bool) {
	o.Active = v
}

// GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field value
func (o *OrganizationSubscriptionResponse) GetCancelAtPeriodEnd() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CancelAtPeriodEnd
}

// GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetCancelAtPeriodEndOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelAtPeriodEnd, true
}

// SetCancelAtPeriodEnd sets field value
func (o *OrganizationSubscriptionResponse) SetCancelAtPeriodEnd(v bool) {
	o.CancelAtPeriodEnd = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationSubscriptionResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationSubscriptionResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCurrency returns the Currency field value
func (o *OrganizationSubscriptionResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *OrganizationSubscriptionResponse) SetCurrency(v string) {
	o.Currency = v
}

// GetCurrentPeriodEnd returns the CurrentPeriodEnd field value
func (o *OrganizationSubscriptionResponse) GetCurrentPeriodEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentPeriodEnd
}

// GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetCurrentPeriodEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPeriodEnd, true
}

// SetCurrentPeriodEnd sets field value
func (o *OrganizationSubscriptionResponse) SetCurrentPeriodEnd(v string) {
	o.CurrentPeriodEnd = v
}

// GetId returns the Id field value
func (o *OrganizationSubscriptionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationSubscriptionResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OrganizationSubscriptionResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationSubscriptionResponse) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *OrganizationSubscriptionResponse) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *OrganizationSubscriptionResponse) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetProductId returns the ProductId field value
func (o *OrganizationSubscriptionResponse) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *OrganizationSubscriptionResponse) SetProductId(v string) {
	o.ProductId = v
}

// GetProductType returns the ProductType field value
func (o *OrganizationSubscriptionResponse) GetProductType() ProductType {
	if o == nil {
		var ret ProductType
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetProductTypeOk() (*ProductType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *OrganizationSubscriptionResponse) SetProductType(v ProductType) {
	o.ProductType = v
}

// GetSeats returns the Seats field value
func (o *OrganizationSubscriptionResponse) GetSeats() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Seats
}

// GetSeatsOk returns a tuple with the Seats field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetSeatsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seats, true
}

// SetSeats sets field value
func (o *OrganizationSubscriptionResponse) SetSeats(v int32) {
	o.Seats = v
}

// GetSeatsTaken returns the SeatsTaken field value
func (o *OrganizationSubscriptionResponse) GetSeatsTaken() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SeatsTaken
}

// GetSeatsTakenOk returns a tuple with the SeatsTaken field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetSeatsTakenOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeatsTaken, true
}

// SetSeatsTaken sets field value
func (o *OrganizationSubscriptionResponse) SetSeatsTaken(v int32) {
	o.SeatsTaken = v
}

// GetStatus returns the Status field value
func (o *OrganizationSubscriptionResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrganizationSubscriptionResponse) SetStatus(v string) {
	o.Status = v
}

// GetTotalAmount returns the TotalAmount field value
func (o *OrganizationSubscriptionResponse) GetTotalAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetTotalAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *OrganizationSubscriptionResponse) SetTotalAmount(v int32) {
	o.TotalAmount = v
}

// GetTrial returns the Trial field value
func (o *OrganizationSubscriptionResponse) GetTrial() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Trial
}

// GetTrialOk returns a tuple with the Trial field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetTrialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trial, true
}

// SetTrial sets field value
func (o *OrganizationSubscriptionResponse) SetTrial(v bool) {
	o.Trial = v
}

// GetTrialEnd returns the TrialEnd field value
func (o *OrganizationSubscriptionResponse) GetTrialEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrialEnd
}

// GetTrialEndOk returns a tuple with the TrialEnd field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetTrialEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrialEnd, true
}

// SetTrialEnd sets field value
func (o *OrganizationSubscriptionResponse) SetTrialEnd(v string) {
	o.TrialEnd = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *OrganizationSubscriptionResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationSubscriptionResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *OrganizationSubscriptionResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o OrganizationSubscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationSubscriptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["cancelAtPeriodEnd"] = o.CancelAtPeriodEnd
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["currency"] = o.Currency
	toSerialize["currentPeriodEnd"] = o.CurrentPeriodEnd
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["productId"] = o.ProductId
	toSerialize["productType"] = o.ProductType
	toSerialize["seats"] = o.Seats
	toSerialize["seatsTaken"] = o.SeatsTaken
	toSerialize["status"] = o.Status
	toSerialize["totalAmount"] = o.TotalAmount
	toSerialize["trial"] = o.Trial
	toSerialize["trialEnd"] = o.TrialEnd
	toSerialize["updatedAt"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationSubscriptionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active",
		"cancelAtPeriodEnd",
		"createdAt",
		"currency",
		"currentPeriodEnd",
		"id",
		"name",
		"organizationId",
		"productId",
		"productType",
		"seats",
		"seatsTaken",
		"status",
		"totalAmount",
		"trial",
		"trialEnd",
		"updatedAt",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrganizationSubscriptionResponse := _OrganizationSubscriptionResponse{}

	err = json.Unmarshal(data, &varOrganizationSubscriptionResponse)

	if err != nil {
		return err
	}

	*o = OrganizationSubscriptionResponse(varOrganizationSubscriptionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "cancelAtPeriodEnd")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "currentPeriodEnd")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "organizationId")
		delete(additionalProperties, "productId")
		delete(additionalProperties, "productType")
		delete(additionalProperties, "seats")
		delete(additionalProperties, "seatsTaken")
		delete(additionalProperties, "status")
		delete(additionalProperties, "totalAmount")
		delete(additionalProperties, "trial")
		delete(additionalProperties, "trialEnd")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationSubscriptionResponse struct {
	value *OrganizationSubscriptionResponse
	isSet bool
}

func (v NullableOrganizationSubscriptionResponse) Get() *OrganizationSubscriptionResponse {
	return v.value
}

func (v *NullableOrganizationSubscriptionResponse) Set(val *OrganizationSubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSubscriptionResponse(val *OrganizationSubscriptionResponse) *NullableOrganizationSubscriptionResponse {
	return &NullableOrganizationSubscriptionResponse{value: val, isSet: true}
}

func (v NullableOrganizationSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


