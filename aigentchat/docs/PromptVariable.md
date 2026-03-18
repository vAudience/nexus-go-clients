# PromptVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Multi** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Options** | Pointer to [**[]PromptVariableOption**](PromptVariableOption.md) |  | [optional] 

## Methods

### NewPromptVariable

`func NewPromptVariable(name string, ) *PromptVariable`

NewPromptVariable instantiates a new PromptVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptVariableWithDefaults

`func NewPromptVariableWithDefaults() *PromptVariable`

NewPromptVariableWithDefaults instantiates a new PromptVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *PromptVariable) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PromptVariable) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PromptVariable) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PromptVariable) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *PromptVariable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PromptVariable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PromptVariable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PromptVariable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMulti

`func (o *PromptVariable) GetMulti() bool`

GetMulti returns the Multi field if non-nil, zero value otherwise.

### GetMultiOk

`func (o *PromptVariable) GetMultiOk() (*bool, bool)`

GetMultiOk returns a tuple with the Multi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulti

`func (o *PromptVariable) SetMulti(v bool)`

SetMulti sets Multi field to given value.

### HasMulti

`func (o *PromptVariable) HasMulti() bool`

HasMulti returns a boolean if a field has been set.

### GetName

`func (o *PromptVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptVariable) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *PromptVariable) GetOptions() []PromptVariableOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PromptVariable) GetOptionsOk() (*[]PromptVariableOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PromptVariable) SetOptions(v []PromptVariableOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PromptVariable) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


