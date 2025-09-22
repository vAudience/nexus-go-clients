# ChannelResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]Channel**](Channel.md) |  | 
**TotalResults** | Pointer to **int64** |  | [optional] 

## Methods

### NewChannelResults

`func NewChannelResults(results []Channel, ) *ChannelResults`

NewChannelResults instantiates a new ChannelResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelResultsWithDefaults

`func NewChannelResultsWithDefaults() *ChannelResults`

NewChannelResultsWithDefaults instantiates a new ChannelResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ChannelResults) GetResults() []Channel`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ChannelResults) GetResultsOk() (*[]Channel, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ChannelResults) SetResults(v []Channel)`

SetResults sets Results field to given value.


### GetTotalResults

`func (o *ChannelResults) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ChannelResults) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ChannelResults) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ChannelResults) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


