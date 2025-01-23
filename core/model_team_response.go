/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// checks if the TeamResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamResponse{}

// TeamResponse struct for TeamResponse
type TeamResponse struct {
	CreatedAt string `json:"createdAt"`
	Id string `json:"id"`
	Members []TeamMemberResponse `json:"members"`
	Name string `json:"name"`
	OrganizationId string `json:"organizationId"`
	Owner string `json:"owner"`
	UpdatedAt string `json:"updatedAt"`
	AdditionalProperties map[string]interface{}
}

type _TeamResponse TeamResponse

// NewTeamResponse instantiates a new TeamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamResponse(createdAt string, id string, members []TeamMemberResponse, name string, organizationId string, owner string, updatedAt string) *TeamResponse {
	this := TeamResponse{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Members = members
	this.Name = name
	this.OrganizationId = organizationId
	this.Owner = owner
	this.UpdatedAt = updatedAt
	return &this
}

// NewTeamResponseWithDefaults instantiates a new TeamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamResponseWithDefaults() *TeamResponse {
	this := TeamResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *TeamResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TeamResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TeamResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *TeamResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeamResponse) SetId(v string) {
	o.Id = v
}

// GetMembers returns the Members field value
func (o *TeamResponse) GetMembers() []TeamMemberResponse {
	if o == nil {
		var ret []TeamMemberResponse
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *TeamResponse) GetMembersOk() ([]TeamMemberResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *TeamResponse) SetMembers(v []TeamMemberResponse) {
	o.Members = v
}

// GetName returns the Name field value
func (o *TeamResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamResponse) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *TeamResponse) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *TeamResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *TeamResponse) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetOwner returns the Owner field value
func (o *TeamResponse) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *TeamResponse) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *TeamResponse) SetOwner(v string) {
	o.Owner = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TeamResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TeamResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TeamResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o TeamResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["members"] = o.Members
	toSerialize["name"] = o.Name
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["owner"] = o.Owner
	toSerialize["updatedAt"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TeamResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"id",
		"members",
		"name",
		"organizationId",
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

	varTeamResponse := _TeamResponse{}

	err = json.Unmarshal(data, &varTeamResponse)

	if err != nil {
		return err
	}

	*o = TeamResponse(varTeamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "id")
		delete(additionalProperties, "members")
		delete(additionalProperties, "name")
		delete(additionalProperties, "organizationId")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTeamResponse struct {
	value *TeamResponse
	isSet bool
}

func (v NullableTeamResponse) Get() *TeamResponse {
	return v.value
}

func (v *NullableTeamResponse) Set(val *TeamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamResponse(val *TeamResponse) *NullableTeamResponse {
	return &NullableTeamResponse{value: val, isSet: true}
}

func (v NullableTeamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


