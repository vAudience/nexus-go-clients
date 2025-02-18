# ImageGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**AttachedTemporaryFiles** | Pointer to **[]string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**MessageReferenceId** | Pointer to **string** |  | [optional] 
**MessageResponseToId** | Pointer to **string** |  | [optional] 
**MissionId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**OutputImageFileFormat** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**TeamIds** | Pointer to **[]string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewImageGenerationRequest

`func NewImageGenerationRequest() *ImageGenerationRequest`

NewImageGenerationRequest instantiates a new ImageGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageGenerationRequestWithDefaults

`func NewImageGenerationRequestWithDefaults() *ImageGenerationRequest`

NewImageGenerationRequestWithDefaults instantiates a new ImageGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *ImageGenerationRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ImageGenerationRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ImageGenerationRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ImageGenerationRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAttachedTemporaryFiles

`func (o *ImageGenerationRequest) GetAttachedTemporaryFiles() []string`

GetAttachedTemporaryFiles returns the AttachedTemporaryFiles field if non-nil, zero value otherwise.

### GetAttachedTemporaryFilesOk

`func (o *ImageGenerationRequest) GetAttachedTemporaryFilesOk() (*[]string, bool)`

GetAttachedTemporaryFilesOk returns a tuple with the AttachedTemporaryFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedTemporaryFiles

`func (o *ImageGenerationRequest) SetAttachedTemporaryFiles(v []string)`

SetAttachedTemporaryFiles sets AttachedTemporaryFiles field to given value.

### HasAttachedTemporaryFiles

`func (o *ImageGenerationRequest) HasAttachedTemporaryFiles() bool`

HasAttachedTemporaryFiles returns a boolean if a field has been set.

### GetChannelId

`func (o *ImageGenerationRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ImageGenerationRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ImageGenerationRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ImageGenerationRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetId

`func (o *ImageGenerationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageGenerationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageGenerationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageGenerationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *ImageGenerationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ImageGenerationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ImageGenerationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ImageGenerationRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageReferenceId

`func (o *ImageGenerationRequest) GetMessageReferenceId() string`

GetMessageReferenceId returns the MessageReferenceId field if non-nil, zero value otherwise.

### GetMessageReferenceIdOk

`func (o *ImageGenerationRequest) GetMessageReferenceIdOk() (*string, bool)`

GetMessageReferenceIdOk returns a tuple with the MessageReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReferenceId

`func (o *ImageGenerationRequest) SetMessageReferenceId(v string)`

SetMessageReferenceId sets MessageReferenceId field to given value.

### HasMessageReferenceId

`func (o *ImageGenerationRequest) HasMessageReferenceId() bool`

HasMessageReferenceId returns a boolean if a field has been set.

### GetMessageResponseToId

`func (o *ImageGenerationRequest) GetMessageResponseToId() string`

GetMessageResponseToId returns the MessageResponseToId field if non-nil, zero value otherwise.

### GetMessageResponseToIdOk

`func (o *ImageGenerationRequest) GetMessageResponseToIdOk() (*string, bool)`

GetMessageResponseToIdOk returns a tuple with the MessageResponseToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageResponseToId

`func (o *ImageGenerationRequest) SetMessageResponseToId(v string)`

SetMessageResponseToId sets MessageResponseToId field to given value.

### HasMessageResponseToId

`func (o *ImageGenerationRequest) HasMessageResponseToId() bool`

HasMessageResponseToId returns a boolean if a field has been set.

### GetMissionId

`func (o *ImageGenerationRequest) GetMissionId() string`

GetMissionId returns the MissionId field if non-nil, zero value otherwise.

### GetMissionIdOk

`func (o *ImageGenerationRequest) GetMissionIdOk() (*string, bool)`

GetMissionIdOk returns a tuple with the MissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionId

`func (o *ImageGenerationRequest) SetMissionId(v string)`

SetMissionId sets MissionId field to given value.

### HasMissionId

`func (o *ImageGenerationRequest) HasMissionId() bool`

HasMissionId returns a boolean if a field has been set.

### GetOrgId

`func (o *ImageGenerationRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ImageGenerationRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ImageGenerationRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ImageGenerationRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOutputImageFileFormat

`func (o *ImageGenerationRequest) GetOutputImageFileFormat() string`

GetOutputImageFileFormat returns the OutputImageFileFormat field if non-nil, zero value otherwise.

### GetOutputImageFileFormatOk

`func (o *ImageGenerationRequest) GetOutputImageFileFormatOk() (*string, bool)`

GetOutputImageFileFormatOk returns a tuple with the OutputImageFileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputImageFileFormat

`func (o *ImageGenerationRequest) SetOutputImageFileFormat(v string)`

SetOutputImageFileFormat sets OutputImageFileFormat field to given value.

### HasOutputImageFileFormat

`func (o *ImageGenerationRequest) HasOutputImageFileFormat() bool`

HasOutputImageFileFormat returns a boolean if a field has been set.

### GetParameters

`func (o *ImageGenerationRequest) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ImageGenerationRequest) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ImageGenerationRequest) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ImageGenerationRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetTeamIds

`func (o *ImageGenerationRequest) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *ImageGenerationRequest) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *ImageGenerationRequest) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *ImageGenerationRequest) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetUserId

`func (o *ImageGenerationRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ImageGenerationRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ImageGenerationRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ImageGenerationRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *ImageGenerationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ImageGenerationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ImageGenerationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ImageGenerationRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


