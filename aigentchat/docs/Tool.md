# Tool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Functions** | [**[]ToolFunction**](ToolFunction.md) |  | 
**HostingLocations** | Pointer to **[]string** |  | [optional] 
**I18n** | Pointer to [**map[string]ToolI18n**](ToolI18n.md) |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**IsPublic** | **bool** |  | 
**Name** | **string** |  | 
**Version** | **string** |  | 

## Methods

### NewTool

`func NewTool(functions []ToolFunction, id string, isPublic bool, name string, version string, ) *Tool`

NewTool instantiates a new Tool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolWithDefaults

`func NewToolWithDefaults() *Tool`

NewToolWithDefaults instantiates a new Tool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Tool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunctions

`func (o *Tool) GetFunctions() []ToolFunction`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *Tool) GetFunctionsOk() (*[]ToolFunction, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *Tool) SetFunctions(v []ToolFunction)`

SetFunctions sets Functions field to given value.


### GetHostingLocations

`func (o *Tool) GetHostingLocations() []string`

GetHostingLocations returns the HostingLocations field if non-nil, zero value otherwise.

### GetHostingLocationsOk

`func (o *Tool) GetHostingLocationsOk() (*[]string, bool)`

GetHostingLocationsOk returns a tuple with the HostingLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingLocations

`func (o *Tool) SetHostingLocations(v []string)`

SetHostingLocations sets HostingLocations field to given value.

### HasHostingLocations

`func (o *Tool) HasHostingLocations() bool`

HasHostingLocations returns a boolean if a field has been set.

### GetI18n

`func (o *Tool) GetI18n() map[string]ToolI18n`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *Tool) GetI18nOk() (*map[string]ToolI18n, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *Tool) SetI18n(v map[string]ToolI18n)`

SetI18n sets I18n field to given value.

### HasI18n

`func (o *Tool) HasI18n() bool`

HasI18n returns a boolean if a field has been set.

### GetIcon

`func (o *Tool) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Tool) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Tool) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Tool) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *Tool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tool) SetId(v string)`

SetId sets Id field to given value.


### GetIsPublic

`func (o *Tool) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Tool) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Tool) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetName

`func (o *Tool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tool) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Tool) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Tool) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Tool) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


