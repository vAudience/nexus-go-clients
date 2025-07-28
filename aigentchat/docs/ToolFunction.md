# ToolFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**HostingLocation** | Pointer to **string** |  | [optional] 
**I18n** | Pointer to [**map[string]ToolI18n**](ToolI18n.md) |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Parameters** | Pointer to **string** |  | [optional] 
**ResponseVisible** | **bool** |  | 
**ToolId** | **string** |  | 
**Version** | **string** |  | 

## Methods

### NewToolFunction

`func NewToolFunction(id string, name string, responseVisible bool, toolId string, version string, ) *ToolFunction`

NewToolFunction instantiates a new ToolFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolFunctionWithDefaults

`func NewToolFunctionWithDefaults() *ToolFunction`

NewToolFunctionWithDefaults instantiates a new ToolFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ToolFunction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolFunction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolFunction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ToolFunction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostingLocation

`func (o *ToolFunction) GetHostingLocation() string`

GetHostingLocation returns the HostingLocation field if non-nil, zero value otherwise.

### GetHostingLocationOk

`func (o *ToolFunction) GetHostingLocationOk() (*string, bool)`

GetHostingLocationOk returns a tuple with the HostingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingLocation

`func (o *ToolFunction) SetHostingLocation(v string)`

SetHostingLocation sets HostingLocation field to given value.

### HasHostingLocation

`func (o *ToolFunction) HasHostingLocation() bool`

HasHostingLocation returns a boolean if a field has been set.

### GetI18n

`func (o *ToolFunction) GetI18n() map[string]ToolI18n`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *ToolFunction) GetI18nOk() (*map[string]ToolI18n, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *ToolFunction) SetI18n(v map[string]ToolI18n)`

SetI18n sets I18n field to given value.

### HasI18n

`func (o *ToolFunction) HasI18n() bool`

HasI18n returns a boolean if a field has been set.

### GetIcon

`func (o *ToolFunction) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ToolFunction) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ToolFunction) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ToolFunction) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *ToolFunction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolFunction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolFunction) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ToolFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolFunction) SetName(v string)`

SetName sets Name field to given value.


### GetParameters

`func (o *ToolFunction) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ToolFunction) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ToolFunction) SetParameters(v string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ToolFunction) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetResponseVisible

`func (o *ToolFunction) GetResponseVisible() bool`

GetResponseVisible returns the ResponseVisible field if non-nil, zero value otherwise.

### GetResponseVisibleOk

`func (o *ToolFunction) GetResponseVisibleOk() (*bool, bool)`

GetResponseVisibleOk returns a tuple with the ResponseVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseVisible

`func (o *ToolFunction) SetResponseVisible(v bool)`

SetResponseVisible sets ResponseVisible field to given value.


### GetToolId

`func (o *ToolFunction) GetToolId() string`

GetToolId returns the ToolId field if non-nil, zero value otherwise.

### GetToolIdOk

`func (o *ToolFunction) GetToolIdOk() (*string, bool)`

GetToolIdOk returns a tuple with the ToolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolId

`func (o *ToolFunction) SetToolId(v string)`

SetToolId sets ToolId field to given value.


### GetVersion

`func (o *ToolFunction) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ToolFunction) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ToolFunction) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


