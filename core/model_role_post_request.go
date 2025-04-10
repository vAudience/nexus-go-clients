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

// checks if the RolePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolePostRequest{}

// RolePostRequest struct for RolePostRequest
type RolePostRequest struct {
	DefaultMemberRole *bool `json:"defaultMemberRole,omitempty"`
	Name string `json:"name"`
	Permissions []string `json:"permissions"`
	AdditionalProperties map[string]interface{}
}

type _RolePostRequest RolePostRequest

// NewRolePostRequest instantiates a new RolePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePostRequest(name string, permissions []string) *RolePostRequest {
	this := RolePostRequest{}
	this.Name = name
	this.Permissions = permissions
	return &this
}

// NewRolePostRequestWithDefaults instantiates a new RolePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePostRequestWithDefaults() *RolePostRequest {
	this := RolePostRequest{}
	return &this
}

// GetDefaultMemberRole returns the DefaultMemberRole field value if set, zero value otherwise.
func (o *RolePostRequest) GetDefaultMemberRole() bool {
	if o == nil || IsNil(o.DefaultMemberRole) {
		var ret bool
		return ret
	}
	return *o.DefaultMemberRole
}

// GetDefaultMemberRoleOk returns a tuple with the DefaultMemberRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePostRequest) GetDefaultMemberRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultMemberRole) {
		return nil, false
	}
	return o.DefaultMemberRole, true
}

// HasDefaultMemberRole returns a boolean if a field has been set.
func (o *RolePostRequest) HasDefaultMemberRole() bool {
	if o != nil && !IsNil(o.DefaultMemberRole) {
		return true
	}

	return false
}

// SetDefaultMemberRole gets a reference to the given bool and assigns it to the DefaultMemberRole field.
func (o *RolePostRequest) SetDefaultMemberRole(v bool) {
	o.DefaultMemberRole = &v
}

// GetName returns the Name field value
func (o *RolePostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RolePostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RolePostRequest) SetName(v string) {
	o.Name = v
}

// GetPermissions returns the Permissions field value
func (o *RolePostRequest) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *RolePostRequest) GetPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *RolePostRequest) SetPermissions(v []string) {
	o.Permissions = v
}

func (o RolePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultMemberRole) {
		toSerialize["defaultMemberRole"] = o.DefaultMemberRole
	}
	toSerialize["name"] = o.Name
	toSerialize["permissions"] = o.Permissions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RolePostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"permissions",
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

	varRolePostRequest := _RolePostRequest{}

	err = json.Unmarshal(data, &varRolePostRequest)

	if err != nil {
		return err
	}

	*o = RolePostRequest(varRolePostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "defaultMemberRole")
		delete(additionalProperties, "name")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRolePostRequest struct {
	value *RolePostRequest
	isSet bool
}

func (v NullableRolePostRequest) Get() *RolePostRequest {
	return v.value
}

func (v *NullableRolePostRequest) Set(val *RolePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePostRequest(val *RolePostRequest) *NullableRolePostRequest {
	return &NullableRolePostRequest{value: val, isSet: true}
}

func (v NullableRolePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


