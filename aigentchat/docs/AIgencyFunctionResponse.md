# AIgencyFunctionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionName** | **string** |  | 
**Id** | **string** |  | 
**IsError** | Pointer to **bool** |  | [optional] 
**Result** | **string** |  | 
**ToolFunctionId** | Pointer to **string** |  | [optional] 

## Methods

### NewAIgencyFunctionResponse

`func NewAIgencyFunctionResponse(functionName string, id string, result string, ) *AIgencyFunctionResponse`

NewAIgencyFunctionResponse instantiates a new AIgencyFunctionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIgencyFunctionResponseWithDefaults

`func NewAIgencyFunctionResponseWithDefaults() *AIgencyFunctionResponse`

NewAIgencyFunctionResponseWithDefaults instantiates a new AIgencyFunctionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionName

`func (o *AIgencyFunctionResponse) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *AIgencyFunctionResponse) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *AIgencyFunctionResponse) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetId

`func (o *AIgencyFunctionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AIgencyFunctionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AIgencyFunctionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsError

`func (o *AIgencyFunctionResponse) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *AIgencyFunctionResponse) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *AIgencyFunctionResponse) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *AIgencyFunctionResponse) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetResult

`func (o *AIgencyFunctionResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AIgencyFunctionResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AIgencyFunctionResponse) SetResult(v string)`

SetResult sets Result field to given value.


### GetToolFunctionId

`func (o *AIgencyFunctionResponse) GetToolFunctionId() string`

GetToolFunctionId returns the ToolFunctionId field if non-nil, zero value otherwise.

### GetToolFunctionIdOk

`func (o *AIgencyFunctionResponse) GetToolFunctionIdOk() (*string, bool)`

GetToolFunctionIdOk returns a tuple with the ToolFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolFunctionId

`func (o *AIgencyFunctionResponse) SetToolFunctionId(v string)`

SetToolFunctionId sets ToolFunctionId field to given value.

### HasToolFunctionId

`func (o *AIgencyFunctionResponse) HasToolFunctionId() bool`

HasToolFunctionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


