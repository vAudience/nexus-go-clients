# Mission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **int64** |  | [optional] 
**CompletionReason** | Pointer to [**MissionCompletionReason**](MissionCompletionReason.md) |  | [optional] 
**Content** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatorName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to [**MissionInstructions**](MissionInstructions.md) |  | [optional] 
**MissionExecutorId** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerOrganizationId** | Pointer to **string** |  | [optional] 
**StatusUpdates** | Pointer to [**MissionStatusUpdateList**](MissionStatusUpdateList.md) |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewMission

`func NewMission() *Mission`

NewMission instantiates a new Mission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissionWithDefaults

`func NewMissionWithDefaults() *Mission`

NewMissionWithDefaults instantiates a new Mission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *Mission) GetCompletedAt() int64`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Mission) GetCompletedAtOk() (*int64, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Mission) SetCompletedAt(v int64)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Mission) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCompletionReason

`func (o *Mission) GetCompletionReason() MissionCompletionReason`

GetCompletionReason returns the CompletionReason field if non-nil, zero value otherwise.

### GetCompletionReasonOk

`func (o *Mission) GetCompletionReasonOk() (*MissionCompletionReason, bool)`

GetCompletionReasonOk returns a tuple with the CompletionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionReason

`func (o *Mission) SetCompletionReason(v MissionCompletionReason)`

SetCompletionReason sets CompletionReason field to given value.

### HasCompletionReason

`func (o *Mission) HasCompletionReason() bool`

HasCompletionReason returns a boolean if a field has been set.

### GetContent

`func (o *Mission) GetContent() map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Mission) GetContentOk() (*map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Mission) SetContent(v map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *Mission) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Mission) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Mission) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Mission) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Mission) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Mission) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Mission) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Mission) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Mission) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatorName

`func (o *Mission) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *Mission) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *Mission) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *Mission) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetDescription

`func (o *Mission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Mission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Mission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Mission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Mission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Mission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Mission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Mission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstructions

`func (o *Mission) GetInstructions() MissionInstructions`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *Mission) GetInstructionsOk() (*MissionInstructions, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *Mission) SetInstructions(v MissionInstructions)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *Mission) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetMissionExecutorId

`func (o *Mission) GetMissionExecutorId() string`

GetMissionExecutorId returns the MissionExecutorId field if non-nil, zero value otherwise.

### GetMissionExecutorIdOk

`func (o *Mission) GetMissionExecutorIdOk() (*string, bool)`

GetMissionExecutorIdOk returns a tuple with the MissionExecutorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionExecutorId

`func (o *Mission) SetMissionExecutorId(v string)`

SetMissionExecutorId sets MissionExecutorId field to given value.

### HasMissionExecutorId

`func (o *Mission) HasMissionExecutorId() bool`

HasMissionExecutorId returns a boolean if a field has been set.

### GetOwnerId

`func (o *Mission) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Mission) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Mission) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Mission) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerOrganizationId

`func (o *Mission) GetOwnerOrganizationId() string`

GetOwnerOrganizationId returns the OwnerOrganizationId field if non-nil, zero value otherwise.

### GetOwnerOrganizationIdOk

`func (o *Mission) GetOwnerOrganizationIdOk() (*string, bool)`

GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrganizationId

`func (o *Mission) SetOwnerOrganizationId(v string)`

SetOwnerOrganizationId sets OwnerOrganizationId field to given value.

### HasOwnerOrganizationId

`func (o *Mission) HasOwnerOrganizationId() bool`

HasOwnerOrganizationId returns a boolean if a field has been set.

### GetStatusUpdates

`func (o *Mission) GetStatusUpdates() MissionStatusUpdateList`

GetStatusUpdates returns the StatusUpdates field if non-nil, zero value otherwise.

### GetStatusUpdatesOk

`func (o *Mission) GetStatusUpdatesOk() (*MissionStatusUpdateList, bool)`

GetStatusUpdatesOk returns a tuple with the StatusUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdates

`func (o *Mission) SetStatusUpdates(v MissionStatusUpdateList)`

SetStatusUpdates sets StatusUpdates field to given value.

### HasStatusUpdates

`func (o *Mission) HasStatusUpdates() bool`

HasStatusUpdates returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Mission) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Mission) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Mission) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Mission) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


