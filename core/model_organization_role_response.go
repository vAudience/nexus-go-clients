/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// checks if the OrganizationRoleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationRoleResponse{}

// OrganizationRoleResponse struct for OrganizationRoleResponse
type OrganizationRoleResponse struct {
	CreatedAt string `json:"createdAt"`
	DefaultMemberRole bool `json:"defaultMemberRole"`
	Id string `json:"id"`
	Name string `json:"name"`
	OrganizationId string `json:"organizationId"`
	Permissions []string `json:"permissions"`
	UpdatedAt string `json:"updatedAt"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationRoleResponse OrganizationRoleResponse

// NewOrganizationRoleResponse instantiates a new OrganizationRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationRoleResponse(createdAt string, defaultMemberRole bool, id string, name string, organizationId string, permissions []string, updatedAt string) *OrganizationRoleResponse {
	this := OrganizationRoleResponse{}
	this.CreatedAt = createdAt
	this.DefaultMemberRole = defaultMemberRole
	this.Id = id
	this.Name = name
	this.OrganizationId = organizationId
	this.Permissions = permissions
	this.UpdatedAt = updatedAt
	return &this
}

// NewOrganizationRoleResponseWithDefaults instantiates a new OrganizationRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationRoleResponseWithDefaults() *OrganizationRoleResponse {
	this := OrganizationRoleResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationRoleResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationRoleResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDefaultMemberRole returns the DefaultMemberRole field value
func (o *OrganizationRoleResponse) GetDefaultMemberRole() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultMemberRole
}

// GetDefaultMemberRoleOk returns a tuple with the DefaultMemberRole field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleResponse) GetDefaultMemberRoleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultMemberRole, true
}

// SetDefaultMemberRole sets field value
func (o *OrganizationRoleResponse) SetDefaultMemberRole(v bool) {
	o.DefaultMemberRole = v
}

// GetId returns the Id field value
func (o *OrganizationRoleResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationRoleResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OrganizationRoleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationRoleResponse) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *OrganizationRoleResponse) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *OrganizationRoleResponse) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetPermissions returns the Permissions field value
func (o *OrganizationRoleResponse) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleResponse) GetPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *OrganizationRoleResponse) SetPermissions(v []string) {
	o.Permissions = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *OrganizationRoleResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationRoleResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *OrganizationRoleResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o OrganizationRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationRoleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["defaultMemberRole"] = o.DefaultMemberRole
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["permissions"] = o.Permissions
	toSerialize["updatedAt"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationRoleResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"defaultMemberRole",
		"id",
		"name",
		"organizationId",
		"permissions",
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

	varOrganizationRoleResponse := _OrganizationRoleResponse{}

	err = json.Unmarshal(data, &varOrganizationRoleResponse)

	if err != nil {
		return err
	}

	*o = OrganizationRoleResponse(varOrganizationRoleResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "defaultMemberRole")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "organizationId")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationRoleResponse struct {
	value *OrganizationRoleResponse
	isSet bool
}

func (v NullableOrganizationRoleResponse) Get() *OrganizationRoleResponse {
	return v.value
}

func (v *NullableOrganizationRoleResponse) Set(val *OrganizationRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationRoleResponse(val *OrganizationRoleResponse) *NullableOrganizationRoleResponse {
	return &NullableOrganizationRoleResponse{value: val, isSet: true}
}

func (v NullableOrganizationRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


