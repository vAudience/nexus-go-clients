# OrganizationSettingsPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioAllowedHostingLocations** | Pointer to **[]string** |  | [optional] 
**ChatAllowedHostingLocations** | Pointer to **[]string** |  | [optional] 
**ChatDefaultAgentId** | Pointer to **string** |  | [optional] 
**EmbeddingAllowedHostingLocations** | Pointer to **[]string** |  | [optional] 
**EmbeddingDefaultAgentId** | Pointer to **string** |  | [optional] 
**ImageAllowedHostingLocations** | Pointer to **[]string** |  | [optional] 
**ImageDefaultAgentId** | Pointer to **string** |  | [optional] 
**ToolsAllowedHostingLocations** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOrganizationSettingsPatchRequest

`func NewOrganizationSettingsPatchRequest() *OrganizationSettingsPatchRequest`

NewOrganizationSettingsPatchRequest instantiates a new OrganizationSettingsPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSettingsPatchRequestWithDefaults

`func NewOrganizationSettingsPatchRequestWithDefaults() *OrganizationSettingsPatchRequest`

NewOrganizationSettingsPatchRequestWithDefaults instantiates a new OrganizationSettingsPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) GetAudioAllowedHostingLocations() []string`

GetAudioAllowedHostingLocations returns the AudioAllowedHostingLocations field if non-nil, zero value otherwise.

### GetAudioAllowedHostingLocationsOk

`func (o *OrganizationSettingsPatchRequest) GetAudioAllowedHostingLocationsOk() (*[]string, bool)`

GetAudioAllowedHostingLocationsOk returns a tuple with the AudioAllowedHostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) SetAudioAllowedHostingLocations(v []string)`

SetAudioAllowedHostingLocations sets AudioAllowedHostingLocations field to given value.

### HasAudioAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) HasAudioAllowedHostingLocations() bool`

HasAudioAllowedHostingLocations returns a boolean if a field has been set.

### GetChatAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) GetChatAllowedHostingLocations() []string`

GetChatAllowedHostingLocations returns the ChatAllowedHostingLocations field if non-nil, zero value otherwise.

### GetChatAllowedHostingLocationsOk

`func (o *OrganizationSettingsPatchRequest) GetChatAllowedHostingLocationsOk() (*[]string, bool)`

GetChatAllowedHostingLocationsOk returns a tuple with the ChatAllowedHostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) SetChatAllowedHostingLocations(v []string)`

SetChatAllowedHostingLocations sets ChatAllowedHostingLocations field to given value.

### HasChatAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) HasChatAllowedHostingLocations() bool`

HasChatAllowedHostingLocations returns a boolean if a field has been set.

### GetChatDefaultAgentId

`func (o *OrganizationSettingsPatchRequest) GetChatDefaultAgentId() string`

GetChatDefaultAgentId returns the ChatDefaultAgentId field if non-nil, zero value otherwise.

### GetChatDefaultAgentIdOk

`func (o *OrganizationSettingsPatchRequest) GetChatDefaultAgentIdOk() (*string, bool)`

GetChatDefaultAgentIdOk returns a tuple with the ChatDefaultAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatDefaultAgentId

`func (o *OrganizationSettingsPatchRequest) SetChatDefaultAgentId(v string)`

SetChatDefaultAgentId sets ChatDefaultAgentId field to given value.

### HasChatDefaultAgentId

`func (o *OrganizationSettingsPatchRequest) HasChatDefaultAgentId() bool`

HasChatDefaultAgentId returns a boolean if a field has been set.

### GetEmbeddingAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) GetEmbeddingAllowedHostingLocations() []string`

GetEmbeddingAllowedHostingLocations returns the EmbeddingAllowedHostingLocations field if non-nil, zero value otherwise.

### GetEmbeddingAllowedHostingLocationsOk

`func (o *OrganizationSettingsPatchRequest) GetEmbeddingAllowedHostingLocationsOk() (*[]string, bool)`

GetEmbeddingAllowedHostingLocationsOk returns a tuple with the EmbeddingAllowedHostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) SetEmbeddingAllowedHostingLocations(v []string)`

SetEmbeddingAllowedHostingLocations sets EmbeddingAllowedHostingLocations field to given value.

### HasEmbeddingAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) HasEmbeddingAllowedHostingLocations() bool`

HasEmbeddingAllowedHostingLocations returns a boolean if a field has been set.

### GetEmbeddingDefaultAgentId

`func (o *OrganizationSettingsPatchRequest) GetEmbeddingDefaultAgentId() string`

GetEmbeddingDefaultAgentId returns the EmbeddingDefaultAgentId field if non-nil, zero value otherwise.

### GetEmbeddingDefaultAgentIdOk

`func (o *OrganizationSettingsPatchRequest) GetEmbeddingDefaultAgentIdOk() (*string, bool)`

GetEmbeddingDefaultAgentIdOk returns a tuple with the EmbeddingDefaultAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingDefaultAgentId

`func (o *OrganizationSettingsPatchRequest) SetEmbeddingDefaultAgentId(v string)`

SetEmbeddingDefaultAgentId sets EmbeddingDefaultAgentId field to given value.

### HasEmbeddingDefaultAgentId

`func (o *OrganizationSettingsPatchRequest) HasEmbeddingDefaultAgentId() bool`

HasEmbeddingDefaultAgentId returns a boolean if a field has been set.

### GetImageAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) GetImageAllowedHostingLocations() []string`

GetImageAllowedHostingLocations returns the ImageAllowedHostingLocations field if non-nil, zero value otherwise.

### GetImageAllowedHostingLocationsOk

`func (o *OrganizationSettingsPatchRequest) GetImageAllowedHostingLocationsOk() (*[]string, bool)`

GetImageAllowedHostingLocationsOk returns a tuple with the ImageAllowedHostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) SetImageAllowedHostingLocations(v []string)`

SetImageAllowedHostingLocations sets ImageAllowedHostingLocations field to given value.

### HasImageAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) HasImageAllowedHostingLocations() bool`

HasImageAllowedHostingLocations returns a boolean if a field has been set.

### GetImageDefaultAgentId

`func (o *OrganizationSettingsPatchRequest) GetImageDefaultAgentId() string`

GetImageDefaultAgentId returns the ImageDefaultAgentId field if non-nil, zero value otherwise.

### GetImageDefaultAgentIdOk

`func (o *OrganizationSettingsPatchRequest) GetImageDefaultAgentIdOk() (*string, bool)`

GetImageDefaultAgentIdOk returns a tuple with the ImageDefaultAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefaultAgentId

`func (o *OrganizationSettingsPatchRequest) SetImageDefaultAgentId(v string)`

SetImageDefaultAgentId sets ImageDefaultAgentId field to given value.

### HasImageDefaultAgentId

`func (o *OrganizationSettingsPatchRequest) HasImageDefaultAgentId() bool`

HasImageDefaultAgentId returns a boolean if a field has been set.

### GetToolsAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) GetToolsAllowedHostingLocations() []string`

GetToolsAllowedHostingLocations returns the ToolsAllowedHostingLocations field if non-nil, zero value otherwise.

### GetToolsAllowedHostingLocationsOk

`func (o *OrganizationSettingsPatchRequest) GetToolsAllowedHostingLocationsOk() (*[]string, bool)`

GetToolsAllowedHostingLocationsOk returns a tuple with the ToolsAllowedHostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) SetToolsAllowedHostingLocations(v []string)`

SetToolsAllowedHostingLocations sets ToolsAllowedHostingLocations field to given value.

### HasToolsAllowedHostingLocations

`func (o *OrganizationSettingsPatchRequest) HasToolsAllowedHostingLocations() bool`

HasToolsAllowedHostingLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


