/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.13.3
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// checks if the AIModelServiceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIModelServiceObject{}

// AIModelServiceObject struct for AIModelServiceObject
type AIModelServiceObject struct {
	// 1.0 is default, we use this to adjust our margin
	CostMultiplier *float32 `json:"cost_multiplier,omitempty"`
	CreatedAt *int64 `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	HostingLocations *map[string]HostingLocation `json:"hosting_locations,omitempty"`
	Id string `json:"id"`
	InternalId *string `json:"internal_id,omitempty"`
	IsPublic *bool `json:"is_public,omitempty"`
	Name string `json:"name"`
	OwnerId string `json:"owner_id"`
	OwnerOrganizationId string `json:"owner_organization_id"`
	// this is used for internal identification!
	ServiceImpl string `json:"service_impl"`
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AIModelServiceObject AIModelServiceObject

// NewAIModelServiceObject instantiates a new AIModelServiceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIModelServiceObject(id string, name string, ownerId string, ownerOrganizationId string, serviceImpl string) *AIModelServiceObject {
	this := AIModelServiceObject{}
	this.Id = id
	this.Name = name
	this.OwnerId = ownerId
	this.OwnerOrganizationId = ownerOrganizationId
	this.ServiceImpl = serviceImpl
	return &this
}

// NewAIModelServiceObjectWithDefaults instantiates a new AIModelServiceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIModelServiceObjectWithDefaults() *AIModelServiceObject {
	this := AIModelServiceObject{}
	return &this
}

// GetCostMultiplier returns the CostMultiplier field value if set, zero value otherwise.
func (o *AIModelServiceObject) GetCostMultiplier() float32 {
	if o == nil || IsNil(o.CostMultiplier) {
		var ret float32
		return ret
	}
	return *o.CostMultiplier
}

// GetCostMultiplierOk returns a tuple with the CostMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetCostMultiplierOk() (*float32, bool) {
	if o == nil || IsNil(o.CostMultiplier) {
		return nil, false
	}
	return o.CostMultiplier, true
}

// HasCostMultiplier returns a boolean if a field has been set.
func (o *AIModelServiceObject) HasCostMultiplier() bool {
	if o != nil && !IsNil(o.CostMultiplier) {
		return true
	}

	return false
}

// SetCostMultiplier gets a reference to the given float32 and assigns it to the CostMultiplier field.
func (o *AIModelServiceObject) SetCostMultiplier(v float32) {
	o.CostMultiplier = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AIModelServiceObject) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AIModelServiceObject) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *AIModelServiceObject) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AIModelServiceObject) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AIModelServiceObject) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AIModelServiceObject) SetDescription(v string) {
	o.Description = &v
}

// GetHostingLocations returns the HostingLocations field value if set, zero value otherwise.
func (o *AIModelServiceObject) GetHostingLocations() map[string]HostingLocation {
	if o == nil || IsNil(o.HostingLocations) {
		var ret map[string]HostingLocation
		return ret
	}
	return *o.HostingLocations
}

// GetHostingLocationsOk returns a tuple with the HostingLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetHostingLocationsOk() (*map[string]HostingLocation, bool) {
	if o == nil || IsNil(o.HostingLocations) {
		return nil, false
	}
	return o.HostingLocations, true
}

// HasHostingLocations returns a boolean if a field has been set.
func (o *AIModelServiceObject) HasHostingLocations() bool {
	if o != nil && !IsNil(o.HostingLocations) {
		return true
	}

	return false
}

// SetHostingLocations gets a reference to the given map[string]HostingLocation and assigns it to the HostingLocations field.
func (o *AIModelServiceObject) SetHostingLocations(v map[string]HostingLocation) {
	o.HostingLocations = &v
}

// GetId returns the Id field value
func (o *AIModelServiceObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AIModelServiceObject) SetId(v string) {
	o.Id = v
}

// GetInternalId returns the InternalId field value if set, zero value otherwise.
func (o *AIModelServiceObject) GetInternalId() string {
	if o == nil || IsNil(o.InternalId) {
		var ret string
		return ret
	}
	return *o.InternalId
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetInternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.InternalId) {
		return nil, false
	}
	return o.InternalId, true
}

// HasInternalId returns a boolean if a field has been set.
func (o *AIModelServiceObject) HasInternalId() bool {
	if o != nil && !IsNil(o.InternalId) {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given string and assigns it to the InternalId field.
func (o *AIModelServiceObject) SetInternalId(v string) {
	o.InternalId = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *AIModelServiceObject) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *AIModelServiceObject) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *AIModelServiceObject) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetName returns the Name field value
func (o *AIModelServiceObject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AIModelServiceObject) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value
func (o *AIModelServiceObject) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *AIModelServiceObject) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetOwnerOrganizationId returns the OwnerOrganizationId field value
func (o *AIModelServiceObject) GetOwnerOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerOrganizationId
}

// GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field value
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetOwnerOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerOrganizationId, true
}

// SetOwnerOrganizationId sets field value
func (o *AIModelServiceObject) SetOwnerOrganizationId(v string) {
	o.OwnerOrganizationId = v
}

// GetServiceImpl returns the ServiceImpl field value
func (o *AIModelServiceObject) GetServiceImpl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceImpl
}

// GetServiceImplOk returns a tuple with the ServiceImpl field value
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetServiceImplOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceImpl, true
}

// SetServiceImpl sets field value
func (o *AIModelServiceObject) SetServiceImpl(v string) {
	o.ServiceImpl = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AIModelServiceObject) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AIModelServiceObject) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *AIModelServiceObject) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *AIModelServiceObject) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIModelServiceObject) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *AIModelServiceObject) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *AIModelServiceObject) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o AIModelServiceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIModelServiceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CostMultiplier) {
		toSerialize["cost_multiplier"] = o.CostMultiplier
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.HostingLocations) {
		toSerialize["hosting_locations"] = o.HostingLocations
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.InternalId) {
		toSerialize["internal_id"] = o.InternalId
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	toSerialize["name"] = o.Name
	toSerialize["owner_id"] = o.OwnerId
	toSerialize["owner_organization_id"] = o.OwnerOrganizationId
	toSerialize["service_impl"] = o.ServiceImpl
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIModelServiceObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"owner_id",
		"owner_organization_id",
		"service_impl",
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

	varAIModelServiceObject := _AIModelServiceObject{}

	err = json.Unmarshal(data, &varAIModelServiceObject)

	if err != nil {
		return err
	}

	*o = AIModelServiceObject(varAIModelServiceObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cost_multiplier")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "description")
		delete(additionalProperties, "hosting_locations")
		delete(additionalProperties, "id")
		delete(additionalProperties, "internal_id")
		delete(additionalProperties, "is_public")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner_id")
		delete(additionalProperties, "owner_organization_id")
		delete(additionalProperties, "service_impl")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "updated_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIModelServiceObject struct {
	value *AIModelServiceObject
	isSet bool
}

func (v NullableAIModelServiceObject) Get() *AIModelServiceObject {
	return v.value
}

func (v *NullableAIModelServiceObject) Set(val *AIModelServiceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAIModelServiceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAIModelServiceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIModelServiceObject(val *AIModelServiceObject) *NullableAIModelServiceObject {
	return &NullableAIModelServiceObject{value: val, isSet: true}
}

func (v NullableAIModelServiceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIModelServiceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


