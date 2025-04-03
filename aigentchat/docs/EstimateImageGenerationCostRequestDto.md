# EstimateImageGenerationCostRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** |  | 
**AttachedTemporaryFiles** | Pointer to **[]string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEstimateImageGenerationCostRequestDto

`func NewEstimateImageGenerationCostRequestDto(agentId string, ) *EstimateImageGenerationCostRequestDto`

NewEstimateImageGenerationCostRequestDto instantiates a new EstimateImageGenerationCostRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateImageGenerationCostRequestDtoWithDefaults

`func NewEstimateImageGenerationCostRequestDtoWithDefaults() *EstimateImageGenerationCostRequestDto`

NewEstimateImageGenerationCostRequestDtoWithDefaults instantiates a new EstimateImageGenerationCostRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *EstimateImageGenerationCostRequestDto) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *EstimateImageGenerationCostRequestDto) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *EstimateImageGenerationCostRequestDto) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetAttachedTemporaryFiles

`func (o *EstimateImageGenerationCostRequestDto) GetAttachedTemporaryFiles() []string`

GetAttachedTemporaryFiles returns the AttachedTemporaryFiles field if non-nil, zero value otherwise.

### GetAttachedTemporaryFilesOk

`func (o *EstimateImageGenerationCostRequestDto) GetAttachedTemporaryFilesOk() (*[]string, bool)`

GetAttachedTemporaryFilesOk returns a tuple with the AttachedTemporaryFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedTemporaryFiles

`func (o *EstimateImageGenerationCostRequestDto) SetAttachedTemporaryFiles(v []string)`

SetAttachedTemporaryFiles sets AttachedTemporaryFiles field to given value.

### HasAttachedTemporaryFiles

`func (o *EstimateImageGenerationCostRequestDto) HasAttachedTemporaryFiles() bool`

HasAttachedTemporaryFiles returns a boolean if a field has been set.

### GetParameters

`func (o *EstimateImageGenerationCostRequestDto) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *EstimateImageGenerationCostRequestDto) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *EstimateImageGenerationCostRequestDto) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *EstimateImageGenerationCostRequestDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


