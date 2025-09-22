# ExecutionCostTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchCostFactor** | Pointer to **float64** |  | [optional] 
**CostPerUnitInEuro** | Pointer to **float64** |  | [optional] 
**CostUnit** | Pointer to [**AIModelCostUnit**](AIModelCostUnit.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewExecutionCostTemplate

`func NewExecutionCostTemplate() *ExecutionCostTemplate`

NewExecutionCostTemplate instantiates a new ExecutionCostTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionCostTemplateWithDefaults

`func NewExecutionCostTemplateWithDefaults() *ExecutionCostTemplate`

NewExecutionCostTemplateWithDefaults instantiates a new ExecutionCostTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchCostFactor

`func (o *ExecutionCostTemplate) GetBatchCostFactor() float64`

GetBatchCostFactor returns the BatchCostFactor field if non-nil, zero value otherwise.

### GetBatchCostFactorOk

`func (o *ExecutionCostTemplate) GetBatchCostFactorOk() (*float64, bool)`

GetBatchCostFactorOk returns a tuple with the BatchCostFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchCostFactor

`func (o *ExecutionCostTemplate) SetBatchCostFactor(v float64)`

SetBatchCostFactor sets BatchCostFactor field to given value.

### HasBatchCostFactor

`func (o *ExecutionCostTemplate) HasBatchCostFactor() bool`

HasBatchCostFactor returns a boolean if a field has been set.

### GetCostPerUnitInEuro

`func (o *ExecutionCostTemplate) GetCostPerUnitInEuro() float64`

GetCostPerUnitInEuro returns the CostPerUnitInEuro field if non-nil, zero value otherwise.

### GetCostPerUnitInEuroOk

`func (o *ExecutionCostTemplate) GetCostPerUnitInEuroOk() (*float64, bool)`

GetCostPerUnitInEuroOk returns a tuple with the CostPerUnitInEuro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerUnitInEuro

`func (o *ExecutionCostTemplate) SetCostPerUnitInEuro(v float64)`

SetCostPerUnitInEuro sets CostPerUnitInEuro field to given value.

### HasCostPerUnitInEuro

`func (o *ExecutionCostTemplate) HasCostPerUnitInEuro() bool`

HasCostPerUnitInEuro returns a boolean if a field has been set.

### GetCostUnit

`func (o *ExecutionCostTemplate) GetCostUnit() AIModelCostUnit`

GetCostUnit returns the CostUnit field if non-nil, zero value otherwise.

### GetCostUnitOk

`func (o *ExecutionCostTemplate) GetCostUnitOk() (*AIModelCostUnit, bool)`

GetCostUnitOk returns a tuple with the CostUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUnit

`func (o *ExecutionCostTemplate) SetCostUnit(v AIModelCostUnit)`

SetCostUnit sets CostUnit field to given value.

### HasCostUnit

`func (o *ExecutionCostTemplate) HasCostUnit() bool`

HasCostUnit returns a boolean if a field has been set.

### GetDescription

`func (o *ExecutionCostTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExecutionCostTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExecutionCostTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExecutionCostTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


