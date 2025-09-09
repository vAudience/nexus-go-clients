# FileMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**FileName** | **string** |  | 
**FileSize** | **int64** |  | 
**Id** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**MimeType** | **string** |  | 
**OrganizationId** | Pointer to **string** |  | [optional] 
**OriginalFileName** | **string** |  | 
**OriginalFileSize** | **int64** |  | 
**OriginalMimeType** | **string** |  | 
**OriginalUrl** | Pointer to **string** |  | [optional] 
**StorageCategory** | **string** | temp, public, privateOrg, privateUser, privateTeam | 
**StoragePath** | **string** |  | 
**TeamId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | **string** |  | 
**UploadCategory** | **string** | temp, image | 
**UploadParams** | Pointer to **map[string]interface{}** |  | [optional] 
**Url** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewFileMetadataResponse

`func NewFileMetadataResponse(createdAt string, fileName string, fileSize int64, id string, mimeType string, originalFileName string, originalFileSize int64, originalMimeType string, storageCategory string, storagePath string, updatedAt string, uploadCategory string, url string, ) *FileMetadataResponse`

NewFileMetadataResponse instantiates a new FileMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileMetadataResponseWithDefaults

`func NewFileMetadataResponseWithDefaults() *FileMetadataResponse`

NewFileMetadataResponseWithDefaults instantiates a new FileMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FileMetadataResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileMetadataResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileMetadataResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *FileMetadataResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *FileMetadataResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *FileMetadataResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *FileMetadataResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFileName

`func (o *FileMetadataResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FileMetadataResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FileMetadataResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileSize

`func (o *FileMetadataResponse) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *FileMetadataResponse) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *FileMetadataResponse) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.


### GetId

`func (o *FileMetadataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileMetadataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileMetadataResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *FileMetadataResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FileMetadataResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FileMetadataResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FileMetadataResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMimeType

`func (o *FileMetadataResponse) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *FileMetadataResponse) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *FileMetadataResponse) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetOrganizationId

`func (o *FileMetadataResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *FileMetadataResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *FileMetadataResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *FileMetadataResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOriginalFileName

`func (o *FileMetadataResponse) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *FileMetadataResponse) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *FileMetadataResponse) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.


### GetOriginalFileSize

`func (o *FileMetadataResponse) GetOriginalFileSize() int64`

GetOriginalFileSize returns the OriginalFileSize field if non-nil, zero value otherwise.

### GetOriginalFileSizeOk

`func (o *FileMetadataResponse) GetOriginalFileSizeOk() (*int64, bool)`

GetOriginalFileSizeOk returns a tuple with the OriginalFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileSize

`func (o *FileMetadataResponse) SetOriginalFileSize(v int64)`

SetOriginalFileSize sets OriginalFileSize field to given value.


### GetOriginalMimeType

`func (o *FileMetadataResponse) GetOriginalMimeType() string`

GetOriginalMimeType returns the OriginalMimeType field if non-nil, zero value otherwise.

### GetOriginalMimeTypeOk

`func (o *FileMetadataResponse) GetOriginalMimeTypeOk() (*string, bool)`

GetOriginalMimeTypeOk returns a tuple with the OriginalMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMimeType

`func (o *FileMetadataResponse) SetOriginalMimeType(v string)`

SetOriginalMimeType sets OriginalMimeType field to given value.


### GetOriginalUrl

`func (o *FileMetadataResponse) GetOriginalUrl() string`

GetOriginalUrl returns the OriginalUrl field if non-nil, zero value otherwise.

### GetOriginalUrlOk

`func (o *FileMetadataResponse) GetOriginalUrlOk() (*string, bool)`

GetOriginalUrlOk returns a tuple with the OriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUrl

`func (o *FileMetadataResponse) SetOriginalUrl(v string)`

SetOriginalUrl sets OriginalUrl field to given value.

### HasOriginalUrl

`func (o *FileMetadataResponse) HasOriginalUrl() bool`

HasOriginalUrl returns a boolean if a field has been set.

### GetStorageCategory

`func (o *FileMetadataResponse) GetStorageCategory() string`

GetStorageCategory returns the StorageCategory field if non-nil, zero value otherwise.

### GetStorageCategoryOk

`func (o *FileMetadataResponse) GetStorageCategoryOk() (*string, bool)`

GetStorageCategoryOk returns a tuple with the StorageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCategory

`func (o *FileMetadataResponse) SetStorageCategory(v string)`

SetStorageCategory sets StorageCategory field to given value.


### GetStoragePath

`func (o *FileMetadataResponse) GetStoragePath() string`

GetStoragePath returns the StoragePath field if non-nil, zero value otherwise.

### GetStoragePathOk

`func (o *FileMetadataResponse) GetStoragePathOk() (*string, bool)`

GetStoragePathOk returns a tuple with the StoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePath

`func (o *FileMetadataResponse) SetStoragePath(v string)`

SetStoragePath sets StoragePath field to given value.


### GetTeamId

`func (o *FileMetadataResponse) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *FileMetadataResponse) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *FileMetadataResponse) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *FileMetadataResponse) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FileMetadataResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FileMetadataResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FileMetadataResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUploadCategory

`func (o *FileMetadataResponse) GetUploadCategory() string`

GetUploadCategory returns the UploadCategory field if non-nil, zero value otherwise.

### GetUploadCategoryOk

`func (o *FileMetadataResponse) GetUploadCategoryOk() (*string, bool)`

GetUploadCategoryOk returns a tuple with the UploadCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadCategory

`func (o *FileMetadataResponse) SetUploadCategory(v string)`

SetUploadCategory sets UploadCategory field to given value.


### GetUploadParams

`func (o *FileMetadataResponse) GetUploadParams() map[string]interface{}`

GetUploadParams returns the UploadParams field if non-nil, zero value otherwise.

### GetUploadParamsOk

`func (o *FileMetadataResponse) GetUploadParamsOk() (*map[string]interface{}, bool)`

GetUploadParamsOk returns a tuple with the UploadParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadParams

`func (o *FileMetadataResponse) SetUploadParams(v map[string]interface{})`

SetUploadParams sets UploadParams field to given value.

### HasUploadParams

`func (o *FileMetadataResponse) HasUploadParams() bool`

HasUploadParams returns a boolean if a field has been set.

### GetUrl

`func (o *FileMetadataResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileMetadataResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileMetadataResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUserId

`func (o *FileMetadataResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FileMetadataResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FileMetadataResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FileMetadataResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


