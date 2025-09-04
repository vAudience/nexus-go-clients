# FileUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionLogId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**ResultingFiles** | [**[]ResultFile**](ResultFile.md) |  | 

## Methods

### NewFileUploadResponse

`func NewFileUploadResponse(id string, resultingFiles []ResultFile, ) *FileUploadResponse`

NewFileUploadResponse instantiates a new FileUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUploadResponseWithDefaults

`func NewFileUploadResponseWithDefaults() *FileUploadResponse`

NewFileUploadResponseWithDefaults instantiates a new FileUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionLogId

`func (o *FileUploadResponse) GetExecutionLogId() string`

GetExecutionLogId returns the ExecutionLogId field if non-nil, zero value otherwise.

### GetExecutionLogIdOk

`func (o *FileUploadResponse) GetExecutionLogIdOk() (*string, bool)`

GetExecutionLogIdOk returns a tuple with the ExecutionLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogId

`func (o *FileUploadResponse) SetExecutionLogId(v string)`

SetExecutionLogId sets ExecutionLogId field to given value.

### HasExecutionLogId

`func (o *FileUploadResponse) HasExecutionLogId() bool`

HasExecutionLogId returns a boolean if a field has been set.

### GetId

`func (o *FileUploadResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileUploadResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileUploadResponse) SetId(v string)`

SetId sets Id field to given value.


### GetResultingFiles

`func (o *FileUploadResponse) GetResultingFiles() []ResultFile`

GetResultingFiles returns the ResultingFiles field if non-nil, zero value otherwise.

### GetResultingFilesOk

`func (o *FileUploadResponse) GetResultingFilesOk() (*[]ResultFile, bool)`

GetResultingFilesOk returns a tuple with the ResultingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultingFiles

`func (o *FileUploadResponse) SetResultingFiles(v []ResultFile)`

SetResultingFiles sets ResultingFiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


