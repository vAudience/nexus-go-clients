# MissionResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Mission**](Mission.md) |  | [optional] 
**TotalResults** | Pointer to **int64** |  | [optional] 

## Methods

### NewMissionResults

`func NewMissionResults() *MissionResults`

NewMissionResults instantiates a new MissionResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissionResultsWithDefaults

`func NewMissionResultsWithDefaults() *MissionResults`

NewMissionResultsWithDefaults instantiates a new MissionResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *MissionResults) GetResults() []Mission`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MissionResults) GetResultsOk() (*[]Mission, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MissionResults) SetResults(v []Mission)`

SetResults sets Results field to given value.

### HasResults

`func (o *MissionResults) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalResults

`func (o *MissionResults) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *MissionResults) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *MissionResults) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *MissionResults) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


