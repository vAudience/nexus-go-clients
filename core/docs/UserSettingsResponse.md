# UserSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColorScheme** | **string** |  | 
**Id** | **string** |  | 
**Language** | Pointer to **string** |  | [optional] 
**OnboardingCompletedTours** | Pointer to **[]string** |  | [optional] 
**OnboardingDone** | Pointer to **bool** |  | [optional] 
**OnboardingHoldUntil** | Pointer to **string** |  | [optional] 
**OnboardingState** | Pointer to **string** |  | [optional] 
**UserId** | **string** |  | 

## Methods

### NewUserSettingsResponse

`func NewUserSettingsResponse(colorScheme string, id string, userId string, ) *UserSettingsResponse`

NewUserSettingsResponse instantiates a new UserSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsResponseWithDefaults

`func NewUserSettingsResponseWithDefaults() *UserSettingsResponse`

NewUserSettingsResponseWithDefaults instantiates a new UserSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColorScheme

`func (o *UserSettingsResponse) GetColorScheme() string`

GetColorScheme returns the ColorScheme field if non-nil, zero value otherwise.

### GetColorSchemeOk

`func (o *UserSettingsResponse) GetColorSchemeOk() (*string, bool)`

GetColorSchemeOk returns a tuple with the ColorScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorScheme

`func (o *UserSettingsResponse) SetColorScheme(v string)`

SetColorScheme sets ColorScheme field to given value.


### GetId

`func (o *UserSettingsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSettingsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSettingsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLanguage

`func (o *UserSettingsResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserSettingsResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserSettingsResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UserSettingsResponse) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetOnboardingCompletedTours

`func (o *UserSettingsResponse) GetOnboardingCompletedTours() []string`

GetOnboardingCompletedTours returns the OnboardingCompletedTours field if non-nil, zero value otherwise.

### GetOnboardingCompletedToursOk

`func (o *UserSettingsResponse) GetOnboardingCompletedToursOk() (*[]string, bool)`

GetOnboardingCompletedToursOk returns a tuple with the OnboardingCompletedTours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingCompletedTours

`func (o *UserSettingsResponse) SetOnboardingCompletedTours(v []string)`

SetOnboardingCompletedTours sets OnboardingCompletedTours field to given value.

### HasOnboardingCompletedTours

`func (o *UserSettingsResponse) HasOnboardingCompletedTours() bool`

HasOnboardingCompletedTours returns a boolean if a field has been set.

### GetOnboardingDone

`func (o *UserSettingsResponse) GetOnboardingDone() bool`

GetOnboardingDone returns the OnboardingDone field if non-nil, zero value otherwise.

### GetOnboardingDoneOk

`func (o *UserSettingsResponse) GetOnboardingDoneOk() (*bool, bool)`

GetOnboardingDoneOk returns a tuple with the OnboardingDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingDone

`func (o *UserSettingsResponse) SetOnboardingDone(v bool)`

SetOnboardingDone sets OnboardingDone field to given value.

### HasOnboardingDone

`func (o *UserSettingsResponse) HasOnboardingDone() bool`

HasOnboardingDone returns a boolean if a field has been set.

### GetOnboardingHoldUntil

`func (o *UserSettingsResponse) GetOnboardingHoldUntil() string`

GetOnboardingHoldUntil returns the OnboardingHoldUntil field if non-nil, zero value otherwise.

### GetOnboardingHoldUntilOk

`func (o *UserSettingsResponse) GetOnboardingHoldUntilOk() (*string, bool)`

GetOnboardingHoldUntilOk returns a tuple with the OnboardingHoldUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingHoldUntil

`func (o *UserSettingsResponse) SetOnboardingHoldUntil(v string)`

SetOnboardingHoldUntil sets OnboardingHoldUntil field to given value.

### HasOnboardingHoldUntil

`func (o *UserSettingsResponse) HasOnboardingHoldUntil() bool`

HasOnboardingHoldUntil returns a boolean if a field has been set.

### GetOnboardingState

`func (o *UserSettingsResponse) GetOnboardingState() string`

GetOnboardingState returns the OnboardingState field if non-nil, zero value otherwise.

### GetOnboardingStateOk

`func (o *UserSettingsResponse) GetOnboardingStateOk() (*string, bool)`

GetOnboardingStateOk returns a tuple with the OnboardingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingState

`func (o *UserSettingsResponse) SetOnboardingState(v string)`

SetOnboardingState sets OnboardingState field to given value.

### HasOnboardingState

`func (o *UserSettingsResponse) HasOnboardingState() bool`

HasOnboardingState returns a boolean if a field has been set.

### GetUserId

`func (o *UserSettingsResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserSettingsResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserSettingsResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


