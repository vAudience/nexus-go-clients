# MemberDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ResourcesDeleted** | Pointer to **[]string** |  | [optional] 
**ResourcesTransferred** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMemberDeleteResponse

`func NewMemberDeleteResponse(id string, ) *MemberDeleteResponse`

NewMemberDeleteResponse instantiates a new MemberDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDeleteResponseWithDefaults

`func NewMemberDeleteResponseWithDefaults() *MemberDeleteResponse`

NewMemberDeleteResponseWithDefaults instantiates a new MemberDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberDeleteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberDeleteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberDeleteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetResourcesDeleted

`func (o *MemberDeleteResponse) GetResourcesDeleted() []string`

GetResourcesDeleted returns the ResourcesDeleted field if non-nil, zero value otherwise.

### GetResourcesDeletedOk

`func (o *MemberDeleteResponse) GetResourcesDeletedOk() (*[]string, bool)`

GetResourcesDeletedOk returns a tuple with the ResourcesDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesDeleted

`func (o *MemberDeleteResponse) SetResourcesDeleted(v []string)`

SetResourcesDeleted sets ResourcesDeleted field to given value.

### HasResourcesDeleted

`func (o *MemberDeleteResponse) HasResourcesDeleted() bool`

HasResourcesDeleted returns a boolean if a field has been set.

### GetResourcesTransferred

`func (o *MemberDeleteResponse) GetResourcesTransferred() []string`

GetResourcesTransferred returns the ResourcesTransferred field if non-nil, zero value otherwise.

### GetResourcesTransferredOk

`func (o *MemberDeleteResponse) GetResourcesTransferredOk() (*[]string, bool)`

GetResourcesTransferredOk returns a tuple with the ResourcesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesTransferred

`func (o *MemberDeleteResponse) SetResourcesTransferred(v []string)`

SetResourcesTransferred sets ResourcesTransferred field to given value.

### HasResourcesTransferred

`func (o *MemberDeleteResponse) HasResourcesTransferred() bool`

HasResourcesTransferred returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


