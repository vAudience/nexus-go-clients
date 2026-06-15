# BulkInviteEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**RoleId** | Pointer to **string** | RoleId validity is enforced by json.Unmarshal (uuid.UUID implements TextUnmarshaler); the validate tag only gates presence. | [optional] 

## Methods

### NewBulkInviteEntry

`func NewBulkInviteEntry(email string, ) *BulkInviteEntry`

NewBulkInviteEntry instantiates a new BulkInviteEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInviteEntryWithDefaults

`func NewBulkInviteEntryWithDefaults() *BulkInviteEntry`

NewBulkInviteEntryWithDefaults instantiates a new BulkInviteEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *BulkInviteEntry) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BulkInviteEntry) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BulkInviteEntry) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRoleId

`func (o *BulkInviteEntry) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *BulkInviteEntry) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *BulkInviteEntry) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *BulkInviteEntry) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


