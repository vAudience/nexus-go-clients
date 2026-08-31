# CompletionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **string** |  | [optional] 
**ClientMessageId** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to [**[]AIgencyMessage**](AIgencyMessage.md) |  | [optional] 
**RetryAfterMs** | Pointer to **int64** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewCompletionStatus

`func NewCompletionStatus() *CompletionStatus`

NewCompletionStatus instantiates a new CompletionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionStatusWithDefaults

`func NewCompletionStatusWithDefaults() *CompletionStatus`

NewCompletionStatusWithDefaults instantiates a new CompletionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *CompletionStatus) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *CompletionStatus) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *CompletionStatus) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *CompletionStatus) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetClientMessageId

`func (o *CompletionStatus) GetClientMessageId() string`

GetClientMessageId returns the ClientMessageId field if non-nil, zero value otherwise.

### GetClientMessageIdOk

`func (o *CompletionStatus) GetClientMessageIdOk() (*string, bool)`

GetClientMessageIdOk returns a tuple with the ClientMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMessageId

`func (o *CompletionStatus) SetClientMessageId(v string)`

SetClientMessageId sets ClientMessageId field to given value.

### HasClientMessageId

`func (o *CompletionStatus) HasClientMessageId() bool`

HasClientMessageId returns a boolean if a field has been set.

### GetMessages

`func (o *CompletionStatus) GetMessages() []AIgencyMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *CompletionStatus) GetMessagesOk() (*[]AIgencyMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *CompletionStatus) SetMessages(v []AIgencyMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *CompletionStatus) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetRetryAfterMs

`func (o *CompletionStatus) GetRetryAfterMs() int64`

GetRetryAfterMs returns the RetryAfterMs field if non-nil, zero value otherwise.

### GetRetryAfterMsOk

`func (o *CompletionStatus) GetRetryAfterMsOk() (*int64, bool)`

GetRetryAfterMsOk returns a tuple with the RetryAfterMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfterMs

`func (o *CompletionStatus) SetRetryAfterMs(v int64)`

SetRetryAfterMs sets RetryAfterMs field to given value.

### HasRetryAfterMs

`func (o *CompletionStatus) HasRetryAfterMs() bool`

HasRetryAfterMs returns a boolean if a field has been set.

### GetState

`func (o *CompletionStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CompletionStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CompletionStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CompletionStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


