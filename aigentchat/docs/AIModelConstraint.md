# AIModelConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**AIModelConstraintDirection**](AIModelConstraintDirection.md) |  | [optional] 
**Max** | Pointer to **int32** |  | [optional] 
**Min** | Pointer to **int32** |  | [optional] 
**Unit** | Pointer to [**AIModelMinMaxUnit**](AIModelMinMaxUnit.md) |  | [optional] 

## Methods

### NewAIModelConstraint

`func NewAIModelConstraint() *AIModelConstraint`

NewAIModelConstraint instantiates a new AIModelConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIModelConstraintWithDefaults

`func NewAIModelConstraintWithDefaults() *AIModelConstraint`

NewAIModelConstraintWithDefaults instantiates a new AIModelConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *AIModelConstraint) GetDirection() AIModelConstraintDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AIModelConstraint) GetDirectionOk() (*AIModelConstraintDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AIModelConstraint) SetDirection(v AIModelConstraintDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *AIModelConstraint) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetMax

`func (o *AIModelConstraint) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *AIModelConstraint) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *AIModelConstraint) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *AIModelConstraint) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *AIModelConstraint) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *AIModelConstraint) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *AIModelConstraint) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *AIModelConstraint) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetUnit

`func (o *AIModelConstraint) GetUnit() AIModelMinMaxUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AIModelConstraint) GetUnitOk() (*AIModelMinMaxUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AIModelConstraint) SetUnit(v AIModelMinMaxUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AIModelConstraint) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


