# RolePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMemberRole** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Permissions** | **[]string** |  | 

## Methods

### NewRolePostRequest

`func NewRolePostRequest(name string, permissions []string, ) *RolePostRequest`

NewRolePostRequest instantiates a new RolePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePostRequestWithDefaults

`func NewRolePostRequestWithDefaults() *RolePostRequest`

NewRolePostRequestWithDefaults instantiates a new RolePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMemberRole

`func (o *RolePostRequest) GetDefaultMemberRole() bool`

GetDefaultMemberRole returns the DefaultMemberRole field if non-nil, zero value otherwise.

### GetDefaultMemberRoleOk

`func (o *RolePostRequest) GetDefaultMemberRoleOk() (*bool, bool)`

GetDefaultMemberRoleOk returns a tuple with the DefaultMemberRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemberRole

`func (o *RolePostRequest) SetDefaultMemberRole(v bool)`

SetDefaultMemberRole sets DefaultMemberRole field to given value.

### HasDefaultMemberRole

`func (o *RolePostRequest) HasDefaultMemberRole() bool`

HasDefaultMemberRole returns a boolean if a field has been set.

### GetName

`func (o *RolePostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RolePostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RolePostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *RolePostRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RolePostRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RolePostRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


