# AIModelServiceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiServiceId** | **string** |  | 
**CostMultiplier** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**HostingLocations** | Pointer to [**map[string]HostingLocation**](HostingLocation.md) |  | [optional] 
**I18n** | Pointer to [**map[string]AIModelServiceI18n**](AIModelServiceI18n.md) |  | [optional] 
**Id** | **string** |  | 
**InternalId** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**OwnerId** | **string** |  | 
**OwnerOrganizationId** | **string** |  | 
**ServiceImpl** | [**AiServiceId**](AiServiceId.md) | Deprecated: use ai_service_id instead | 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewAIModelServiceObject

`func NewAIModelServiceObject(aiServiceId string, id string, name string, ownerId string, ownerOrganizationId string, serviceImpl AiServiceId, ) *AIModelServiceObject`

NewAIModelServiceObject instantiates a new AIModelServiceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIModelServiceObjectWithDefaults

`func NewAIModelServiceObjectWithDefaults() *AIModelServiceObject`

NewAIModelServiceObjectWithDefaults instantiates a new AIModelServiceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiServiceId

`func (o *AIModelServiceObject) GetAiServiceId() string`

GetAiServiceId returns the AiServiceId field if non-nil, zero value otherwise.

### GetAiServiceIdOk

`func (o *AIModelServiceObject) GetAiServiceIdOk() (*string, bool)`

GetAiServiceIdOk returns a tuple with the AiServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiServiceId

`func (o *AIModelServiceObject) SetAiServiceId(v string)`

SetAiServiceId sets AiServiceId field to given value.


### GetCostMultiplier

`func (o *AIModelServiceObject) GetCostMultiplier() float32`

GetCostMultiplier returns the CostMultiplier field if non-nil, zero value otherwise.

### GetCostMultiplierOk

`func (o *AIModelServiceObject) GetCostMultiplierOk() (*float32, bool)`

GetCostMultiplierOk returns a tuple with the CostMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostMultiplier

`func (o *AIModelServiceObject) SetCostMultiplier(v float32)`

SetCostMultiplier sets CostMultiplier field to given value.

### HasCostMultiplier

`func (o *AIModelServiceObject) HasCostMultiplier() bool`

HasCostMultiplier returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AIModelServiceObject) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AIModelServiceObject) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AIModelServiceObject) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AIModelServiceObject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *AIModelServiceObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIModelServiceObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIModelServiceObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AIModelServiceObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostingLocations

`func (o *AIModelServiceObject) GetHostingLocations() map[string]HostingLocation`

GetHostingLocations returns the HostingLocations field if non-nil, zero value otherwise.

### GetHostingLocationsOk

`func (o *AIModelServiceObject) GetHostingLocationsOk() (*map[string]HostingLocation, bool)`

GetHostingLocationsOk returns a tuple with the HostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingLocations

`func (o *AIModelServiceObject) SetHostingLocations(v map[string]HostingLocation)`

SetHostingLocations sets HostingLocations field to given value.

### HasHostingLocations

`func (o *AIModelServiceObject) HasHostingLocations() bool`

HasHostingLocations returns a boolean if a field has been set.

### GetI18n

`func (o *AIModelServiceObject) GetI18n() map[string]AIModelServiceI18n`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *AIModelServiceObject) GetI18nOk() (*map[string]AIModelServiceI18n, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *AIModelServiceObject) SetI18n(v map[string]AIModelServiceI18n)`

SetI18n sets I18n field to given value.

### HasI18n

`func (o *AIModelServiceObject) HasI18n() bool`

HasI18n returns a boolean if a field has been set.

### GetId

`func (o *AIModelServiceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AIModelServiceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AIModelServiceObject) SetId(v string)`

SetId sets Id field to given value.


### GetInternalId

`func (o *AIModelServiceObject) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *AIModelServiceObject) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *AIModelServiceObject) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *AIModelServiceObject) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetIsPublic

`func (o *AIModelServiceObject) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AIModelServiceObject) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AIModelServiceObject) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *AIModelServiceObject) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetName

`func (o *AIModelServiceObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIModelServiceObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIModelServiceObject) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *AIModelServiceObject) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AIModelServiceObject) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AIModelServiceObject) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerOrganizationId

`func (o *AIModelServiceObject) GetOwnerOrganizationId() string`

GetOwnerOrganizationId returns the OwnerOrganizationId field if non-nil, zero value otherwise.

### GetOwnerOrganizationIdOk

`func (o *AIModelServiceObject) GetOwnerOrganizationIdOk() (*string, bool)`

GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrganizationId

`func (o *AIModelServiceObject) SetOwnerOrganizationId(v string)`

SetOwnerOrganizationId sets OwnerOrganizationId field to given value.


### GetServiceImpl

`func (o *AIModelServiceObject) GetServiceImpl() AiServiceId`

GetServiceImpl returns the ServiceImpl field if non-nil, zero value otherwise.

### GetServiceImplOk

`func (o *AIModelServiceObject) GetServiceImplOk() (*AiServiceId, bool)`

GetServiceImplOk returns a tuple with the ServiceImpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceImpl

`func (o *AIModelServiceObject) SetServiceImpl(v AiServiceId)`

SetServiceImpl sets ServiceImpl field to given value.


### GetUpdatedAt

`func (o *AIModelServiceObject) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AIModelServiceObject) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AIModelServiceObject) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AIModelServiceObject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AIModelServiceObject) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AIModelServiceObject) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AIModelServiceObject) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AIModelServiceObject) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


