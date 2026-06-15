# BulkInviteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**ErrorCode** | Pointer to **string** |  | [optional] 
**InviteId** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewBulkInviteResult

`func NewBulkInviteResult(email string, status string, ) *BulkInviteResult`

NewBulkInviteResult instantiates a new BulkInviteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInviteResultWithDefaults

`func NewBulkInviteResultWithDefaults() *BulkInviteResult`

NewBulkInviteResultWithDefaults instantiates a new BulkInviteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *BulkInviteResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BulkInviteResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BulkInviteResult) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetErrorCode

`func (o *BulkInviteResult) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BulkInviteResult) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BulkInviteResult) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BulkInviteResult) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetInviteId

`func (o *BulkInviteResult) GetInviteId() string`

GetInviteId returns the InviteId field if non-nil, zero value otherwise.

### GetInviteIdOk

`func (o *BulkInviteResult) GetInviteIdOk() (*string, bool)`

GetInviteIdOk returns a tuple with the InviteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteId

`func (o *BulkInviteResult) SetInviteId(v string)`

SetInviteId sets InviteId field to given value.

### HasInviteId

`func (o *BulkInviteResult) HasInviteId() bool`

HasInviteId returns a boolean if a field has been set.

### GetStatus

`func (o *BulkInviteResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkInviteResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkInviteResult) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


