/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// checks if the OrganizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationResponse{}

// OrganizationResponse struct for OrganizationResponse
type OrganizationResponse struct {
	CreatedAt string `json:"createdAt"`
	Id string `json:"id"`
	Name string `json:"name"`
	Owner string `json:"owner"`
	UpdatedAt string `json:"updatedAt"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationResponse OrganizationResponse

// NewOrganizationResponse instantiates a new OrganizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationResponse(createdAt string, id string, name string, owner string, updatedAt string) *OrganizationResponse {
	this := OrganizationResponse{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Name = name
	this.Owner = owner
	this.UpdatedAt = updatedAt
	return &this
}

// NewOrganizationResponseWithDefaults instantiates a new OrganizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationResponseWithDefaults() *OrganizationResponse {
	this := OrganizationResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *OrganizationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OrganizationResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationResponse) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value
func (o *OrganizationResponse) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *OrganizationResponse) SetOwner(v string) {
	o.Owner = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *OrganizationResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *OrganizationResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o OrganizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["owner"] = o.Owner
	toSerialize["updatedAt"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"id",
		"name",
		"owner",
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

	varOrganizationResponse := _OrganizationResponse{}

	err = json.Unmarshal(data, &varOrganizationResponse)

	if err != nil {
		return err
	}

	*o = OrganizationResponse(varOrganizationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationResponse struct {
	value *OrganizationResponse
	isSet bool
}

func (v NullableOrganizationResponse) Get() *OrganizationResponse {
	return v.value
}

func (v *NullableOrganizationResponse) Set(val *OrganizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationResponse(val *OrganizationResponse) *NullableOrganizationResponse {
	return &NullableOrganizationResponse{value: val, isSet: true}
}

func (v NullableOrganizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


