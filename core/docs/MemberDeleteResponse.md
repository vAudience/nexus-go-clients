# MemberDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ResourceIds** | **[]string** |  | 
**ResourcesAction** | **string** |  | 

## Methods

### NewMemberDeleteResponse

`func NewMemberDeleteResponse(id string, resourceIds []string, resourcesAction string, ) *MemberDeleteResponse`

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


### GetResourceIds

`func (o *MemberDeleteResponse) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *MemberDeleteResponse) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *MemberDeleteResponse) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.


### GetResourcesAction

`func (o *MemberDeleteResponse) GetResourcesAction() string`

GetResourcesAction returns the ResourcesAction field if non-nil, zero value otherwise.

### GetResourcesActionOk

`func (o *MemberDeleteResponse) GetResourcesActionOk() (*string, bool)`

GetResourcesActionOk returns a tuple with the ResourcesAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesAction

`func (o *MemberDeleteResponse) SetResourcesAction(v string)`

SetResourcesAction sets ResourcesAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


