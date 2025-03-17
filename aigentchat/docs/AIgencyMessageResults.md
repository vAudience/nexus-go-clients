# AIgencyMessageResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]AIgencyMessage**](AIgencyMessage.md) |  | 
**TotalResults** | **int64** |  | 

## Methods

### NewAIgencyMessageResults

`func NewAIgencyMessageResults(results []AIgencyMessage, totalResults int64, ) *AIgencyMessageResults`

NewAIgencyMessageResults instantiates a new AIgencyMessageResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIgencyMessageResultsWithDefaults

`func NewAIgencyMessageResultsWithDefaults() *AIgencyMessageResults`

NewAIgencyMessageResultsWithDefaults instantiates a new AIgencyMessageResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *AIgencyMessageResults) GetResults() []AIgencyMessage`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AIgencyMessageResults) GetResultsOk() (*[]AIgencyMessage, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AIgencyMessageResults) SetResults(v []AIgencyMessage)`

SetResults sets Results field to given value.


### GetTotalResults

`func (o *AIgencyMessageResults) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *AIgencyMessageResults) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *AIgencyMessageResults) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


