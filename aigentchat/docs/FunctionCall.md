# FunctionCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**HostingLocation** | Pointer to **string** |  | [optional] 
**I18n** | Pointer to [**map[string]FunctionCallI18n**](FunctionCallI18n.md) |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **string** |  | [optional] 
**ResponseVisible** | Pointer to **bool** |  | [optional] 
**ToolId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewFunctionCall

`func NewFunctionCall() *FunctionCall`

NewFunctionCall instantiates a new FunctionCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionCallWithDefaults

`func NewFunctionCallWithDefaults() *FunctionCall`

NewFunctionCallWithDefaults instantiates a new FunctionCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FunctionCall) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FunctionCall) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FunctionCall) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FunctionCall) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostingLocation

`func (o *FunctionCall) GetHostingLocation() string`

GetHostingLocation returns the HostingLocation field if non-nil, zero value otherwise.

### GetHostingLocationOk

`func (o *FunctionCall) GetHostingLocationOk() (*string, bool)`

GetHostingLocationOk returns a tuple with the HostingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingLocation

`func (o *FunctionCall) SetHostingLocation(v string)`

SetHostingLocation sets HostingLocation field to given value.

### HasHostingLocation

`func (o *FunctionCall) HasHostingLocation() bool`

HasHostingLocation returns a boolean if a field has been set.

### GetI18n

`func (o *FunctionCall) GetI18n() map[string]FunctionCallI18n`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *FunctionCall) GetI18nOk() (*map[string]FunctionCallI18n, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *FunctionCall) SetI18n(v map[string]FunctionCallI18n)`

SetI18n sets I18n field to given value.

### HasI18n

`func (o *FunctionCall) HasI18n() bool`

HasI18n returns a boolean if a field has been set.

### GetIcon

`func (o *FunctionCall) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *FunctionCall) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *FunctionCall) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *FunctionCall) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *FunctionCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FunctionCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FunctionCall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FunctionCall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FunctionCall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionCall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionCall) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionCall) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *FunctionCall) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *FunctionCall) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *FunctionCall) SetParameters(v string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *FunctionCall) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetResponseVisible

`func (o *FunctionCall) GetResponseVisible() bool`

GetResponseVisible returns the ResponseVisible field if non-nil, zero value otherwise.

### GetResponseVisibleOk

`func (o *FunctionCall) GetResponseVisibleOk() (*bool, bool)`

GetResponseVisibleOk returns a tuple with the ResponseVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseVisible

`func (o *FunctionCall) SetResponseVisible(v bool)`

SetResponseVisible sets ResponseVisible field to given value.

### HasResponseVisible

`func (o *FunctionCall) HasResponseVisible() bool`

HasResponseVisible returns a boolean if a field has been set.

### GetToolId

`func (o *FunctionCall) GetToolId() string`

GetToolId returns the ToolId field if non-nil, zero value otherwise.

### GetToolIdOk

`func (o *FunctionCall) GetToolIdOk() (*string, bool)`

GetToolIdOk returns a tuple with the ToolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolId

`func (o *FunctionCall) SetToolId(v string)`

SetToolId sets ToolId field to given value.

### HasToolId

`func (o *FunctionCall) HasToolId() bool`

HasToolId returns a boolean if a field has been set.

### GetVersion

`func (o *FunctionCall) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FunctionCall) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FunctionCall) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FunctionCall) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


