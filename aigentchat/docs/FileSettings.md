# FileSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | Pointer to [**FileChatSettings**](FileChatSettings.md) |  | [optional] 
**Upload** | Pointer to [**FileUploadSettings**](FileUploadSettings.md) |  | [optional] 

## Methods

### NewFileSettings

`func NewFileSettings() *FileSettings`

NewFileSettings instantiates a new FileSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSettingsWithDefaults

`func NewFileSettingsWithDefaults() *FileSettings`

NewFileSettingsWithDefaults instantiates a new FileSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *FileSettings) GetChat() FileChatSettings`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *FileSettings) GetChatOk() (*FileChatSettings, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *FileSettings) SetChat(v FileChatSettings)`

SetChat sets Chat field to given value.

### HasChat

`func (o *FileSettings) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetUpload

`func (o *FileSettings) GetUpload() FileUploadSettings`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *FileSettings) GetUploadOk() (*FileUploadSettings, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *FileSettings) SetUpload(v FileUploadSettings)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *FileSettings) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


