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

// checks if the OrganizationMemberResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationMemberResponse{}

// OrganizationMemberResponse struct for OrganizationMemberResponse
type OrganizationMemberResponse struct {
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	CreatedByName string `json:"createdByName"`
	Id string `json:"id"`
	InvitedAt string `json:"invitedAt"`
	OrganizationId string `json:"organizationId"`
	OrganizationOwner bool `json:"organizationOwner"`
	OrganizationProductRoles []RoleReducedResponse `json:"organizationProductRoles"`
	OrganizationRole OrganizationRoleReducedResponse `json:"organizationRole"`
	UpdatedAt string `json:"updatedAt"`
	User UserResponse `json:"user"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationMemberResponse OrganizationMemberResponse

// NewOrganizationMemberResponse instantiates a new OrganizationMemberResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationMemberResponse(createdAt string, createdBy string, createdByName string, id string, invitedAt string, organizationId string, organizationOwner bool, organizationProductRoles []RoleReducedResponse, organizationRole OrganizationRoleReducedResponse, updatedAt string, user UserResponse) *OrganizationMemberResponse {
	this := OrganizationMemberResponse{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.CreatedByName = createdByName
	this.Id = id
	this.InvitedAt = invitedAt
	this.OrganizationId = organizationId
	this.OrganizationOwner = organizationOwner
	this.OrganizationProductRoles = organizationProductRoles
	this.OrganizationRole = organizationRole
	this.UpdatedAt = updatedAt
	this.User = user
	return &this
}

// NewOrganizationMemberResponseWithDefaults instantiates a new OrganizationMemberResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationMemberResponseWithDefaults() *OrganizationMemberResponse {
	this := OrganizationMemberResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationMemberResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationMemberResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationMemberResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *OrganizationMemberResponse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *OrganizationMemberResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *OrganizationMemberResponse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedByName returns the CreatedByName field value
func (o *OrganizationMemberResponse) GetCreatedByName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value
// and a boolean to check if the value has been set.
func (o *OrganizationMemberResponse) GetCreatedByNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByName, true
}

// SetCreatedByName sets field value
func (o *OrganizationMemberResponse) SetCreatedByName(v string) {
	o.CreatedByName = v
}

// GetId returns the Id field value
func (o *OrganizationMemberResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationMemberResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationMemberResponse) SetId(v string) {
	o.Id = v
}

// GetInvitedAt returns the InvitedAt field value
func (o *OrganizationMemberResponse) GetInvitedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvitedAt
}

// GetInvitedAtOk returns a tuple with the InvitedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationMemberResponse) GetInvitedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvitedAt, true
}

// SetInvitedAt sets field value
func (o *OrganizationMemberResponse) SetInvitedAt(v string) {
	o.InvitedAt = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *OrganizationMemberResponse) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *OrganizationMemberResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *OrganizationMemberResponse) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetOrganizationOwner returns the OrganizationOwner field value
func (o *OrganizationMemberResponse) GetOrganizationOwner() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OrganizationOwner
}

// GetOrganizationOwnerOk returns a tuple with the OrganizationOwner field value
// and a boolean to check if the value has been set.
func (o *OrganizationMemberResponse) GetOrganizationOwnerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationOwner, true
}

// SetOrganizationOwner sets field value
func (o *OrganizationMemberResponse) SetOrganizationOwner(v bool) {
	o.OrganizationOwner = v
}

// GetOrganizationProductRoles returns the OrganizationProductRoles field value
func (o *OrganizationMemberResponse) GetOrganizationProductRoles() []RoleReducedResponse {
	if o == nil {
		var ret []RoleReducedResponse
		return ret
	}

	return o.OrganizationProductRoles
}

// GetOrganizationProductRolesOk returns a tuple with the OrganizationProductRoles field value
// and a boolean to check if the value has been set.
func (o *OrganizationMemberResponse) GetOrganizationProductRolesOk() ([]RoleReducedResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationProductRoles, true
}

// SetOrganizationProductRoles sets field value
func (o *OrganizationMemberResponse) SetOrganizationProductRoles(v []RoleReducedResponse) {
	o.OrganizationProductRoles = v
}

// GetOrganizationRole returns the OrganizationRole field value
func (o *OrganizationMemberResponse) GetOrganizationRole() OrganizationRoleReducedResponse {
	if o == nil {
		var ret OrganizationRoleReducedResponse
		return ret
	}

	return o.OrganizationRole
}

// GetOrganizationRoleOk returns a tuple with the OrganizationRole field value
// and a boolean to check if the value has been set.
func (o *OrganizationMemberResponse) GetOrganizationRoleOk() (*OrganizationRoleReducedResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationRole, true
}

// SetOrganizationRole sets field value
func (o *OrganizationMemberResponse) SetOrganizationRole(v OrganizationRoleReducedResponse) {
	o.OrganizationRole = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *OrganizationMemberResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationMemberResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *OrganizationMemberResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUser returns the User field value
func (o *OrganizationMemberResponse) GetUser() UserResponse {
	if o == nil {
		var ret UserResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *OrganizationMemberResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *OrganizationMemberResponse) SetUser(v UserResponse) {
	o.User = v
}

func (o OrganizationMemberResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationMemberResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["createdByName"] = o.CreatedByName
	toSerialize["id"] = o.Id
	toSerialize["invitedAt"] = o.InvitedAt
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["organizationOwner"] = o.OrganizationOwner
	toSerialize["organizationProductRoles"] = o.OrganizationProductRoles
	toSerialize["organizationRole"] = o.OrganizationRole
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationMemberResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"createdBy",
		"createdByName",
		"id",
		"invitedAt",
		"organizationId",
		"organizationOwner",
		"organizationProductRoles",
		"organizationRole",
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

	varOrganizationMemberResponse := _OrganizationMemberResponse{}

	err = json.Unmarshal(data, &varOrganizationMemberResponse)

	if err != nil {
		return err
	}

	*o = OrganizationMemberResponse(varOrganizationMemberResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "createdByName")
		delete(additionalProperties, "id")
		delete(additionalProperties, "invitedAt")
		delete(additionalProperties, "organizationId")
		delete(additionalProperties, "organizationOwner")
		delete(additionalProperties, "organizationProductRoles")
		delete(additionalProperties, "organizationRole")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationMemberResponse struct {
	value *OrganizationMemberResponse
	isSet bool
}

func (v NullableOrganizationMemberResponse) Get() *OrganizationMemberResponse {
	return v.value
}

func (v *NullableOrganizationMemberResponse) Set(val *OrganizationMemberResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationMemberResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationMemberResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationMemberResponse(val *OrganizationMemberResponse) *NullableOrganizationMemberResponse {
	return &NullableOrganizationMemberResponse{value: val, isSet: true}
}

func (v NullableOrganizationMemberResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationMemberResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


