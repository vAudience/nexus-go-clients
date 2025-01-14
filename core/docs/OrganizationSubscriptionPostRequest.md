# OrganizationSubscriptionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodEnd** | Pointer to **string** |  | [optional] 
**ProductId** | **string** |  | 
**Seats** | **int32** |  | 

## Methods

### NewOrganizationSubscriptionPostRequest

`func NewOrganizationSubscriptionPostRequest(productId string, seats int32, ) *OrganizationSubscriptionPostRequest`

NewOrganizationSubscriptionPostRequest instantiates a new OrganizationSubscriptionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSubscriptionPostRequestWithDefaults

`func NewOrganizationSubscriptionPostRequestWithDefaults() *OrganizationSubscriptionPostRequest`

NewOrganizationSubscriptionPostRequestWithDefaults instantiates a new OrganizationSubscriptionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodEnd

`func (o *OrganizationSubscriptionPostRequest) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *OrganizationSubscriptionPostRequest) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *OrganizationSubscriptionPostRequest) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *OrganizationSubscriptionPostRequest) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetProductId

`func (o *OrganizationSubscriptionPostRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *OrganizationSubscriptionPostRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *OrganizationSubscriptionPostRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetSeats

`func (o *OrganizationSubscriptionPostRequest) GetSeats() int32`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *OrganizationSubscriptionPostRequest) GetSeatsOk() (*int32, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *OrganizationSubscriptionPostRequest) SetSeats(v int32)`

SetSeats sets Seats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


