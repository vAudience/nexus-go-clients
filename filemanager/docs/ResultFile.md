# ResultFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** |  | 
**FileSize** | **int64** |  | 
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**MimeType** | **string** |  | 
**OriginalFileMimeType** | **string** |  | 
**OriginalFileName** | **string** |  | 
**OriginalFileSize** | **int64** |  | 
**OriginalFileUrl** | Pointer to **string** |  | [optional] 
**StoragePath** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewResultFile

`func NewResultFile(fileName string, fileSize int64, mimeType string, originalFileMimeType string, originalFileName string, originalFileSize int64, storagePath string, url string, ) *ResultFile`

NewResultFile instantiates a new ResultFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultFileWithDefaults

`func NewResultFileWithDefaults() *ResultFile`

NewResultFileWithDefaults instantiates a new ResultFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *ResultFile) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ResultFile) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ResultFile) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileSize

`func (o *ResultFile) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *ResultFile) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *ResultFile) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.


### GetMetaData

`func (o *ResultFile) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *ResultFile) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *ResultFile) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *ResultFile) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetMimeType

`func (o *ResultFile) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResultFile) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResultFile) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetOriginalFileMimeType

`func (o *ResultFile) GetOriginalFileMimeType() string`

GetOriginalFileMimeType returns the OriginalFileMimeType field if non-nil, zero value otherwise.

### GetOriginalFileMimeTypeOk

`func (o *ResultFile) GetOriginalFileMimeTypeOk() (*string, bool)`

GetOriginalFileMimeTypeOk returns a tuple with the OriginalFileMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileMimeType

`func (o *ResultFile) SetOriginalFileMimeType(v string)`

SetOriginalFileMimeType sets OriginalFileMimeType field to given value.


### GetOriginalFileName

`func (o *ResultFile) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *ResultFile) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *ResultFile) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.


### GetOriginalFileSize

`func (o *ResultFile) GetOriginalFileSize() int64`

GetOriginalFileSize returns the OriginalFileSize field if non-nil, zero value otherwise.

### GetOriginalFileSizeOk

`func (o *ResultFile) GetOriginalFileSizeOk() (*int64, bool)`

GetOriginalFileSizeOk returns a tuple with the OriginalFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileSize

`func (o *ResultFile) SetOriginalFileSize(v int64)`

SetOriginalFileSize sets OriginalFileSize field to given value.


### GetOriginalFileUrl

`func (o *ResultFile) GetOriginalFileUrl() string`

GetOriginalFileUrl returns the OriginalFileUrl field if non-nil, zero value otherwise.

### GetOriginalFileUrlOk

`func (o *ResultFile) GetOriginalFileUrlOk() (*string, bool)`

GetOriginalFileUrlOk returns a tuple with the OriginalFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileUrl

`func (o *ResultFile) SetOriginalFileUrl(v string)`

SetOriginalFileUrl sets OriginalFileUrl field to given value.

### HasOriginalFileUrl

`func (o *ResultFile) HasOriginalFileUrl() bool`

HasOriginalFileUrl returns a boolean if a field has been set.

### GetStoragePath

`func (o *ResultFile) GetStoragePath() string`

GetStoragePath returns the StoragePath field if non-nil, zero value otherwise.

### GetStoragePathOk

`func (o *ResultFile) GetStoragePathOk() (*string, bool)`

GetStoragePathOk returns a tuple with the StoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePath

`func (o *ResultFile) SetStoragePath(v string)`

SetStoragePath sets StoragePath field to given value.


### GetUrl

`func (o *ResultFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResultFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResultFile) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


