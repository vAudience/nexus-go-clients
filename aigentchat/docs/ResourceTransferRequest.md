# ResourceTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceTypes** | [**[]TransferableResourceType**](TransferableResourceType.md) |  | 
**SourceOwnerId** | **string** |  | 
**TargetOrganizationId** | **string** |  | 
**TargetOwnerId** | Pointer to **string** |  | [optional] 

## Methods

### NewResourceTransferRequest

`func NewResourceTransferRequest(resourceTypes []TransferableResourceType, sourceOwnerId string, targetOrganizationId string, ) *ResourceTransferRequest`

NewResourceTransferRequest instantiates a new ResourceTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTransferRequestWithDefaults

`func NewResourceTransferRequestWithDefaults() *ResourceTransferRequest`

NewResourceTransferRequestWithDefaults instantiates a new ResourceTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceTypes

`func (o *ResourceTransferRequest) GetResourceTypes() []TransferableResourceType`

GetResourceTypes returns the ResourceTypes field if non-nil, zero value otherwise.

### GetResourceTypesOk

`func (o *ResourceTransferRequest) GetResourceTypesOk() (*[]TransferableResourceType, bool)`

GetResourceTypesOk returns a tuple with the ResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypes

`func (o *ResourceTransferRequest) SetResourceTypes(v []TransferableResourceType)`

SetResourceTypes sets ResourceTypes field to given value.


### GetSourceOwnerId

`func (o *ResourceTransferRequest) GetSourceOwnerId() string`

GetSourceOwnerId returns the SourceOwnerId field if non-nil, zero value otherwise.

### GetSourceOwnerIdOk

`func (o *ResourceTransferRequest) GetSourceOwnerIdOk() (*string, bool)`

GetSourceOwnerIdOk returns a tuple with the SourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOwnerId

`func (o *ResourceTransferRequest) SetSourceOwnerId(v string)`

SetSourceOwnerId sets SourceOwnerId field to given value.


### GetTargetOrganizationId

`func (o *ResourceTransferRequest) GetTargetOrganizationId() string`

GetTargetOrganizationId returns the TargetOrganizationId field if non-nil, zero value otherwise.

### GetTargetOrganizationIdOk

`func (o *ResourceTransferRequest) GetTargetOrganizationIdOk() (*string, bool)`

GetTargetOrganizationIdOk returns a tuple with the TargetOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOrganizationId

`func (o *ResourceTransferRequest) SetTargetOrganizationId(v string)`

SetTargetOrganizationId sets TargetOrganizationId field to given value.


### GetTargetOwnerId

`func (o *ResourceTransferRequest) GetTargetOwnerId() string`

GetTargetOwnerId returns the TargetOwnerId field if non-nil, zero value otherwise.

### GetTargetOwnerIdOk

`func (o *ResourceTransferRequest) GetTargetOwnerIdOk() (*string, bool)`

GetTargetOwnerIdOk returns a tuple with the TargetOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOwnerId

`func (o *ResourceTransferRequest) SetTargetOwnerId(v string)`

SetTargetOwnerId sets TargetOwnerId field to given value.

### HasTargetOwnerId

`func (o *ResourceTransferRequest) HasTargetOwnerId() bool`

HasTargetOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


