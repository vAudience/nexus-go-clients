# ChatCompletionCostTrackingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batch** | Pointer to **bool** |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 
**InputTokens** | Pointer to **int32** |  | [optional] 
**ModelInternalId** | **string** |  | 
**OutputTokens** | Pointer to **int32** |  | [optional] 
**ServiceInternalId** | **string** |  | 
**Streaming** | Pointer to **bool** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewChatCompletionCostTrackingRequest

`func NewChatCompletionCostTrackingRequest(modelInternalId string, serviceInternalId string, ) *ChatCompletionCostTrackingRequest`

NewChatCompletionCostTrackingRequest instantiates a new ChatCompletionCostTrackingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionCostTrackingRequestWithDefaults

`func NewChatCompletionCostTrackingRequestWithDefaults() *ChatCompletionCostTrackingRequest`

NewChatCompletionCostTrackingRequestWithDefaults instantiates a new ChatCompletionCostTrackingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatch

`func (o *ChatCompletionCostTrackingRequest) GetBatch() bool`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *ChatCompletionCostTrackingRequest) GetBatchOk() (*bool, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *ChatCompletionCostTrackingRequest) SetBatch(v bool)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *ChatCompletionCostTrackingRequest) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetExecutionId

`func (o *ChatCompletionCostTrackingRequest) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ChatCompletionCostTrackingRequest) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ChatCompletionCostTrackingRequest) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ChatCompletionCostTrackingRequest) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetInputTokens

`func (o *ChatCompletionCostTrackingRequest) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *ChatCompletionCostTrackingRequest) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *ChatCompletionCostTrackingRequest) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *ChatCompletionCostTrackingRequest) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetModelInternalId

`func (o *ChatCompletionCostTrackingRequest) GetModelInternalId() string`

GetModelInternalId returns the ModelInternalId field if non-nil, zero value otherwise.

### GetModelInternalIdOk

`func (o *ChatCompletionCostTrackingRequest) GetModelInternalIdOk() (*string, bool)`

GetModelInternalIdOk returns a tuple with the ModelInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelInternalId

`func (o *ChatCompletionCostTrackingRequest) SetModelInternalId(v string)`

SetModelInternalId sets ModelInternalId field to given value.


### GetOutputTokens

`func (o *ChatCompletionCostTrackingRequest) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *ChatCompletionCostTrackingRequest) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *ChatCompletionCostTrackingRequest) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *ChatCompletionCostTrackingRequest) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetServiceInternalId

`func (o *ChatCompletionCostTrackingRequest) GetServiceInternalId() string`

GetServiceInternalId returns the ServiceInternalId field if non-nil, zero value otherwise.

### GetServiceInternalIdOk

`func (o *ChatCompletionCostTrackingRequest) GetServiceInternalIdOk() (*string, bool)`

GetServiceInternalIdOk returns a tuple with the ServiceInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInternalId

`func (o *ChatCompletionCostTrackingRequest) SetServiceInternalId(v string)`

SetServiceInternalId sets ServiceInternalId field to given value.


### GetStreaming

`func (o *ChatCompletionCostTrackingRequest) GetStreaming() bool`

GetStreaming returns the Streaming field if non-nil, zero value otherwise.

### GetStreamingOk

`func (o *ChatCompletionCostTrackingRequest) GetStreamingOk() (*bool, bool)`

GetStreamingOk returns a tuple with the Streaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreaming

`func (o *ChatCompletionCostTrackingRequest) SetStreaming(v bool)`

SetStreaming sets Streaming field to given value.

### HasStreaming

`func (o *ChatCompletionCostTrackingRequest) HasStreaming() bool`

HasStreaming returns a boolean if a field has been set.

### GetUserId

`func (o *ChatCompletionCostTrackingRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ChatCompletionCostTrackingRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ChatCompletionCostTrackingRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ChatCompletionCostTrackingRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


