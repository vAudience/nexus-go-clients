# CollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**FileCount** | Pointer to **int64** |  | [optional] 
**Id** | **string** |  | 
**LastFileAddedAt** | Pointer to **string** | LastFileAddedAt is the creation time of the most recently added file in the collection. Omitted when the collection has no files. | [optional] 
**Name** | **string** |  | 
**Processors** | Pointer to **map[string]interface{}** | Processors maps each enabled processor to its resource id (read-only), e.g. {\&quot;deepr\&quot;: \&quot;corpus-abc\&quot;}. | [optional] 
**StatusBreakdown** | Pointer to **map[string]map[string]int64** | StatusBreakdown is keyed by processor name (v1: \&quot;deepr\&quot;). Each value is a per-status file count (e.g. {\&quot;pending\&quot;: 5, \&quot;completed\&quot;: 10}), derived from corpus file statuses. Omitted when no processor is enabled. | [optional] 
**TokenCount** | Pointer to **int64** | TokenCount is the collection&#39;s deepr ingested-token usage. Fail-soft: 0 when the collection has no provisioned corpus or the usage read is unavailable. | [optional] 
**TotalSize** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | **string** |  | 

## Methods

### NewCollectionResponse

`func NewCollectionResponse(createdAt string, id string, name string, updatedAt string, ) *CollectionResponse`

NewCollectionResponse instantiates a new CollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithDefaults

`func NewCollectionResponseWithDefaults() *CollectionResponse`

NewCollectionResponseWithDefaults instantiates a new CollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CollectionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *CollectionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileCount

`func (o *CollectionResponse) GetFileCount() int64`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *CollectionResponse) GetFileCountOk() (*int64, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *CollectionResponse) SetFileCount(v int64)`

SetFileCount sets FileCount field to given value.

### HasFileCount

`func (o *CollectionResponse) HasFileCount() bool`

HasFileCount returns a boolean if a field has been set.

### GetId

`func (o *CollectionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastFileAddedAt

`func (o *CollectionResponse) GetLastFileAddedAt() string`

GetLastFileAddedAt returns the LastFileAddedAt field if non-nil, zero value otherwise.

### GetLastFileAddedAtOk

`func (o *CollectionResponse) GetLastFileAddedAtOk() (*string, bool)`

GetLastFileAddedAtOk returns a tuple with the LastFileAddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFileAddedAt

`func (o *CollectionResponse) SetLastFileAddedAt(v string)`

SetLastFileAddedAt sets LastFileAddedAt field to given value.

### HasLastFileAddedAt

`func (o *CollectionResponse) HasLastFileAddedAt() bool`

HasLastFileAddedAt returns a boolean if a field has been set.

### GetName

`func (o *CollectionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProcessors

`func (o *CollectionResponse) GetProcessors() map[string]interface{}`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *CollectionResponse) GetProcessorsOk() (*map[string]interface{}, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *CollectionResponse) SetProcessors(v map[string]interface{})`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *CollectionResponse) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetStatusBreakdown

`func (o *CollectionResponse) GetStatusBreakdown() map[string]map[string]int64`

GetStatusBreakdown returns the StatusBreakdown field if non-nil, zero value otherwise.

### GetStatusBreakdownOk

`func (o *CollectionResponse) GetStatusBreakdownOk() (*map[string]map[string]int64, bool)`

GetStatusBreakdownOk returns a tuple with the StatusBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusBreakdown

`func (o *CollectionResponse) SetStatusBreakdown(v map[string]map[string]int64)`

SetStatusBreakdown sets StatusBreakdown field to given value.

### HasStatusBreakdown

`func (o *CollectionResponse) HasStatusBreakdown() bool`

HasStatusBreakdown returns a boolean if a field has been set.

### GetTokenCount

`func (o *CollectionResponse) GetTokenCount() int64`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *CollectionResponse) GetTokenCountOk() (*int64, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *CollectionResponse) SetTokenCount(v int64)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *CollectionResponse) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetTotalSize

`func (o *CollectionResponse) GetTotalSize() int64`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *CollectionResponse) GetTotalSizeOk() (*int64, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *CollectionResponse) SetTotalSize(v int64)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *CollectionResponse) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CollectionResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CollectionResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CollectionResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


