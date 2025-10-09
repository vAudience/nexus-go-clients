# ImageGenerationRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** |  | 
**AttachedFiles** | Pointer to **[]string** |  | [optional] 
**Message** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewImageGenerationRequestDto

`func NewImageGenerationRequestDto(agentId string, message string, ) *ImageGenerationRequestDto`

NewImageGenerationRequestDto instantiates a new ImageGenerationRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageGenerationRequestDtoWithDefaults

`func NewImageGenerationRequestDtoWithDefaults() *ImageGenerationRequestDto`

NewImageGenerationRequestDtoWithDefaults instantiates a new ImageGenerationRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *ImageGenerationRequestDto) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ImageGenerationRequestDto) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ImageGenerationRequestDto) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetAttachedFiles

`func (o *ImageGenerationRequestDto) GetAttachedFiles() []string`

GetAttachedFiles returns the AttachedFiles field if non-nil, zero value otherwise.

### GetAttachedFilesOk

`func (o *ImageGenerationRequestDto) GetAttachedFilesOk() (*[]string, bool)`

GetAttachedFilesOk returns a tuple with the AttachedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedFiles

`func (o *ImageGenerationRequestDto) SetAttachedFiles(v []string)`

SetAttachedFiles sets AttachedFiles field to given value.

### HasAttachedFiles

`func (o *ImageGenerationRequestDto) HasAttachedFiles() bool`

HasAttachedFiles returns a boolean if a field has been set.

### GetMessage

`func (o *ImageGenerationRequestDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ImageGenerationRequestDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ImageGenerationRequestDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetParameters

`func (o *ImageGenerationRequestDto) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ImageGenerationRequestDto) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ImageGenerationRequestDto) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ImageGenerationRequestDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


