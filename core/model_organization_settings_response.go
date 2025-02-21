/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// checks if the OrganizationSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationSettingsResponse{}

// OrganizationSettingsResponse struct for OrganizationSettingsResponse
type OrganizationSettingsResponse struct {
	CreatedAt string `json:"createdAt"`
	DefaultAgentId string `json:"defaultAgentId"`
	Id string `json:"id"`
	OrganizationId string `json:"organizationId"`
	UpdatedAt string `json:"updatedAt"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationSettingsResponse OrganizationSettingsResponse

// NewOrganizationSettingsResponse instantiates a new OrganizationSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSettingsResponse(createdAt string, defaultAgentId string, id string, organizationId string, updatedAt string) *OrganizationSettingsResponse {
	this := OrganizationSettingsResponse{}
	this.CreatedAt = createdAt
	this.DefaultAgentId = defaultAgentId
	this.Id = id
	this.OrganizationId = organizationId
	this.UpdatedAt = updatedAt
	return &this
}

// NewOrganizationSettingsResponseWithDefaults instantiates a new OrganizationSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSettingsResponseWithDefaults() *OrganizationSettingsResponse {
	this := OrganizationSettingsResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationSettingsResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationSettingsResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationSettingsResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDefaultAgentId returns the DefaultAgentId field value
func (o *OrganizationSettingsResponse) GetDefaultAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultAgentId
}

// GetDefaultAgentIdOk returns a tuple with the DefaultAgentId field value
// and a boolean to check if the value has been set.
func (o *OrganizationSettingsResponse) GetDefaultAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultAgentId, true
}

// SetDefaultAgentId sets field value
func (o *OrganizationSettingsResponse) SetDefaultAgentId(v string) {
	o.DefaultAgentId = v
}

// GetId returns the Id field value
func (o *OrganizationSettingsResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationSettingsResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationSettingsResponse) SetId(v string) {
	o.Id = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *OrganizationSettingsResponse) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *OrganizationSettingsResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *OrganizationSettingsResponse) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *OrganizationSettingsResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationSettingsResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *OrganizationSettingsResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o OrganizationSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["defaultAgentId"] = o.DefaultAgentId
	toSerialize["id"] = o.Id
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["updatedAt"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"defaultAgentId",
		"id",
		"organizationId",
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

	varOrganizationSettingsResponse := _OrganizationSettingsResponse{}

	err = json.Unmarshal(data, &varOrganizationSettingsResponse)

	if err != nil {
		return err
	}

	*o = OrganizationSettingsResponse(varOrganizationSettingsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "defaultAgentId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "organizationId")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationSettingsResponse struct {
	value *OrganizationSettingsResponse
	isSet bool
}

func (v NullableOrganizationSettingsResponse) Get() *OrganizationSettingsResponse {
	return v.value
}

func (v *NullableOrganizationSettingsResponse) Set(val *OrganizationSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSettingsResponse(val *OrganizationSettingsResponse) *NullableOrganizationSettingsResponse {
	return &NullableOrganizationSettingsResponse{value: val, isSet: true}
}

func (v NullableOrganizationSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


