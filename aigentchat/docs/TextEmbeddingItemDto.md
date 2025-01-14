# TextEmbeddingItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**Text** | **string** |  | 

## Methods

### NewTextEmbeddingItemDto

`func NewTextEmbeddingItemDto(text string, ) *TextEmbeddingItemDto`

NewTextEmbeddingItemDto instantiates a new TextEmbeddingItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextEmbeddingItemDtoWithDefaults

`func NewTextEmbeddingItemDtoWithDefaults() *TextEmbeddingItemDto`

NewTextEmbeddingItemDtoWithDefaults instantiates a new TextEmbeddingItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaData

`func (o *TextEmbeddingItemDto) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *TextEmbeddingItemDto) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *TextEmbeddingItemDto) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *TextEmbeddingItemDto) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetText

`func (o *TextEmbeddingItemDto) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TextEmbeddingItemDto) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TextEmbeddingItemDto) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


