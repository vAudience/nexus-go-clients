# AIModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedFileMimetypes** | Pointer to **[]string** |  | [optional] 
**Actions** | **[]string** |  | 
**Capabilities** | Pointer to **[]string** | Note: only set when returning the model (not stored at model level), derived from features | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]AIModelFeature**](AIModelFeature.md) |  | [optional] 
**I18n** | Pointer to [**map[string]AIModelI18n**](AIModelI18n.md) |  | [optional] 
**Id** | **string** |  | 
**InternalId** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Lifecycle** | Pointer to **string** |  | [optional] 
**ModelCategory** | Pointer to **string** |  | [optional] 
**ModelId** | **string** |  | 
**ModelReleaseDate** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**OwnerId** | **string** |  | 
**OwnerOrganizationId** | **string** |  | 
**ParameterDefinitions** | Pointer to **map[string]interface{}** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**ServiceHostLocations** | [**[]HostingLocation**](HostingLocation.md) |  | 
**ServiceId** | **string** |  | 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewAIModel

`func NewAIModel(actions []string, id string, modelId string, name string, ownerId string, ownerOrganizationId string, serviceHostLocations []HostingLocation, serviceId string, ) *AIModel`

NewAIModel instantiates a new AIModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIModelWithDefaults

`func NewAIModelWithDefaults() *AIModel`

NewAIModelWithDefaults instantiates a new AIModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedFileMimetypes

`func (o *AIModel) GetAcceptedFileMimetypes() []string`

GetAcceptedFileMimetypes returns the AcceptedFileMimetypes field if non-nil, zero value otherwise.

### GetAcceptedFileMimetypesOk

`func (o *AIModel) GetAcceptedFileMimetypesOk() (*[]string, bool)`

GetAcceptedFileMimetypesOk returns a tuple with the AcceptedFileMimetypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedFileMimetypes

`func (o *AIModel) SetAcceptedFileMimetypes(v []string)`

SetAcceptedFileMimetypes sets AcceptedFileMimetypes field to given value.

### HasAcceptedFileMimetypes

`func (o *AIModel) HasAcceptedFileMimetypes() bool`

HasAcceptedFileMimetypes returns a boolean if a field has been set.

### GetActions

`func (o *AIModel) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AIModel) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AIModel) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetCapabilities

`func (o *AIModel) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AIModel) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AIModel) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AIModel) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AIModel) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AIModel) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AIModel) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AIModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *AIModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AIModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *AIModel) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *AIModel) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *AIModel) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *AIModel) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetFeatures

`func (o *AIModel) GetFeatures() []AIModelFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AIModel) GetFeaturesOk() (*[]AIModelFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AIModel) SetFeatures(v []AIModelFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AIModel) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetI18n

`func (o *AIModel) GetI18n() map[string]AIModelI18n`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *AIModel) GetI18nOk() (*map[string]AIModelI18n, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *AIModel) SetI18n(v map[string]AIModelI18n)`

SetI18n sets I18n field to given value.

### HasI18n

`func (o *AIModel) HasI18n() bool`

HasI18n returns a boolean if a field has been set.

### GetId

`func (o *AIModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AIModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AIModel) SetId(v string)`

SetId sets Id field to given value.


### GetInternalId

`func (o *AIModel) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *AIModel) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *AIModel) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *AIModel) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetIsPublic

`func (o *AIModel) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AIModel) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AIModel) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *AIModel) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetLifecycle

`func (o *AIModel) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *AIModel) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *AIModel) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *AIModel) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetModelCategory

`func (o *AIModel) GetModelCategory() string`

GetModelCategory returns the ModelCategory field if non-nil, zero value otherwise.

### GetModelCategoryOk

`func (o *AIModel) GetModelCategoryOk() (*string, bool)`

GetModelCategoryOk returns a tuple with the ModelCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCategory

`func (o *AIModel) SetModelCategory(v string)`

SetModelCategory sets ModelCategory field to given value.

### HasModelCategory

`func (o *AIModel) HasModelCategory() bool`

HasModelCategory returns a boolean if a field has been set.

### GetModelId

`func (o *AIModel) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AIModel) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AIModel) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetModelReleaseDate

`func (o *AIModel) GetModelReleaseDate() int64`

GetModelReleaseDate returns the ModelReleaseDate field if non-nil, zero value otherwise.

### GetModelReleaseDateOk

`func (o *AIModel) GetModelReleaseDateOk() (*int64, bool)`

GetModelReleaseDateOk returns a tuple with the ModelReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelReleaseDate

`func (o *AIModel) SetModelReleaseDate(v int64)`

SetModelReleaseDate sets ModelReleaseDate field to given value.

### HasModelReleaseDate

`func (o *AIModel) HasModelReleaseDate() bool`

HasModelReleaseDate returns a boolean if a field has been set.

### GetName

`func (o *AIModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIModel) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *AIModel) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AIModel) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AIModel) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerOrganizationId

`func (o *AIModel) GetOwnerOrganizationId() string`

GetOwnerOrganizationId returns the OwnerOrganizationId field if non-nil, zero value otherwise.

### GetOwnerOrganizationIdOk

`func (o *AIModel) GetOwnerOrganizationIdOk() (*string, bool)`

GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrganizationId

`func (o *AIModel) SetOwnerOrganizationId(v string)`

SetOwnerOrganizationId sets OwnerOrganizationId field to given value.


### GetParameterDefinitions

`func (o *AIModel) GetParameterDefinitions() map[string]interface{}`

GetParameterDefinitions returns the ParameterDefinitions field if non-nil, zero value otherwise.

### GetParameterDefinitionsOk

`func (o *AIModel) GetParameterDefinitionsOk() (*map[string]interface{}, bool)`

GetParameterDefinitionsOk returns a tuple with the ParameterDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterDefinitions

`func (o *AIModel) SetParameterDefinitions(v map[string]interface{})`

SetParameterDefinitions sets ParameterDefinitions field to given value.

### HasParameterDefinitions

`func (o *AIModel) HasParameterDefinitions() bool`

HasParameterDefinitions returns a boolean if a field has been set.

### GetParameters

`func (o *AIModel) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AIModel) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AIModel) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AIModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetServiceHostLocations

`func (o *AIModel) GetServiceHostLocations() []HostingLocation`

GetServiceHostLocations returns the ServiceHostLocations field if non-nil, zero value otherwise.

### GetServiceHostLocationsOk

`func (o *AIModel) GetServiceHostLocationsOk() (*[]HostingLocation, bool)`

GetServiceHostLocationsOk returns a tuple with the ServiceHostLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHostLocations

`func (o *AIModel) SetServiceHostLocations(v []HostingLocation)`

SetServiceHostLocations sets ServiceHostLocations field to given value.


### GetServiceId

`func (o *AIModel) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AIModel) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AIModel) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetUpdatedAt

`func (o *AIModel) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AIModel) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AIModel) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AIModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AIModel) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AIModel) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AIModel) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AIModel) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


