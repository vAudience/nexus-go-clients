# AIgencyImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiModelId** | **string** |  | 
**AiServiceId** | **string** |  | 
**CreatedAt** | **int64** |  | 
**ExecutionId** | **string** |  | 
**Id** | **string** |  | 
**Message** | **string** |  | 
**MimeType** | **string** |  | 
**OwnerId** | **string** |  | 
**OwnerOrganizationId** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Thumbnails** | Pointer to [**[]AIgencyThumbnail**](AIgencyThumbnail.md) |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewAIgencyImage

`func NewAIgencyImage(aiModelId string, aiServiceId string, createdAt int64, executionId string, id string, message string, mimeType string, ownerId string, ownerOrganizationId string, url string, ) *AIgencyImage`

NewAIgencyImage instantiates a new AIgencyImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIgencyImageWithDefaults

`func NewAIgencyImageWithDefaults() *AIgencyImage`

NewAIgencyImageWithDefaults instantiates a new AIgencyImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiModelId

`func (o *AIgencyImage) GetAiModelId() string`

GetAiModelId returns the AiModelId field if non-nil, zero value otherwise.

### GetAiModelIdOk

`func (o *AIgencyImage) GetAiModelIdOk() (*string, bool)`

GetAiModelIdOk returns a tuple with the AiModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiModelId

`func (o *AIgencyImage) SetAiModelId(v string)`

SetAiModelId sets AiModelId field to given value.


### GetAiServiceId

`func (o *AIgencyImage) GetAiServiceId() string`

GetAiServiceId returns the AiServiceId field if non-nil, zero value otherwise.

### GetAiServiceIdOk

`func (o *AIgencyImage) GetAiServiceIdOk() (*string, bool)`

GetAiServiceIdOk returns a tuple with the AiServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiServiceId

`func (o *AIgencyImage) SetAiServiceId(v string)`

SetAiServiceId sets AiServiceId field to given value.


### GetCreatedAt

`func (o *AIgencyImage) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AIgencyImage) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AIgencyImage) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetExecutionId

`func (o *AIgencyImage) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *AIgencyImage) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *AIgencyImage) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetId

`func (o *AIgencyImage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AIgencyImage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AIgencyImage) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *AIgencyImage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AIgencyImage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AIgencyImage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMimeType

`func (o *AIgencyImage) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AIgencyImage) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AIgencyImage) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetOwnerId

`func (o *AIgencyImage) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AIgencyImage) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AIgencyImage) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerOrganizationId

`func (o *AIgencyImage) GetOwnerOrganizationId() string`

GetOwnerOrganizationId returns the OwnerOrganizationId field if non-nil, zero value otherwise.

### GetOwnerOrganizationIdOk

`func (o *AIgencyImage) GetOwnerOrganizationIdOk() (*string, bool)`

GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrganizationId

`func (o *AIgencyImage) SetOwnerOrganizationId(v string)`

SetOwnerOrganizationId sets OwnerOrganizationId field to given value.


### GetParameters

`func (o *AIgencyImage) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AIgencyImage) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AIgencyImage) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AIgencyImage) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetThumbnails

`func (o *AIgencyImage) GetThumbnails() []AIgencyThumbnail`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *AIgencyImage) GetThumbnailsOk() (*[]AIgencyThumbnail, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *AIgencyImage) SetThumbnails(v []AIgencyThumbnail)`

SetThumbnails sets Thumbnails field to given value.

### HasThumbnails

`func (o *AIgencyImage) HasThumbnails() bool`

HasThumbnails returns a boolean if a field has been set.

### GetUrl

`func (o *AIgencyImage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AIgencyImage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AIgencyImage) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


