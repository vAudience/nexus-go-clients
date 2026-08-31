# CollectionFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | Pointer to **string** |  | [optional] 
**CreatedAt** | **string** |  | 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**FileLlmInputType** | **string** |  | 
**FileName** | **string** |  | 
**FileSize** | **int64** |  | 
**FileStorageType** | **string** |  | 
**Id** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**MimeType** | **string** |  | 
**OrganizationId** | Pointer to **string** |  | [optional] 
**OriginalFileMimeType** | **string** |  | 
**OriginalFileName** | **string** |  | 
**OriginalFileSize** | **int64** |  | 
**OriginalFileStoragePath** | Pointer to **string** |  | [optional] 
**OriginalFileUrl** | Pointer to **string** |  | [optional] 
**ProcessingStatus** | Pointer to [**map[string]ProcessingStatus**](ProcessingStatus.md) | ProcessingStatus is keyed by processor name (v1: \&quot;deepr\&quot;). Omitted when the collection has no processor enabled. Each processor&#39;s status carries the per-file deepr token_count. | [optional] 
**StorageBackend** | Pointer to **string** | StorageBackend is the backend that owns the file (local, s3). Additive and omitempty so existing generated clients are unaffected. | [optional] 
**StoragePath** | **string** |  | 
**TeamId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | **string** |  | 
**UploadCategory** | **string** |  | 
**Url** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewCollectionFileResponse

`func NewCollectionFileResponse(createdAt string, fileLlmInputType string, fileName string, fileSize int64, fileStorageType string, id string, mimeType string, originalFileMimeType string, originalFileName string, originalFileSize int64, storagePath string, updatedAt string, uploadCategory string, url string, ) *CollectionFileResponse`

NewCollectionFileResponse instantiates a new CollectionFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionFileResponseWithDefaults

`func NewCollectionFileResponseWithDefaults() *CollectionFileResponse`

NewCollectionFileResponseWithDefaults instantiates a new CollectionFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *CollectionFileResponse) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CollectionFileResponse) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CollectionFileResponse) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *CollectionFileResponse) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CollectionFileResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionFileResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionFileResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *CollectionFileResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CollectionFileResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CollectionFileResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CollectionFileResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFileLlmInputType

`func (o *CollectionFileResponse) GetFileLlmInputType() string`

GetFileLlmInputType returns the FileLlmInputType field if non-nil, zero value otherwise.

### GetFileLlmInputTypeOk

`func (o *CollectionFileResponse) GetFileLlmInputTypeOk() (*string, bool)`

GetFileLlmInputTypeOk returns a tuple with the FileLlmInputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLlmInputType

`func (o *CollectionFileResponse) SetFileLlmInputType(v string)`

SetFileLlmInputType sets FileLlmInputType field to given value.


### GetFileName

`func (o *CollectionFileResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CollectionFileResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CollectionFileResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileSize

`func (o *CollectionFileResponse) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *CollectionFileResponse) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *CollectionFileResponse) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.


### GetFileStorageType

`func (o *CollectionFileResponse) GetFileStorageType() string`

GetFileStorageType returns the FileStorageType field if non-nil, zero value otherwise.

### GetFileStorageTypeOk

`func (o *CollectionFileResponse) GetFileStorageTypeOk() (*string, bool)`

GetFileStorageTypeOk returns a tuple with the FileStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileStorageType

`func (o *CollectionFileResponse) SetFileStorageType(v string)`

SetFileStorageType sets FileStorageType field to given value.


### GetId

`func (o *CollectionFileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionFileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionFileResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *CollectionFileResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CollectionFileResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CollectionFileResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CollectionFileResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMimeType

`func (o *CollectionFileResponse) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *CollectionFileResponse) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *CollectionFileResponse) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetOrganizationId

`func (o *CollectionFileResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CollectionFileResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CollectionFileResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CollectionFileResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOriginalFileMimeType

`func (o *CollectionFileResponse) GetOriginalFileMimeType() string`

GetOriginalFileMimeType returns the OriginalFileMimeType field if non-nil, zero value otherwise.

### GetOriginalFileMimeTypeOk

`func (o *CollectionFileResponse) GetOriginalFileMimeTypeOk() (*string, bool)`

GetOriginalFileMimeTypeOk returns a tuple with the OriginalFileMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileMimeType

`func (o *CollectionFileResponse) SetOriginalFileMimeType(v string)`

SetOriginalFileMimeType sets OriginalFileMimeType field to given value.


### GetOriginalFileName

`func (o *CollectionFileResponse) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *CollectionFileResponse) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *CollectionFileResponse) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.


### GetOriginalFileSize

`func (o *CollectionFileResponse) GetOriginalFileSize() int64`

GetOriginalFileSize returns the OriginalFileSize field if non-nil, zero value otherwise.

### GetOriginalFileSizeOk

`func (o *CollectionFileResponse) GetOriginalFileSizeOk() (*int64, bool)`

GetOriginalFileSizeOk returns a tuple with the OriginalFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileSize

`func (o *CollectionFileResponse) SetOriginalFileSize(v int64)`

SetOriginalFileSize sets OriginalFileSize field to given value.


### GetOriginalFileStoragePath

`func (o *CollectionFileResponse) GetOriginalFileStoragePath() string`

GetOriginalFileStoragePath returns the OriginalFileStoragePath field if non-nil, zero value otherwise.

### GetOriginalFileStoragePathOk

`func (o *CollectionFileResponse) GetOriginalFileStoragePathOk() (*string, bool)`

GetOriginalFileStoragePathOk returns a tuple with the OriginalFileStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileStoragePath

`func (o *CollectionFileResponse) SetOriginalFileStoragePath(v string)`

SetOriginalFileStoragePath sets OriginalFileStoragePath field to given value.

### HasOriginalFileStoragePath

`func (o *CollectionFileResponse) HasOriginalFileStoragePath() bool`

HasOriginalFileStoragePath returns a boolean if a field has been set.

### GetOriginalFileUrl

`func (o *CollectionFileResponse) GetOriginalFileUrl() string`

GetOriginalFileUrl returns the OriginalFileUrl field if non-nil, zero value otherwise.

### GetOriginalFileUrlOk

`func (o *CollectionFileResponse) GetOriginalFileUrlOk() (*string, bool)`

GetOriginalFileUrlOk returns a tuple with the OriginalFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileUrl

`func (o *CollectionFileResponse) SetOriginalFileUrl(v string)`

SetOriginalFileUrl sets OriginalFileUrl field to given value.

### HasOriginalFileUrl

`func (o *CollectionFileResponse) HasOriginalFileUrl() bool`

HasOriginalFileUrl returns a boolean if a field has been set.

### GetProcessingStatus

`func (o *CollectionFileResponse) GetProcessingStatus() map[string]ProcessingStatus`

GetProcessingStatus returns the ProcessingStatus field if non-nil, zero value otherwise.

### GetProcessingStatusOk

`func (o *CollectionFileResponse) GetProcessingStatusOk() (*map[string]ProcessingStatus, bool)`

GetProcessingStatusOk returns a tuple with the ProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStatus

`func (o *CollectionFileResponse) SetProcessingStatus(v map[string]ProcessingStatus)`

SetProcessingStatus sets ProcessingStatus field to given value.

### HasProcessingStatus

`func (o *CollectionFileResponse) HasProcessingStatus() bool`

HasProcessingStatus returns a boolean if a field has been set.

### GetStorageBackend

`func (o *CollectionFileResponse) GetStorageBackend() string`

GetStorageBackend returns the StorageBackend field if non-nil, zero value otherwise.

### GetStorageBackendOk

`func (o *CollectionFileResponse) GetStorageBackendOk() (*string, bool)`

GetStorageBackendOk returns a tuple with the StorageBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBackend

`func (o *CollectionFileResponse) SetStorageBackend(v string)`

SetStorageBackend sets StorageBackend field to given value.

### HasStorageBackend

`func (o *CollectionFileResponse) HasStorageBackend() bool`

HasStorageBackend returns a boolean if a field has been set.

### GetStoragePath

`func (o *CollectionFileResponse) GetStoragePath() string`

GetStoragePath returns the StoragePath field if non-nil, zero value otherwise.

### GetStoragePathOk

`func (o *CollectionFileResponse) GetStoragePathOk() (*string, bool)`

GetStoragePathOk returns a tuple with the StoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePath

`func (o *CollectionFileResponse) SetStoragePath(v string)`

SetStoragePath sets StoragePath field to given value.


### GetTeamId

`func (o *CollectionFileResponse) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CollectionFileResponse) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CollectionFileResponse) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *CollectionFileResponse) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CollectionFileResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CollectionFileResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CollectionFileResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUploadCategory

`func (o *CollectionFileResponse) GetUploadCategory() string`

GetUploadCategory returns the UploadCategory field if non-nil, zero value otherwise.

### GetUploadCategoryOk

`func (o *CollectionFileResponse) GetUploadCategoryOk() (*string, bool)`

GetUploadCategoryOk returns a tuple with the UploadCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadCategory

`func (o *CollectionFileResponse) SetUploadCategory(v string)`

SetUploadCategory sets UploadCategory field to given value.


### GetUrl

`func (o *CollectionFileResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CollectionFileResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CollectionFileResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUserId

`func (o *CollectionFileResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CollectionFileResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CollectionFileResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CollectionFileResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


