# OrganizationSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioAllowedHostingLocations** | Pointer to **[]string** |  | [optional] 
**ChatAgentQuickFilters** | Pointer to [**[]AgentQuickFilter**](AgentQuickFilter.md) |  | [optional] 
**ChatAllowedHostingLocations** | Pointer to **[]string** |  | [optional] 
**ChatDefaultAgentId** | Pointer to **string** |  | [optional] 
**CreatedAt** | **string** |  | 
**EmbeddingAllowedHostingLocations** | Pointer to **[]string** |  | [optional] 
**EmbeddingDefaultAgentId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**ImageAllowedHostingLocations** | Pointer to **[]string** |  | [optional] 
**ImageDefaultAgentId** | Pointer to **string** |  | [optional] 
**OrganizationId** | **string** |  | 
**ToolsAllowedHostingLocations** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | **string** |  | 

## Methods

### NewOrganizationSettingsResponse

`func NewOrganizationSettingsResponse(createdAt string, id string, organizationId string, updatedAt string, ) *OrganizationSettingsResponse`

NewOrganizationSettingsResponse instantiates a new OrganizationSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSettingsResponseWithDefaults

`func NewOrganizationSettingsResponseWithDefaults() *OrganizationSettingsResponse`

NewOrganizationSettingsResponseWithDefaults instantiates a new OrganizationSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioAllowedHostingLocations

`func (o *OrganizationSettingsResponse) GetAudioAllowedHostingLocations() []string`

GetAudioAllowedHostingLocations returns the AudioAllowedHostingLocations field if non-nil, zero value otherwise.

### GetAudioAllowedHostingLocationsOk

`func (o *OrganizationSettingsResponse) GetAudioAllowedHostingLocationsOk() (*[]string, bool)`

GetAudioAllowedHostingLocationsOk returns a tuple with the AudioAllowedHostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioAllowedHostingLocations

`func (o *OrganizationSettingsResponse) SetAudioAllowedHostingLocations(v []string)`

SetAudioAllowedHostingLocations sets AudioAllowedHostingLocations field to given value.

### HasAudioAllowedHostingLocations

`func (o *OrganizationSettingsResponse) HasAudioAllowedHostingLocations() bool`

HasAudioAllowedHostingLocations returns a boolean if a field has been set.

### GetChatAgentQuickFilters

`func (o *OrganizationSettingsResponse) GetChatAgentQuickFilters() []AgentQuickFilter`

GetChatAgentQuickFilters returns the ChatAgentQuickFilters field if non-nil, zero value otherwise.

### GetChatAgentQuickFiltersOk

`func (o *OrganizationSettingsResponse) GetChatAgentQuickFiltersOk() (*[]AgentQuickFilter, bool)`

GetChatAgentQuickFiltersOk returns a tuple with the ChatAgentQuickFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatAgentQuickFilters

`func (o *OrganizationSettingsResponse) SetChatAgentQuickFilters(v []AgentQuickFilter)`

SetChatAgentQuickFilters sets ChatAgentQuickFilters field to given value.

### HasChatAgentQuickFilters

`func (o *OrganizationSettingsResponse) HasChatAgentQuickFilters() bool`

HasChatAgentQuickFilters returns a boolean if a field has been set.

### GetChatAllowedHostingLocations

`func (o *OrganizationSettingsResponse) GetChatAllowedHostingLocations() []string`

GetChatAllowedHostingLocations returns the ChatAllowedHostingLocations field if non-nil, zero value otherwise.

### GetChatAllowedHostingLocationsOk

`func (o *OrganizationSettingsResponse) GetChatAllowedHostingLocationsOk() (*[]string, bool)`

GetChatAllowedHostingLocationsOk returns a tuple with the ChatAllowedHostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatAllowedHostingLocations

`func (o *OrganizationSettingsResponse) SetChatAllowedHostingLocations(v []string)`

SetChatAllowedHostingLocations sets ChatAllowedHostingLocations field to given value.

### HasChatAllowedHostingLocations

`func (o *OrganizationSettingsResponse) HasChatAllowedHostingLocations() bool`

HasChatAllowedHostingLocations returns a boolean if a field has been set.

### GetChatDefaultAgentId

`func (o *OrganizationSettingsResponse) GetChatDefaultAgentId() string`

GetChatDefaultAgentId returns the ChatDefaultAgentId field if non-nil, zero value otherwise.

### GetChatDefaultAgentIdOk

`func (o *OrganizationSettingsResponse) GetChatDefaultAgentIdOk() (*string, bool)`

GetChatDefaultAgentIdOk returns a tuple with the ChatDefaultAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatDefaultAgentId

`func (o *OrganizationSettingsResponse) SetChatDefaultAgentId(v string)`

SetChatDefaultAgentId sets ChatDefaultAgentId field to given value.

### HasChatDefaultAgentId

`func (o *OrganizationSettingsResponse) HasChatDefaultAgentId() bool`

HasChatDefaultAgentId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrganizationSettingsResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationSettingsResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationSettingsResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmbeddingAllowedHostingLocations

`func (o *OrganizationSettingsResponse) GetEmbeddingAllowedHostingLocations() []string`

GetEmbeddingAllowedHostingLocations returns the EmbeddingAllowedHostingLocations field if non-nil, zero value otherwise.

### GetEmbeddingAllowedHostingLocationsOk

`func (o *OrganizationSettingsResponse) GetEmbeddingAllowedHostingLocationsOk() (*[]string, bool)`

GetEmbeddingAllowedHostingLocationsOk returns a tuple with the EmbeddingAllowedHostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingAllowedHostingLocations

`func (o *OrganizationSettingsResponse) SetEmbeddingAllowedHostingLocations(v []string)`

SetEmbeddingAllowedHostingLocations sets EmbeddingAllowedHostingLocations field to given value.

### HasEmbeddingAllowedHostingLocations

`func (o *OrganizationSettingsResponse) HasEmbeddingAllowedHostingLocations() bool`

HasEmbeddingAllowedHostingLocations returns a boolean if a field has been set.

### GetEmbeddingDefaultAgentId

`func (o *OrganizationSettingsResponse) GetEmbeddingDefaultAgentId() string`

GetEmbeddingDefaultAgentId returns the EmbeddingDefaultAgentId field if non-nil, zero value otherwise.

### GetEmbeddingDefaultAgentIdOk

`func (o *OrganizationSettingsResponse) GetEmbeddingDefaultAgentIdOk() (*string, bool)`

GetEmbeddingDefaultAgentIdOk returns a tuple with the EmbeddingDefaultAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingDefaultAgentId

`func (o *OrganizationSettingsResponse) SetEmbeddingDefaultAgentId(v string)`

SetEmbeddingDefaultAgentId sets EmbeddingDefaultAgentId field to given value.

### HasEmbeddingDefaultAgentId

`func (o *OrganizationSettingsResponse) HasEmbeddingDefaultAgentId() bool`

HasEmbeddingDefaultAgentId returns a boolean if a field has been set.

### GetId

`func (o *OrganizationSettingsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationSettingsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationSettingsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetImageAllowedHostingLocations

`func (o *OrganizationSettingsResponse) GetImageAllowedHostingLocations() []string`

GetImageAllowedHostingLocations returns the ImageAllowedHostingLocations field if non-nil, zero value otherwise.

### GetImageAllowedHostingLocationsOk

`func (o *OrganizationSettingsResponse) GetImageAllowedHostingLocationsOk() (*[]string, bool)`

GetImageAllowedHostingLocationsOk returns a tuple with the ImageAllowedHostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAllowedHostingLocations

`func (o *OrganizationSettingsResponse) SetImageAllowedHostingLocations(v []string)`

SetImageAllowedHostingLocations sets ImageAllowedHostingLocations field to given value.

### HasImageAllowedHostingLocations

`func (o *OrganizationSettingsResponse) HasImageAllowedHostingLocations() bool`

HasImageAllowedHostingLocations returns a boolean if a field has been set.

### GetImageDefaultAgentId

`func (o *OrganizationSettingsResponse) GetImageDefaultAgentId() string`

GetImageDefaultAgentId returns the ImageDefaultAgentId field if non-nil, zero value otherwise.

### GetImageDefaultAgentIdOk

`func (o *OrganizationSettingsResponse) GetImageDefaultAgentIdOk() (*string, bool)`

GetImageDefaultAgentIdOk returns a tuple with the ImageDefaultAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefaultAgentId

`func (o *OrganizationSettingsResponse) SetImageDefaultAgentId(v string)`

SetImageDefaultAgentId sets ImageDefaultAgentId field to given value.

### HasImageDefaultAgentId

`func (o *OrganizationSettingsResponse) HasImageDefaultAgentId() bool`

HasImageDefaultAgentId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationSettingsResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationSettingsResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationSettingsResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetToolsAllowedHostingLocations

`func (o *OrganizationSettingsResponse) GetToolsAllowedHostingLocations() []string`

GetToolsAllowedHostingLocations returns the ToolsAllowedHostingLocations field if non-nil, zero value otherwise.

### GetToolsAllowedHostingLocationsOk

`func (o *OrganizationSettingsResponse) GetToolsAllowedHostingLocationsOk() (*[]string, bool)`

GetToolsAllowedHostingLocationsOk returns a tuple with the ToolsAllowedHostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsAllowedHostingLocations

`func (o *OrganizationSettingsResponse) SetToolsAllowedHostingLocations(v []string)`

SetToolsAllowedHostingLocations sets ToolsAllowedHostingLocations field to given value.

### HasToolsAllowedHostingLocations

`func (o *OrganizationSettingsResponse) HasToolsAllowedHostingLocations() bool`

HasToolsAllowedHostingLocations returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrganizationSettingsResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationSettingsResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationSettingsResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


