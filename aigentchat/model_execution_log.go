/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.14.0
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
)

// checks if the ExecutionLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionLog{}

// ExecutionLog struct for ExecutionLog
type ExecutionLog struct {
	AgentId *string `json:"agent_id,omitempty"`
	AiModelId *string `json:"ai_model_id,omitempty"`
	AiModelServiceId *string `json:"ai_model_service_id,omitempty"`
	ChannelId *string `json:"channel_id,omitempty"`
	CreatedAt *int64 `json:"created_at,omitempty"`
	FinalCostInEuro *float64 `json:"final_cost_in_euro,omitempty"`
	Id *string `json:"id,omitempty"`
	MessageId *string `json:"message_id,omitempty"`
	OwnerId *string `json:"owner_id,omitempty"`
	OwnerOrganizationId *string `json:"owner_organization_id,omitempty"`
	Usage []ExecutionUsageCost `json:"usage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExecutionLog ExecutionLog

// NewExecutionLog instantiates a new ExecutionLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionLog() *ExecutionLog {
	this := ExecutionLog{}
	return &this
}

// NewExecutionLogWithDefaults instantiates a new ExecutionLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionLogWithDefaults() *ExecutionLog {
	this := ExecutionLog{}
	return &this
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *ExecutionLog) GetAgentId() string {
	if o == nil || IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetAgentIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *ExecutionLog) HasAgentId() bool {
	if o != nil && !IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *ExecutionLog) SetAgentId(v string) {
	o.AgentId = &v
}

// GetAiModelId returns the AiModelId field value if set, zero value otherwise.
func (o *ExecutionLog) GetAiModelId() string {
	if o == nil || IsNil(o.AiModelId) {
		var ret string
		return ret
	}
	return *o.AiModelId
}

// GetAiModelIdOk returns a tuple with the AiModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetAiModelIdOk() (*string, bool) {
	if o == nil || IsNil(o.AiModelId) {
		return nil, false
	}
	return o.AiModelId, true
}

// HasAiModelId returns a boolean if a field has been set.
func (o *ExecutionLog) HasAiModelId() bool {
	if o != nil && !IsNil(o.AiModelId) {
		return true
	}

	return false
}

// SetAiModelId gets a reference to the given string and assigns it to the AiModelId field.
func (o *ExecutionLog) SetAiModelId(v string) {
	o.AiModelId = &v
}

// GetAiModelServiceId returns the AiModelServiceId field value if set, zero value otherwise.
func (o *ExecutionLog) GetAiModelServiceId() string {
	if o == nil || IsNil(o.AiModelServiceId) {
		var ret string
		return ret
	}
	return *o.AiModelServiceId
}

// GetAiModelServiceIdOk returns a tuple with the AiModelServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetAiModelServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AiModelServiceId) {
		return nil, false
	}
	return o.AiModelServiceId, true
}

// HasAiModelServiceId returns a boolean if a field has been set.
func (o *ExecutionLog) HasAiModelServiceId() bool {
	if o != nil && !IsNil(o.AiModelServiceId) {
		return true
	}

	return false
}

// SetAiModelServiceId gets a reference to the given string and assigns it to the AiModelServiceId field.
func (o *ExecutionLog) SetAiModelServiceId(v string) {
	o.AiModelServiceId = &v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *ExecutionLog) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *ExecutionLog) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *ExecutionLog) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ExecutionLog) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ExecutionLog) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ExecutionLog) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetFinalCostInEuro returns the FinalCostInEuro field value if set, zero value otherwise.
func (o *ExecutionLog) GetFinalCostInEuro() float64 {
	if o == nil || IsNil(o.FinalCostInEuro) {
		var ret float64
		return ret
	}
	return *o.FinalCostInEuro
}

// GetFinalCostInEuroOk returns a tuple with the FinalCostInEuro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetFinalCostInEuroOk() (*float64, bool) {
	if o == nil || IsNil(o.FinalCostInEuro) {
		return nil, false
	}
	return o.FinalCostInEuro, true
}

// HasFinalCostInEuro returns a boolean if a field has been set.
func (o *ExecutionLog) HasFinalCostInEuro() bool {
	if o != nil && !IsNil(o.FinalCostInEuro) {
		return true
	}

	return false
}

// SetFinalCostInEuro gets a reference to the given float64 and assigns it to the FinalCostInEuro field.
func (o *ExecutionLog) SetFinalCostInEuro(v float64) {
	o.FinalCostInEuro = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExecutionLog) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExecutionLog) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExecutionLog) SetId(v string) {
	o.Id = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *ExecutionLog) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *ExecutionLog) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *ExecutionLog) SetMessageId(v string) {
	o.MessageId = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ExecutionLog) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ExecutionLog) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ExecutionLog) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerOrganizationId returns the OwnerOrganizationId field value if set, zero value otherwise.
func (o *ExecutionLog) GetOwnerOrganizationId() string {
	if o == nil || IsNil(o.OwnerOrganizationId) {
		var ret string
		return ret
	}
	return *o.OwnerOrganizationId
}

// GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetOwnerOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerOrganizationId) {
		return nil, false
	}
	return o.OwnerOrganizationId, true
}

// HasOwnerOrganizationId returns a boolean if a field has been set.
func (o *ExecutionLog) HasOwnerOrganizationId() bool {
	if o != nil && !IsNil(o.OwnerOrganizationId) {
		return true
	}

	return false
}

// SetOwnerOrganizationId gets a reference to the given string and assigns it to the OwnerOrganizationId field.
func (o *ExecutionLog) SetOwnerOrganizationId(v string) {
	o.OwnerOrganizationId = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ExecutionLog) GetUsage() []ExecutionUsageCost {
	if o == nil || IsNil(o.Usage) {
		var ret []ExecutionUsageCost
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionLog) GetUsageOk() ([]ExecutionUsageCost, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ExecutionLog) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []ExecutionUsageCost and assigns it to the Usage field.
func (o *ExecutionLog) SetUsage(v []ExecutionUsageCost) {
	o.Usage = v
}

func (o ExecutionLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentId) {
		toSerialize["agent_id"] = o.AgentId
	}
	if !IsNil(o.AiModelId) {
		toSerialize["ai_model_id"] = o.AiModelId
	}
	if !IsNil(o.AiModelServiceId) {
		toSerialize["ai_model_service_id"] = o.AiModelServiceId
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.FinalCostInEuro) {
		toSerialize["final_cost_in_euro"] = o.FinalCostInEuro
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MessageId) {
		toSerialize["message_id"] = o.MessageId
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.OwnerOrganizationId) {
		toSerialize["owner_organization_id"] = o.OwnerOrganizationId
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExecutionLog) UnmarshalJSON(data []byte) (err error) {
	varExecutionLog := _ExecutionLog{}

	err = json.Unmarshal(data, &varExecutionLog)

	if err != nil {
		return err
	}

	*o = ExecutionLog(varExecutionLog)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "agent_id")
		delete(additionalProperties, "ai_model_id")
		delete(additionalProperties, "ai_model_service_id")
		delete(additionalProperties, "channel_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "final_cost_in_euro")
		delete(additionalProperties, "id")
		delete(additionalProperties, "message_id")
		delete(additionalProperties, "owner_id")
		delete(additionalProperties, "owner_organization_id")
		delete(additionalProperties, "usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecutionLog struct {
	value *ExecutionLog
	isSet bool
}

func (v NullableExecutionLog) Get() *ExecutionLog {
	return v.value
}

func (v *NullableExecutionLog) Set(val *ExecutionLog) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionLog) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionLog(val *ExecutionLog) *NullableExecutionLog {
	return &NullableExecutionLog{value: val, isSet: true}
}

func (v NullableExecutionLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


