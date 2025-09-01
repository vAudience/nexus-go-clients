# AIgencyMessageContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**AIgencyMessageFile**](AIgencyMessageFile.md) | Currently unused in favor of attachment list | [optional] 
**FunctionCall** | Pointer to [**AIgencyFunctionCall**](AIgencyFunctionCall.md) |  | [optional] 
**FunctionResponses** | Pointer to [**AIgencyFunctionResponse**](AIgencyFunctionResponse.md) |  | [optional] 
**FunctionStatusUpdate** | Pointer to [**AIgencyFunctionStatusUpdate**](AIgencyFunctionStatusUpdate.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Thinking** | Pointer to [**AIgencyThinking**](AIgencyThinking.md) |  | [optional] 
**Type** | [**AIgencyMessageContentType**](AIgencyMessageContentType.md) |  | 

## Methods

### NewAIgencyMessageContent

`func NewAIgencyMessageContent(type_ AIgencyMessageContentType, ) *AIgencyMessageContent`

NewAIgencyMessageContent instantiates a new AIgencyMessageContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIgencyMessageContentWithDefaults

`func NewAIgencyMessageContentWithDefaults() *AIgencyMessageContent`

NewAIgencyMessageContentWithDefaults instantiates a new AIgencyMessageContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *AIgencyMessageContent) GetFile() AIgencyMessageFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AIgencyMessageContent) GetFileOk() (*AIgencyMessageFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AIgencyMessageContent) SetFile(v AIgencyMessageFile)`

SetFile sets File field to given value.

### HasFile

`func (o *AIgencyMessageContent) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFunctionCall

`func (o *AIgencyMessageContent) GetFunctionCall() AIgencyFunctionCall`

GetFunctionCall returns the FunctionCall field if non-nil, zero value otherwise.

### GetFunctionCallOk

`func (o *AIgencyMessageContent) GetFunctionCallOk() (*AIgencyFunctionCall, bool)`

GetFunctionCallOk returns a tuple with the FunctionCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCall

`func (o *AIgencyMessageContent) SetFunctionCall(v AIgencyFunctionCall)`

SetFunctionCall sets FunctionCall field to given value.

### HasFunctionCall

`func (o *AIgencyMessageContent) HasFunctionCall() bool`

HasFunctionCall returns a boolean if a field has been set.

### GetFunctionResponses

`func (o *AIgencyMessageContent) GetFunctionResponses() AIgencyFunctionResponse`

GetFunctionResponses returns the FunctionResponses field if non-nil, zero value otherwise.

### GetFunctionResponsesOk

`func (o *AIgencyMessageContent) GetFunctionResponsesOk() (*AIgencyFunctionResponse, bool)`

GetFunctionResponsesOk returns a tuple with the FunctionResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionResponses

`func (o *AIgencyMessageContent) SetFunctionResponses(v AIgencyFunctionResponse)`

SetFunctionResponses sets FunctionResponses field to given value.

### HasFunctionResponses

`func (o *AIgencyMessageContent) HasFunctionResponses() bool`

HasFunctionResponses returns a boolean if a field has been set.

### GetFunctionStatusUpdate

`func (o *AIgencyMessageContent) GetFunctionStatusUpdate() AIgencyFunctionStatusUpdate`

GetFunctionStatusUpdate returns the FunctionStatusUpdate field if non-nil, zero value otherwise.

### GetFunctionStatusUpdateOk

`func (o *AIgencyMessageContent) GetFunctionStatusUpdateOk() (*AIgencyFunctionStatusUpdate, bool)`

GetFunctionStatusUpdateOk returns a tuple with the FunctionStatusUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionStatusUpdate

`func (o *AIgencyMessageContent) SetFunctionStatusUpdate(v AIgencyFunctionStatusUpdate)`

SetFunctionStatusUpdate sets FunctionStatusUpdate field to given value.

### HasFunctionStatusUpdate

`func (o *AIgencyMessageContent) HasFunctionStatusUpdate() bool`

HasFunctionStatusUpdate returns a boolean if a field has been set.

### GetText

`func (o *AIgencyMessageContent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AIgencyMessageContent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AIgencyMessageContent) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *AIgencyMessageContent) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThinking

`func (o *AIgencyMessageContent) GetThinking() AIgencyThinking`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *AIgencyMessageContent) GetThinkingOk() (*AIgencyThinking, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *AIgencyMessageContent) SetThinking(v AIgencyThinking)`

SetThinking sets Thinking field to given value.

### HasThinking

`func (o *AIgencyMessageContent) HasThinking() bool`

HasThinking returns a boolean if a field has been set.

### GetType

`func (o *AIgencyMessageContent) GetType() AIgencyMessageContentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AIgencyMessageContent) GetTypeOk() (*AIgencyMessageContentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AIgencyMessageContent) SetType(v AIgencyMessageContentType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


