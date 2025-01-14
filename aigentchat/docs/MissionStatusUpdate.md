# MissionStatusUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MissionId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**MissionStatus**](MissionStatus.md) |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 

## Methods

### NewMissionStatusUpdate

`func NewMissionStatusUpdate() *MissionStatusUpdate`

NewMissionStatusUpdate instantiates a new MissionStatusUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissionStatusUpdateWithDefaults

`func NewMissionStatusUpdateWithDefaults() *MissionStatusUpdate`

NewMissionStatusUpdateWithDefaults instantiates a new MissionStatusUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MissionStatusUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MissionStatusUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MissionStatusUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MissionStatusUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *MissionStatusUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MissionStatusUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MissionStatusUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MissionStatusUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMissionId

`func (o *MissionStatusUpdate) GetMissionId() string`

GetMissionId returns the MissionId field if non-nil, zero value otherwise.

### GetMissionIdOk

`func (o *MissionStatusUpdate) GetMissionIdOk() (*string, bool)`

GetMissionIdOk returns a tuple with the MissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionId

`func (o *MissionStatusUpdate) SetMissionId(v string)`

SetMissionId sets MissionId field to given value.

### HasMissionId

`func (o *MissionStatusUpdate) HasMissionId() bool`

HasMissionId returns a boolean if a field has been set.

### GetStatus

`func (o *MissionStatusUpdate) GetStatus() MissionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MissionStatusUpdate) GetStatusOk() (*MissionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MissionStatusUpdate) SetStatus(v MissionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MissionStatusUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *MissionStatusUpdate) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MissionStatusUpdate) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MissionStatusUpdate) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MissionStatusUpdate) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


