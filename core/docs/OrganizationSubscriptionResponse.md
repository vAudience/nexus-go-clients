# OrganizationSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**CancelAtPeriodEnd** | **bool** |  | 
**CreatedAt** | **string** |  | 
**Currency** | **string** |  | 
**CurrentPeriodEnd** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**OrganizationId** | **string** |  | 
**ProductId** | **string** |  | 
**ProductType** | [**ProductType**](ProductType.md) |  | 
**Seats** | **int32** |  | 
**SeatsTaken** | **int32** |  | 
**Status** | **string** |  | 
**TotalAmount** | **int32** | total amount in cents | 
**Trial** | **bool** |  | 
**TrialEnd** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewOrganizationSubscriptionResponse

`func NewOrganizationSubscriptionResponse(active bool, cancelAtPeriodEnd bool, createdAt string, currency string, currentPeriodEnd string, id string, name string, organizationId string, productId string, productType ProductType, seats int32, seatsTaken int32, status string, totalAmount int32, trial bool, trialEnd string, updatedAt string, ) *OrganizationSubscriptionResponse`

NewOrganizationSubscriptionResponse instantiates a new OrganizationSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSubscriptionResponseWithDefaults

`func NewOrganizationSubscriptionResponseWithDefaults() *OrganizationSubscriptionResponse`

NewOrganizationSubscriptionResponseWithDefaults instantiates a new OrganizationSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *OrganizationSubscriptionResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OrganizationSubscriptionResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OrganizationSubscriptionResponse) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCancelAtPeriodEnd

`func (o *OrganizationSubscriptionResponse) GetCancelAtPeriodEnd() bool`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *OrganizationSubscriptionResponse) GetCancelAtPeriodEndOk() (*bool, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *OrganizationSubscriptionResponse) SetCancelAtPeriodEnd(v bool)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.


### GetCreatedAt

`func (o *OrganizationSubscriptionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationSubscriptionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationSubscriptionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *OrganizationSubscriptionResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrganizationSubscriptionResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrganizationSubscriptionResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurrentPeriodEnd

`func (o *OrganizationSubscriptionResponse) GetCurrentPeriodEnd() string`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *OrganizationSubscriptionResponse) GetCurrentPeriodEndOk() (*string, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *OrganizationSubscriptionResponse) SetCurrentPeriodEnd(v string)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.


### GetId

`func (o *OrganizationSubscriptionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationSubscriptionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationSubscriptionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OrganizationSubscriptionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationSubscriptionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationSubscriptionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *OrganizationSubscriptionResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationSubscriptionResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationSubscriptionResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetProductId

`func (o *OrganizationSubscriptionResponse) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *OrganizationSubscriptionResponse) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *OrganizationSubscriptionResponse) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetProductType

`func (o *OrganizationSubscriptionResponse) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *OrganizationSubscriptionResponse) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *OrganizationSubscriptionResponse) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetSeats

`func (o *OrganizationSubscriptionResponse) GetSeats() int32`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *OrganizationSubscriptionResponse) GetSeatsOk() (*int32, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *OrganizationSubscriptionResponse) SetSeats(v int32)`

SetSeats sets Seats field to given value.


### GetSeatsTaken

`func (o *OrganizationSubscriptionResponse) GetSeatsTaken() int32`

GetSeatsTaken returns the SeatsTaken field if non-nil, zero value otherwise.

### GetSeatsTakenOk

`func (o *OrganizationSubscriptionResponse) GetSeatsTakenOk() (*int32, bool)`

GetSeatsTakenOk returns a tuple with the SeatsTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatsTaken

`func (o *OrganizationSubscriptionResponse) SetSeatsTaken(v int32)`

SetSeatsTaken sets SeatsTaken field to given value.


### GetStatus

`func (o *OrganizationSubscriptionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationSubscriptionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationSubscriptionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalAmount

`func (o *OrganizationSubscriptionResponse) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *OrganizationSubscriptionResponse) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *OrganizationSubscriptionResponse) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetTrial

`func (o *OrganizationSubscriptionResponse) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *OrganizationSubscriptionResponse) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *OrganizationSubscriptionResponse) SetTrial(v bool)`

SetTrial sets Trial field to given value.


### GetTrialEnd

`func (o *OrganizationSubscriptionResponse) GetTrialEnd() string`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *OrganizationSubscriptionResponse) GetTrialEndOk() (*string, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *OrganizationSubscriptionResponse) SetTrialEnd(v string)`

SetTrialEnd sets TrialEnd field to given value.


### GetUpdatedAt

`func (o *OrganizationSubscriptionResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationSubscriptionResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationSubscriptionResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


