# MissionWriteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Instructions** | [**MissionInstructions**](MissionInstructions.md) |  | 

## Methods

### NewMissionWriteDto

`func NewMissionWriteDto(instructions MissionInstructions, ) *MissionWriteDto`

NewMissionWriteDto instantiates a new MissionWriteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissionWriteDtoWithDefaults

`func NewMissionWriteDtoWithDefaults() *MissionWriteDto`

NewMissionWriteDtoWithDefaults instantiates a new MissionWriteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MissionWriteDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MissionWriteDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MissionWriteDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MissionWriteDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstructions

`func (o *MissionWriteDto) GetInstructions() MissionInstructions`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *MissionWriteDto) GetInstructionsOk() (*MissionInstructions, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *MissionWriteDto) SetInstructions(v MissionInstructions)`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


