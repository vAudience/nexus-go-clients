# UserSettingsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColorScheme** | **string** |  | 
**Language** | Pointer to **string** |  | [optional] 
**OnboardingDone** | Pointer to **bool** |  | [optional] 
**OnboardingHoldUntil** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewUserSettingsPostRequest

`func NewUserSettingsPostRequest(colorScheme string, ) *UserSettingsPostRequest`

NewUserSettingsPostRequest instantiates a new UserSettingsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsPostRequestWithDefaults

`func NewUserSettingsPostRequestWithDefaults() *UserSettingsPostRequest`

NewUserSettingsPostRequestWithDefaults instantiates a new UserSettingsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColorScheme

`func (o *UserSettingsPostRequest) GetColorScheme() string`

GetColorScheme returns the ColorScheme field if non-nil, zero value otherwise.

### GetColorSchemeOk

`func (o *UserSettingsPostRequest) GetColorSchemeOk() (*string, bool)`

GetColorSchemeOk returns a tuple with the ColorScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorScheme

`func (o *UserSettingsPostRequest) SetColorScheme(v string)`

SetColorScheme sets ColorScheme field to given value.


### GetLanguage

`func (o *UserSettingsPostRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserSettingsPostRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserSettingsPostRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UserSettingsPostRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetOnboardingDone

`func (o *UserSettingsPostRequest) GetOnboardingDone() bool`

GetOnboardingDone returns the OnboardingDone field if non-nil, zero value otherwise.

### GetOnboardingDoneOk

`func (o *UserSettingsPostRequest) GetOnboardingDoneOk() (*bool, bool)`

GetOnboardingDoneOk returns a tuple with the OnboardingDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingDone

`func (o *UserSettingsPostRequest) SetOnboardingDone(v bool)`

SetOnboardingDone sets OnboardingDone field to given value.

### HasOnboardingDone

`func (o *UserSettingsPostRequest) HasOnboardingDone() bool`

HasOnboardingDone returns a boolean if a field has been set.

### GetOnboardingHoldUntil

`func (o *UserSettingsPostRequest) GetOnboardingHoldUntil() string`

GetOnboardingHoldUntil returns the OnboardingHoldUntil field if non-nil, zero value otherwise.

### GetOnboardingHoldUntilOk

`func (o *UserSettingsPostRequest) GetOnboardingHoldUntilOk() (*string, bool)`

GetOnboardingHoldUntilOk returns a tuple with the OnboardingHoldUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingHoldUntil

`func (o *UserSettingsPostRequest) SetOnboardingHoldUntil(v string)`

SetOnboardingHoldUntil sets OnboardingHoldUntil field to given value.

### HasOnboardingHoldUntil

`func (o *UserSettingsPostRequest) HasOnboardingHoldUntil() bool`

HasOnboardingHoldUntil returns a boolean if a field has been set.

### GetUserId

`func (o *UserSettingsPostRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserSettingsPostRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserSettingsPostRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserSettingsPostRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


