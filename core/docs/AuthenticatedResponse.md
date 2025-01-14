# AuthenticatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticated** | **bool** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TokenExpiresInSec** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticatedResponse

`func NewAuthenticatedResponse(authenticated bool, ) *AuthenticatedResponse`

NewAuthenticatedResponse instantiates a new AuthenticatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatedResponseWithDefaults

`func NewAuthenticatedResponseWithDefaults() *AuthenticatedResponse`

NewAuthenticatedResponseWithDefaults instantiates a new AuthenticatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticated

`func (o *AuthenticatedResponse) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *AuthenticatedResponse) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *AuthenticatedResponse) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.


### GetEmail

`func (o *AuthenticatedResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthenticatedResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthenticatedResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthenticatedResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *AuthenticatedResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatedResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatedResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticatedResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTokenExpiresInSec

`func (o *AuthenticatedResponse) GetTokenExpiresInSec() int32`

GetTokenExpiresInSec returns the TokenExpiresInSec field if non-nil, zero value otherwise.

### GetTokenExpiresInSecOk

`func (o *AuthenticatedResponse) GetTokenExpiresInSecOk() (*int32, bool)`

GetTokenExpiresInSecOk returns a tuple with the TokenExpiresInSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiresInSec

`func (o *AuthenticatedResponse) SetTokenExpiresInSec(v int32)`

SetTokenExpiresInSec sets TokenExpiresInSec field to given value.

### HasTokenExpiresInSec

`func (o *AuthenticatedResponse) HasTokenExpiresInSec() bool`

HasTokenExpiresInSec returns a boolean if a field has been set.

### GetUserId

`func (o *AuthenticatedResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthenticatedResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthenticatedResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuthenticatedResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *AuthenticatedResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuthenticatedResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuthenticatedResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AuthenticatedResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


