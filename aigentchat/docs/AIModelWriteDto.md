# AIModelWriteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedFileMimetypes** | Pointer to **[]string** |  | [optional] 
**Actions** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]AIModelFeature**](AIModelFeature.md) |  | [optional] 
**I18n** | Pointer to [**map[string]AIModelI18n**](AIModelI18n.md) |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Lifecycle** | Pointer to **string** |  | [optional] 
**ModelCategory** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**ModelReleaseDate** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParameterDefinitions** | Pointer to **map[string]interface{}** |  | [optional] 
**Parameters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ServiceHostLocations** | Pointer to [**[]HostingLocation**](HostingLocation.md) |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 

## Methods

### NewAIModelWriteDto

`func NewAIModelWriteDto() *AIModelWriteDto`

NewAIModelWriteDto instantiates a new AIModelWriteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIModelWriteDtoWithDefaults

`func NewAIModelWriteDtoWithDefaults() *AIModelWriteDto`

NewAIModelWriteDtoWithDefaults instantiates a new AIModelWriteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedFileMimetypes

`func (o *AIModelWriteDto) GetAcceptedFileMimetypes() []string`

GetAcceptedFileMimetypes returns the AcceptedFileMimetypes field if non-nil, zero value otherwise.

### GetAcceptedFileMimetypesOk

`func (o *AIModelWriteDto) GetAcceptedFileMimetypesOk() (*[]string, bool)`

GetAcceptedFileMimetypesOk returns a tuple with the AcceptedFileMimetypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedFileMimetypes

`func (o *AIModelWriteDto) SetAcceptedFileMimetypes(v []string)`

SetAcceptedFileMimetypes sets AcceptedFileMimetypes field to given value.

### HasAcceptedFileMimetypes

`func (o *AIModelWriteDto) HasAcceptedFileMimetypes() bool`

HasAcceptedFileMimetypes returns a boolean if a field has been set.

### GetActions

`func (o *AIModelWriteDto) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AIModelWriteDto) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AIModelWriteDto) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AIModelWriteDto) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDescription

`func (o *AIModelWriteDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIModelWriteDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIModelWriteDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AIModelWriteDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *AIModelWriteDto) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *AIModelWriteDto) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *AIModelWriteDto) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *AIModelWriteDto) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetFeatures

`func (o *AIModelWriteDto) GetFeatures() []AIModelFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AIModelWriteDto) GetFeaturesOk() (*[]AIModelFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AIModelWriteDto) SetFeatures(v []AIModelFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AIModelWriteDto) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetI18n

`func (o *AIModelWriteDto) GetI18n() map[string]AIModelI18n`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *AIModelWriteDto) GetI18nOk() (*map[string]AIModelI18n, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *AIModelWriteDto) SetI18n(v map[string]AIModelI18n)`

SetI18n sets I18n field to given value.

### HasI18n

`func (o *AIModelWriteDto) HasI18n() bool`

HasI18n returns a boolean if a field has been set.

### GetInternalId

`func (o *AIModelWriteDto) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *AIModelWriteDto) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *AIModelWriteDto) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *AIModelWriteDto) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetIsPublic

`func (o *AIModelWriteDto) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AIModelWriteDto) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AIModelWriteDto) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *AIModelWriteDto) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetLifecycle

`func (o *AIModelWriteDto) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *AIModelWriteDto) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *AIModelWriteDto) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *AIModelWriteDto) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetModelCategory

`func (o *AIModelWriteDto) GetModelCategory() string`

GetModelCategory returns the ModelCategory field if non-nil, zero value otherwise.

### GetModelCategoryOk

`func (o *AIModelWriteDto) GetModelCategoryOk() (*string, bool)`

GetModelCategoryOk returns a tuple with the ModelCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCategory

`func (o *AIModelWriteDto) SetModelCategory(v string)`

SetModelCategory sets ModelCategory field to given value.

### HasModelCategory

`func (o *AIModelWriteDto) HasModelCategory() bool`

HasModelCategory returns a boolean if a field has been set.

### GetModelId

`func (o *AIModelWriteDto) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AIModelWriteDto) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AIModelWriteDto) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *AIModelWriteDto) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetModelReleaseDate

`func (o *AIModelWriteDto) GetModelReleaseDate() int32`

GetModelReleaseDate returns the ModelReleaseDate field if non-nil, zero value otherwise.

### GetModelReleaseDateOk

`func (o *AIModelWriteDto) GetModelReleaseDateOk() (*int32, bool)`

GetModelReleaseDateOk returns a tuple with the ModelReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelReleaseDate

`func (o *AIModelWriteDto) SetModelReleaseDate(v int32)`

SetModelReleaseDate sets ModelReleaseDate field to given value.

### HasModelReleaseDate

`func (o *AIModelWriteDto) HasModelReleaseDate() bool`

HasModelReleaseDate returns a boolean if a field has been set.

### GetName

`func (o *AIModelWriteDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIModelWriteDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIModelWriteDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AIModelWriteDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameterDefinitions

`func (o *AIModelWriteDto) GetParameterDefinitions() map[string]interface{}`

GetParameterDefinitions returns the ParameterDefinitions field if non-nil, zero value otherwise.

### GetParameterDefinitionsOk

`func (o *AIModelWriteDto) GetParameterDefinitionsOk() (*map[string]interface{}, bool)`

GetParameterDefinitionsOk returns a tuple with the ParameterDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterDefinitions

`func (o *AIModelWriteDto) SetParameterDefinitions(v map[string]interface{})`

SetParameterDefinitions sets ParameterDefinitions field to given value.

### HasParameterDefinitions

`func (o *AIModelWriteDto) HasParameterDefinitions() bool`

HasParameterDefinitions returns a boolean if a field has been set.

### GetParameters

`func (o *AIModelWriteDto) GetParameters() map[string]map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AIModelWriteDto) GetParametersOk() (*map[string]map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AIModelWriteDto) SetParameters(v map[string]map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AIModelWriteDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetServiceHostLocations

`func (o *AIModelWriteDto) GetServiceHostLocations() []HostingLocation`

GetServiceHostLocations returns the ServiceHostLocations field if non-nil, zero value otherwise.

### GetServiceHostLocationsOk

`func (o *AIModelWriteDto) GetServiceHostLocationsOk() (*[]HostingLocation, bool)`

GetServiceHostLocationsOk returns a tuple with the ServiceHostLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHostLocations

`func (o *AIModelWriteDto) SetServiceHostLocations(v []HostingLocation)`

SetServiceHostLocations sets ServiceHostLocations field to given value.

### HasServiceHostLocations

`func (o *AIModelWriteDto) HasServiceHostLocations() bool`

HasServiceHostLocations returns a boolean if a field has been set.

### GetServiceId

`func (o *AIModelWriteDto) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AIModelWriteDto) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AIModelWriteDto) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *AIModelWriteDto) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


