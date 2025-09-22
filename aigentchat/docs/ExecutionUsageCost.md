# ExecutionUsageCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchCostFactor** | Pointer to **float64** |  | [optional] 
**CostPerUnitInEuro** | Pointer to **float64** |  | [optional] 
**CostUnit** | Pointer to [**AIModelCostUnit**](AIModelCostUnit.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ResultingCostInEuro** | Pointer to **float64** |  | [optional] 
**UsedUnits** | Pointer to **float64** |  | [optional] 

## Methods

### NewExecutionUsageCost

`func NewExecutionUsageCost() *ExecutionUsageCost`

NewExecutionUsageCost instantiates a new ExecutionUsageCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionUsageCostWithDefaults

`func NewExecutionUsageCostWithDefaults() *ExecutionUsageCost`

NewExecutionUsageCostWithDefaults instantiates a new ExecutionUsageCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchCostFactor

`func (o *ExecutionUsageCost) GetBatchCostFactor() float64`

GetBatchCostFactor returns the BatchCostFactor field if non-nil, zero value otherwise.

### GetBatchCostFactorOk

`func (o *ExecutionUsageCost) GetBatchCostFactorOk() (*float64, bool)`

GetBatchCostFactorOk returns a tuple with the BatchCostFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchCostFactor

`func (o *ExecutionUsageCost) SetBatchCostFactor(v float64)`

SetBatchCostFactor sets BatchCostFactor field to given value.

### HasBatchCostFactor

`func (o *ExecutionUsageCost) HasBatchCostFactor() bool`

HasBatchCostFactor returns a boolean if a field has been set.

### GetCostPerUnitInEuro

`func (o *ExecutionUsageCost) GetCostPerUnitInEuro() float64`

GetCostPerUnitInEuro returns the CostPerUnitInEuro field if non-nil, zero value otherwise.

### GetCostPerUnitInEuroOk

`func (o *ExecutionUsageCost) GetCostPerUnitInEuroOk() (*float64, bool)`

GetCostPerUnitInEuroOk returns a tuple with the CostPerUnitInEuro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerUnitInEuro

`func (o *ExecutionUsageCost) SetCostPerUnitInEuro(v float64)`

SetCostPerUnitInEuro sets CostPerUnitInEuro field to given value.

### HasCostPerUnitInEuro

`func (o *ExecutionUsageCost) HasCostPerUnitInEuro() bool`

HasCostPerUnitInEuro returns a boolean if a field has been set.

### GetCostUnit

`func (o *ExecutionUsageCost) GetCostUnit() AIModelCostUnit`

GetCostUnit returns the CostUnit field if non-nil, zero value otherwise.

### GetCostUnitOk

`func (o *ExecutionUsageCost) GetCostUnitOk() (*AIModelCostUnit, bool)`

GetCostUnitOk returns a tuple with the CostUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUnit

`func (o *ExecutionUsageCost) SetCostUnit(v AIModelCostUnit)`

SetCostUnit sets CostUnit field to given value.

### HasCostUnit

`func (o *ExecutionUsageCost) HasCostUnit() bool`

HasCostUnit returns a boolean if a field has been set.

### GetDescription

`func (o *ExecutionUsageCost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExecutionUsageCost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExecutionUsageCost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExecutionUsageCost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResultingCostInEuro

`func (o *ExecutionUsageCost) GetResultingCostInEuro() float64`

GetResultingCostInEuro returns the ResultingCostInEuro field if non-nil, zero value otherwise.

### GetResultingCostInEuroOk

`func (o *ExecutionUsageCost) GetResultingCostInEuroOk() (*float64, bool)`

GetResultingCostInEuroOk returns a tuple with the ResultingCostInEuro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultingCostInEuro

`func (o *ExecutionUsageCost) SetResultingCostInEuro(v float64)`

SetResultingCostInEuro sets ResultingCostInEuro field to given value.

### HasResultingCostInEuro

`func (o *ExecutionUsageCost) HasResultingCostInEuro() bool`

HasResultingCostInEuro returns a boolean if a field has been set.

### GetUsedUnits

`func (o *ExecutionUsageCost) GetUsedUnits() float64`

GetUsedUnits returns the UsedUnits field if non-nil, zero value otherwise.

### GetUsedUnitsOk

`func (o *ExecutionUsageCost) GetUsedUnitsOk() (*float64, bool)`

GetUsedUnitsOk returns a tuple with the UsedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedUnits

`func (o *ExecutionUsageCost) SetUsedUnits(v float64)`

SetUsedUnits sets UsedUnits field to given value.

### HasUsedUnits

`func (o *ExecutionUsageCost) HasUsedUnits() bool`

HasUsedUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


