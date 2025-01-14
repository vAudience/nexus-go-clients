# AIModelServiceWithModels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | Pointer to [**[]AIModel**](AIModel.md) |  | [optional] 
**Service** | Pointer to [**AIModelServiceObject**](AIModelServiceObject.md) |  | [optional] 

## Methods

### NewAIModelServiceWithModels

`func NewAIModelServiceWithModels() *AIModelServiceWithModels`

NewAIModelServiceWithModels instantiates a new AIModelServiceWithModels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIModelServiceWithModelsWithDefaults

`func NewAIModelServiceWithModelsWithDefaults() *AIModelServiceWithModels`

NewAIModelServiceWithModelsWithDefaults instantiates a new AIModelServiceWithModels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *AIModelServiceWithModels) GetModels() []AIModel`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *AIModelServiceWithModels) GetModelsOk() (*[]AIModel, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *AIModelServiceWithModels) SetModels(v []AIModel)`

SetModels sets Models field to given value.

### HasModels

`func (o *AIModelServiceWithModels) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetService

`func (o *AIModelServiceWithModels) GetService() AIModelServiceObject`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AIModelServiceWithModels) GetServiceOk() (*AIModelServiceObject, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AIModelServiceWithModels) SetService(v AIModelServiceObject)`

SetService sets Service field to given value.

### HasService

`func (o *AIModelServiceWithModels) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


