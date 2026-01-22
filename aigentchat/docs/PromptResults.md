# PromptResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]Prompt**](Prompt.md) |  | 
**TotalResults** | **int64** |  | 

## Methods

### NewPromptResults

`func NewPromptResults(results []Prompt, totalResults int64, ) *PromptResults`

NewPromptResults instantiates a new PromptResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptResultsWithDefaults

`func NewPromptResultsWithDefaults() *PromptResults`

NewPromptResultsWithDefaults instantiates a new PromptResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


