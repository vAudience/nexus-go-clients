# AudioGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**AttachedFiles** | Pointer to **[]string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**Dictionary** | Pointer to **map[string]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**MessageReferenceId** | Pointer to **string** |  | [optional] 
**MessageResponseToId** | Pointer to **string** |  | [optional] 
**MissionId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**OutputAudioFileFormat** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Speed** | Pointer to **float64** |  | [optional] 
**StreamAudio** | Pointer to **bool** |  | [optional] 
**TeamIds** | Pointer to **[]string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**VoiceId** | Pointer to **string** |  | [optional] 
**VoiceName** | Pointer to **string** |  | [optional] 
**WithTimestamps** | Pointer to **bool** |  | [optional] 

## Methods

### NewAudioGenerationRequest

`func NewAudioGenerationRequest() *AudioGenerationRequest`

NewAudioGenerationRequest instantiates a new AudioGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioGenerationRequestWithDefaults

`func NewAudioGenerationRequestWithDefaults() *AudioGenerationRequest`

NewAudioGenerationRequestWithDefaults instantiates a new AudioGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *AudioGenerationRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AudioGenerationRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AudioGenerationRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *AudioGenerationRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAttachedFiles

`func (o *AudioGenerationRequest) GetAttachedFiles() []string`

GetAttachedFiles returns the AttachedFiles field if non-nil, zero value otherwise.

### GetAttachedFilesOk

`func (o *AudioGenerationRequest) GetAttachedFilesOk() (*[]string, bool)`

GetAttachedFilesOk returns a tuple with the AttachedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedFiles

`func (o *AudioGenerationRequest) SetAttachedFiles(v []string)`

SetAttachedFiles sets AttachedFiles field to given value.

### HasAttachedFiles

`func (o *AudioGenerationRequest) HasAttachedFiles() bool`

HasAttachedFiles returns a boolean if a field has been set.

### GetChannelId

`func (o *AudioGenerationRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *AudioGenerationRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *AudioGenerationRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *AudioGenerationRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetDictionary

`func (o *AudioGenerationRequest) GetDictionary() map[string]string`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *AudioGenerationRequest) GetDictionaryOk() (*map[string]string, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *AudioGenerationRequest) SetDictionary(v map[string]string)`

SetDictionary sets Dictionary field to given value.

### HasDictionary

`func (o *AudioGenerationRequest) HasDictionary() bool`

HasDictionary returns a boolean if a field has been set.

### GetId

`func (o *AudioGenerationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudioGenerationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudioGenerationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AudioGenerationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *AudioGenerationRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AudioGenerationRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AudioGenerationRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AudioGenerationRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMessage

`func (o *AudioGenerationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AudioGenerationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AudioGenerationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AudioGenerationRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageReferenceId

`func (o *AudioGenerationRequest) GetMessageReferenceId() string`

GetMessageReferenceId returns the MessageReferenceId field if non-nil, zero value otherwise.

### GetMessageReferenceIdOk

`func (o *AudioGenerationRequest) GetMessageReferenceIdOk() (*string, bool)`

GetMessageReferenceIdOk returns a tuple with the MessageReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReferenceId

`func (o *AudioGenerationRequest) SetMessageReferenceId(v string)`

SetMessageReferenceId sets MessageReferenceId field to given value.

### HasMessageReferenceId

`func (o *AudioGenerationRequest) HasMessageReferenceId() bool`

HasMessageReferenceId returns a boolean if a field has been set.

### GetMessageResponseToId

`func (o *AudioGenerationRequest) GetMessageResponseToId() string`

GetMessageResponseToId returns the MessageResponseToId field if non-nil, zero value otherwise.

### GetMessageResponseToIdOk

`func (o *AudioGenerationRequest) GetMessageResponseToIdOk() (*string, bool)`

GetMessageResponseToIdOk returns a tuple with the MessageResponseToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageResponseToId

`func (o *AudioGenerationRequest) SetMessageResponseToId(v string)`

SetMessageResponseToId sets MessageResponseToId field to given value.

### HasMessageResponseToId

`func (o *AudioGenerationRequest) HasMessageResponseToId() bool`

HasMessageResponseToId returns a boolean if a field has been set.

### GetMissionId

`func (o *AudioGenerationRequest) GetMissionId() string`

GetMissionId returns the MissionId field if non-nil, zero value otherwise.

### GetMissionIdOk

`func (o *AudioGenerationRequest) GetMissionIdOk() (*string, bool)`

GetMissionIdOk returns a tuple with the MissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionId

`func (o *AudioGenerationRequest) SetMissionId(v string)`

SetMissionId sets MissionId field to given value.

### HasMissionId

`func (o *AudioGenerationRequest) HasMissionId() bool`

HasMissionId returns a boolean if a field has been set.

### GetOrgId

`func (o *AudioGenerationRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AudioGenerationRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AudioGenerationRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AudioGenerationRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOutputAudioFileFormat

`func (o *AudioGenerationRequest) GetOutputAudioFileFormat() string`

GetOutputAudioFileFormat returns the OutputAudioFileFormat field if non-nil, zero value otherwise.

### GetOutputAudioFileFormatOk

`func (o *AudioGenerationRequest) GetOutputAudioFileFormatOk() (*string, bool)`

GetOutputAudioFileFormatOk returns a tuple with the OutputAudioFileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputAudioFileFormat

`func (o *AudioGenerationRequest) SetOutputAudioFileFormat(v string)`

SetOutputAudioFileFormat sets OutputAudioFileFormat field to given value.

### HasOutputAudioFileFormat

`func (o *AudioGenerationRequest) HasOutputAudioFileFormat() bool`

HasOutputAudioFileFormat returns a boolean if a field has been set.

### GetParameters

`func (o *AudioGenerationRequest) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AudioGenerationRequest) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AudioGenerationRequest) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AudioGenerationRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSpeed

`func (o *AudioGenerationRequest) GetSpeed() float64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *AudioGenerationRequest) GetSpeedOk() (*float64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *AudioGenerationRequest) SetSpeed(v float64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *AudioGenerationRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStreamAudio

`func (o *AudioGenerationRequest) GetStreamAudio() bool`

GetStreamAudio returns the StreamAudio field if non-nil, zero value otherwise.

### GetStreamAudioOk

`func (o *AudioGenerationRequest) GetStreamAudioOk() (*bool, bool)`

GetStreamAudioOk returns a tuple with the StreamAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamAudio

`func (o *AudioGenerationRequest) SetStreamAudio(v bool)`

SetStreamAudio sets StreamAudio field to given value.

### HasStreamAudio

`func (o *AudioGenerationRequest) HasStreamAudio() bool`

HasStreamAudio returns a boolean if a field has been set.

### GetTeamIds

`func (o *AudioGenerationRequest) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *AudioGenerationRequest) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *AudioGenerationRequest) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *AudioGenerationRequest) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetUserId

`func (o *AudioGenerationRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AudioGenerationRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AudioGenerationRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AudioGenerationRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *AudioGenerationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AudioGenerationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AudioGenerationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AudioGenerationRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVoiceId

`func (o *AudioGenerationRequest) GetVoiceId() string`

GetVoiceId returns the VoiceId field if non-nil, zero value otherwise.

### GetVoiceIdOk

`func (o *AudioGenerationRequest) GetVoiceIdOk() (*string, bool)`

GetVoiceIdOk returns a tuple with the VoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceId

`func (o *AudioGenerationRequest) SetVoiceId(v string)`

SetVoiceId sets VoiceId field to given value.

### HasVoiceId

`func (o *AudioGenerationRequest) HasVoiceId() bool`

HasVoiceId returns a boolean if a field has been set.

### GetVoiceName

`func (o *AudioGenerationRequest) GetVoiceName() string`

GetVoiceName returns the VoiceName field if non-nil, zero value otherwise.

### GetVoiceNameOk

`func (o *AudioGenerationRequest) GetVoiceNameOk() (*string, bool)`

GetVoiceNameOk returns a tuple with the VoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceName

`func (o *AudioGenerationRequest) SetVoiceName(v string)`

SetVoiceName sets VoiceName field to given value.

### HasVoiceName

`func (o *AudioGenerationRequest) HasVoiceName() bool`

HasVoiceName returns a boolean if a field has been set.

### GetWithTimestamps

`func (o *AudioGenerationRequest) GetWithTimestamps() bool`

GetWithTimestamps returns the WithTimestamps field if non-nil, zero value otherwise.

### GetWithTimestampsOk

`func (o *AudioGenerationRequest) GetWithTimestampsOk() (*bool, bool)`

GetWithTimestampsOk returns a tuple with the WithTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithTimestamps

`func (o *AudioGenerationRequest) SetWithTimestamps(v bool)`

SetWithTimestamps sets WithTimestamps field to given value.

### HasWithTimestamps

`func (o *AudioGenerationRequest) HasWithTimestamps() bool`

HasWithTimestamps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


