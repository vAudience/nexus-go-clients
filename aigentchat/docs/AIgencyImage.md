# AIgencyImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int32** |  | [optional] 
**LocalFilePath** | Pointer to **string** |  | [optional] 
**MessageReferenceId** | Pointer to **string** |  | [optional] 
**MessageResponseToId** | Pointer to **string** |  | [optional] 
**MetaData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Mimetype** | Pointer to **string** |  | [optional] 
**MissionId** | Pointer to **string** |  | [optional] 
**OrgOwnerId** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**ProcessingErrors** | Pointer to **[]string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewAIgencyImage

`func NewAIgencyImage() *AIgencyImage`

NewAIgencyImage instantiates a new AIgencyImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIgencyImageWithDefaults

`func NewAIgencyImageWithDefaults() *AIgencyImage`

NewAIgencyImageWithDefaults instantiates a new AIgencyImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *AIgencyImage) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AIgencyImage) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AIgencyImage) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *AIgencyImage) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetChannelId

`func (o *AIgencyImage) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *AIgencyImage) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *AIgencyImage) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *AIgencyImage) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetFileName

`func (o *AIgencyImage) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AIgencyImage) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AIgencyImage) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AIgencyImage) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *AIgencyImage) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *AIgencyImage) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *AIgencyImage) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *AIgencyImage) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetLocalFilePath

`func (o *AIgencyImage) GetLocalFilePath() string`

GetLocalFilePath returns the LocalFilePath field if non-nil, zero value otherwise.

### GetLocalFilePathOk

`func (o *AIgencyImage) GetLocalFilePathOk() (*string, bool)`

GetLocalFilePathOk returns a tuple with the LocalFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalFilePath

`func (o *AIgencyImage) SetLocalFilePath(v string)`

SetLocalFilePath sets LocalFilePath field to given value.

### HasLocalFilePath

`func (o *AIgencyImage) HasLocalFilePath() bool`

HasLocalFilePath returns a boolean if a field has been set.

### GetMessageReferenceId

`func (o *AIgencyImage) GetMessageReferenceId() string`

GetMessageReferenceId returns the MessageReferenceId field if non-nil, zero value otherwise.

### GetMessageReferenceIdOk

`func (o *AIgencyImage) GetMessageReferenceIdOk() (*string, bool)`

GetMessageReferenceIdOk returns a tuple with the MessageReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReferenceId

`func (o *AIgencyImage) SetMessageReferenceId(v string)`

SetMessageReferenceId sets MessageReferenceId field to given value.

### HasMessageReferenceId

`func (o *AIgencyImage) HasMessageReferenceId() bool`

HasMessageReferenceId returns a boolean if a field has been set.

### GetMessageResponseToId

`func (o *AIgencyImage) GetMessageResponseToId() string`

GetMessageResponseToId returns the MessageResponseToId field if non-nil, zero value otherwise.

### GetMessageResponseToIdOk

`func (o *AIgencyImage) GetMessageResponseToIdOk() (*string, bool)`

GetMessageResponseToIdOk returns a tuple with the MessageResponseToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageResponseToId

`func (o *AIgencyImage) SetMessageResponseToId(v string)`

SetMessageResponseToId sets MessageResponseToId field to given value.

### HasMessageResponseToId

`func (o *AIgencyImage) HasMessageResponseToId() bool`

HasMessageResponseToId returns a boolean if a field has been set.

### GetMetaData

`func (o *AIgencyImage) GetMetaData() map[string]map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AIgencyImage) GetMetaDataOk() (*map[string]map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AIgencyImage) SetMetaData(v map[string]map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AIgencyImage) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetMimetype

`func (o *AIgencyImage) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *AIgencyImage) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *AIgencyImage) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.

### HasMimetype

`func (o *AIgencyImage) HasMimetype() bool`

HasMimetype returns a boolean if a field has been set.

### GetMissionId

`func (o *AIgencyImage) GetMissionId() string`

GetMissionId returns the MissionId field if non-nil, zero value otherwise.

### GetMissionIdOk

`func (o *AIgencyImage) GetMissionIdOk() (*string, bool)`

GetMissionIdOk returns a tuple with the MissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionId

`func (o *AIgencyImage) SetMissionId(v string)`

SetMissionId sets MissionId field to given value.

### HasMissionId

`func (o *AIgencyImage) HasMissionId() bool`

HasMissionId returns a boolean if a field has been set.

### GetOrgOwnerId

`func (o *AIgencyImage) GetOrgOwnerId() string`

GetOrgOwnerId returns the OrgOwnerId field if non-nil, zero value otherwise.

### GetOrgOwnerIdOk

`func (o *AIgencyImage) GetOrgOwnerIdOk() (*string, bool)`

GetOrgOwnerIdOk returns a tuple with the OrgOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgOwnerId

`func (o *AIgencyImage) SetOrgOwnerId(v string)`

SetOrgOwnerId sets OrgOwnerId field to given value.

### HasOrgOwnerId

`func (o *AIgencyImage) HasOrgOwnerId() bool`

HasOrgOwnerId returns a boolean if a field has been set.

### GetParameters

`func (o *AIgencyImage) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AIgencyImage) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AIgencyImage) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AIgencyImage) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetProcessingErrors

`func (o *AIgencyImage) GetProcessingErrors() []string`

GetProcessingErrors returns the ProcessingErrors field if non-nil, zero value otherwise.

### GetProcessingErrorsOk

`func (o *AIgencyImage) GetProcessingErrorsOk() (*[]string, bool)`

GetProcessingErrorsOk returns a tuple with the ProcessingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingErrors

`func (o *AIgencyImage) SetProcessingErrors(v []string)`

SetProcessingErrors sets ProcessingErrors field to given value.

### HasProcessingErrors

`func (o *AIgencyImage) HasProcessingErrors() bool`

HasProcessingErrors returns a boolean if a field has been set.

### GetUrl

`func (o *AIgencyImage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AIgencyImage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AIgencyImage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AIgencyImage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserId

`func (o *AIgencyImage) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AIgencyImage) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AIgencyImage) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AIgencyImage) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *AIgencyImage) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AIgencyImage) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AIgencyImage) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AIgencyImage) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


