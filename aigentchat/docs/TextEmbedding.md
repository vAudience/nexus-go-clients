# TextEmbedding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embeddings** | [**[]EmbeddingItem**](EmbeddingItem.md) |  | 
**ExecutionId** | **string** |  | 
**ModelId** | **string** |  | 
**Parameters** | **map[string]interface{}** |  | 
**ServiceId** | **string** |  | 
**Tokens** | **int32** |  | 

## Methods

### NewTextEmbedding

`func NewTextEmbedding(embeddings []EmbeddingItem, executionId string, modelId string, parameters map[string]interface{}, serviceId string, tokens int32, ) *TextEmbedding`

NewTextEmbedding instantiates a new TextEmbedding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextEmbeddingWithDefaults

`func NewTextEmbeddingWithDefaults() *TextEmbedding`

NewTextEmbeddingWithDefaults instantiates a new TextEmbedding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbeddings

`func (o *TextEmbedding) GetEmbeddings() []EmbeddingItem`

GetEmbeddings returns the Embeddings field if non-nil, zero value otherwise.

### GetEmbeddingsOk

`func (o *TextEmbedding) GetEmbeddingsOk() (*[]EmbeddingItem, bool)`

GetEmbeddingsOk returns a tuple with the Embeddings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddings

`func (o *TextEmbedding) SetEmbeddings(v []EmbeddingItem)`

SetEmbeddings sets Embeddings field to given value.


### GetExecutionId

`func (o *TextEmbedding) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *TextEmbedding) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *TextEmbedding) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetModelId

`func (o *TextEmbedding) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *TextEmbedding) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *TextEmbedding) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetParameters

`func (o *TextEmbedding) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TextEmbedding) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TextEmbedding) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.


### GetServiceId

`func (o *TextEmbedding) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *TextEmbedding) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *TextEmbedding) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTokens

`func (o *TextEmbedding) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TextEmbedding) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TextEmbedding) SetTokens(v int32)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


