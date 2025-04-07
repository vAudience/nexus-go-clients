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

// checks if the TeamMemberResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamMemberResponse{}

// TeamMemberResponse struct for TeamMemberResponse
type TeamMemberResponse struct {
	CreatedAt string `json:"createdAt"`
	Id string `json:"id"`
	TeamId string `json:"teamId"`
	TeamOwner bool `json:"teamOwner"`
	TeamRole *TeamRoleReducedResponse `json:"teamRole,omitempty"`
	UpdatedAt string `json:"updatedAt"`
	User UserResponse `json:"user"`
	AdditionalProperties map[string]interface{}
}

type _TeamMemberResponse TeamMemberResponse

// NewTeamMemberResponse instantiates a new TeamMemberResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamMemberResponse(createdAt string, id string, teamId string, teamOwner bool, updatedAt string, user UserResponse) *TeamMemberResponse {
	this := TeamMemberResponse{}
	this.CreatedAt = createdAt
	this.Id = id
	this.TeamId = teamId
	this.TeamOwner = teamOwner
	this.UpdatedAt = updatedAt
	this.User = user
	return &this
}

// NewTeamMemberResponseWithDefaults instantiates a new TeamMemberResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamMemberResponseWithDefaults() *TeamMemberResponse {
	this := TeamMemberResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *TeamMemberResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TeamMemberResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TeamMemberResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *TeamMemberResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamMemberResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeamMemberResponse) SetId(v string) {
	o.Id = v
}

// GetTeamId returns the TeamId field value
func (o *TeamMemberResponse) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *TeamMemberResponse) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *TeamMemberResponse) SetTeamId(v string) {
	o.TeamId = v
}

// GetTeamOwner returns the TeamOwner field value
func (o *TeamMemberResponse) GetTeamOwner() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TeamOwner
}

// GetTeamOwnerOk returns a tuple with the TeamOwner field value
// and a boolean to check if the value has been set.
func (o *TeamMemberResponse) GetTeamOwnerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamOwner, true
}

// SetTeamOwner sets field value
func (o *TeamMemberResponse) SetTeamOwner(v bool) {
	o.TeamOwner = v
}

// GetTeamRole returns the TeamRole field value if set, zero value otherwise.
func (o *TeamMemberResponse) GetTeamRole() TeamRoleReducedResponse {
	if o == nil || IsNil(o.TeamRole) {
		var ret TeamRoleReducedResponse
		return ret
	}
	return *o.TeamRole
}

// GetTeamRoleOk returns a tuple with the TeamRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamMemberResponse) GetTeamRoleOk() (*TeamRoleReducedResponse, bool) {
	if o == nil || IsNil(o.TeamRole) {
		return nil, false
	}
	return o.TeamRole, true
}

// HasTeamRole returns a boolean if a field has been set.
func (o *TeamMemberResponse) HasTeamRole() bool {
	if o != nil && !IsNil(o.TeamRole) {
		return true
	}

	return false
}

// SetTeamRole gets a reference to the given TeamRoleReducedResponse and assigns it to the TeamRole field.
func (o *TeamMemberResponse) SetTeamRole(v TeamRoleReducedResponse) {
	o.TeamRole = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TeamMemberResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TeamMemberResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TeamMemberResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUser returns the User field value
func (o *TeamMemberResponse) GetUser() UserResponse {
	if o == nil {
		var ret UserResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TeamMemberResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TeamMemberResponse) SetUser(v UserResponse) {
	o.User = v
}

func (o TeamMemberResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamMemberResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["teamId"] = o.TeamId
	toSerialize["teamOwner"] = o.TeamOwner
	if !IsNil(o.TeamRole) {
		toSerialize["teamRole"] = o.TeamRole
	}
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TeamMemberResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"id",
		"teamId",
		"teamOwner",
		"updatedAt",
		"user",
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

	varTeamMemberResponse := _TeamMemberResponse{}

	err = json.Unmarshal(data, &varTeamMemberResponse)

	if err != nil {
		return err
	}

	*o = TeamMemberResponse(varTeamMemberResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "id")
		delete(additionalProperties, "teamId")
		delete(additionalProperties, "teamOwner")
		delete(additionalProperties, "teamRole")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTeamMemberResponse struct {
	value *TeamMemberResponse
	isSet bool
}

func (v NullableTeamMemberResponse) Get() *TeamMemberResponse {
	return v.value
}

func (v *NullableTeamMemberResponse) Set(val *TeamMemberResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMemberResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMemberResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMemberResponse(val *TeamMemberResponse) *NullableTeamMemberResponse {
	return &NullableTeamMemberResponse{value: val, isSet: true}
}

func (v NullableTeamMemberResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMemberResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


