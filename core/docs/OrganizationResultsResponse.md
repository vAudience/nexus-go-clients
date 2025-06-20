# OrganizationResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]OrganizationResponse**](OrganizationResponse.md) |  | 
**TotalResults** | **int32** |  | 

## Methods

### NewOrganizationResultsResponse

`func NewOrganizationResultsResponse(results []OrganizationResponse, totalResults int32, ) *OrganizationResultsResponse`

NewOrganizationResultsResponse instantiates a new OrganizationResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationResultsResponseWithDefaults

`func NewOrganizationResultsResponseWithDefaults() *OrganizationResultsResponse`

NewOrganizationResultsResponseWithDefaults instantiates a new OrganizationResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *OrganizationResultsResponse) GetResults() []OrganizationResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OrganizationResultsResponse) GetResultsOk() (*[]OrganizationResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OrganizationResultsResponse) SetResults(v []OrganizationResponse)`

SetResults sets Results field to given value.


### GetTotalResults

`func (o *OrganizationResultsResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *OrganizationResultsResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *OrganizationResultsResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


