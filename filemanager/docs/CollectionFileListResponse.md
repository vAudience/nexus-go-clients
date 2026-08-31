# CollectionFileListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**CollectionResponse**](CollectionResponse.md) |  | [optional] 
**Files** | Pointer to [**[]CollectionFileResponse**](CollectionFileResponse.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewCollectionFileListResponse

`func NewCollectionFileListResponse() *CollectionFileListResponse`

NewCollectionFileListResponse instantiates a new CollectionFileListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionFileListResponseWithDefaults

`func NewCollectionFileListResponseWithDefaults() *CollectionFileListResponse`

NewCollectionFileListResponseWithDefaults instantiates a new CollectionFileListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *CollectionFileListResponse) GetCollection() CollectionResponse`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *CollectionFileListResponse) GetCollectionOk() (*CollectionResponse, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *CollectionFileListResponse) SetCollection(v CollectionResponse)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *CollectionFileListResponse) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetFiles

`func (o *CollectionFileListResponse) GetFiles() []CollectionFileResponse`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CollectionFileListResponse) GetFilesOk() (*[]CollectionFileResponse, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CollectionFileListResponse) SetFiles(v []CollectionFileResponse)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CollectionFileListResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetLimit

`func (o *CollectionFileListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CollectionFileListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CollectionFileListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CollectionFileListResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *CollectionFileListResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CollectionFileListResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CollectionFileListResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CollectionFileListResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotal

`func (o *CollectionFileListResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionFileListResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionFileListResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CollectionFileListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


