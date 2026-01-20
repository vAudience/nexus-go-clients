# OrganizationDetailsResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]OrganizationDetailsResponse**](OrganizationDetailsResponse.md) |  | 
**TotalResults** | **int32** |  | 

## Methods

### NewOrganizationDetailsResultsResponse

`func NewOrganizationDetailsResultsResponse(results []OrganizationDetailsResponse, totalResults int32, ) *OrganizationDetailsResultsResponse`

NewOrganizationDetailsResultsResponse instantiates a new OrganizationDetailsResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDetailsResultsResponseWithDefaults

`func NewOrganizationDetailsResultsResponseWithDefaults() *OrganizationDetailsResultsResponse`

NewOrganizationDetailsResultsResponseWithDefaults instantiates a new OrganizationDetailsResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *OrganizationDetailsResultsResponse) GetResults() []OrganizationDetailsResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OrganizationDetailsResultsResponse) GetResultsOk() (*[]OrganizationDetailsResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OrganizationDetailsResultsResponse) SetResults(v []OrganizationDetailsResponse)`

SetResults sets Results field to given value.


### GetTotalResults

`func (o *OrganizationDetailsResultsResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *OrganizationDetailsResultsResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *OrganizationDetailsResultsResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


