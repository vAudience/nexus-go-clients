# AuditTrail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** |  | 
**Messages** | Pointer to [**[]AuditTrailMessage**](AuditTrailMessage.md) |  | [optional] 

## Methods

### NewAuditTrail

`func NewAuditTrail(channelId string, ) *AuditTrail`

NewAuditTrail instantiates a new AuditTrail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditTrailWithDefaults

`func NewAuditTrailWithDefaults() *AuditTrail`

NewAuditTrailWithDefaults instantiates a new AuditTrail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *AuditTrail) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *AuditTrail) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *AuditTrail) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetMessages

`func (o *AuditTrail) GetMessages() []AuditTrailMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AuditTrail) GetMessagesOk() (*[]AuditTrailMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AuditTrail) SetMessages(v []AuditTrailMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *AuditTrail) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


