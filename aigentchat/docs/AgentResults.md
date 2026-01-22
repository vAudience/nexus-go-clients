# AgentResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Results** | [**[]Agent**](Agent.md) |  | 
**Stats** | Pointer to [**AgentStatistics**](AgentStatistics.md) |  | [optional] 
**TotalResults** | Pointer to **int64** |  | [optional] 

## Methods

### NewAgentResults

`func NewAgentResults(results []Agent, ) *AgentResults`

NewAgentResults instantiates a new AgentResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentResultsWithDefaults

`func NewAgentResultsWithDefaults() *AgentResults`

NewAgentResultsWithDefaults instantiates a new AgentResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *AgentResults) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AgentResults) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AgentResults) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AgentResults) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *AgentResults) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AgentResults) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AgentResults) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AgentResults) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetResults

`func (o *AgentResults) GetResults() []Agent`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AgentResults) GetResultsOk() (*[]Agent, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AgentResults) SetResults(v []Agent)`

SetResults sets Results field to given value.


### GetStats

`func (o *AgentResults) GetStats() AgentStatistics`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AgentResults) GetStatsOk() (*AgentStatistics, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AgentResults) SetStats(v AgentStatistics)`

SetStats sets Stats field to given value.

### HasStats

`func (o *AgentResults) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetTotalResults

`func (o *AgentResults) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *AgentResults) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *AgentResults) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *AgentResults) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


