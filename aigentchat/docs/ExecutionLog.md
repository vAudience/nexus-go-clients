# ExecutionLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**AiModelId** | Pointer to **string** |  | [optional] 
**AiModelServiceId** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**FinalCostInEuro** | Pointer to **float64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerOrganizationId** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**[]ExecutionUsageCost**](ExecutionUsageCost.md) |  | [optional] 

## Methods

### NewExecutionLog

`func NewExecutionLog() *ExecutionLog`

NewExecutionLog instantiates a new ExecutionLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionLogWithDefaults

`func NewExecutionLogWithDefaults() *ExecutionLog`

NewExecutionLogWithDefaults instantiates a new ExecutionLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *ExecutionLog) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ExecutionLog) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ExecutionLog) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ExecutionLog) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAiModelId

`func (o *ExecutionLog) GetAiModelId() string`

GetAiModelId returns the AiModelId field if non-nil, zero value otherwise.

### GetAiModelIdOk

`func (o *ExecutionLog) GetAiModelIdOk() (*string, bool)`

GetAiModelIdOk returns a tuple with the AiModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiModelId

`func (o *ExecutionLog) SetAiModelId(v string)`

SetAiModelId sets AiModelId field to given value.

### HasAiModelId

`func (o *ExecutionLog) HasAiModelId() bool`

HasAiModelId returns a boolean if a field has been set.

### GetAiModelServiceId

`func (o *ExecutionLog) GetAiModelServiceId() string`

GetAiModelServiceId returns the AiModelServiceId field if non-nil, zero value otherwise.

### GetAiModelServiceIdOk

`func (o *ExecutionLog) GetAiModelServiceIdOk() (*string, bool)`

GetAiModelServiceIdOk returns a tuple with the AiModelServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiModelServiceId

`func (o *ExecutionLog) SetAiModelServiceId(v string)`

SetAiModelServiceId sets AiModelServiceId field to given value.

### HasAiModelServiceId

`func (o *ExecutionLog) HasAiModelServiceId() bool`

HasAiModelServiceId returns a boolean if a field has been set.

### GetChannelId

`func (o *ExecutionLog) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ExecutionLog) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ExecutionLog) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ExecutionLog) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExecutionLog) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExecutionLog) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExecutionLog) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExecutionLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFinalCostInEuro

`func (o *ExecutionLog) GetFinalCostInEuro() float64`

GetFinalCostInEuro returns the FinalCostInEuro field if non-nil, zero value otherwise.

### GetFinalCostInEuroOk

`func (o *ExecutionLog) GetFinalCostInEuroOk() (*float64, bool)`

GetFinalCostInEuroOk returns a tuple with the FinalCostInEuro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalCostInEuro

`func (o *ExecutionLog) SetFinalCostInEuro(v float64)`

SetFinalCostInEuro sets FinalCostInEuro field to given value.

### HasFinalCostInEuro

`func (o *ExecutionLog) HasFinalCostInEuro() bool`

HasFinalCostInEuro returns a boolean if a field has been set.

### GetId

`func (o *ExecutionLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExecutionLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessageId

`func (o *ExecutionLog) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ExecutionLog) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ExecutionLog) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ExecutionLog) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetOwnerId

`func (o *ExecutionLog) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ExecutionLog) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ExecutionLog) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ExecutionLog) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerOrganizationId

`func (o *ExecutionLog) GetOwnerOrganizationId() string`

GetOwnerOrganizationId returns the OwnerOrganizationId field if non-nil, zero value otherwise.

### GetOwnerOrganizationIdOk

`func (o *ExecutionLog) GetOwnerOrganizationIdOk() (*string, bool)`

GetOwnerOrganizationIdOk returns a tuple with the OwnerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrganizationId

`func (o *ExecutionLog) SetOwnerOrganizationId(v string)`

SetOwnerOrganizationId sets OwnerOrganizationId field to given value.

### HasOwnerOrganizationId

`func (o *ExecutionLog) HasOwnerOrganizationId() bool`

HasOwnerOrganizationId returns a boolean if a field has been set.

### GetUsage

`func (o *ExecutionLog) GetUsage() []ExecutionUsageCost`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ExecutionLog) GetUsageOk() (*[]ExecutionUsageCost, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ExecutionLog) SetUsage(v []ExecutionUsageCost)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ExecutionLog) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


