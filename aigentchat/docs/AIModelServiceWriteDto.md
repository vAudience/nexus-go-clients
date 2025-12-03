# AIModelServiceWriteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiServiceId** | **string** |  | 
**CostMultiplier** | Pointer to **float64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**HostingLocations** | Pointer to [**map[string]HostingLocation**](HostingLocation.md) |  | [optional] 
**I18n** | Pointer to [**map[string]AIModelServiceI18n**](AIModelServiceI18n.md) |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**ServiceImpl** | [**AiServiceId**](AiServiceId.md) | Deprecated: use ai_service_id instead | 

## Methods

### NewAIModelServiceWriteDto

`func NewAIModelServiceWriteDto(aiServiceId string, name string, serviceImpl AiServiceId, ) *AIModelServiceWriteDto`

NewAIModelServiceWriteDto instantiates a new AIModelServiceWriteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIModelServiceWriteDtoWithDefaults

`func NewAIModelServiceWriteDtoWithDefaults() *AIModelServiceWriteDto`

NewAIModelServiceWriteDtoWithDefaults instantiates a new AIModelServiceWriteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiServiceId

`func (o *AIModelServiceWriteDto) GetAiServiceId() string`

GetAiServiceId returns the AiServiceId field if non-nil, zero value otherwise.

### GetAiServiceIdOk

`func (o *AIModelServiceWriteDto) GetAiServiceIdOk() (*string, bool)`

GetAiServiceIdOk returns a tuple with the AiServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiServiceId

`func (o *AIModelServiceWriteDto) SetAiServiceId(v string)`

SetAiServiceId sets AiServiceId field to given value.


### GetCostMultiplier

`func (o *AIModelServiceWriteDto) GetCostMultiplier() float64`

GetCostMultiplier returns the CostMultiplier field if non-nil, zero value otherwise.

### GetCostMultiplierOk

`func (o *AIModelServiceWriteDto) GetCostMultiplierOk() (*float64, bool)`

GetCostMultiplierOk returns a tuple with the CostMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostMultiplier

`func (o *AIModelServiceWriteDto) SetCostMultiplier(v float64)`

SetCostMultiplier sets CostMultiplier field to given value.

### HasCostMultiplier

`func (o *AIModelServiceWriteDto) HasCostMultiplier() bool`

HasCostMultiplier returns a boolean if a field has been set.

### GetDescription

`func (o *AIModelServiceWriteDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIModelServiceWriteDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIModelServiceWriteDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AIModelServiceWriteDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostingLocations

`func (o *AIModelServiceWriteDto) GetHostingLocations() map[string]HostingLocation`

GetHostingLocations returns the HostingLocations field if non-nil, zero value otherwise.

### GetHostingLocationsOk

`func (o *AIModelServiceWriteDto) GetHostingLocationsOk() (*map[string]HostingLocation, bool)`

GetHostingLocationsOk returns a tuple with the HostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingLocations

`func (o *AIModelServiceWriteDto) SetHostingLocations(v map[string]HostingLocation)`

SetHostingLocations sets HostingLocations field to given value.

### HasHostingLocations

`func (o *AIModelServiceWriteDto) HasHostingLocations() bool`

HasHostingLocations returns a boolean if a field has been set.

### GetI18n

`func (o *AIModelServiceWriteDto) GetI18n() map[string]AIModelServiceI18n`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *AIModelServiceWriteDto) GetI18nOk() (*map[string]AIModelServiceI18n, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *AIModelServiceWriteDto) SetI18n(v map[string]AIModelServiceI18n)`

SetI18n sets I18n field to given value.

### HasI18n

`func (o *AIModelServiceWriteDto) HasI18n() bool`

HasI18n returns a boolean if a field has been set.

### GetInternalId

`func (o *AIModelServiceWriteDto) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *AIModelServiceWriteDto) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *AIModelServiceWriteDto) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *AIModelServiceWriteDto) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetIsPublic

`func (o *AIModelServiceWriteDto) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AIModelServiceWriteDto) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AIModelServiceWriteDto) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *AIModelServiceWriteDto) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetName

`func (o *AIModelServiceWriteDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIModelServiceWriteDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIModelServiceWriteDto) SetName(v string)`

SetName sets Name field to given value.


### GetServiceImpl

`func (o *AIModelServiceWriteDto) GetServiceImpl() AiServiceId`

GetServiceImpl returns the ServiceImpl field if non-nil, zero value otherwise.

### GetServiceImplOk

`func (o *AIModelServiceWriteDto) GetServiceImplOk() (*AiServiceId, bool)`

GetServiceImplOk returns a tuple with the ServiceImpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceImpl

`func (o *AIModelServiceWriteDto) SetServiceImpl(v AiServiceId)`

SetServiceImpl sets ServiceImpl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


