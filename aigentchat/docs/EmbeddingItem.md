# EmbeddingItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedding** | **[]float32** |  | 
**Index** | **int32** |  | 
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEmbeddingItem

`func NewEmbeddingItem(embedding []float32, index int32, ) *EmbeddingItem`

NewEmbeddingItem instantiates a new EmbeddingItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingItemWithDefaults

`func NewEmbeddingItemWithDefaults() *EmbeddingItem`

NewEmbeddingItemWithDefaults instantiates a new EmbeddingItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedding

`func (o *EmbeddingItem) GetEmbedding() []float32`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *EmbeddingItem) GetEmbeddingOk() (*[]float32, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *EmbeddingItem) SetEmbedding(v []float32)`

SetEmbedding sets Embedding field to given value.


### GetIndex

`func (o *EmbeddingItem) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *EmbeddingItem) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *EmbeddingItem) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetMetaData

`func (o *EmbeddingItem) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *EmbeddingItem) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *EmbeddingItem) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *EmbeddingItem) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


