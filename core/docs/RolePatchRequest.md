# RolePatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMemberRole** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRolePatchRequest

`func NewRolePatchRequest() *RolePatchRequest`

NewRolePatchRequest instantiates a new RolePatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePatchRequestWithDefaults

`func NewRolePatchRequestWithDefaults() *RolePatchRequest`

NewRolePatchRequestWithDefaults instantiates a new RolePatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMemberRole

`func (o *RolePatchRequest) GetDefaultMemberRole() bool`

GetDefaultMemberRole returns the DefaultMemberRole field if non-nil, zero value otherwise.

### GetDefaultMemberRoleOk

`func (o *RolePatchRequest) GetDefaultMemberRoleOk() (*bool, bool)`

GetDefaultMemberRoleOk returns a tuple with the DefaultMemberRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemberRole

`func (o *RolePatchRequest) SetDefaultMemberRole(v bool)`

SetDefaultMemberRole sets DefaultMemberRole field to given value.

### HasDefaultMemberRole

`func (o *RolePatchRequest) HasDefaultMemberRole() bool`

HasDefaultMemberRole returns a boolean if a field has been set.

### GetName

`func (o *RolePatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RolePatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RolePatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RolePatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *RolePatchRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RolePatchRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RolePatchRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RolePatchRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


