# AiServiceError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | [**ApiErrorCode**](ApiErrorCode.md) |  | 
**Message** | **string** |  | 
**OriginalStatusCode** | **int32** |  | 
**StatusCode** | **int32** |  | 

## Methods

### NewAiServiceError

`func NewAiServiceError(errorCode ApiErrorCode, message string, originalStatusCode int32, statusCode int32, ) *AiServiceError`

NewAiServiceError instantiates a new AiServiceError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiServiceErrorWithDefaults

`func NewAiServiceErrorWithDefaults() *AiServiceError`

NewAiServiceErrorWithDefaults instantiates a new AiServiceError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *AiServiceError) GetErrorCode() ApiErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AiServiceError) GetErrorCodeOk() (*ApiErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AiServiceError) SetErrorCode(v ApiErrorCode)`

SetErrorCode sets ErrorCode field to given value.


### GetMessage

`func (o *AiServiceError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiServiceError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiServiceError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOriginalStatusCode

`func (o *AiServiceError) GetOriginalStatusCode() int32`

GetOriginalStatusCode returns the OriginalStatusCode field if non-nil, zero value otherwise.

### GetOriginalStatusCodeOk

`func (o *AiServiceError) GetOriginalStatusCodeOk() (*int32, bool)`

GetOriginalStatusCodeOk returns a tuple with the OriginalStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStatusCode

`func (o *AiServiceError) SetOriginalStatusCode(v int32)`

SetOriginalStatusCode sets OriginalStatusCode field to given value.


### GetStatusCode

`func (o *AiServiceError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *AiServiceError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *AiServiceError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


