# PromptContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constant** | Pointer to [**PromptConstant**](PromptConstant.md) |  | [optional] 
**Prompt** | Pointer to [**PromptReference**](PromptReference.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Type** | [**PromptContentType**](PromptContentType.md) |  | 
**Variable** | Pointer to [**PromptVariable**](PromptVariable.md) |  | [optional] 

## Methods

### NewPromptContent

`func NewPromptContent(type_ PromptContentType, ) *PromptContent`

NewPromptContent instantiates a new PromptContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptContentWithDefaults

`func NewPromptContentWithDefaults() *PromptContent`

NewPromptContentWithDefaults instantiates a new PromptContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstant

`func (o *PromptContent) GetConstant() PromptConstant`

GetConstant returns the Constant field if non-nil, zero value otherwise.

### GetConstantOk

`func (o *PromptContent) GetConstantOk() (*PromptConstant, bool)`

GetConstantOk returns a tuple with the Constant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstant

`func (o *PromptContent) SetConstant(v PromptConstant)`

SetConstant sets Constant field to given value.

### HasConstant

`func (o *PromptContent) HasConstant() bool`

HasConstant returns a boolean if a field has been set.

### GetPrompt

`func (o *PromptContent) GetPrompt() PromptReference`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *PromptContent) GetPromptOk() (*PromptReference, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *PromptContent) SetPrompt(v PromptReference)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *PromptContent) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetText

`func (o *PromptContent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PromptContent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PromptContent) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PromptContent) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *PromptContent) GetType() PromptContentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptContent) GetTypeOk() (*PromptContentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptContent) SetType(v PromptContentType)`

SetType sets Type field to given value.


### GetVariable

`func (o *PromptContent) GetVariable() PromptVariable`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *PromptContent) GetVariableOk() (*PromptVariable, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *PromptContent) SetVariable(v PromptVariable)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *PromptContent) HasVariable() bool`

HasVariable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


