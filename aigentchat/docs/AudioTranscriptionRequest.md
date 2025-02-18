# AudioTranscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**AudioFile** | Pointer to [**AIgencyMessageFile**](AIgencyMessageFile.md) |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Language** | **string** |  | 
**MessageReferenceId** | Pointer to **string** |  | [optional] 
**MessageResponseToId** | Pointer to **string** |  | [optional] 
**MissionId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**SkipAiAgencyMessageGeneration** | Pointer to **bool** |  | [optional] 
**TeamIds** | Pointer to **[]string** |  | [optional] 
**TranscriptionFormat** | **string** |  | 
**TriggerChatCompletion** | Pointer to **bool** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewAudioTranscriptionRequest

`func NewAudioTranscriptionRequest(language string, transcriptionFormat string, ) *AudioTranscriptionRequest`

NewAudioTranscriptionRequest instantiates a new AudioTranscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioTranscriptionRequestWithDefaults

`func NewAudioTranscriptionRequestWithDefaults() *AudioTranscriptionRequest`

NewAudioTranscriptionRequestWithDefaults instantiates a new AudioTranscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *AudioTranscriptionRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AudioTranscriptionRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AudioTranscriptionRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *AudioTranscriptionRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAudioFile

`func (o *AudioTranscriptionRequest) GetAudioFile() AIgencyMessageFile`

GetAudioFile returns the AudioFile field if non-nil, zero value otherwise.

### GetAudioFileOk

`func (o *AudioTranscriptionRequest) GetAudioFileOk() (*AIgencyMessageFile, bool)`

GetAudioFileOk returns a tuple with the AudioFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioFile

`func (o *AudioTranscriptionRequest) SetAudioFile(v AIgencyMessageFile)`

SetAudioFile sets AudioFile field to given value.

### HasAudioFile

`func (o *AudioTranscriptionRequest) HasAudioFile() bool`

HasAudioFile returns a boolean if a field has been set.

### GetChannelId

`func (o *AudioTranscriptionRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *AudioTranscriptionRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *AudioTranscriptionRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *AudioTranscriptionRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetId

`func (o *AudioTranscriptionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudioTranscriptionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudioTranscriptionRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AudioTranscriptionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *AudioTranscriptionRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AudioTranscriptionRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AudioTranscriptionRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetMessageReferenceId

`func (o *AudioTranscriptionRequest) GetMessageReferenceId() string`

GetMessageReferenceId returns the MessageReferenceId field if non-nil, zero value otherwise.

### GetMessageReferenceIdOk

`func (o *AudioTranscriptionRequest) GetMessageReferenceIdOk() (*string, bool)`

GetMessageReferenceIdOk returns a tuple with the MessageReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReferenceId

`func (o *AudioTranscriptionRequest) SetMessageReferenceId(v string)`

SetMessageReferenceId sets MessageReferenceId field to given value.

### HasMessageReferenceId

`func (o *AudioTranscriptionRequest) HasMessageReferenceId() bool`

HasMessageReferenceId returns a boolean if a field has been set.

### GetMessageResponseToId

`func (o *AudioTranscriptionRequest) GetMessageResponseToId() string`

GetMessageResponseToId returns the MessageResponseToId field if non-nil, zero value otherwise.

### GetMessageResponseToIdOk

`func (o *AudioTranscriptionRequest) GetMessageResponseToIdOk() (*string, bool)`

GetMessageResponseToIdOk returns a tuple with the MessageResponseToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageResponseToId

`func (o *AudioTranscriptionRequest) SetMessageResponseToId(v string)`

SetMessageResponseToId sets MessageResponseToId field to given value.

### HasMessageResponseToId

`func (o *AudioTranscriptionRequest) HasMessageResponseToId() bool`

HasMessageResponseToId returns a boolean if a field has been set.

### GetMissionId

`func (o *AudioTranscriptionRequest) GetMissionId() string`

GetMissionId returns the MissionId field if non-nil, zero value otherwise.

### GetMissionIdOk

`func (o *AudioTranscriptionRequest) GetMissionIdOk() (*string, bool)`

GetMissionIdOk returns a tuple with the MissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionId

`func (o *AudioTranscriptionRequest) SetMissionId(v string)`

SetMissionId sets MissionId field to given value.

### HasMissionId

`func (o *AudioTranscriptionRequest) HasMissionId() bool`

HasMissionId returns a boolean if a field has been set.

### GetOrgId

`func (o *AudioTranscriptionRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AudioTranscriptionRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AudioTranscriptionRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AudioTranscriptionRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetParameters

`func (o *AudioTranscriptionRequest) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AudioTranscriptionRequest) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AudioTranscriptionRequest) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AudioTranscriptionRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSkipAiAgencyMessageGeneration

`func (o *AudioTranscriptionRequest) GetSkipAiAgencyMessageGeneration() bool`

GetSkipAiAgencyMessageGeneration returns the SkipAiAgencyMessageGeneration field if non-nil, zero value otherwise.

### GetSkipAiAgencyMessageGenerationOk

`func (o *AudioTranscriptionRequest) GetSkipAiAgencyMessageGenerationOk() (*bool, bool)`

GetSkipAiAgencyMessageGenerationOk returns a tuple with the SkipAiAgencyMessageGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipAiAgencyMessageGeneration

`func (o *AudioTranscriptionRequest) SetSkipAiAgencyMessageGeneration(v bool)`

SetSkipAiAgencyMessageGeneration sets SkipAiAgencyMessageGeneration field to given value.

### HasSkipAiAgencyMessageGeneration

`func (o *AudioTranscriptionRequest) HasSkipAiAgencyMessageGeneration() bool`

HasSkipAiAgencyMessageGeneration returns a boolean if a field has been set.

### GetTeamIds

`func (o *AudioTranscriptionRequest) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *AudioTranscriptionRequest) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *AudioTranscriptionRequest) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *AudioTranscriptionRequest) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetTranscriptionFormat

`func (o *AudioTranscriptionRequest) GetTranscriptionFormat() string`

GetTranscriptionFormat returns the TranscriptionFormat field if non-nil, zero value otherwise.

### GetTranscriptionFormatOk

`func (o *AudioTranscriptionRequest) GetTranscriptionFormatOk() (*string, bool)`

GetTranscriptionFormatOk returns a tuple with the TranscriptionFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionFormat

`func (o *AudioTranscriptionRequest) SetTranscriptionFormat(v string)`

SetTranscriptionFormat sets TranscriptionFormat field to given value.


### GetTriggerChatCompletion

`func (o *AudioTranscriptionRequest) GetTriggerChatCompletion() bool`

GetTriggerChatCompletion returns the TriggerChatCompletion field if non-nil, zero value otherwise.

### GetTriggerChatCompletionOk

`func (o *AudioTranscriptionRequest) GetTriggerChatCompletionOk() (*bool, bool)`

GetTriggerChatCompletionOk returns a tuple with the TriggerChatCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerChatCompletion

`func (o *AudioTranscriptionRequest) SetTriggerChatCompletion(v bool)`

SetTriggerChatCompletion sets TriggerChatCompletion field to given value.

### HasTriggerChatCompletion

`func (o *AudioTranscriptionRequest) HasTriggerChatCompletion() bool`

HasTriggerChatCompletion returns a boolean if a field has been set.

### GetUserId

`func (o *AudioTranscriptionRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AudioTranscriptionRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AudioTranscriptionRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AudioTranscriptionRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *AudioTranscriptionRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AudioTranscriptionRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AudioTranscriptionRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AudioTranscriptionRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


