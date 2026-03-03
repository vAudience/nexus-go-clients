# PromptReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Vars** | Pointer to **map[string]string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewPromptReference

`func NewPromptReference(id string, ) *PromptReference`

NewPromptReference instantiates a new PromptReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptReferenceWithDefaults

`func NewPromptReferenceWithDefaults() *PromptReference`

NewPromptReferenceWithDefaults instantiates a new PromptReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromptReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromptReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromptReference) SetId(v string)`

SetId sets Id field to given value.


### GetVars

`func (o *PromptReference) GetVars() map[string]string`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *PromptReference) GetVarsOk() (*map[string]string, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *PromptReference) SetVars(v map[string]string)`

SetVars sets Vars field to given value.

### HasVars

`func (o *PromptReference) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetVersion

`func (o *PromptReference) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PromptReference) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PromptReference) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PromptReference) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


