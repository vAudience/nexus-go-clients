# AgentQuickFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**I18n** | Pointer to [**map[string]AgentQuickFilterI18n**](AgentQuickFilterI18n.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**AgentFilterOptions**](AgentFilterOptions.md) |  | [optional] 

## Methods

### NewAgentQuickFilter

`func NewAgentQuickFilter() *AgentQuickFilter`

NewAgentQuickFilter instantiates a new AgentQuickFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentQuickFilterWithDefaults

`func NewAgentQuickFilterWithDefaults() *AgentQuickFilter`

NewAgentQuickFilterWithDefaults instantiates a new AgentQuickFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetI18n

`func (o *AgentQuickFilter) GetI18n() map[string]AgentQuickFilterI18n`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *AgentQuickFilter) GetI18nOk() (*map[string]AgentQuickFilterI18n, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *AgentQuickFilter) SetI18n(v map[string]AgentQuickFilterI18n)`

SetI18n sets I18n field to given value.

### HasI18n

`func (o *AgentQuickFilter) HasI18n() bool`

HasI18n returns a boolean if a field has been set.

### GetName

`func (o *AgentQuickFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentQuickFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentQuickFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentQuickFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *AgentQuickFilter) GetOptions() AgentFilterOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AgentQuickFilter) GetOptionsOk() (*AgentFilterOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AgentQuickFilter) SetOptions(v AgentFilterOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AgentQuickFilter) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


