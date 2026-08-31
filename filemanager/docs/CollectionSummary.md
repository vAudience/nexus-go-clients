# CollectionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCollections** | Pointer to **int64** |  | [optional] 
**TotalCollectionsPerUser** | Pointer to **int64** | TotalCollectionsPerUser counts the user&#39;s collections across all orgs, so it pairs with max_collections_count_per_user from /collections/settings. Every other total here is scoped to this org. | [optional] 
**TotalFiles** | Pointer to **int64** |  | [optional] 
**TotalSize** | Pointer to **int64** |  | [optional] 
**TotalTokens** | Pointer to **int64** |  | [optional] 

## Methods

### NewCollectionSummary

`func NewCollectionSummary() *CollectionSummary`

NewCollectionSummary instantiates a new CollectionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionSummaryWithDefaults

`func NewCollectionSummaryWithDefaults() *CollectionSummary`

NewCollectionSummaryWithDefaults instantiates a new CollectionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCollections

`func (o *CollectionSummary) GetTotalCollections() int64`

GetTotalCollections returns the TotalCollections field if non-nil, zero value otherwise.

### GetTotalCollectionsOk

`func (o *CollectionSummary) GetTotalCollectionsOk() (*int64, bool)`

GetTotalCollectionsOk returns a tuple with the TotalCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCollections

`func (o *CollectionSummary) SetTotalCollections(v int64)`

SetTotalCollections sets TotalCollections field to given value.

### HasTotalCollections

`func (o *CollectionSummary) HasTotalCollections() bool`

HasTotalCollections returns a boolean if a field has been set.

### GetTotalCollectionsPerUser

`func (o *CollectionSummary) GetTotalCollectionsPerUser() int64`

GetTotalCollectionsPerUser returns the TotalCollectionsPerUser field if non-nil, zero value otherwise.

### GetTotalCollectionsPerUserOk

`func (o *CollectionSummary) GetTotalCollectionsPerUserOk() (*int64, bool)`

GetTotalCollectionsPerUserOk returns a tuple with the TotalCollectionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCollectionsPerUser

`func (o *CollectionSummary) SetTotalCollectionsPerUser(v int64)`

SetTotalCollectionsPerUser sets TotalCollectionsPerUser field to given value.

### HasTotalCollectionsPerUser

`func (o *CollectionSummary) HasTotalCollectionsPerUser() bool`

HasTotalCollectionsPerUser returns a boolean if a field has been set.

### GetTotalFiles

`func (o *CollectionSummary) GetTotalFiles() int64`

GetTotalFiles returns the TotalFiles field if non-nil, zero value otherwise.

### GetTotalFilesOk

`func (o *CollectionSummary) GetTotalFilesOk() (*int64, bool)`

GetTotalFilesOk returns a tuple with the TotalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFiles

`func (o *CollectionSummary) SetTotalFiles(v int64)`

SetTotalFiles sets TotalFiles field to given value.

### HasTotalFiles

`func (o *CollectionSummary) HasTotalFiles() bool`

HasTotalFiles returns a boolean if a field has been set.

### GetTotalSize

`func (o *CollectionSummary) GetTotalSize() int64`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *CollectionSummary) GetTotalSizeOk() (*int64, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *CollectionSummary) SetTotalSize(v int64)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *CollectionSummary) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetTotalTokens

`func (o *CollectionSummary) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *CollectionSummary) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *CollectionSummary) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *CollectionSummary) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


