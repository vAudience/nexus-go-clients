# ChatCompletionFileSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxFiles** | Pointer to **int32** |  | [optional] 
**MaxTotalFileSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewChatCompletionFileSettings

`func NewChatCompletionFileSettings() *ChatCompletionFileSettings`

NewChatCompletionFileSettings instantiates a new ChatCompletionFileSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionFileSettingsWithDefaults

`func NewChatCompletionFileSettingsWithDefaults() *ChatCompletionFileSettings`

NewChatCompletionFileSettingsWithDefaults instantiates a new ChatCompletionFileSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxFiles

`func (o *ChatCompletionFileSettings) GetMaxFiles() int32`

GetMaxFiles returns the MaxFiles field if non-nil, zero value otherwise.

### GetMaxFilesOk

`func (o *ChatCompletionFileSettings) GetMaxFilesOk() (*int32, bool)`

GetMaxFilesOk returns a tuple with the MaxFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFiles

`func (o *ChatCompletionFileSettings) SetMaxFiles(v int32)`

SetMaxFiles sets MaxFiles field to given value.

### HasMaxFiles

`func (o *ChatCompletionFileSettings) HasMaxFiles() bool`

HasMaxFiles returns a boolean if a field has been set.

### GetMaxTotalFileSize

`func (o *ChatCompletionFileSettings) GetMaxTotalFileSize() int64`

GetMaxTotalFileSize returns the MaxTotalFileSize field if non-nil, zero value otherwise.

### GetMaxTotalFileSizeOk

`func (o *ChatCompletionFileSettings) GetMaxTotalFileSizeOk() (*int64, bool)`

GetMaxTotalFileSizeOk returns a tuple with the MaxTotalFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalFileSize

`func (o *ChatCompletionFileSettings) SetMaxTotalFileSize(v int64)`

SetMaxTotalFileSize sets MaxTotalFileSize field to given value.

### HasMaxTotalFileSize

`func (o *ChatCompletionFileSettings) HasMaxTotalFileSize() bool`

HasMaxTotalFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


