# HealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to [**map[string]ServiceStatus**](ServiceStatus.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthResponse

`func NewHealthResponse() *HealthResponse`

NewHealthResponse instantiates a new HealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthResponseWithDefaults

`func NewHealthResponseWithDefaults() *HealthResponse`

NewHealthResponseWithDefaults instantiates a new HealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *HealthResponse) GetServices() map[string]ServiceStatus`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *HealthResponse) GetServicesOk() (*map[string]ServiceStatus, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *HealthResponse) SetServices(v map[string]ServiceStatus)`

SetServices sets Services field to given value.

### HasServices

`func (o *HealthResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetStatus

`func (o *HealthResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


