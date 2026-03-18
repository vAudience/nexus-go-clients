# AIModelFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchCostFactor** | Pointer to **float64** |  | [optional] 
**Capability** | Pointer to [**AIModelCapability**](AIModelCapability.md) |  | [optional] 
**Constraints** | Pointer to [**[]AIModelConstraint**](AIModelConstraint.md) |  | [optional] 
**CostItemTemplates** | Pointer to [**[]ExecutionCostTemplate**](ExecutionCostTemplate.md) |  | [optional] 
**CostItems** | Pointer to [**[]ExecutionUsageCost**](ExecutionUsageCost.md) |  | [optional] 
**LongContextThresholdTokens** | Pointer to **int32** |  | [optional] 
**NonDefaultHostingLocationCostFactor** | Pointer to **float64** |  | [optional] 

## Methods

### NewAIModelFeature

`func NewAIModelFeature() *AIModelFeature`

NewAIModelFeature instantiates a new AIModelFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIModelFeatureWithDefaults

`func NewAIModelFeatureWithDefaults() *AIModelFeature`

NewAIModelFeatureWithDefaults instantiates a new AIModelFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchCostFactor

`func (o *AIModelFeature) GetBatchCostFactor() float64`

GetBatchCostFactor returns the BatchCostFactor field if non-nil, zero value otherwise.

### GetBatchCostFactorOk

`func (o *AIModelFeature) GetBatchCostFactorOk() (*float64, bool)`

GetBatchCostFactorOk returns a tuple with the BatchCostFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchCostFactor

`func (o *AIModelFeature) SetBatchCostFactor(v float64)`

SetBatchCostFactor sets BatchCostFactor field to given value.

### HasBatchCostFactor

`func (o *AIModelFeature) HasBatchCostFactor() bool`

HasBatchCostFactor returns a boolean if a field has been set.

### GetCapability

`func (o *AIModelFeature) GetCapability() AIModelCapability`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *AIModelFeature) GetCapabilityOk() (*AIModelCapability, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *AIModelFeature) SetCapability(v AIModelCapability)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *AIModelFeature) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetConstraints

`func (o *AIModelFeature) GetConstraints() []AIModelConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *AIModelFeature) GetConstraintsOk() (*[]AIModelConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *AIModelFeature) SetConstraints(v []AIModelConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *AIModelFeature) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetCostItemTemplates

`func (o *AIModelFeature) GetCostItemTemplates() []ExecutionCostTemplate`

GetCostItemTemplates returns the CostItemTemplates field if non-nil, zero value otherwise.

### GetCostItemTemplatesOk

`func (o *AIModelFeature) GetCostItemTemplatesOk() (*[]ExecutionCostTemplate, bool)`

GetCostItemTemplatesOk returns a tuple with the CostItemTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostItemTemplates

`func (o *AIModelFeature) SetCostItemTemplates(v []ExecutionCostTemplate)`

SetCostItemTemplates sets CostItemTemplates field to given value.

### HasCostItemTemplates

`func (o *AIModelFeature) HasCostItemTemplates() bool`

HasCostItemTemplates returns a boolean if a field has been set.

### GetCostItems

`func (o *AIModelFeature) GetCostItems() []ExecutionUsageCost`

GetCostItems returns the CostItems field if non-nil, zero value otherwise.

### GetCostItemsOk

`func (o *AIModelFeature) GetCostItemsOk() (*[]ExecutionUsageCost, bool)`

GetCostItemsOk returns a tuple with the CostItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostItems

`func (o *AIModelFeature) SetCostItems(v []ExecutionUsageCost)`

SetCostItems sets CostItems field to given value.

### HasCostItems

`func (o *AIModelFeature) HasCostItems() bool`

HasCostItems returns a boolean if a field has been set.

### GetLongContextThresholdTokens

`func (o *AIModelFeature) GetLongContextThresholdTokens() int32`

GetLongContextThresholdTokens returns the LongContextThresholdTokens field if non-nil, zero value otherwise.

### GetLongContextThresholdTokensOk

`func (o *AIModelFeature) GetLongContextThresholdTokensOk() (*int32, bool)`

GetLongContextThresholdTokensOk returns a tuple with the LongContextThresholdTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongContextThresholdTokens

`func (o *AIModelFeature) SetLongContextThresholdTokens(v int32)`

SetLongContextThresholdTokens sets LongContextThresholdTokens field to given value.

### HasLongContextThresholdTokens

`func (o *AIModelFeature) HasLongContextThresholdTokens() bool`

HasLongContextThresholdTokens returns a boolean if a field has been set.

### GetNonDefaultHostingLocationCostFactor

`func (o *AIModelFeature) GetNonDefaultHostingLocationCostFactor() float64`

GetNonDefaultHostingLocationCostFactor returns the NonDefaultHostingLocationCostFactor field if non-nil, zero value otherwise.

### GetNonDefaultHostingLocationCostFactorOk

`func (o *AIModelFeature) GetNonDefaultHostingLocationCostFactorOk() (*float64, bool)`

GetNonDefaultHostingLocationCostFactorOk returns a tuple with the NonDefaultHostingLocationCostFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDefaultHostingLocationCostFactor

`func (o *AIModelFeature) SetNonDefaultHostingLocationCostFactor(v float64)`

SetNonDefaultHostingLocationCostFactor sets NonDefaultHostingLocationCostFactor field to given value.

### HasNonDefaultHostingLocationCostFactor

`func (o *AIModelFeature) HasNonDefaultHostingLocationCostFactor() bool`

HasNonDefaultHostingLocationCostFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


