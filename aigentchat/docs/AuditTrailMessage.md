# AuditTrailMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiModelId** | **string** |  | 
**AiServiceId** | **string** |  | 
**ChannelId** | **string** |  | 
**ChannelName** | **string** |  | 
**Content** | [**AIgencyMessageContentList**](AIgencyMessageContentList.md) |  | 
**CreatedAt** | **int64** |  | 
**Error** | Pointer to [**AiServiceError**](AiServiceError.md) |  | [optional] 
**FinishReason** | Pointer to [**FinishReason**](FinishReason.md) |  | [optional] 
**Id** | **string** |  | 
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**OwnerOrganizationId** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**ResponseToId** | Pointer to **string** |  | [optional] 
**SenderConversationRole** | [**ConversationRole**](ConversationRole.md) |  | 
**SenderId** | **string** |  | 
**SenderName** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**TokenCount** | Pointer to **int32** |  | [optional] 
**TokenDirection** | [**TokenDirection**](TokenDirection.md) |  | 
**Type** | [**AIgencyMessageType**](AIgencyMessageType.md) |  | 
**UpdatedAt** | **int64** |  | 

## Methods

### NewAuditTrailMessage

`func NewAuditTrailMessage(aiModelId string, aiServiceId string, channelId string, channelName string, content AIgencyMessageContentList, createdAt int64, id string, ownerOrganizationId string, senderConversationRole ConversationRole, senderId string, senderName string, tokenDirection TokenDirection, type_ AIgencyMessageType, updatedAt int64, ) *AuditTrailMessage`

NewAuditTrailMessage instantiates a new AuditTrailMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditTrailMessageWithDefaults

`func NewAuditTrailMessageWithDefaults() *AuditTrailMessage`

NewAuditTrailMessageWithDefaults instantiates a new AuditTrailMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiModelId

`func (o *AuditTrailMessage) GetAiModelId() string`

GetAiModelId returns the AiModelId field if non-nil, zero value otherwise.

### GetAiModelIdOk

`func (o *AuditTrailMessage) GetAiModelIdOk() (*string, bool)`

GetAiModelIdOk returns a tuple with the AiModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiModelId

`func (o *AuditTrailMessage) SetAiModelId(v string)`

SetAiModelId sets AiModelId field to given value.


### GetAiServiceId

`func (o *AuditTrailMessage) GetAiServiceId() string`

GetAiServiceId returns the AiServiceId field if non-nil, zero value otherwise.

### GetAiServiceIdOk

`func (o *AuditTrailMessage) GetAiServiceIdOk() (*string, bool)`

GetAiServiceIdOk returns a tuple with the AiServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiServiceId

`func (o *AuditTrailMessage) SetAiServiceId(v string)`

SetAiServiceId sets AiServiceId field to given value.


### GetChannelId

`func (o *AuditTrailMessage) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *AuditTrailMessage) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *AuditTrailMessage) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetChannelName

`func (o *AuditTrailMessage) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *AuditTrailMessage) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *AuditTrailMessage) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.


### GetContent

`func (o *AuditTrailMessage) GetContent() AIgencyMessageContentList`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AuditTrailMessage) GetContentOk() (*AIgencyMessageContentList, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AuditTrailMessage) SetContent(v AIgencyMessageContentList)`

SetContent sets Content field to given value.


### GetCreatedAt

`func (o *AuditTrailMessage) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditTrailMessage) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditTrailMessage) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetError

`func (o *AuditTrailMessage) GetError() AiServiceError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AuditTrailMessage) GetErrorOk() (*AiServiceError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AuditTrailMessage) SetError(v AiServiceError)`

SetError sets Error field to given value.

### HasError

`func (o *AuditTrailMessage) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFinishReason

`func (o *AuditTrailMessage) GetFinishReason() FinishReason`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *AuditTrailMessage) GetFinishReasonOk() (*FinishReason, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *AuditTrailMessage) SetFinishReason(v FinishReason)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *AuditTrailMessage) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.

### GetId

`func (o *AuditTrailMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditTrailMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditTrailMessage) SetId(v string)`

SetId sets Id field to given value.


### GetMetaData

`func (o *AuditTrailMessage) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AuditTrailMessage) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AuditTrailMessage) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AuditTrailMessage) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetOwnerOrganizationId

`func (o *AuditTrailMessage) GetOwnerOrganizationId() string`

GetOwnerOrganizationId returns the OwnerOrganizationId field if non-nil, zero value otherwise.

### GetOwnerOrganizationIdOk

`func (o *AuditTrailMessage) GetOwnerOrganizationIdOk() (*string, bool)`

GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrganizationId

`func (o *AuditTrailMessage) SetOwnerOrganizationId(v string)`

SetOwnerOrganizationId sets OwnerOrganizationId field to given value.


### GetParameters

`func (o *AuditTrailMessage) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AuditTrailMessage) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AuditTrailMessage) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AuditTrailMessage) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetReferenceId

`func (o *AuditTrailMessage) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AuditTrailMessage) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AuditTrailMessage) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *AuditTrailMessage) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetResponseToId

`func (o *AuditTrailMessage) GetResponseToId() string`

GetResponseToId returns the ResponseToId field if non-nil, zero value otherwise.

### GetResponseToIdOk

`func (o *AuditTrailMessage) GetResponseToIdOk() (*string, bool)`

GetResponseToIdOk returns a tuple with the ResponseToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseToId

`func (o *AuditTrailMessage) SetResponseToId(v string)`

SetResponseToId sets ResponseToId field to given value.

### HasResponseToId

`func (o *AuditTrailMessage) HasResponseToId() bool`

HasResponseToId returns a boolean if a field has been set.

### GetSenderConversationRole

`func (o *AuditTrailMessage) GetSenderConversationRole() ConversationRole`

GetSenderConversationRole returns the SenderConversationRole field if non-nil, zero value otherwise.

### GetSenderConversationRoleOk

`func (o *AuditTrailMessage) GetSenderConversationRoleOk() (*ConversationRole, bool)`

GetSenderConversationRoleOk returns a tuple with the SenderConversationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderConversationRole

`func (o *AuditTrailMessage) SetSenderConversationRole(v ConversationRole)`

SetSenderConversationRole sets SenderConversationRole field to given value.


### GetSenderId

`func (o *AuditTrailMessage) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *AuditTrailMessage) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *AuditTrailMessage) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.


### GetSenderName

`func (o *AuditTrailMessage) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *AuditTrailMessage) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *AuditTrailMessage) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.


### GetState

`func (o *AuditTrailMessage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AuditTrailMessage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AuditTrailMessage) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AuditTrailMessage) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTokenCount

`func (o *AuditTrailMessage) GetTokenCount() int32`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *AuditTrailMessage) GetTokenCountOk() (*int32, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *AuditTrailMessage) SetTokenCount(v int32)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *AuditTrailMessage) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetTokenDirection

`func (o *AuditTrailMessage) GetTokenDirection() TokenDirection`

GetTokenDirection returns the TokenDirection field if non-nil, zero value otherwise.

### GetTokenDirectionOk

`func (o *AuditTrailMessage) GetTokenDirectionOk() (*TokenDirection, bool)`

GetTokenDirectionOk returns a tuple with the TokenDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDirection

`func (o *AuditTrailMessage) SetTokenDirection(v TokenDirection)`

SetTokenDirection sets TokenDirection field to given value.


### GetType

`func (o *AuditTrailMessage) GetType() AIgencyMessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditTrailMessage) GetTypeOk() (*AIgencyMessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditTrailMessage) SetType(v AIgencyMessageType)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *AuditTrailMessage) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuditTrailMessage) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuditTrailMessage) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


