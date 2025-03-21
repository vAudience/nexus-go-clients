/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.17.3
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the FunctionCall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionCall{}

// FunctionCall struct for FunctionCall
type FunctionCall struct {
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Parameters *string `json:"parameters,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FunctionCall FunctionCall

// NewFunctionCall instantiates a new FunctionCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionCall() *FunctionCall {
	this := FunctionCall{}
	return &this
}

// NewFunctionCallWithDefaults instantiates a new FunctionCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionCallWithDefaults() *FunctionCall {
	this := FunctionCall{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FunctionCall) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCall) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FunctionCall) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FunctionCall) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FunctionCall) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCall) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FunctionCall) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FunctionCall) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *FunctionCall) GetParameters() string {
	if o == nil || IsNil(o.Parameters) {
		var ret string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCall) GetParametersOk() (*string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *FunctionCall) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given string and assigns it to the Parameters field.
func (o *FunctionCall) SetParameters(v string) {
	o.Parameters = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FunctionCall) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionCall) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FunctionCall) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FunctionCall) SetVersion(v string) {
	o.Version = &v
}

func (o FunctionCall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FunctionCall) UnmarshalJSON(data []byte) (err error) {
	varFunctionCall := _FunctionCall{}

	err = json.Unmarshal(data, &varFunctionCall)

	if err != nil {
		return err
	}

	*o = FunctionCall(varFunctionCall)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFunctionCall struct {
	value *FunctionCall
	isSet bool
}

func (v NullableFunctionCall) Get() *FunctionCall {
	return v.value
}

func (v *NullableFunctionCall) Set(val *FunctionCall) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionCall) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionCall(val *FunctionCall) *NullableFunctionCall {
	return &NullableFunctionCall{value: val, isSet: true}
}

func (v NullableFunctionCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


