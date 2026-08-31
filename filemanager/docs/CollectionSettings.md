# CollectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedMimeTypes** | Pointer to **[]string** |  | [optional] 
**FileUploadCategory** | Pointer to **string** |  | [optional] 
**MaxCollectionsCount** | Pointer to **int64** |  | [optional] 
**MaxCollectionsCountPerUser** | Pointer to **int64** | MaxCollectionsCountPerUser caps the user&#39;s collections across all orgs, checked on create alongside the per-org max_collections_count. | [optional] 
**MaxCollectionsSize** | Pointer to **int64** |  | [optional] 
**MaxCollectionsTokens** | Pointer to **int64** |  | [optional] 
**MaxFileSize** | Pointer to **int64** |  | [optional] 
**MinFileSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewCollectionSettings

`func NewCollectionSettings() *CollectionSettings`

NewCollectionSettings instantiates a new CollectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionSettingsWithDefaults

`func NewCollectionSettingsWithDefaults() *CollectionSettings`

NewCollectionSettingsWithDefaults instantiates a new CollectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedMimeTypes

`func (o *CollectionSettings) GetAcceptedMimeTypes() []string`

GetAcceptedMimeTypes returns the AcceptedMimeTypes field if non-nil, zero value otherwise.

### GetAcceptedMimeTypesOk

`func (o *CollectionSettings) GetAcceptedMimeTypesOk() (*[]string, bool)`

GetAcceptedMimeTypesOk returns a tuple with the AcceptedMimeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedMimeTypes

`func (o *CollectionSettings) SetAcceptedMimeTypes(v []string)`

SetAcceptedMimeTypes sets AcceptedMimeTypes field to given value.

### HasAcceptedMimeTypes

`func (o *CollectionSettings) HasAcceptedMimeTypes() bool`

HasAcceptedMimeTypes returns a boolean if a field has been set.

### GetFileUploadCategory

`func (o *CollectionSettings) GetFileUploadCategory() string`

GetFileUploadCategory returns the FileUploadCategory field if non-nil, zero value otherwise.

### GetFileUploadCategoryOk

`func (o *CollectionSettings) GetFileUploadCategoryOk() (*string, bool)`

GetFileUploadCategoryOk returns a tuple with the FileUploadCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadCategory

`func (o *CollectionSettings) SetFileUploadCategory(v string)`

SetFileUploadCategory sets FileUploadCategory field to given value.

### HasFileUploadCategory

`func (o *CollectionSettings) HasFileUploadCategory() bool`

HasFileUploadCategory returns a boolean if a field has been set.

### GetMaxCollectionsCount

`func (o *CollectionSettings) GetMaxCollectionsCount() int64`

GetMaxCollectionsCount returns the MaxCollectionsCount field if non-nil, zero value otherwise.

### GetMaxCollectionsCountOk

`func (o *CollectionSettings) GetMaxCollectionsCountOk() (*int64, bool)`

GetMaxCollectionsCountOk returns a tuple with the MaxCollectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCollectionsCount

`func (o *CollectionSettings) SetMaxCollectionsCount(v int64)`

SetMaxCollectionsCount sets MaxCollectionsCount field to given value.

### HasMaxCollectionsCount

`func (o *CollectionSettings) HasMaxCollectionsCount() bool`

HasMaxCollectionsCount returns a boolean if a field has been set.

### GetMaxCollectionsCountPerUser

`func (o *CollectionSettings) GetMaxCollectionsCountPerUser() int64`

GetMaxCollectionsCountPerUser returns the MaxCollectionsCountPerUser field if non-nil, zero value otherwise.

### GetMaxCollectionsCountPerUserOk

`func (o *CollectionSettings) GetMaxCollectionsCountPerUserOk() (*int64, bool)`

GetMaxCollectionsCountPerUserOk returns a tuple with the MaxCollectionsCountPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCollectionsCountPerUser

`func (o *CollectionSettings) SetMaxCollectionsCountPerUser(v int64)`

SetMaxCollectionsCountPerUser sets MaxCollectionsCountPerUser field to given value.

### HasMaxCollectionsCountPerUser

`func (o *CollectionSettings) HasMaxCollectionsCountPerUser() bool`

HasMaxCollectionsCountPerUser returns a boolean if a field has been set.

### GetMaxCollectionsSize

`func (o *CollectionSettings) GetMaxCollectionsSize() int64`

GetMaxCollectionsSize returns the MaxCollectionsSize field if non-nil, zero value otherwise.

### GetMaxCollectionsSizeOk

`func (o *CollectionSettings) GetMaxCollectionsSizeOk() (*int64, bool)`

GetMaxCollectionsSizeOk returns a tuple with the MaxCollectionsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCollectionsSize

`func (o *CollectionSettings) SetMaxCollectionsSize(v int64)`

SetMaxCollectionsSize sets MaxCollectionsSize field to given value.

### HasMaxCollectionsSize

`func (o *CollectionSettings) HasMaxCollectionsSize() bool`

HasMaxCollectionsSize returns a boolean if a field has been set.

### GetMaxCollectionsTokens

`func (o *CollectionSettings) GetMaxCollectionsTokens() int64`

GetMaxCollectionsTokens returns the MaxCollectionsTokens field if non-nil, zero value otherwise.

### GetMaxCollectionsTokensOk

`func (o *CollectionSettings) GetMaxCollectionsTokensOk() (*int64, bool)`

GetMaxCollectionsTokensOk returns a tuple with the MaxCollectionsTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCollectionsTokens

`func (o *CollectionSettings) SetMaxCollectionsTokens(v int64)`

SetMaxCollectionsTokens sets MaxCollectionsTokens field to given value.

### HasMaxCollectionsTokens

`func (o *CollectionSettings) HasMaxCollectionsTokens() bool`

HasMaxCollectionsTokens returns a boolean if a field has been set.

### GetMaxFileSize

`func (o *CollectionSettings) GetMaxFileSize() int64`

GetMaxFileSize returns the MaxFileSize field if non-nil, zero value otherwise.

### GetMaxFileSizeOk

`func (o *CollectionSettings) GetMaxFileSizeOk() (*int64, bool)`

GetMaxFileSizeOk returns a tuple with the MaxFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSize

`func (o *CollectionSettings) SetMaxFileSize(v int64)`

SetMaxFileSize sets MaxFileSize field to given value.

### HasMaxFileSize

`func (o *CollectionSettings) HasMaxFileSize() bool`

HasMaxFileSize returns a boolean if a field has been set.

### GetMinFileSize

`func (o *CollectionSettings) GetMinFileSize() int64`

GetMinFileSize returns the MinFileSize field if non-nil, zero value otherwise.

### GetMinFileSizeOk

`func (o *CollectionSettings) GetMinFileSizeOk() (*int64, bool)`

GetMinFileSizeOk returns a tuple with the MinFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFileSize

`func (o *CollectionSettings) SetMinFileSize(v int64)`

SetMinFileSize sets MinFileSize field to given value.

### HasMinFileSize

`func (o *CollectionSettings) HasMinFileSize() bool`

HasMinFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


