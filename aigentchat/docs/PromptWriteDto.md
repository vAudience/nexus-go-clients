# PromptWriteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**DefaultAgentId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ThumbnailUrl** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to [**PromptVisibilityStates**](PromptVisibilityStates.md) |  | [optional] 

## Methods

### NewPromptWriteDto

`func NewPromptWriteDto() *PromptWriteDto`

NewPromptWriteDto instantiates a new PromptWriteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptWriteDtoWithDefaults

`func NewPromptWriteDtoWithDefaults() *PromptWriteDto`

NewPromptWriteDtoWithDefaults instantiates a new PromptWriteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PromptWriteDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PromptWriteDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PromptWriteDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *PromptWriteDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDefaultAgentId

`func (o *PromptWriteDto) GetDefaultAgentId() string`

GetDefaultAgentId returns the DefaultAgentId field if non-nil, zero value otherwise.

### GetDefaultAgentIdOk

`func (o *PromptWriteDto) GetDefaultAgentIdOk() (*string, bool)`

GetDefaultAgentIdOk returns a tuple with the DefaultAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAgentId

`func (o *PromptWriteDto) SetDefaultAgentId(v string)`

SetDefaultAgentId sets DefaultAgentId field to given value.

### HasDefaultAgentId

`func (o *PromptWriteDto) HasDefaultAgentId() bool`

HasDefaultAgentId returns a boolean if a field has been set.

### GetDescription

`func (o *PromptWriteDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PromptWriteDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PromptWriteDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PromptWriteDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PromptWriteDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromptWriteDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromptWriteDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromptWriteDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *PromptWriteDto) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *PromptWriteDto) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *PromptWriteDto) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *PromptWriteDto) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetTitle

`func (o *PromptWriteDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PromptWriteDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PromptWriteDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PromptWriteDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVisibility

`func (o *PromptWriteDto) GetVisibility() PromptVisibilityStates`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PromptWriteDto) GetVisibilityOk() (*PromptVisibilityStates, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PromptWriteDto) SetVisibility(v PromptVisibilityStates)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *PromptWriteDto) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


