# ExecutionResultAudioGeneration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | LocalFilePaths []string &#x60;json:\&quot;media_file_paths\&quot;&#x60; | [optional] 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**ExecutionId** | Pointer to **string** | maps to CompletionID for Text2Text | [optional] 
**FeaturesUsed** | Pointer to [**[]AIModelFeature**](AIModelFeature.md) |  | [optional] 
**FinishReason** | Pointer to **string** |  | [optional] 
**InputTokens** | Pointer to **int32** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**ResultingFiles** | Pointer to [**[]FilemanagerManagedFile**](FilemanagerManagedFile.md) |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**TimeNeeded** | Pointer to **int64** | in Milliseconds | [optional] 
**Timestamp** | Pointer to **int64** | in Milliseconds unix epoch | [optional] 

## Methods

### NewExecutionResultAudioGeneration

`func NewExecutionResultAudioGeneration() *ExecutionResultAudioGeneration`

NewExecutionResultAudioGeneration instantiates a new ExecutionResultAudioGeneration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionResultAudioGenerationWithDefaults

`func NewExecutionResultAudioGenerationWithDefaults() *ExecutionResultAudioGeneration`

NewExecutionResultAudioGenerationWithDefaults instantiates a new ExecutionResultAudioGeneration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ExecutionResultAudioGeneration) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ExecutionResultAudioGeneration) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ExecutionResultAudioGeneration) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ExecutionResultAudioGeneration) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetError

`func (o *ExecutionResultAudioGeneration) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExecutionResultAudioGeneration) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExecutionResultAudioGeneration) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *ExecutionResultAudioGeneration) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExecutionId

`func (o *ExecutionResultAudioGeneration) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ExecutionResultAudioGeneration) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ExecutionResultAudioGeneration) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ExecutionResultAudioGeneration) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetFeaturesUsed

`func (o *ExecutionResultAudioGeneration) GetFeaturesUsed() []AIModelFeature`

GetFeaturesUsed returns the FeaturesUsed field if non-nil, zero value otherwise.

### GetFeaturesUsedOk

`func (o *ExecutionResultAudioGeneration) GetFeaturesUsedOk() (*[]AIModelFeature, bool)`

GetFeaturesUsedOk returns a tuple with the FeaturesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesUsed

`func (o *ExecutionResultAudioGeneration) SetFeaturesUsed(v []AIModelFeature)`

SetFeaturesUsed sets FeaturesUsed field to given value.

### HasFeaturesUsed

`func (o *ExecutionResultAudioGeneration) HasFeaturesUsed() bool`

HasFeaturesUsed returns a boolean if a field has been set.

### GetFinishReason

`func (o *ExecutionResultAudioGeneration) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *ExecutionResultAudioGeneration) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *ExecutionResultAudioGeneration) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *ExecutionResultAudioGeneration) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.

### GetInputTokens

`func (o *ExecutionResultAudioGeneration) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *ExecutionResultAudioGeneration) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *ExecutionResultAudioGeneration) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *ExecutionResultAudioGeneration) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetModelId

`func (o *ExecutionResultAudioGeneration) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ExecutionResultAudioGeneration) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ExecutionResultAudioGeneration) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *ExecutionResultAudioGeneration) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetResultingFiles

`func (o *ExecutionResultAudioGeneration) GetResultingFiles() []FilemanagerManagedFile`

GetResultingFiles returns the ResultingFiles field if non-nil, zero value otherwise.

### GetResultingFilesOk

`func (o *ExecutionResultAudioGeneration) GetResultingFilesOk() (*[]FilemanagerManagedFile, bool)`

GetResultingFilesOk returns a tuple with the ResultingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultingFiles

`func (o *ExecutionResultAudioGeneration) SetResultingFiles(v []FilemanagerManagedFile)`

SetResultingFiles sets ResultingFiles field to given value.

### HasResultingFiles

`func (o *ExecutionResultAudioGeneration) HasResultingFiles() bool`

HasResultingFiles returns a boolean if a field has been set.

### GetServiceId

`func (o *ExecutionResultAudioGeneration) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ExecutionResultAudioGeneration) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ExecutionResultAudioGeneration) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ExecutionResultAudioGeneration) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetTimeNeeded

`func (o *ExecutionResultAudioGeneration) GetTimeNeeded() int64`

GetTimeNeeded returns the TimeNeeded field if non-nil, zero value otherwise.

### GetTimeNeededOk

`func (o *ExecutionResultAudioGeneration) GetTimeNeededOk() (*int64, bool)`

GetTimeNeededOk returns a tuple with the TimeNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeNeeded

`func (o *ExecutionResultAudioGeneration) SetTimeNeeded(v int64)`

SetTimeNeeded sets TimeNeeded field to given value.

### HasTimeNeeded

`func (o *ExecutionResultAudioGeneration) HasTimeNeeded() bool`

HasTimeNeeded returns a boolean if a field has been set.

### GetTimestamp

`func (o *ExecutionResultAudioGeneration) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExecutionResultAudioGeneration) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExecutionResultAudioGeneration) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExecutionResultAudioGeneration) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


