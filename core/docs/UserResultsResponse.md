# UserResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]UserResponse**](UserResponse.md) |  | 
**TotalResults** | **int32** |  | 

## Methods

### NewUserResultsResponse

`func NewUserResultsResponse(results []UserResponse, totalResults int32, ) *UserResultsResponse`

NewUserResultsResponse instantiates a new UserResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResultsResponseWithDefaults

`func NewUserResultsResponseWithDefaults() *UserResultsResponse`

NewUserResultsResponseWithDefaults instantiates a new UserResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *UserResultsResponse) GetResults() []UserResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UserResultsResponse) GetResultsOk() (*[]UserResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UserResultsResponse) SetResults(v []UserResponse)`

SetResults sets Results field to given value.


### GetTotalResults

`func (o *UserResultsResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *UserResultsResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *UserResultsResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


