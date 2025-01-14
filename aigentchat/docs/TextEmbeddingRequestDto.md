# TextEmbeddingRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** |  | 
**Items** | [**[]TextEmbeddingItemDto**](TextEmbeddingItemDto.md) |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTextEmbeddingRequestDto

`func NewTextEmbeddingRequestDto(agentId string, items []TextEmbeddingItemDto, ) *TextEmbeddingRequestDto`

NewTextEmbeddingRequestDto instantiates a new TextEmbeddingRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextEmbeddingRequestDtoWithDefaults

`func NewTextEmbeddingRequestDtoWithDefaults() *TextEmbeddingRequestDto`

NewTextEmbeddingRequestDtoWithDefaults instantiates a new TextEmbeddingRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *TextEmbeddingRequestDto) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *TextEmbeddingRequestDto) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *TextEmbeddingRequestDto) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetItems

`func (o *TextEmbeddingRequestDto) GetItems() []TextEmbeddingItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TextEmbeddingRequestDto) GetItemsOk() (*[]TextEmbeddingItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TextEmbeddingRequestDto) SetItems(v []TextEmbeddingItemDto)`

SetItems sets Items field to given value.


### GetParameters

`func (o *TextEmbeddingRequestDto) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TextEmbeddingRequestDto) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TextEmbeddingRequestDto) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TextEmbeddingRequestDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


