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

// checks if the TeamRoleReducedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamRoleReducedResponse{}

// TeamRoleReducedResponse struct for TeamRoleReducedResponse
type TeamRoleReducedResponse struct {
	CreatedAt string `json:"createdAt"`
	DefaultMemberRole bool `json:"defaultMemberRole"`
	Id string `json:"id"`
	Name string `json:"name"`
	TeamId string `json:"teamId"`
	UpdatedAt string `json:"updatedAt"`
	AdditionalProperties map[string]interface{}
}

type _TeamRoleReducedResponse TeamRoleReducedResponse

// NewTeamRoleReducedResponse instantiates a new TeamRoleReducedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamRoleReducedResponse(createdAt string, defaultMemberRole bool, id string, name string, teamId string, updatedAt string) *TeamRoleReducedResponse {
	this := TeamRoleReducedResponse{}
	this.CreatedAt = createdAt
	this.DefaultMemberRole = defaultMemberRole
	this.Id = id
	this.Name = name
	this.TeamId = teamId
	this.UpdatedAt = updatedAt
	return &this
}

// NewTeamRoleReducedResponseWithDefaults instantiates a new TeamRoleReducedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamRoleReducedResponseWithDefaults() *TeamRoleReducedResponse {
	this := TeamRoleReducedResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *TeamRoleReducedResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TeamRoleReducedResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TeamRoleReducedResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDefaultMemberRole returns the DefaultMemberRole field value
func (o *TeamRoleReducedResponse) GetDefaultMemberRole() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultMemberRole
}

// GetDefaultMemberRoleOk returns a tuple with the DefaultMemberRole field value
// and a boolean to check if the value has been set.
func (o *TeamRoleReducedResponse) GetDefaultMemberRoleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultMemberRole, true
}

// SetDefaultMemberRole sets field value
func (o *TeamRoleReducedResponse) SetDefaultMemberRole(v bool) {
	o.DefaultMemberRole = v
}

// GetId returns the Id field value
func (o *TeamRoleReducedResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamRoleReducedResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeamRoleReducedResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TeamRoleReducedResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamRoleReducedResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamRoleReducedResponse) SetName(v string) {
	o.Name = v
}

// GetTeamId returns the TeamId field value
func (o *TeamRoleReducedResponse) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *TeamRoleReducedResponse) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *TeamRoleReducedResponse) SetTeamId(v string) {
	o.TeamId = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TeamRoleReducedResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TeamRoleReducedResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TeamRoleReducedResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o TeamRoleReducedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamRoleReducedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["defaultMemberRole"] = o.DefaultMemberRole
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["teamId"] = o.TeamId
	toSerialize["updatedAt"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TeamRoleReducedResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"defaultMemberRole",
		"id",
		"name",
		"teamId",
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

	varTeamRoleReducedResponse := _TeamRoleReducedResponse{}

	err = json.Unmarshal(data, &varTeamRoleReducedResponse)

	if err != nil {
		return err
	}

	*o = TeamRoleReducedResponse(varTeamRoleReducedResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "defaultMemberRole")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "teamId")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTeamRoleReducedResponse struct {
	value *TeamRoleReducedResponse
	isSet bool
}

func (v NullableTeamRoleReducedResponse) Get() *TeamRoleReducedResponse {
	return v.value
}

func (v *NullableTeamRoleReducedResponse) Set(val *TeamRoleReducedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamRoleReducedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamRoleReducedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamRoleReducedResponse(val *TeamRoleReducedResponse) *NullableTeamRoleReducedResponse {
	return &NullableTeamRoleReducedResponse{value: val, isSet: true}
}

func (v NullableTeamRoleReducedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamRoleReducedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


