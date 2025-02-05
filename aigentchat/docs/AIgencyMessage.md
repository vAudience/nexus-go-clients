# AIgencyMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiModelId** | **string** |  | 
**AiServiceId** | **string** |  | 
**AigentThreadId** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to [**AIgencyMessageFileList**](AIgencyMessageFileList.md) |  | [optional] 
**ChannelId** | **string** |  | 
**ChannelName** | **string** |  | 
**Content** | Pointer to [**AIgencyMessageContentList**](AIgencyMessageContentList.md) |  | [optional] 
**CreatedAt** | **int64** |  | 
**CreatedForFeature** | Pointer to [**AIModelFeature**](AIModelFeature.md) |  | [optional] 
**Error** | Pointer to [**AiServiceError**](AiServiceError.md) |  | [optional] 
**FinishReason** | Pointer to [**FinishReason**](FinishReason.md) |  | [optional] 
**Id** | **string** |  | 
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**MissionId** | Pointer to **string** |  | [optional] 
**OwnerOrganizationId** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**ResponseToId** | Pointer to **string** |  | [optional] 
**SenderConversationRole** | [**ConversationRole**](ConversationRole.md) |  | 
**SenderId** | **string** |  | 
**SenderName** | **string** |  | 
**TokenCount** | Pointer to **int32** |  | [optional] 
**TokenDirection** | [**TokenDirection**](TokenDirection.md) |  | 
**Type** | [**AIgencyMessageType**](AIgencyMessageType.md) |  | 
**UpdatedAt** | **int64** |  | 

## Methods

### NewAIgencyMessage

`func NewAIgencyMessage(aiModelId string, aiServiceId string, channelId string, channelName string, createdAt int64, id string, ownerOrganizationId string, senderConversationRole ConversationRole, senderId string, senderName string, tokenDirection TokenDirection, type_ AIgencyMessageType, updatedAt int64, ) *AIgencyMessage`

NewAIgencyMessage instantiates a new AIgencyMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIgencyMessageWithDefaults

`func NewAIgencyMessageWithDefaults() *AIgencyMessage`

NewAIgencyMessageWithDefaults instantiates a new AIgencyMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiModelId

`func (o *AIgencyMessage) GetAiModelId() string`

GetAiModelId returns the AiModelId field if non-nil, zero value otherwise.

### GetAiModelIdOk

`func (o *AIgencyMessage) GetAiModelIdOk() (*string, bool)`

GetAiModelIdOk returns a tuple with the AiModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiModelId

`func (o *AIgencyMessage) SetAiModelId(v string)`

SetAiModelId sets AiModelId field to given value.


### GetAiServiceId

`func (o *AIgencyMessage) GetAiServiceId() string`

GetAiServiceId returns the AiServiceId field if non-nil, zero value otherwise.

### GetAiServiceIdOk

`func (o *AIgencyMessage) GetAiServiceIdOk() (*string, bool)`

GetAiServiceIdOk returns a tuple with the AiServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiServiceId

`func (o *AIgencyMessage) SetAiServiceId(v string)`

SetAiServiceId sets AiServiceId field to given value.


### GetAigentThreadId

`func (o *AIgencyMessage) GetAigentThreadId() string`

GetAigentThreadId returns the AigentThreadId field if non-nil, zero value otherwise.

### GetAigentThreadIdOk

`func (o *AIgencyMessage) GetAigentThreadIdOk() (*string, bool)`

GetAigentThreadIdOk returns a tuple with the AigentThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAigentThreadId

`func (o *AIgencyMessage) SetAigentThreadId(v string)`

SetAigentThreadId sets AigentThreadId field to given value.

### HasAigentThreadId

`func (o *AIgencyMessage) HasAigentThreadId() bool`

HasAigentThreadId returns a boolean if a field has been set.

### GetAttachments

`func (o *AIgencyMessage) GetAttachments() AIgencyMessageFileList`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *AIgencyMessage) GetAttachmentsOk() (*AIgencyMessageFileList, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *AIgencyMessage) SetAttachments(v AIgencyMessageFileList)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *AIgencyMessage) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetChannelId

`func (o *AIgencyMessage) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *AIgencyMessage) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *AIgencyMessage) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetChannelName

`func (o *AIgencyMessage) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *AIgencyMessage) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *AIgencyMessage) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.


### GetContent

`func (o *AIgencyMessage) GetContent() AIgencyMessageContentList`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AIgencyMessage) GetContentOk() (*AIgencyMessageContentList, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AIgencyMessage) SetContent(v AIgencyMessageContentList)`

SetContent sets Content field to given value.

### HasContent

`func (o *AIgencyMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AIgencyMessage) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AIgencyMessage) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AIgencyMessage) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedForFeature

`func (o *AIgencyMessage) GetCreatedForFeature() AIModelFeature`

GetCreatedForFeature returns the CreatedForFeature field if non-nil, zero value otherwise.

### GetCreatedForFeatureOk

`func (o *AIgencyMessage) GetCreatedForFeatureOk() (*AIModelFeature, bool)`

GetCreatedForFeatureOk returns a tuple with the CreatedForFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedForFeature

`func (o *AIgencyMessage) SetCreatedForFeature(v AIModelFeature)`

SetCreatedForFeature sets CreatedForFeature field to given value.

### HasCreatedForFeature

`func (o *AIgencyMessage) HasCreatedForFeature() bool`

HasCreatedForFeature returns a boolean if a field has been set.

### GetError

`func (o *AIgencyMessage) GetError() AiServiceError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AIgencyMessage) GetErrorOk() (*AiServiceError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AIgencyMessage) SetError(v AiServiceError)`

SetError sets Error field to given value.

### HasError

`func (o *AIgencyMessage) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFinishReason

`func (o *AIgencyMessage) GetFinishReason() FinishReason`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *AIgencyMessage) GetFinishReasonOk() (*FinishReason, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *AIgencyMessage) SetFinishReason(v FinishReason)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *AIgencyMessage) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.

### GetId

`func (o *AIgencyMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AIgencyMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AIgencyMessage) SetId(v string)`

SetId sets Id field to given value.


### GetMetaData

`func (o *AIgencyMessage) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AIgencyMessage) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AIgencyMessage) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AIgencyMessage) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetMissionId

`func (o *AIgencyMessage) GetMissionId() string`

GetMissionId returns the MissionId field if non-nil, zero value otherwise.

### GetMissionIdOk

`func (o *AIgencyMessage) GetMissionIdOk() (*string, bool)`

GetMissionIdOk returns a tuple with the MissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionId

`func (o *AIgencyMessage) SetMissionId(v string)`

SetMissionId sets MissionId field to given value.

### HasMissionId

`func (o *AIgencyMessage) HasMissionId() bool`

HasMissionId returns a boolean if a field has been set.

### GetOwnerOrganizationId

`func (o *AIgencyMessage) GetOwnerOrganizationId() string`

GetOwnerOrganizationId returns the OwnerOrganizationId field if non-nil, zero value otherwise.

### GetOwnerOrganizationIdOk

`func (o *AIgencyMessage) GetOwnerOrganizationIdOk() (*string, bool)`

GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrganizationId

`func (o *AIgencyMessage) SetOwnerOrganizationId(v string)`

SetOwnerOrganizationId sets OwnerOrganizationId field to given value.


### GetParameters

`func (o *AIgencyMessage) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AIgencyMessage) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AIgencyMessage) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AIgencyMessage) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetReferenceId

`func (o *AIgencyMessage) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AIgencyMessage) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AIgencyMessage) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *AIgencyMessage) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetResponseToId

`func (o *AIgencyMessage) GetResponseToId() string`

GetResponseToId returns the ResponseToId field if non-nil, zero value otherwise.

### GetResponseToIdOk

`func (o *AIgencyMessage) GetResponseToIdOk() (*string, bool)`

GetResponseToIdOk returns a tuple with the ResponseToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseToId

`func (o *AIgencyMessage) SetResponseToId(v string)`

SetResponseToId sets ResponseToId field to given value.

### HasResponseToId

`func (o *AIgencyMessage) HasResponseToId() bool`

HasResponseToId returns a boolean if a field has been set.

### GetSenderConversationRole

`func (o *AIgencyMessage) GetSenderConversationRole() ConversationRole`

GetSenderConversationRole returns the SenderConversationRole field if non-nil, zero value otherwise.

### GetSenderConversationRoleOk

`func (o *AIgencyMessage) GetSenderConversationRoleOk() (*ConversationRole, bool)`

GetSenderConversationRoleOk returns a tuple with the SenderConversationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderConversationRole

`func (o *AIgencyMessage) SetSenderConversationRole(v ConversationRole)`

SetSenderConversationRole sets SenderConversationRole field to given value.


### GetSenderId

`func (o *AIgencyMessage) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *AIgencyMessage) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *AIgencyMessage) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.


### GetSenderName

`func (o *AIgencyMessage) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *AIgencyMessage) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *AIgencyMessage) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.


### GetTokenCount

`func (o *AIgencyMessage) GetTokenCount() int32`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *AIgencyMessage) GetTokenCountOk() (*int32, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *AIgencyMessage) SetTokenCount(v int32)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *AIgencyMessage) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetTokenDirection

`func (o *AIgencyMessage) GetTokenDirection() TokenDirection`

GetTokenDirection returns the TokenDirection field if non-nil, zero value otherwise.

### GetTokenDirectionOk

`func (o *AIgencyMessage) GetTokenDirectionOk() (*TokenDirection, bool)`

GetTokenDirectionOk returns a tuple with the TokenDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDirection

`func (o *AIgencyMessage) SetTokenDirection(v TokenDirection)`

SetTokenDirection sets TokenDirection field to given value.


### GetType

`func (o *AIgencyMessage) GetType() AIgencyMessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AIgencyMessage) GetTypeOk() (*AIgencyMessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AIgencyMessage) SetType(v AIgencyMessageType)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *AIgencyMessage) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AIgencyMessage) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AIgencyMessage) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


