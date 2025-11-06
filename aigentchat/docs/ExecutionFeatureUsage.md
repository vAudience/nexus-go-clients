# ExecutionFeatureUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | Pointer to [**AIModelCapability**](AIModelCapability.md) |  | [optional] 
**CostItems** | Pointer to [**[]ExecutionUsageCost**](ExecutionUsageCost.md) |  | [optional] 

## Methods

### NewExecutionFeatureUsage

`func NewExecutionFeatureUsage() *ExecutionFeatureUsage`

NewExecutionFeatureUsage instantiates a new ExecutionFeatureUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionFeatureUsageWithDefaults

`func NewExecutionFeatureUsageWithDefaults() *ExecutionFeatureUsage`

NewExecutionFeatureUsageWithDefaults instantiates a new ExecutionFeatureUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *ExecutionFeatureUsage) GetCapability() AIModelCapability`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *ExecutionFeatureUsage) GetCapabilityOk() (*AIModelCapability, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *ExecutionFeatureUsage) SetCapability(v AIModelCapability)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *ExecutionFeatureUsage) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetCostItems

`func (o *ExecutionFeatureUsage) GetCostItems() []ExecutionUsageCost`

GetCostItems returns the CostItems field if non-nil, zero value otherwise.

### GetCostItemsOk

`func (o *ExecutionFeatureUsage) GetCostItemsOk() (*[]ExecutionUsageCost, bool)`

GetCostItemsOk returns a tuple with the CostItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostItems

`func (o *ExecutionFeatureUsage) SetCostItems(v []ExecutionUsageCost)`

SetCostItems sets CostItems field to given value.

### HasCostItems

`func (o *ExecutionFeatureUsage) HasCostItems() bool`

HasCostItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


