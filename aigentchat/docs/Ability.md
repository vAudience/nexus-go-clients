# Ability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to [**[]AIModelConstraint**](AIModelConstraint.md) |  | [optional] 
**Type** | Pointer to [**AbilityType**](AbilityType.md) |  | [optional] 

## Methods

### NewAbility

`func NewAbility() *Ability`

NewAbility instantiates a new Ability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbilityWithDefaults

`func NewAbilityWithDefaults() *Ability`

NewAbilityWithDefaults instantiates a new Ability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *Ability) GetConstraints() []AIModelConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *Ability) GetConstraintsOk() (*[]AIModelConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *Ability) SetConstraints(v []AIModelConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *Ability) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetType

`func (o *Ability) GetType() AbilityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ability) GetTypeOk() (*AbilityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Ability) SetType(v AbilityType)`

SetType sets Type field to given value.

### HasType

`func (o *Ability) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


