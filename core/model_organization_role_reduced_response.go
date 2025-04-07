/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// checks if the OrganizationRoleReducedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationRoleReducedResponse{}

// OrganizationRoleReducedResponse struct for OrganizationRoleReducedResponse
type OrganizationRoleReducedResponse struct {
	CreatedAt string `json:"createdAt"`
	DefaultMemberRole bool `json:"defaultMemberRole"`
	Id string `json:"id"`
	Name string `json:"name"`
	OrganizationId string `json:"organizationId"`
	UpdatedAt string `json:"updatedAt"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationRoleReducedResponse OrganizationRoleReducedResponse

// NewOrganizationRoleReducedResponse instantiates a new OrganizationRoleReducedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationRoleReducedResponse(createdAt string, defaultMemberRole bool, id string, name string, organizationId string, updatedAt string) *OrganizationRoleReducedResponse {
	this := OrganizationRoleReducedResponse{}
	this.CreatedAt = createdAt
	this.DefaultMemberRole = defaultMemberRole
	this.Id = id
	this.Name = name
	this.OrganizationId = organizationId
	this.UpdatedAt = updatedAt
	return &this
}

// NewOrganizationRoleReducedResponseWithDefaults instantiates a new OrganizationRoleReducedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationRoleReducedResponseWithDefaults() *OrganizationRoleReducedResponse {
	this := OrganizationRoleReducedResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationRoleReducedResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleReducedResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationRoleReducedResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDefaultMemberRole returns the DefaultMemberRole field value
func (o *OrganizationRoleReducedResponse) GetDefaultMemberRole() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultMemberRole
}

// GetDefaultMemberRoleOk returns a tuple with the DefaultMemberRole field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleReducedResponse) GetDefaultMemberRoleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultMemberRole, true
}

// SetDefaultMemberRole sets field value
func (o *OrganizationRoleReducedResponse) SetDefaultMemberRole(v bool) {
	o.DefaultMemberRole = v
}

// GetId returns the Id field value
func (o *OrganizationRoleReducedResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleReducedResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationRoleReducedResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OrganizationRoleReducedResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleReducedResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationRoleReducedResponse) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *OrganizationRoleReducedResponse) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleReducedResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *OrganizationRoleReducedResponse) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *OrganizationRoleReducedResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleReducedResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *OrganizationRoleReducedResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o OrganizationRoleReducedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationRoleReducedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["defaultMemberRole"] = o.DefaultMemberRole
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["updatedAt"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationRoleReducedResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"defaultMemberRole",
		"id",
		"name",
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

	varOrganizationRoleReducedResponse := _OrganizationRoleReducedResponse{}

	err = json.Unmarshal(data, &varOrganizationRoleReducedResponse)

	if err != nil {
		return err
	}

	*o = OrganizationRoleReducedResponse(varOrganizationRoleReducedResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "defaultMemberRole")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "organizationId")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationRoleReducedResponse struct {
	value *OrganizationRoleReducedResponse
	isSet bool
}

func (v NullableOrganizationRoleReducedResponse) Get() *OrganizationRoleReducedResponse {
	return v.value
}

func (v *NullableOrganizationRoleReducedResponse) Set(val *OrganizationRoleReducedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationRoleReducedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationRoleReducedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationRoleReducedResponse(val *OrganizationRoleReducedResponse) *NullableOrganizationRoleReducedResponse {
	return &NullableOrganizationRoleReducedResponse{value: val, isSet: true}
}

func (v NullableOrganizationRoleReducedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationRoleReducedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


