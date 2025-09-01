# MetadataFileMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**OriginalFileName** | Pointer to **string** |  | [optional] 
**OriginalFileSize** | Pointer to **int32** |  | [optional] 
**OriginalMimeType** | Pointer to **string** |  | [optional] 
**OriginalUrl** | Pointer to **string** |  | [optional] 
**StorageCategory** | Pointer to **string** | temp, public, privateOrg, privateUser, privateTeam | [optional] 
**StoragePath** | Pointer to **string** |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UploadCategory** | Pointer to **string** | temp, image | [optional] 
**UploadParams** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewMetadataFileMetadataResponse

`func NewMetadataFileMetadataResponse() *MetadataFileMetadataResponse`

NewMetadataFileMetadataResponse instantiates a new MetadataFileMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataFileMetadataResponseWithDefaults

`func NewMetadataFileMetadataResponseWithDefaults() *MetadataFileMetadataResponse`

NewMetadataFileMetadataResponseWithDefaults instantiates a new MetadataFileMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MetadataFileMetadataResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetadataFileMetadataResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetadataFileMetadataResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MetadataFileMetadataResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *MetadataFileMetadataResponse) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *MetadataFileMetadataResponse) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *MetadataFileMetadataResponse) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *MetadataFileMetadataResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *MetadataFileMetadataResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *MetadataFileMetadataResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *MetadataFileMetadataResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *MetadataFileMetadataResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFileName

`func (o *MetadataFileMetadataResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MetadataFileMetadataResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MetadataFileMetadataResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *MetadataFileMetadataResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *MetadataFileMetadataResponse) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *MetadataFileMetadataResponse) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *MetadataFileMetadataResponse) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *MetadataFileMetadataResponse) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetId

`func (o *MetadataFileMetadataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataFileMetadataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataFileMetadataResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataFileMetadataResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *MetadataFileMetadataResponse) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MetadataFileMetadataResponse) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MetadataFileMetadataResponse) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MetadataFileMetadataResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMimeType

`func (o *MetadataFileMetadataResponse) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *MetadataFileMetadataResponse) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *MetadataFileMetadataResponse) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *MetadataFileMetadataResponse) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetOrganizationId

`func (o *MetadataFileMetadataResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *MetadataFileMetadataResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *MetadataFileMetadataResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *MetadataFileMetadataResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOriginalFileName

`func (o *MetadataFileMetadataResponse) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *MetadataFileMetadataResponse) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *MetadataFileMetadataResponse) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *MetadataFileMetadataResponse) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### GetOriginalFileSize

`func (o *MetadataFileMetadataResponse) GetOriginalFileSize() int32`

GetOriginalFileSize returns the OriginalFileSize field if non-nil, zero value otherwise.

### GetOriginalFileSizeOk

`func (o *MetadataFileMetadataResponse) GetOriginalFileSizeOk() (*int32, bool)`

GetOriginalFileSizeOk returns a tuple with the OriginalFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileSize

`func (o *MetadataFileMetadataResponse) SetOriginalFileSize(v int32)`

SetOriginalFileSize sets OriginalFileSize field to given value.

### HasOriginalFileSize

`func (o *MetadataFileMetadataResponse) HasOriginalFileSize() bool`

HasOriginalFileSize returns a boolean if a field has been set.

### GetOriginalMimeType

`func (o *MetadataFileMetadataResponse) GetOriginalMimeType() string`

GetOriginalMimeType returns the OriginalMimeType field if non-nil, zero value otherwise.

### GetOriginalMimeTypeOk

`func (o *MetadataFileMetadataResponse) GetOriginalMimeTypeOk() (*string, bool)`

GetOriginalMimeTypeOk returns a tuple with the OriginalMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMimeType

`func (o *MetadataFileMetadataResponse) SetOriginalMimeType(v string)`

SetOriginalMimeType sets OriginalMimeType field to given value.

### HasOriginalMimeType

`func (o *MetadataFileMetadataResponse) HasOriginalMimeType() bool`

HasOriginalMimeType returns a boolean if a field has been set.

### GetOriginalUrl

`func (o *MetadataFileMetadataResponse) GetOriginalUrl() string`

GetOriginalUrl returns the OriginalUrl field if non-nil, zero value otherwise.

### GetOriginalUrlOk

`func (o *MetadataFileMetadataResponse) GetOriginalUrlOk() (*string, bool)`

GetOriginalUrlOk returns a tuple with the OriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUrl

`func (o *MetadataFileMetadataResponse) SetOriginalUrl(v string)`

SetOriginalUrl sets OriginalUrl field to given value.

### HasOriginalUrl

`func (o *MetadataFileMetadataResponse) HasOriginalUrl() bool`

HasOriginalUrl returns a boolean if a field has been set.

### GetStorageCategory

`func (o *MetadataFileMetadataResponse) GetStorageCategory() string`

GetStorageCategory returns the StorageCategory field if non-nil, zero value otherwise.

### GetStorageCategoryOk

`func (o *MetadataFileMetadataResponse) GetStorageCategoryOk() (*string, bool)`

GetStorageCategoryOk returns a tuple with the StorageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCategory

`func (o *MetadataFileMetadataResponse) SetStorageCategory(v string)`

SetStorageCategory sets StorageCategory field to given value.

### HasStorageCategory

`func (o *MetadataFileMetadataResponse) HasStorageCategory() bool`

HasStorageCategory returns a boolean if a field has been set.

### GetStoragePath

`func (o *MetadataFileMetadataResponse) GetStoragePath() string`

GetStoragePath returns the StoragePath field if non-nil, zero value otherwise.

### GetStoragePathOk

`func (o *MetadataFileMetadataResponse) GetStoragePathOk() (*string, bool)`

GetStoragePathOk returns a tuple with the StoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePath

`func (o *MetadataFileMetadataResponse) SetStoragePath(v string)`

SetStoragePath sets StoragePath field to given value.

### HasStoragePath

`func (o *MetadataFileMetadataResponse) HasStoragePath() bool`

HasStoragePath returns a boolean if a field has been set.

### GetTeamId

`func (o *MetadataFileMetadataResponse) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *MetadataFileMetadataResponse) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *MetadataFileMetadataResponse) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *MetadataFileMetadataResponse) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MetadataFileMetadataResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MetadataFileMetadataResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MetadataFileMetadataResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MetadataFileMetadataResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUploadCategory

`func (o *MetadataFileMetadataResponse) GetUploadCategory() string`

GetUploadCategory returns the UploadCategory field if non-nil, zero value otherwise.

### GetUploadCategoryOk

`func (o *MetadataFileMetadataResponse) GetUploadCategoryOk() (*string, bool)`

GetUploadCategoryOk returns a tuple with the UploadCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadCategory

`func (o *MetadataFileMetadataResponse) SetUploadCategory(v string)`

SetUploadCategory sets UploadCategory field to given value.

### HasUploadCategory

`func (o *MetadataFileMetadataResponse) HasUploadCategory() bool`

HasUploadCategory returns a boolean if a field has been set.

### GetUploadParams

`func (o *MetadataFileMetadataResponse) GetUploadParams() map[string]map[string]interface{}`

GetUploadParams returns the UploadParams field if non-nil, zero value otherwise.

### GetUploadParamsOk

`func (o *MetadataFileMetadataResponse) GetUploadParamsOk() (*map[string]map[string]interface{}, bool)`

GetUploadParamsOk returns a tuple with the UploadParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadParams

`func (o *MetadataFileMetadataResponse) SetUploadParams(v map[string]map[string]interface{})`

SetUploadParams sets UploadParams field to given value.

### HasUploadParams

`func (o *MetadataFileMetadataResponse) HasUploadParams() bool`

HasUploadParams returns a boolean if a field has been set.

### GetUrl

`func (o *MetadataFileMetadataResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MetadataFileMetadataResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MetadataFileMetadataResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MetadataFileMetadataResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserId

`func (o *MetadataFileMetadataResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MetadataFileMetadataResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MetadataFileMetadataResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MetadataFileMetadataResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


