# PriceTierResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitAmount** | Pointer to **int32** | price per seat in smallest currency unit (e.g. cents) | [optional] 
**UpTo** | Pointer to **int32** | max seats for this tier; 0 means unlimited (last tier) | [optional] 

## Methods

### NewPriceTierResponse

`func NewPriceTierResponse() *PriceTierResponse`

NewPriceTierResponse instantiates a new PriceTierResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTierResponseWithDefaults

`func NewPriceTierResponseWithDefaults() *PriceTierResponse`

NewPriceTierResponseWithDefaults instantiates a new PriceTierResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitAmount

`func (o *PriceTierResponse) GetUnitAmount() int32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *PriceTierResponse) GetUnitAmountOk() (*int32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *PriceTierResponse) SetUnitAmount(v int32)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmount

`func (o *PriceTierResponse) HasUnitAmount() bool`

HasUnitAmount returns a boolean if a field has been set.

### GetUpTo

`func (o *PriceTierResponse) GetUpTo() int32`

GetUpTo returns the UpTo field if non-nil, zero value otherwise.

### GetUpToOk

`func (o *PriceTierResponse) GetUpToOk() (*int32, bool)`

GetUpToOk returns a tuple with the UpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTo

`func (o *PriceTierResponse) SetUpTo(v int32)`

SetUpTo sets UpTo field to given value.

### HasUpTo

`func (o *PriceTierResponse) HasUpTo() bool`

HasUpTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


