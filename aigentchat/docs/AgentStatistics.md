# AgentStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelOccurrences** | Pointer to **map[string]int32** |  | [optional] 
**TagOccurrences** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewAgentStatistics

`func NewAgentStatistics() *AgentStatistics`

NewAgentStatistics instantiates a new AgentStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentStatisticsWithDefaults

`func NewAgentStatisticsWithDefaults() *AgentStatistics`

NewAgentStatisticsWithDefaults instantiates a new AgentStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelOccurrences

`func (o *AgentStatistics) GetModelOccurrences() map[string]int32`

GetModelOccurrences returns the ModelOccurrences field if non-nil, zero value otherwise.

### GetModelOccurrencesOk

`func (o *AgentStatistics) GetModelOccurrencesOk() (*map[string]int32, bool)`

GetModelOccurrencesOk returns a tuple with the ModelOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelOccurrences

`func (o *AgentStatistics) SetModelOccurrences(v map[string]int32)`

SetModelOccurrences sets ModelOccurrences field to given value.

### HasModelOccurrences

`func (o *AgentStatistics) HasModelOccurrences() bool`

HasModelOccurrences returns a boolean if a field has been set.

### GetTagOccurrences

`func (o *AgentStatistics) GetTagOccurrences() map[string]int32`

GetTagOccurrences returns the TagOccurrences field if non-nil, zero value otherwise.

### GetTagOccurrencesOk

`func (o *AgentStatistics) GetTagOccurrencesOk() (*map[string]int32, bool)`

GetTagOccurrencesOk returns a tuple with the TagOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagOccurrences

`func (o *AgentStatistics) SetTagOccurrences(v map[string]int32)`

SetTagOccurrences sets TagOccurrences field to given value.

### HasTagOccurrences

`func (o *AgentStatistics) HasTagOccurrences() bool`

HasTagOccurrences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


