# TeamPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewTeamPostRequest

`func NewTeamPostRequest(name string, ) *TeamPostRequest`

NewTeamPostRequest instantiates a new TeamPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPostRequestWithDefaults

`func NewTeamPostRequestWithDefaults() *TeamPostRequest`

NewTeamPostRequestWithDefaults instantiates a new TeamPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaId

`func (o *TeamPostRequest) GetMetaId() string`

GetMetaId returns the MetaId field if non-nil, zero value otherwise.

### GetMetaIdOk

`func (o *TeamPostRequest) GetMetaIdOk() (*string, bool)`

GetMetaIdOk returns a tuple with the MetaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaId

`func (o *TeamPostRequest) SetMetaId(v string)`

SetMetaId sets MetaId field to given value.

### HasMetaId

`func (o *TeamPostRequest) HasMetaId() bool`

HasMetaId returns a boolean if a field has been set.

### GetName

`func (o *TeamPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamPostRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


