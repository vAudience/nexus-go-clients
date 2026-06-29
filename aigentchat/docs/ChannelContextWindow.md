# ChannelContextWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnchorMessageId** | Pointer to **string** | assistant msg that produced a verified count (empty if estimated) | [optional] 
**IsVerified** | Pointer to **bool** | true &#x3D; provider-reported; false &#x3D; local estimate | [optional] 
**MaxInputTokens** | Pointer to **int32** | model&#39;s input limit (registry lookup at write time) | [optional] 
**ModelId** | Pointer to **string** | AIModel.ID (mirrors AIgencyMessage.AIModelID) | [optional] 
**ModelInternalId** | Pointer to **string** | AIModel.InternalId (provider-agnostic) | [optional] 
**OverheadTokens** | Pointer to **int32** | overhead baked into a provider Tokens (the gate re-bases it); 0 for estimates | [optional] 
**Source** | Pointer to **string** | provider (Tokens include overhead) or estimate (exclude) | [optional] 
**Tokens** | Pointer to **int32** | total input-token cost as of the last anchor | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewChannelContextWindow

`func NewChannelContextWindow() *ChannelContextWindow`

NewChannelContextWindow instantiates a new ChannelContextWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelContextWindowWithDefaults

`func NewChannelContextWindowWithDefaults() *ChannelContextWindow`

NewChannelContextWindowWithDefaults instantiates a new ChannelContextWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnchorMessageId

`func (o *ChannelContextWindow) GetAnchorMessageId() string`

GetAnchorMessageId returns the AnchorMessageId field if non-nil, zero value otherwise.

### GetAnchorMessageIdOk

`func (o *ChannelContextWindow) GetAnchorMessageIdOk() (*string, bool)`

GetAnchorMessageIdOk returns a tuple with the AnchorMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorMessageId

`func (o *ChannelContextWindow) SetAnchorMessageId(v string)`

SetAnchorMessageId sets AnchorMessageId field to given value.

### HasAnchorMessageId

`func (o *ChannelContextWindow) HasAnchorMessageId() bool`

HasAnchorMessageId returns a boolean if a field has been set.

### GetIsVerified

`func (o *ChannelContextWindow) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *ChannelContextWindow) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *ChannelContextWindow) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *ChannelContextWindow) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetMaxInputTokens

`func (o *ChannelContextWindow) GetMaxInputTokens() int32`

GetMaxInputTokens returns the MaxInputTokens field if non-nil, zero value otherwise.

### GetMaxInputTokensOk

`func (o *ChannelContextWindow) GetMaxInputTokensOk() (*int32, bool)`

GetMaxInputTokensOk returns a tuple with the MaxInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInputTokens

`func (o *ChannelContextWindow) SetMaxInputTokens(v int32)`

SetMaxInputTokens sets MaxInputTokens field to given value.

### HasMaxInputTokens

`func (o *ChannelContextWindow) HasMaxInputTokens() bool`

HasMaxInputTokens returns a boolean if a field has been set.

### GetModelId

`func (o *ChannelContextWindow) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ChannelContextWindow) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ChannelContextWindow) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *ChannelContextWindow) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetModelInternalId

`func (o *ChannelContextWindow) GetModelInternalId() string`

GetModelInternalId returns the ModelInternalId field if non-nil, zero value otherwise.

### GetModelInternalIdOk

`func (o *ChannelContextWindow) GetModelInternalIdOk() (*string, bool)`

GetModelInternalIdOk returns a tuple with the ModelInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelInternalId

`func (o *ChannelContextWindow) SetModelInternalId(v string)`

SetModelInternalId sets ModelInternalId field to given value.

### HasModelInternalId

`func (o *ChannelContextWindow) HasModelInternalId() bool`

HasModelInternalId returns a boolean if a field has been set.

### GetOverheadTokens

`func (o *ChannelContextWindow) GetOverheadTokens() int32`

GetOverheadTokens returns the OverheadTokens field if non-nil, zero value otherwise.

### GetOverheadTokensOk

`func (o *ChannelContextWindow) GetOverheadTokensOk() (*int32, bool)`

GetOverheadTokensOk returns a tuple with the OverheadTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverheadTokens

`func (o *ChannelContextWindow) SetOverheadTokens(v int32)`

SetOverheadTokens sets OverheadTokens field to given value.

### HasOverheadTokens

`func (o *ChannelContextWindow) HasOverheadTokens() bool`

HasOverheadTokens returns a boolean if a field has been set.

### GetSource

`func (o *ChannelContextWindow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ChannelContextWindow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ChannelContextWindow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ChannelContextWindow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTokens

`func (o *ChannelContextWindow) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ChannelContextWindow) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ChannelContextWindow) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *ChannelContextWindow) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ChannelContextWindow) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChannelContextWindow) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChannelContextWindow) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ChannelContextWindow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


