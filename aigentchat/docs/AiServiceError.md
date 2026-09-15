# AiServiceError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | [**ApiErrorCode**](ApiErrorCode.md) |  | 
**Message** | **string** | Message is the client-facing summary of the failure, derived from ErrorCode by ClientMessageForAiServiceErrorCode. It never carries provider text: it is persisted on the message, fanned out over SSE and written into HTTP error bodies. | 
**MidStream** | Pointer to **bool** | MidStream is true when the provider failed after at least one delta of this call had already been delivered to the streaming callback. The content the client saw cannot be retracted, so AIgent never retries such an error, whatever its StatusCode. An in-band provider error on an already-open stream carries the stream&#39;s own HTTP status (typically 200) in OriginalStatusCode. | [optional] 
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


### GetMidStream

`func (o *AiServiceError) GetMidStream() bool`

GetMidStream returns the MidStream field if non-nil, zero value otherwise.

### GetMidStreamOk

`func (o *AiServiceError) GetMidStreamOk() (*bool, bool)`

GetMidStreamOk returns a tuple with the MidStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidStream

`func (o *AiServiceError) SetMidStream(v bool)`

SetMidStream sets MidStream field to given value.

### HasMidStream

`func (o *AiServiceError) HasMidStream() bool`

HasMidStream returns a boolean if a field has been set.

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


