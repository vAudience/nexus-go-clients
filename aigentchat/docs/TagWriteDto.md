# TagWriteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**I18n** | Pointer to [**map[string]TagI18n**](TagI18n.md) |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewTagWriteDto

`func NewTagWriteDto() *TagWriteDto`

NewTagWriteDto instantiates a new TagWriteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWriteDtoWithDefaults

`func NewTagWriteDtoWithDefaults() *TagWriteDto`

NewTagWriteDtoWithDefaults instantiates a new TagWriteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetI18n

`func (o *TagWriteDto) GetI18n() map[string]TagI18n`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *TagWriteDto) GetI18nOk() (*map[string]TagI18n, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *TagWriteDto) SetI18n(v map[string]TagI18n)`

SetI18n sets I18n field to given value.

### HasI18n

`func (o *TagWriteDto) HasI18n() bool`

HasI18n returns a boolean if a field has been set.

### GetInternalId

`func (o *TagWriteDto) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *TagWriteDto) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *TagWriteDto) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *TagWriteDto) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetIsPublic

`func (o *TagWriteDto) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *TagWriteDto) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *TagWriteDto) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *TagWriteDto) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetName

`func (o *TagWriteDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagWriteDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagWriteDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagWriteDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *TagWriteDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagWriteDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagWriteDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TagWriteDto) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


