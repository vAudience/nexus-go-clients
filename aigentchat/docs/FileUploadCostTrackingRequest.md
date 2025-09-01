# FileUploadCostTrackingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | Pointer to **string** |  | [optional] 
**NumFiles** | **int32** |  | 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewFileUploadCostTrackingRequest

`func NewFileUploadCostTrackingRequest(numFiles int32, ) *FileUploadCostTrackingRequest`

NewFileUploadCostTrackingRequest instantiates a new FileUploadCostTrackingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUploadCostTrackingRequestWithDefaults

`func NewFileUploadCostTrackingRequestWithDefaults() *FileUploadCostTrackingRequest`

NewFileUploadCostTrackingRequestWithDefaults instantiates a new FileUploadCostTrackingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *FileUploadCostTrackingRequest) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *FileUploadCostTrackingRequest) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *FileUploadCostTrackingRequest) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *FileUploadCostTrackingRequest) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetNumFiles

`func (o *FileUploadCostTrackingRequest) GetNumFiles() int32`

GetNumFiles returns the NumFiles field if non-nil, zero value otherwise.

### GetNumFilesOk

`func (o *FileUploadCostTrackingRequest) GetNumFilesOk() (*int32, bool)`

GetNumFilesOk returns a tuple with the NumFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFiles

`func (o *FileUploadCostTrackingRequest) SetNumFiles(v int32)`

SetNumFiles sets NumFiles field to given value.


### GetUserId

`func (o *FileUploadCostTrackingRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FileUploadCostTrackingRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FileUploadCostTrackingRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FileUploadCostTrackingRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


