# TagResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Results** | [**[]Tag**](Tag.md) |  | 
**TotalResults** | Pointer to **int64** |  | [optional] 

## Methods

### NewTagResults

`func NewTagResults(results []Tag, ) *TagResults`

NewTagResults instantiates a new TagResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagResultsWithDefaults

`func NewTagResultsWithDefaults() *TagResults`

NewTagResultsWithDefaults instantiates a new TagResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *TagResults) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TagResults) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TagResults) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TagResults) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *TagResults) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TagResults) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TagResults) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TagResults) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetResults

`func (o *TagResults) GetResults() []Tag`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TagResults) GetResultsOk() (*[]Tag, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TagResults) SetResults(v []Tag)`

SetResults sets Results field to given value.


### GetTotalResults

`func (o *TagResults) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *TagResults) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *TagResults) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *TagResults) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


