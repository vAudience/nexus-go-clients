/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.17.8
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// checks if the AIgencyMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIgencyMessage{}

// AIgencyMessage struct for AIgencyMessage
type AIgencyMessage struct {
	AiModelId string `json:"ai_model_id"`
	AiServiceId string `json:"ai_service_id"`
	Attachments AIgencyMessageFileList `json:"attachments"`
	ChannelId string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Content AIgencyMessageContentList `json:"content"`
	CreatedAt int64 `json:"created_at"`
	CreatedForFeature *AIModelFeature `json:"created_for_feature,omitempty"`
	Error *AiServiceError `json:"error,omitempty"`
	ExecutionId *string `json:"execution_id,omitempty"`
	FinishReason *FinishReason `json:"finish_reason,omitempty"`
	Id string `json:"id"`
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
	MissionId *string `json:"mission_id,omitempty"`
	OwnerOrganizationId string `json:"owner_organization_id"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
	ResponseToId *string `json:"response_to_id,omitempty"`
	SenderConversationRole ConversationRole `json:"sender_conversation_role"`
	SenderId string `json:"sender_id"`
	SenderName string `json:"sender_name"`
	TokenCount *int32 `json:"token_count,omitempty"`
	TokenDirection TokenDirection `json:"token_direction"`
	Type AIgencyMessageType `json:"type"`
	UpdatedAt int64 `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _AIgencyMessage AIgencyMessage

// NewAIgencyMessage instantiates a new AIgencyMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIgencyMessage(aiModelId string, aiServiceId string, attachments AIgencyMessageFileList, channelId string, channelName string, content AIgencyMessageContentList, createdAt int64, id string, ownerOrganizationId string, senderConversationRole ConversationRole, senderId string, senderName string, tokenDirection TokenDirection, type_ AIgencyMessageType, updatedAt int64) *AIgencyMessage {
	this := AIgencyMessage{}
	this.AiModelId = aiModelId
	this.AiServiceId = aiServiceId
	this.Attachments = attachments
	this.ChannelId = channelId
	this.ChannelName = channelName
	this.Content = content
	this.CreatedAt = createdAt
	this.Id = id
	this.OwnerOrganizationId = ownerOrganizationId
	this.SenderConversationRole = senderConversationRole
	this.SenderId = senderId
	this.SenderName = senderName
	this.TokenDirection = tokenDirection
	this.Type = type_
	this.UpdatedAt = updatedAt
	return &this
}

// NewAIgencyMessageWithDefaults instantiates a new AIgencyMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIgencyMessageWithDefaults() *AIgencyMessage {
	this := AIgencyMessage{}
	return &this
}

// GetAiModelId returns the AiModelId field value
func (o *AIgencyMessage) GetAiModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AiModelId
}

// GetAiModelIdOk returns a tuple with the AiModelId field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetAiModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AiModelId, true
}

// SetAiModelId sets field value
func (o *AIgencyMessage) SetAiModelId(v string) {
	o.AiModelId = v
}

// GetAiServiceId returns the AiServiceId field value
func (o *AIgencyMessage) GetAiServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AiServiceId
}

// GetAiServiceIdOk returns a tuple with the AiServiceId field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetAiServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AiServiceId, true
}

// SetAiServiceId sets field value
func (o *AIgencyMessage) SetAiServiceId(v string) {
	o.AiServiceId = v
}

// GetAttachments returns the Attachments field value
func (o *AIgencyMessage) GetAttachments() AIgencyMessageFileList {
	if o == nil {
		var ret AIgencyMessageFileList
		return ret
	}

	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetAttachmentsOk() (*AIgencyMessageFileList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attachments, true
}

// SetAttachments sets field value
func (o *AIgencyMessage) SetAttachments(v AIgencyMessageFileList) {
	o.Attachments = v
}

// GetChannelId returns the ChannelId field value
func (o *AIgencyMessage) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *AIgencyMessage) SetChannelId(v string) {
	o.ChannelId = v
}

// GetChannelName returns the ChannelName field value
func (o *AIgencyMessage) GetChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelName, true
}

// SetChannelName sets field value
func (o *AIgencyMessage) SetChannelName(v string) {
	o.ChannelName = v
}

// GetContent returns the Content field value
func (o *AIgencyMessage) GetContent() AIgencyMessageContentList {
	if o == nil {
		var ret AIgencyMessageContentList
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetContentOk() (*AIgencyMessageContentList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *AIgencyMessage) SetContent(v AIgencyMessageContentList) {
	o.Content = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AIgencyMessage) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AIgencyMessage) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCreatedForFeature returns the CreatedForFeature field value if set, zero value otherwise.
func (o *AIgencyMessage) GetCreatedForFeature() AIModelFeature {
	if o == nil || IsNil(o.CreatedForFeature) {
		var ret AIModelFeature
		return ret
	}
	return *o.CreatedForFeature
}

// GetCreatedForFeatureOk returns a tuple with the CreatedForFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetCreatedForFeatureOk() (*AIModelFeature, bool) {
	if o == nil || IsNil(o.CreatedForFeature) {
		return nil, false
	}
	return o.CreatedForFeature, true
}

// HasCreatedForFeature returns a boolean if a field has been set.
func (o *AIgencyMessage) HasCreatedForFeature() bool {
	if o != nil && !IsNil(o.CreatedForFeature) {
		return true
	}

	return false
}

// SetCreatedForFeature gets a reference to the given AIModelFeature and assigns it to the CreatedForFeature field.
func (o *AIgencyMessage) SetCreatedForFeature(v AIModelFeature) {
	o.CreatedForFeature = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AIgencyMessage) GetError() AiServiceError {
	if o == nil || IsNil(o.Error) {
		var ret AiServiceError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetErrorOk() (*AiServiceError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AIgencyMessage) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given AiServiceError and assigns it to the Error field.
func (o *AIgencyMessage) SetError(v AiServiceError) {
	o.Error = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *AIgencyMessage) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId) {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *AIgencyMessage) HasExecutionId() bool {
	if o != nil && !IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *AIgencyMessage) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetFinishReason returns the FinishReason field value if set, zero value otherwise.
func (o *AIgencyMessage) GetFinishReason() FinishReason {
	if o == nil || IsNil(o.FinishReason) {
		var ret FinishReason
		return ret
	}
	return *o.FinishReason
}

// GetFinishReasonOk returns a tuple with the FinishReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetFinishReasonOk() (*FinishReason, bool) {
	if o == nil || IsNil(o.FinishReason) {
		return nil, false
	}
	return o.FinishReason, true
}

// HasFinishReason returns a boolean if a field has been set.
func (o *AIgencyMessage) HasFinishReason() bool {
	if o != nil && !IsNil(o.FinishReason) {
		return true
	}

	return false
}

// SetFinishReason gets a reference to the given FinishReason and assigns it to the FinishReason field.
func (o *AIgencyMessage) SetFinishReason(v FinishReason) {
	o.FinishReason = &v
}

// GetId returns the Id field value
func (o *AIgencyMessage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AIgencyMessage) SetId(v string) {
	o.Id = v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *AIgencyMessage) GetMetaData() map[string]interface{} {
	if o == nil || IsNil(o.MetaData) {
		var ret map[string]interface{}
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetaData) {
		return map[string]interface{}{}, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *AIgencyMessage) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *AIgencyMessage) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

// GetMissionId returns the MissionId field value if set, zero value otherwise.
func (o *AIgencyMessage) GetMissionId() string {
	if o == nil || IsNil(o.MissionId) {
		var ret string
		return ret
	}
	return *o.MissionId
}

// GetMissionIdOk returns a tuple with the MissionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetMissionIdOk() (*string, bool) {
	if o == nil || IsNil(o.MissionId) {
		return nil, false
	}
	return o.MissionId, true
}

// HasMissionId returns a boolean if a field has been set.
func (o *AIgencyMessage) HasMissionId() bool {
	if o != nil && !IsNil(o.MissionId) {
		return true
	}

	return false
}

// SetMissionId gets a reference to the given string and assigns it to the MissionId field.
func (o *AIgencyMessage) SetMissionId(v string) {
	o.MissionId = &v
}

// GetOwnerOrganizationId returns the OwnerOrganizationId field value
func (o *AIgencyMessage) GetOwnerOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerOrganizationId
}

// GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetOwnerOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerOrganizationId, true
}

// SetOwnerOrganizationId sets field value
func (o *AIgencyMessage) SetOwnerOrganizationId(v string) {
	o.OwnerOrganizationId = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *AIgencyMessage) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *AIgencyMessage) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *AIgencyMessage) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *AIgencyMessage) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *AIgencyMessage) HasReferenceId() bool {
	if o != nil && !IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *AIgencyMessage) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetResponseToId returns the ResponseToId field value if set, zero value otherwise.
func (o *AIgencyMessage) GetResponseToId() string {
	if o == nil || IsNil(o.ResponseToId) {
		var ret string
		return ret
	}
	return *o.ResponseToId
}

// GetResponseToIdOk returns a tuple with the ResponseToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetResponseToIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseToId) {
		return nil, false
	}
	return o.ResponseToId, true
}

// HasResponseToId returns a boolean if a field has been set.
func (o *AIgencyMessage) HasResponseToId() bool {
	if o != nil && !IsNil(o.ResponseToId) {
		return true
	}

	return false
}

// SetResponseToId gets a reference to the given string and assigns it to the ResponseToId field.
func (o *AIgencyMessage) SetResponseToId(v string) {
	o.ResponseToId = &v
}

// GetSenderConversationRole returns the SenderConversationRole field value
func (o *AIgencyMessage) GetSenderConversationRole() ConversationRole {
	if o == nil {
		var ret ConversationRole
		return ret
	}

	return o.SenderConversationRole
}

// GetSenderConversationRoleOk returns a tuple with the SenderConversationRole field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetSenderConversationRoleOk() (*ConversationRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderConversationRole, true
}

// SetSenderConversationRole sets field value
func (o *AIgencyMessage) SetSenderConversationRole(v ConversationRole) {
	o.SenderConversationRole = v
}

// GetSenderId returns the SenderId field value
func (o *AIgencyMessage) GetSenderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderId
}

// GetSenderIdOk returns a tuple with the SenderId field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetSenderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderId, true
}

// SetSenderId sets field value
func (o *AIgencyMessage) SetSenderId(v string) {
	o.SenderId = v
}

// GetSenderName returns the SenderName field value
func (o *AIgencyMessage) GetSenderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderName
}

// GetSenderNameOk returns a tuple with the SenderName field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetSenderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderName, true
}

// SetSenderName sets field value
func (o *AIgencyMessage) SetSenderName(v string) {
	o.SenderName = v
}

// GetTokenCount returns the TokenCount field value if set, zero value otherwise.
func (o *AIgencyMessage) GetTokenCount() int32 {
	if o == nil || IsNil(o.TokenCount) {
		var ret int32
		return ret
	}
	return *o.TokenCount
}

// GetTokenCountOk returns a tuple with the TokenCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetTokenCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TokenCount) {
		return nil, false
	}
	return o.TokenCount, true
}

// HasTokenCount returns a boolean if a field has been set.
func (o *AIgencyMessage) HasTokenCount() bool {
	if o != nil && !IsNil(o.TokenCount) {
		return true
	}

	return false
}

// SetTokenCount gets a reference to the given int32 and assigns it to the TokenCount field.
func (o *AIgencyMessage) SetTokenCount(v int32) {
	o.TokenCount = &v
}

// GetTokenDirection returns the TokenDirection field value
func (o *AIgencyMessage) GetTokenDirection() TokenDirection {
	if o == nil {
		var ret TokenDirection
		return ret
	}

	return o.TokenDirection
}

// GetTokenDirectionOk returns a tuple with the TokenDirection field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetTokenDirectionOk() (*TokenDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenDirection, true
}

// SetTokenDirection sets field value
func (o *AIgencyMessage) SetTokenDirection(v TokenDirection) {
	o.TokenDirection = v
}

// GetType returns the Type field value
func (o *AIgencyMessage) GetType() AIgencyMessageType {
	if o == nil {
		var ret AIgencyMessageType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetTypeOk() (*AIgencyMessageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AIgencyMessage) SetType(v AIgencyMessageType) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AIgencyMessage) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AIgencyMessage) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AIgencyMessage) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

func (o AIgencyMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIgencyMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ai_model_id"] = o.AiModelId
	toSerialize["ai_service_id"] = o.AiServiceId
	toSerialize["attachments"] = o.Attachments
	toSerialize["channel_id"] = o.ChannelId
	toSerialize["channel_name"] = o.ChannelName
	toSerialize["content"] = o.Content
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.CreatedForFeature) {
		toSerialize["created_for_feature"] = o.CreatedForFeature
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ExecutionId) {
		toSerialize["execution_id"] = o.ExecutionId
	}
	if !IsNil(o.FinishReason) {
		toSerialize["finish_reason"] = o.FinishReason
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.MetaData) {
		toSerialize["meta_data"] = o.MetaData
	}
	if !IsNil(o.MissionId) {
		toSerialize["mission_id"] = o.MissionId
	}
	toSerialize["owner_organization_id"] = o.OwnerOrganizationId
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if !IsNil(o.ResponseToId) {
		toSerialize["response_to_id"] = o.ResponseToId
	}
	toSerialize["sender_conversation_role"] = o.SenderConversationRole
	toSerialize["sender_id"] = o.SenderId
	toSerialize["sender_name"] = o.SenderName
	if !IsNil(o.TokenCount) {
		toSerialize["token_count"] = o.TokenCount
	}
	toSerialize["token_direction"] = o.TokenDirection
	toSerialize["type"] = o.Type
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIgencyMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ai_model_id",
		"ai_service_id",
		"attachments",
		"channel_id",
		"channel_name",
		"content",
		"created_at",
		"id",
		"owner_organization_id",
		"sender_conversation_role",
		"sender_id",
		"sender_name",
		"token_direction",
		"type",
		"updated_at",
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

	varAIgencyMessage := _AIgencyMessage{}

	err = json.Unmarshal(data, &varAIgencyMessage)

	if err != nil {
		return err
	}

	*o = AIgencyMessage(varAIgencyMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ai_model_id")
		delete(additionalProperties, "ai_service_id")
		delete(additionalProperties, "attachments")
		delete(additionalProperties, "channel_id")
		delete(additionalProperties, "channel_name")
		delete(additionalProperties, "content")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "created_for_feature")
		delete(additionalProperties, "error")
		delete(additionalProperties, "execution_id")
		delete(additionalProperties, "finish_reason")
		delete(additionalProperties, "id")
		delete(additionalProperties, "meta_data")
		delete(additionalProperties, "mission_id")
		delete(additionalProperties, "owner_organization_id")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "reference_id")
		delete(additionalProperties, "response_to_id")
		delete(additionalProperties, "sender_conversation_role")
		delete(additionalProperties, "sender_id")
		delete(additionalProperties, "sender_name")
		delete(additionalProperties, "token_count")
		delete(additionalProperties, "token_direction")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIgencyMessage struct {
	value *AIgencyMessage
	isSet bool
}

func (v NullableAIgencyMessage) Get() *AIgencyMessage {
	return v.value
}

func (v *NullableAIgencyMessage) Set(val *AIgencyMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableAIgencyMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableAIgencyMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIgencyMessage(val *AIgencyMessage) *NullableAIgencyMessage {
	return &NullableAIgencyMessage{value: val, isSet: true}
}

func (v NullableAIgencyMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIgencyMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


