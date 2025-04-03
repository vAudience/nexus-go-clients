# CostEstimation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostInEuro** | **float64** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCostEstimation

`func NewCostEstimation(costInEuro float64, ) *CostEstimation`

NewCostEstimation instantiates a new CostEstimation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostEstimationWithDefaults

`func NewCostEstimationWithDefaults() *CostEstimation`

NewCostEstimationWithDefaults instantiates a new CostEstimation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostInEuro

`func (o *CostEstimation) GetCostInEuro() float64`

GetCostInEuro returns the CostInEuro field if non-nil, zero value otherwise.

### GetCostInEuroOk

`func (o *CostEstimation) GetCostInEuroOk() (*float64, bool)`

GetCostInEuroOk returns a tuple with the CostInEuro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostInEuro

`func (o *CostEstimation) SetCostInEuro(v float64)`

SetCostInEuro sets CostInEuro field to given value.


### GetParameters

`func (o *CostEstimation) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CostEstimation) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CostEstimation) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CostEstimation) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


