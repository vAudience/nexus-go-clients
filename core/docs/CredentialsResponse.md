# CredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpAllowed** | **bool** |  | 
**OtpEnabled** | **bool** |  | 

## Methods

### NewCredentialsResponse

`func NewCredentialsResponse(otpAllowed bool, otpEnabled bool, ) *CredentialsResponse`

NewCredentialsResponse instantiates a new CredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsResponseWithDefaults

`func NewCredentialsResponseWithDefaults() *CredentialsResponse`

NewCredentialsResponseWithDefaults instantiates a new CredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpAllowed

`func (o *CredentialsResponse) GetOtpAllowed() bool`

GetOtpAllowed returns the OtpAllowed field if non-nil, zero value otherwise.

### GetOtpAllowedOk

`func (o *CredentialsResponse) GetOtpAllowedOk() (*bool, bool)`

GetOtpAllowedOk returns a tuple with the OtpAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpAllowed

`func (o *CredentialsResponse) SetOtpAllowed(v bool)`

SetOtpAllowed sets OtpAllowed field to given value.


### GetOtpEnabled

`func (o *CredentialsResponse) GetOtpEnabled() bool`

GetOtpEnabled returns the OtpEnabled field if non-nil, zero value otherwise.

### GetOtpEnabledOk

`func (o *CredentialsResponse) GetOtpEnabledOk() (*bool, bool)`

GetOtpEnabledOk returns a tuple with the OtpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpEnabled

`func (o *CredentialsResponse) SetOtpEnabled(v bool)`

SetOtpEnabled sets OtpEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


