# AIModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedFileMimetypes** | Pointer to **[]string** |  | [optional] 
**Constraints** | Pointer to [**[]AIModelConstraint**](AIModelConstraint.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]AIModelFeature**](AIModelFeature.md) |  | [optional] 
**Id** | **string** |  | 
**InternalId** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**MaxInputTokens** | Pointer to **int32** |  | [optional] 
**MaxOutputTokens** | Pointer to **int32** |  | [optional] 
**ModelId** | **string** |  | 
**Name** | **string** |  | 
**OwnerId** | **string** |  | 
**OwnerOrganizationId** | **string** |  | 
**ParameterDefinitions** | Pointer to **map[string]interface{}** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**ServiceHostLocations** | Pointer to [**[]HostingLocation**](HostingLocation.md) |  | [optional] 
**ServiceId** | **string** |  | 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewAIModel

`func NewAIModel(id string, modelId string, name string, ownerId string, ownerOrganizationId string, serviceId string, ) *AIModel`

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

### GetConstraints

`func (o *AIModel) GetConstraints() []AIModelConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *AIModel) GetConstraintsOk() (*[]AIModelConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *AIModel) SetConstraints(v []AIModelConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *AIModel) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

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

### GetMaxInputTokens

`func (o *AIModel) GetMaxInputTokens() int32`

GetMaxInputTokens returns the MaxInputTokens field if non-nil, zero value otherwise.

### GetMaxInputTokensOk

`func (o *AIModel) GetMaxInputTokensOk() (*int32, bool)`

GetMaxInputTokensOk returns a tuple with the MaxInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInputTokens

`func (o *AIModel) SetMaxInputTokens(v int32)`

SetMaxInputTokens sets MaxInputTokens field to given value.

### HasMaxInputTokens

`func (o *AIModel) HasMaxInputTokens() bool`

HasMaxInputTokens returns a boolean if a field has been set.

### GetMaxOutputTokens

`func (o *AIModel) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *AIModel) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *AIModel) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *AIModel) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

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

### HasServiceHostLocations

`func (o *AIModel) HasServiceHostLocations() bool`

HasServiceHostLocations returns a boolean if a field has been set.

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


