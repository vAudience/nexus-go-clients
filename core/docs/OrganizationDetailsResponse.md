# OrganizationDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | [**OrganizationResponse**](OrganizationResponse.md) |  | 
**Subscription** | Pointer to [**OrganizationSubscriptionResponse**](OrganizationSubscriptionResponse.md) |  | [optional] 

## Methods

### NewOrganizationDetailsResponse

`func NewOrganizationDetailsResponse(organization OrganizationResponse, ) *OrganizationDetailsResponse`

NewOrganizationDetailsResponse instantiates a new OrganizationDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDetailsResponseWithDefaults

`func NewOrganizationDetailsResponseWithDefaults() *OrganizationDetailsResponse`

NewOrganizationDetailsResponseWithDefaults instantiates a new OrganizationDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *OrganizationDetailsResponse) GetOrganization() OrganizationResponse`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationDetailsResponse) GetOrganizationOk() (*OrganizationResponse, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationDetailsResponse) SetOrganization(v OrganizationResponse)`

SetOrganization sets Organization field to given value.


### GetSubscription

`func (o *OrganizationDetailsResponse) GetSubscription() OrganizationSubscriptionResponse`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *OrganizationDetailsResponse) GetSubscriptionOk() (*OrganizationSubscriptionResponse, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *OrganizationDetailsResponse) SetSubscription(v OrganizationSubscriptionResponse)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *OrganizationDetailsResponse) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


