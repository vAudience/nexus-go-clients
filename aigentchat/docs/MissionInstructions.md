# MissionInstructions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTokens** | Pointer to **int32** |  | [optional] 
**SystemPromptInjectionMode** | Pointer to [**PromptInjectionMode**](PromptInjectionMode.md) |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**Text** | **string** |  | 
**UserPromptInjectionMode** | Pointer to [**PromptInjectionMode**](PromptInjectionMode.md) |  | [optional] 
**VarReplacements** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMissionInstructions

`func NewMissionInstructions(text string, ) *MissionInstructions`

NewMissionInstructions instantiates a new MissionInstructions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissionInstructionsWithDefaults

`func NewMissionInstructionsWithDefaults() *MissionInstructions`

NewMissionInstructionsWithDefaults instantiates a new MissionInstructions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxTokens

`func (o *MissionInstructions) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *MissionInstructions) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *MissionInstructions) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *MissionInstructions) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetSystemPromptInjectionMode

`func (o *MissionInstructions) GetSystemPromptInjectionMode() PromptInjectionMode`

GetSystemPromptInjectionMode returns the SystemPromptInjectionMode field if non-nil, zero value otherwise.

### GetSystemPromptInjectionModeOk

`func (o *MissionInstructions) GetSystemPromptInjectionModeOk() (*PromptInjectionMode, bool)`

GetSystemPromptInjectionModeOk returns a tuple with the SystemPromptInjectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPromptInjectionMode

`func (o *MissionInstructions) SetSystemPromptInjectionMode(v PromptInjectionMode)`

SetSystemPromptInjectionMode sets SystemPromptInjectionMode field to given value.

### HasSystemPromptInjectionMode

`func (o *MissionInstructions) HasSystemPromptInjectionMode() bool`

HasSystemPromptInjectionMode returns a boolean if a field has been set.

### GetTemperature

`func (o *MissionInstructions) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *MissionInstructions) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *MissionInstructions) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *MissionInstructions) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetText

`func (o *MissionInstructions) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MissionInstructions) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MissionInstructions) SetText(v string)`

SetText sets Text field to given value.


### GetUserPromptInjectionMode

`func (o *MissionInstructions) GetUserPromptInjectionMode() PromptInjectionMode`

GetUserPromptInjectionMode returns the UserPromptInjectionMode field if non-nil, zero value otherwise.

### GetUserPromptInjectionModeOk

`func (o *MissionInstructions) GetUserPromptInjectionModeOk() (*PromptInjectionMode, bool)`

GetUserPromptInjectionModeOk returns a tuple with the UserPromptInjectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPromptInjectionMode

`func (o *MissionInstructions) SetUserPromptInjectionMode(v PromptInjectionMode)`

SetUserPromptInjectionMode sets UserPromptInjectionMode field to given value.

### HasUserPromptInjectionMode

`func (o *MissionInstructions) HasUserPromptInjectionMode() bool`

HasUserPromptInjectionMode returns a boolean if a field has been set.

### GetVarReplacements

`func (o *MissionInstructions) GetVarReplacements() map[string]string`

GetVarReplacements returns the VarReplacements field if non-nil, zero value otherwise.

### GetVarReplacementsOk

`func (o *MissionInstructions) GetVarReplacementsOk() (*map[string]string, bool)`

GetVarReplacementsOk returns a tuple with the VarReplacements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarReplacements

`func (o *MissionInstructions) SetVarReplacements(v map[string]string)`

SetVarReplacements sets VarReplacements field to given value.

### HasVarReplacements

`func (o *MissionInstructions) HasVarReplacements() bool`

HasVarReplacements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


