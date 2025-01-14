# OrganizationSubscriptionPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodEnd** | Pointer to **string** |  | [optional] 
**Seats** | **int32** |  | 

## Methods

### NewOrganizationSubscriptionPatchRequest

`func NewOrganizationSubscriptionPatchRequest(seats int32, ) *OrganizationSubscriptionPatchRequest`

NewOrganizationSubscriptionPatchRequest instantiates a new OrganizationSubscriptionPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSubscriptionPatchRequestWithDefaults

`func NewOrganizationSubscriptionPatchRequestWithDefaults() *OrganizationSubscriptionPatchRequest`

NewOrganizationSubscriptionPatchRequestWithDefaults instantiates a new OrganizationSubscriptionPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodEnd

`func (o *OrganizationSubscriptionPatchRequest) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *OrganizationSubscriptionPatchRequest) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *OrganizationSubscriptionPatchRequest) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *OrganizationSubscriptionPatchRequest) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetSeats

`func (o *OrganizationSubscriptionPatchRequest) GetSeats() int32`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *OrganizationSubscriptionPatchRequest) GetSeatsOk() (*int32, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *OrganizationSubscriptionPatchRequest) SetSeats(v int32)`

SetSeats sets Seats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


