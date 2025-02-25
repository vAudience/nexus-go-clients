/*
vAudience AIgentChat API

chat and api server for AIgents

API version: 0.15.12
Contact: contact@vaudience.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aigentchat

import (
	"encoding/json"
	"fmt"
)

// checks if the AudioTranscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudioTranscriptionRequest{}

// AudioTranscriptionRequest struct for AudioTranscriptionRequest
type AudioTranscriptionRequest struct {
	AgentId *string `json:"agent_id,omitempty"`
	AudioFile *AIgencyMessageFile `json:"audio_file,omitempty"`
	ChannelId *string `json:"channel_id,omitempty"`
	Id *string `json:"id,omitempty"`
	Language string `json:"language"`
	MessageReferenceId *string `json:"message_reference_id,omitempty"`
	MessageResponseToId *string `json:"message_response_to_id,omitempty"`
	MissionId *string `json:"mission_id,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	SkipAiAgencyMessageGeneration *bool `json:"skip_ai_agency_message_generation,omitempty"`
	TeamIds []string `json:"team_ids,omitempty"`
	TranscriptionFormat string `json:"transcription_format"`
	TriggerChatCompletion *bool `json:"trigger_chat_completion,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AudioTranscriptionRequest AudioTranscriptionRequest

// NewAudioTranscriptionRequest instantiates a new AudioTranscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudioTranscriptionRequest(language string, transcriptionFormat string) *AudioTranscriptionRequest {
	this := AudioTranscriptionRequest{}
	this.Language = language
	this.TranscriptionFormat = transcriptionFormat
	return &this
}

// NewAudioTranscriptionRequestWithDefaults instantiates a new AudioTranscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudioTranscriptionRequestWithDefaults() *AudioTranscriptionRequest {
	this := AudioTranscriptionRequest{}
	return &this
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetAgentId() string {
	if o == nil || IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetAgentIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasAgentId() bool {
	if o != nil && !IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *AudioTranscriptionRequest) SetAgentId(v string) {
	o.AgentId = &v
}

// GetAudioFile returns the AudioFile field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetAudioFile() AIgencyMessageFile {
	if o == nil || IsNil(o.AudioFile) {
		var ret AIgencyMessageFile
		return ret
	}
	return *o.AudioFile
}

// GetAudioFileOk returns a tuple with the AudioFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetAudioFileOk() (*AIgencyMessageFile, bool) {
	if o == nil || IsNil(o.AudioFile) {
		return nil, false
	}
	return o.AudioFile, true
}

// HasAudioFile returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasAudioFile() bool {
	if o != nil && !IsNil(o.AudioFile) {
		return true
	}

	return false
}

// SetAudioFile gets a reference to the given AIgencyMessageFile and assigns it to the AudioFile field.
func (o *AudioTranscriptionRequest) SetAudioFile(v AIgencyMessageFile) {
	o.AudioFile = &v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *AudioTranscriptionRequest) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AudioTranscriptionRequest) SetId(v string) {
	o.Id = &v
}

// GetLanguage returns the Language field value
func (o *AudioTranscriptionRequest) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *AudioTranscriptionRequest) SetLanguage(v string) {
	o.Language = v
}

// GetMessageReferenceId returns the MessageReferenceId field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetMessageReferenceId() string {
	if o == nil || IsNil(o.MessageReferenceId) {
		var ret string
		return ret
	}
	return *o.MessageReferenceId
}

// GetMessageReferenceIdOk returns a tuple with the MessageReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetMessageReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageReferenceId) {
		return nil, false
	}
	return o.MessageReferenceId, true
}

// HasMessageReferenceId returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasMessageReferenceId() bool {
	if o != nil && !IsNil(o.MessageReferenceId) {
		return true
	}

	return false
}

// SetMessageReferenceId gets a reference to the given string and assigns it to the MessageReferenceId field.
func (o *AudioTranscriptionRequest) SetMessageReferenceId(v string) {
	o.MessageReferenceId = &v
}

// GetMessageResponseToId returns the MessageResponseToId field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetMessageResponseToId() string {
	if o == nil || IsNil(o.MessageResponseToId) {
		var ret string
		return ret
	}
	return *o.MessageResponseToId
}

// GetMessageResponseToIdOk returns a tuple with the MessageResponseToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetMessageResponseToIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageResponseToId) {
		return nil, false
	}
	return o.MessageResponseToId, true
}

// HasMessageResponseToId returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasMessageResponseToId() bool {
	if o != nil && !IsNil(o.MessageResponseToId) {
		return true
	}

	return false
}

// SetMessageResponseToId gets a reference to the given string and assigns it to the MessageResponseToId field.
func (o *AudioTranscriptionRequest) SetMessageResponseToId(v string) {
	o.MessageResponseToId = &v
}

// GetMissionId returns the MissionId field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetMissionId() string {
	if o == nil || IsNil(o.MissionId) {
		var ret string
		return ret
	}
	return *o.MissionId
}

// GetMissionIdOk returns a tuple with the MissionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetMissionIdOk() (*string, bool) {
	if o == nil || IsNil(o.MissionId) {
		return nil, false
	}
	return o.MissionId, true
}

// HasMissionId returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasMissionId() bool {
	if o != nil && !IsNil(o.MissionId) {
		return true
	}

	return false
}

// SetMissionId gets a reference to the given string and assigns it to the MissionId field.
func (o *AudioTranscriptionRequest) SetMissionId(v string) {
	o.MissionId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *AudioTranscriptionRequest) SetOrgId(v string) {
	o.OrgId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *AudioTranscriptionRequest) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetSkipAiAgencyMessageGeneration returns the SkipAiAgencyMessageGeneration field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetSkipAiAgencyMessageGeneration() bool {
	if o == nil || IsNil(o.SkipAiAgencyMessageGeneration) {
		var ret bool
		return ret
	}
	return *o.SkipAiAgencyMessageGeneration
}

// GetSkipAiAgencyMessageGenerationOk returns a tuple with the SkipAiAgencyMessageGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetSkipAiAgencyMessageGenerationOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipAiAgencyMessageGeneration) {
		return nil, false
	}
	return o.SkipAiAgencyMessageGeneration, true
}

// HasSkipAiAgencyMessageGeneration returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasSkipAiAgencyMessageGeneration() bool {
	if o != nil && !IsNil(o.SkipAiAgencyMessageGeneration) {
		return true
	}

	return false
}

// SetSkipAiAgencyMessageGeneration gets a reference to the given bool and assigns it to the SkipAiAgencyMessageGeneration field.
func (o *AudioTranscriptionRequest) SetSkipAiAgencyMessageGeneration(v bool) {
	o.SkipAiAgencyMessageGeneration = &v
}

// GetTeamIds returns the TeamIds field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetTeamIds() []string {
	if o == nil || IsNil(o.TeamIds) {
		var ret []string
		return ret
	}
	return o.TeamIds
}

// GetTeamIdsOk returns a tuple with the TeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetTeamIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TeamIds) {
		return nil, false
	}
	return o.TeamIds, true
}

// HasTeamIds returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasTeamIds() bool {
	if o != nil && !IsNil(o.TeamIds) {
		return true
	}

	return false
}

// SetTeamIds gets a reference to the given []string and assigns it to the TeamIds field.
func (o *AudioTranscriptionRequest) SetTeamIds(v []string) {
	o.TeamIds = v
}

// GetTranscriptionFormat returns the TranscriptionFormat field value
func (o *AudioTranscriptionRequest) GetTranscriptionFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TranscriptionFormat
}

// GetTranscriptionFormatOk returns a tuple with the TranscriptionFormat field value
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetTranscriptionFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TranscriptionFormat, true
}

// SetTranscriptionFormat sets field value
func (o *AudioTranscriptionRequest) SetTranscriptionFormat(v string) {
	o.TranscriptionFormat = v
}

// GetTriggerChatCompletion returns the TriggerChatCompletion field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetTriggerChatCompletion() bool {
	if o == nil || IsNil(o.TriggerChatCompletion) {
		var ret bool
		return ret
	}
	return *o.TriggerChatCompletion
}

// GetTriggerChatCompletionOk returns a tuple with the TriggerChatCompletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetTriggerChatCompletionOk() (*bool, bool) {
	if o == nil || IsNil(o.TriggerChatCompletion) {
		return nil, false
	}
	return o.TriggerChatCompletion, true
}

// HasTriggerChatCompletion returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasTriggerChatCompletion() bool {
	if o != nil && !IsNil(o.TriggerChatCompletion) {
		return true
	}

	return false
}

// SetTriggerChatCompletion gets a reference to the given bool and assigns it to the TriggerChatCompletion field.
func (o *AudioTranscriptionRequest) SetTriggerChatCompletion(v bool) {
	o.TriggerChatCompletion = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AudioTranscriptionRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AudioTranscriptionRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudioTranscriptionRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AudioTranscriptionRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AudioTranscriptionRequest) SetUsername(v string) {
	o.Username = &v
}

func (o AudioTranscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudioTranscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentId) {
		toSerialize["agent_id"] = o.AgentId
	}
	if !IsNil(o.AudioFile) {
		toSerialize["audio_file"] = o.AudioFile
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["language"] = o.Language
	if !IsNil(o.MessageReferenceId) {
		toSerialize["message_reference_id"] = o.MessageReferenceId
	}
	if !IsNil(o.MessageResponseToId) {
		toSerialize["message_response_to_id"] = o.MessageResponseToId
	}
	if !IsNil(o.MissionId) {
		toSerialize["mission_id"] = o.MissionId
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.SkipAiAgencyMessageGeneration) {
		toSerialize["skip_ai_agency_message_generation"] = o.SkipAiAgencyMessageGeneration
	}
	if !IsNil(o.TeamIds) {
		toSerialize["team_ids"] = o.TeamIds
	}
	toSerialize["transcription_format"] = o.TranscriptionFormat
	if !IsNil(o.TriggerChatCompletion) {
		toSerialize["trigger_chat_completion"] = o.TriggerChatCompletion
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AudioTranscriptionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"language",
		"transcription_format",
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

	varAudioTranscriptionRequest := _AudioTranscriptionRequest{}

	err = json.Unmarshal(data, &varAudioTranscriptionRequest)

	if err != nil {
		return err
	}

	*o = AudioTranscriptionRequest(varAudioTranscriptionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "agent_id")
		delete(additionalProperties, "audio_file")
		delete(additionalProperties, "channel_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "language")
		delete(additionalProperties, "message_reference_id")
		delete(additionalProperties, "message_response_to_id")
		delete(additionalProperties, "mission_id")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "skip_ai_agency_message_generation")
		delete(additionalProperties, "team_ids")
		delete(additionalProperties, "transcription_format")
		delete(additionalProperties, "trigger_chat_completion")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAudioTranscriptionRequest struct {
	value *AudioTranscriptionRequest
	isSet bool
}

func (v NullableAudioTranscriptionRequest) Get() *AudioTranscriptionRequest {
	return v.value
}

func (v *NullableAudioTranscriptionRequest) Set(val *AudioTranscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAudioTranscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAudioTranscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudioTranscriptionRequest(val *AudioTranscriptionRequest) *NullableAudioTranscriptionRequest {
	return &NullableAudioTranscriptionRequest{value: val, isSet: true}
}

func (v NullableAudioTranscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudioTranscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


