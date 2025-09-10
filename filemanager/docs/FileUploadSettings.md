# FileUploadSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedMimeTypes** | Pointer to **[]string** |  | [optional] 
**MaxFileSize** | Pointer to **int64** |  | [optional] 
**MinFileSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewFileUploadSettings

`func NewFileUploadSettings() *FileUploadSettings`

NewFileUploadSettings instantiates a new FileUploadSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUploadSettingsWithDefaults

`func NewFileUploadSettingsWithDefaults() *FileUploadSettings`

NewFileUploadSettingsWithDefaults instantiates a new FileUploadSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedMimeTypes

`func (o *FileUploadSettings) GetAcceptedMimeTypes() []string`

GetAcceptedMimeTypes returns the AcceptedMimeTypes field if non-nil, zero value otherwise.

### GetAcceptedMimeTypesOk

`func (o *FileUploadSettings) GetAcceptedMimeTypesOk() (*[]string, bool)`

GetAcceptedMimeTypesOk returns a tuple with the AcceptedMimeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedMimeTypes

`func (o *FileUploadSettings) SetAcceptedMimeTypes(v []string)`

SetAcceptedMimeTypes sets AcceptedMimeTypes field to given value.

### HasAcceptedMimeTypes

`func (o *FileUploadSettings) HasAcceptedMimeTypes() bool`

HasAcceptedMimeTypes returns a boolean if a field has been set.

### GetMaxFileSize

`func (o *FileUploadSettings) GetMaxFileSize() int64`

GetMaxFileSize returns the MaxFileSize field if non-nil, zero value otherwise.

### GetMaxFileSizeOk

`func (o *FileUploadSettings) GetMaxFileSizeOk() (*int64, bool)`

GetMaxFileSizeOk returns a tuple with the MaxFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSize

`func (o *FileUploadSettings) SetMaxFileSize(v int64)`

SetMaxFileSize sets MaxFileSize field to given value.

### HasMaxFileSize

`func (o *FileUploadSettings) HasMaxFileSize() bool`

HasMaxFileSize returns a boolean if a field has been set.

### GetMinFileSize

`func (o *FileUploadSettings) GetMinFileSize() int64`

GetMinFileSize returns the MinFileSize field if non-nil, zero value otherwise.

### GetMinFileSizeOk

`func (o *FileUploadSettings) GetMinFileSizeOk() (*int64, bool)`

GetMinFileSizeOk returns a tuple with the MinFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFileSize

`func (o *FileUploadSettings) SetMinFileSize(v int64)`

SetMinFileSize sets MinFileSize field to given value.

### HasMinFileSize

`func (o *FileUploadSettings) HasMinFileSize() bool`

HasMinFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


