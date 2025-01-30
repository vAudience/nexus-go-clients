# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedTools** | Pointer to **[]string** |  | [optional] 
**AttachedFileIds** | Pointer to **[]string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Capabilities** | Pointer to [**[]AgentCapability**](AgentCapability.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**InitialUserMessages** | Pointer to **[]string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**ModelHostLocation** | Pointer to [**HostingLocation**](HostingLocation.md) |  | [optional] 
**ModelId** | **string** |  | 
**Name** | **string** |  | 
**OwnerId** | **string** |  | 
**OwnerOrganizationId** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**SystemMessages** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**UseTools** | Pointer to **bool** |  | [optional] 

## Methods

### NewAgent

`func NewAgent(id string, modelId string, name string, ownerId string, ownerOrganizationId string, ) *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedTools

`func (o *Agent) GetAssignedTools() []string`

GetAssignedTools returns the AssignedTools field if non-nil, zero value otherwise.

### GetAssignedToolsOk

`func (o *Agent) GetAssignedToolsOk() (*[]string, bool)`

GetAssignedToolsOk returns a tuple with the AssignedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTools

`func (o *Agent) SetAssignedTools(v []string)`

SetAssignedTools sets AssignedTools field to given value.

### HasAssignedTools

`func (o *Agent) HasAssignedTools() bool`

HasAssignedTools returns a boolean if a field has been set.

### GetAttachedFileIds

`func (o *Agent) GetAttachedFileIds() []string`

GetAttachedFileIds returns the AttachedFileIds field if non-nil, zero value otherwise.

### GetAttachedFileIdsOk

`func (o *Agent) GetAttachedFileIdsOk() (*[]string, bool)`

GetAttachedFileIdsOk returns a tuple with the AttachedFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedFileIds

`func (o *Agent) SetAttachedFileIds(v []string)`

SetAttachedFileIds sets AttachedFileIds field to given value.

### HasAttachedFileIds

`func (o *Agent) HasAttachedFileIds() bool`

HasAttachedFileIds returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *Agent) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Agent) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Agent) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *Agent) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetCapabilities

`func (o *Agent) GetCapabilities() []AgentCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Agent) GetCapabilitiesOk() (*[]AgentCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Agent) SetCapabilities(v []AgentCapability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Agent) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Agent) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Agent) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Agent) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Agent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Agent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Agent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Agent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Agent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Agent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agent) SetId(v string)`

SetId sets Id field to given value.


### GetInitialUserMessages

`func (o *Agent) GetInitialUserMessages() []string`

GetInitialUserMessages returns the InitialUserMessages field if non-nil, zero value otherwise.

### GetInitialUserMessagesOk

`func (o *Agent) GetInitialUserMessagesOk() (*[]string, bool)`

GetInitialUserMessagesOk returns a tuple with the InitialUserMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUserMessages

`func (o *Agent) SetInitialUserMessages(v []string)`

SetInitialUserMessages sets InitialUserMessages field to given value.

### HasInitialUserMessages

`func (o *Agent) HasInitialUserMessages() bool`

HasInitialUserMessages returns a boolean if a field has been set.

### GetInternalId

`func (o *Agent) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *Agent) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *Agent) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *Agent) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetIsPublic

`func (o *Agent) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Agent) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Agent) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *Agent) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetMetaData

`func (o *Agent) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *Agent) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *Agent) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *Agent) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetModelHostLocation

`func (o *Agent) GetModelHostLocation() HostingLocation`

GetModelHostLocation returns the ModelHostLocation field if non-nil, zero value otherwise.

### GetModelHostLocationOk

`func (o *Agent) GetModelHostLocationOk() (*HostingLocation, bool)`

GetModelHostLocationOk returns a tuple with the ModelHostLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelHostLocation

`func (o *Agent) SetModelHostLocation(v HostingLocation)`

SetModelHostLocation sets ModelHostLocation field to given value.

### HasModelHostLocation

`func (o *Agent) HasModelHostLocation() bool`

HasModelHostLocation returns a boolean if a field has been set.

### GetModelId

`func (o *Agent) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *Agent) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *Agent) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetName

`func (o *Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agent) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *Agent) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Agent) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Agent) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerOrganizationId

`func (o *Agent) GetOwnerOrganizationId() string`

GetOwnerOrganizationId returns the OwnerOrganizationId field if non-nil, zero value otherwise.

### GetOwnerOrganizationIdOk

`func (o *Agent) GetOwnerOrganizationIdOk() (*string, bool)`

GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrganizationId

`func (o *Agent) SetOwnerOrganizationId(v string)`

SetOwnerOrganizationId sets OwnerOrganizationId field to given value.


### GetParameters

`func (o *Agent) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Agent) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Agent) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Agent) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSystemMessages

`func (o *Agent) GetSystemMessages() []string`

GetSystemMessages returns the SystemMessages field if non-nil, zero value otherwise.

### GetSystemMessagesOk

`func (o *Agent) GetSystemMessagesOk() (*[]string, bool)`

GetSystemMessagesOk returns a tuple with the SystemMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMessages

`func (o *Agent) SetSystemMessages(v []string)`

SetSystemMessages sets SystemMessages field to given value.

### HasSystemMessages

`func (o *Agent) HasSystemMessages() bool`

HasSystemMessages returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Agent) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Agent) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Agent) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Agent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Agent) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Agent) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Agent) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Agent) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUseTools

`func (o *Agent) GetUseTools() bool`

GetUseTools returns the UseTools field if non-nil, zero value otherwise.

### GetUseToolsOk

`func (o *Agent) GetUseToolsOk() (*bool, bool)`

GetUseToolsOk returns a tuple with the UseTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTools

`func (o *Agent) SetUseTools(v bool)`

SetUseTools sets UseTools field to given value.

### HasUseTools

`func (o *Agent) HasUseTools() bool`

HasUseTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


