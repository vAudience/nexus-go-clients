# FileUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomingFileName** | Pointer to **string** |  | [optional] 
**ProcessingError** | Pointer to **string** |  | [optional] 
**ResultingFiles** | Pointer to [**[]FilemanagerProcessingResultFile**](FilemanagerProcessingResultFile.md) |  | [optional] 

## Methods

### NewFileUploadResponse

`func NewFileUploadResponse() *FileUploadResponse`

NewFileUploadResponse instantiates a new FileUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUploadResponseWithDefaults

`func NewFileUploadResponseWithDefaults() *FileUploadResponse`

NewFileUploadResponseWithDefaults instantiates a new FileUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomingFileName

`func (o *FileUploadResponse) GetIncomingFileName() string`

GetIncomingFileName returns the IncomingFileName field if non-nil, zero value otherwise.

### GetIncomingFileNameOk

`func (o *FileUploadResponse) GetIncomingFileNameOk() (*string, bool)`

GetIncomingFileNameOk returns a tuple with the IncomingFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingFileName

`func (o *FileUploadResponse) SetIncomingFileName(v string)`

SetIncomingFileName sets IncomingFileName field to given value.

### HasIncomingFileName

`func (o *FileUploadResponse) HasIncomingFileName() bool`

HasIncomingFileName returns a boolean if a field has been set.

### GetProcessingError

`func (o *FileUploadResponse) GetProcessingError() string`

GetProcessingError returns the ProcessingError field if non-nil, zero value otherwise.

### GetProcessingErrorOk

`func (o *FileUploadResponse) GetProcessingErrorOk() (*string, bool)`

GetProcessingErrorOk returns a tuple with the ProcessingError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingError

`func (o *FileUploadResponse) SetProcessingError(v string)`

SetProcessingError sets ProcessingError field to given value.

### HasProcessingError

`func (o *FileUploadResponse) HasProcessingError() bool`

HasProcessingError returns a boolean if a field has been set.

### GetResultingFiles

`func (o *FileUploadResponse) GetResultingFiles() []FilemanagerProcessingResultFile`

GetResultingFiles returns the ResultingFiles field if non-nil, zero value otherwise.

### GetResultingFilesOk

`func (o *FileUploadResponse) GetResultingFilesOk() (*[]FilemanagerProcessingResultFile, bool)`

GetResultingFilesOk returns a tuple with the ResultingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultingFiles

`func (o *FileUploadResponse) SetResultingFiles(v []FilemanagerProcessingResultFile)`

SetResultingFiles sets ResultingFiles field to given value.

### HasResultingFiles

`func (o *FileUploadResponse) HasResultingFiles() bool`

HasResultingFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


