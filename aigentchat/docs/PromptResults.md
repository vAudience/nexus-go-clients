# PromptResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Results** | [**[]Prompt**](Prompt.md) |  | 
**Stats** | Pointer to [**PromptStatistics**](PromptStatistics.md) |  | [optional] 
**TotalResults** | Pointer to **int64** |  | [optional] 

## Methods

### NewPromptResults

`func NewPromptResults(results []Prompt, ) *PromptResults`

NewPromptResults instantiates a new PromptResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptResultsWithDefaults

`func NewPromptResultsWithDefaults() *PromptResults`

NewPromptResultsWithDefaults instantiates a new PromptResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PromptResults) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PromptResults) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PromptResults) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PromptResults) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *PromptResults) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PromptResults) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PromptResults) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PromptResults) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetResults

`func (o *PromptResults) GetResults() []Prompt`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PromptResults) GetResultsOk() (*[]Prompt, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PromptResults) SetResults(v []Prompt)`

SetResults sets Results field to given value.


### GetStats

`func (o *PromptResults) GetStats() PromptStatistics`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *PromptResults) GetStatsOk() (*PromptStatistics, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *PromptResults) SetStats(v PromptStatistics)`

SetStats sets Stats field to given value.

### HasStats

`func (o *PromptResults) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetTotalResults

`func (o *PromptResults) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *PromptResults) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *PromptResults) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *PromptResults) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


