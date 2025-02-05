# AIgencyMessageWriteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiModelId** | **string** |  | 
**AiServiceId** | **string** |  | 
**AigentThreadId** | **string** |  | 
**Attachments** | Pointer to [**AIgencyMessageFileList**](AIgencyMessageFileList.md) |  | [optional] 
**ChannelId** | **string** |  | 
**ChannelName** | **string** |  | 
**Content** | Pointer to [**AIgencyMessageContentList**](AIgencyMessageContentList.md) |  | [optional] 
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**MissionId** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**ResponseToId** | Pointer to **string** |  | [optional] 
**SenderConversationRole** | [**ConversationRole**](ConversationRole.md) |  | 
**SenderName** | **string** |  | 
**Type** | [**AIgencyMessageType**](AIgencyMessageType.md) |  | 

## Methods

### NewAIgencyMessageWriteDto

`func NewAIgencyMessageWriteDto(aiModelId string, aiServiceId string, aigentThreadId string, channelId string, channelName string, missionId string, senderConversationRole ConversationRole, senderName string, type_ AIgencyMessageType, ) *AIgencyMessageWriteDto`

NewAIgencyMessageWriteDto instantiates a new AIgencyMessageWriteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIgencyMessageWriteDtoWithDefaults

`func NewAIgencyMessageWriteDtoWithDefaults() *AIgencyMessageWriteDto`

NewAIgencyMessageWriteDtoWithDefaults instantiates a new AIgencyMessageWriteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiModelId

`func (o *AIgencyMessageWriteDto) GetAiModelId() string`

GetAiModelId returns the AiModelId field if non-nil, zero value otherwise.

### GetAiModelIdOk

`func (o *AIgencyMessageWriteDto) GetAiModelIdOk() (*string, bool)`

GetAiModelIdOk returns a tuple with the AiModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiModelId

`func (o *AIgencyMessageWriteDto) SetAiModelId(v string)`

SetAiModelId sets AiModelId field to given value.


### GetAiServiceId

`func (o *AIgencyMessageWriteDto) GetAiServiceId() string`

GetAiServiceId returns the AiServiceId field if non-nil, zero value otherwise.

### GetAiServiceIdOk

`func (o *AIgencyMessageWriteDto) GetAiServiceIdOk() (*string, bool)`

GetAiServiceIdOk returns a tuple with the AiServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiServiceId

`func (o *AIgencyMessageWriteDto) SetAiServiceId(v string)`

SetAiServiceId sets AiServiceId field to given value.


### GetAigentThreadId

`func (o *AIgencyMessageWriteDto) GetAigentThreadId() string`

GetAigentThreadId returns the AigentThreadId field if non-nil, zero value otherwise.

### GetAigentThreadIdOk

`func (o *AIgencyMessageWriteDto) GetAigentThreadIdOk() (*string, bool)`

GetAigentThreadIdOk returns a tuple with the AigentThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAigentThreadId

`func (o *AIgencyMessageWriteDto) SetAigentThreadId(v string)`

SetAigentThreadId sets AigentThreadId field to given value.


### GetAttachments

`func (o *AIgencyMessageWriteDto) GetAttachments() AIgencyMessageFileList`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *AIgencyMessageWriteDto) GetAttachmentsOk() (*AIgencyMessageFileList, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *AIgencyMessageWriteDto) SetAttachments(v AIgencyMessageFileList)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *AIgencyMessageWriteDto) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetChannelId

`func (o *AIgencyMessageWriteDto) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *AIgencyMessageWriteDto) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *AIgencyMessageWriteDto) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetChannelName

`func (o *AIgencyMessageWriteDto) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *AIgencyMessageWriteDto) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *AIgencyMessageWriteDto) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.


### GetContent

`func (o *AIgencyMessageWriteDto) GetContent() AIgencyMessageContentList`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AIgencyMessageWriteDto) GetContentOk() (*AIgencyMessageContentList, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AIgencyMessageWriteDto) SetContent(v AIgencyMessageContentList)`

SetContent sets Content field to given value.

### HasContent

`func (o *AIgencyMessageWriteDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMetaData

`func (o *AIgencyMessageWriteDto) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AIgencyMessageWriteDto) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AIgencyMessageWriteDto) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AIgencyMessageWriteDto) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetMissionId

`func (o *AIgencyMessageWriteDto) GetMissionId() string`

GetMissionId returns the MissionId field if non-nil, zero value otherwise.

### GetMissionIdOk

`func (o *AIgencyMessageWriteDto) GetMissionIdOk() (*string, bool)`

GetMissionIdOk returns a tuple with the MissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionId

`func (o *AIgencyMessageWriteDto) SetMissionId(v string)`

SetMissionId sets MissionId field to given value.


### GetParameters

`func (o *AIgencyMessageWriteDto) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AIgencyMessageWriteDto) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AIgencyMessageWriteDto) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AIgencyMessageWriteDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetReferenceId

`func (o *AIgencyMessageWriteDto) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AIgencyMessageWriteDto) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AIgencyMessageWriteDto) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *AIgencyMessageWriteDto) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetResponseToId

`func (o *AIgencyMessageWriteDto) GetResponseToId() string`

GetResponseToId returns the ResponseToId field if non-nil, zero value otherwise.

### GetResponseToIdOk

`func (o *AIgencyMessageWriteDto) GetResponseToIdOk() (*string, bool)`

GetResponseToIdOk returns a tuple with the ResponseToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseToId

`func (o *AIgencyMessageWriteDto) SetResponseToId(v string)`

SetResponseToId sets ResponseToId field to given value.

### HasResponseToId

`func (o *AIgencyMessageWriteDto) HasResponseToId() bool`

HasResponseToId returns a boolean if a field has been set.

### GetSenderConversationRole

`func (o *AIgencyMessageWriteDto) GetSenderConversationRole() ConversationRole`

GetSenderConversationRole returns the SenderConversationRole field if non-nil, zero value otherwise.

### GetSenderConversationRoleOk

`func (o *AIgencyMessageWriteDto) GetSenderConversationRoleOk() (*ConversationRole, bool)`

GetSenderConversationRoleOk returns a tuple with the SenderConversationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderConversationRole

`func (o *AIgencyMessageWriteDto) SetSenderConversationRole(v ConversationRole)`

SetSenderConversationRole sets SenderConversationRole field to given value.


### GetSenderName

`func (o *AIgencyMessageWriteDto) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *AIgencyMessageWriteDto) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *AIgencyMessageWriteDto) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.


### GetType

`func (o *AIgencyMessageWriteDto) GetType() AIgencyMessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AIgencyMessageWriteDto) GetTypeOk() (*AIgencyMessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AIgencyMessageWriteDto) SetType(v AIgencyMessageType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


