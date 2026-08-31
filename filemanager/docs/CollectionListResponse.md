# CollectionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]CollectionResponse**](CollectionResponse.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Summary** | Pointer to [**CollectionSummary**](CollectionSummary.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewCollectionListResponse

`func NewCollectionListResponse() *CollectionListResponse`

NewCollectionListResponse instantiates a new CollectionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionListResponseWithDefaults

`func NewCollectionListResponseWithDefaults() *CollectionListResponse`

NewCollectionListResponseWithDefaults instantiates a new CollectionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *CollectionListResponse) GetCollections() []CollectionResponse`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *CollectionListResponse) GetCollectionsOk() (*[]CollectionResponse, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *CollectionListResponse) SetCollections(v []CollectionResponse)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *CollectionListResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetLimit

`func (o *CollectionListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CollectionListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CollectionListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CollectionListResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *CollectionListResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CollectionListResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CollectionListResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CollectionListResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSummary

`func (o *CollectionListResponse) GetSummary() CollectionSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CollectionListResponse) GetSummaryOk() (*CollectionSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CollectionListResponse) SetSummary(v CollectionSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CollectionListResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTotal

`func (o *CollectionListResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionListResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionListResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CollectionListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


