# AdapterExecutionResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdapterName** | Pointer to **string** |  | [optional] 
**Err** | Pointer to **map[string]interface{}** |  | [optional] 
**FinalState** | Pointer to [**AdapterExecutionState**](AdapterExecutionState.md) |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**ResultFiles** | Pointer to [**[]AdapterFileInfo**](AdapterFileInfo.md) |  | [optional] 
**ResultTexts** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdapterExecutionResults

`func NewAdapterExecutionResults() *AdapterExecutionResults`

NewAdapterExecutionResults instantiates a new AdapterExecutionResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterExecutionResultsWithDefaults

`func NewAdapterExecutionResultsWithDefaults() *AdapterExecutionResults`

NewAdapterExecutionResultsWithDefaults instantiates a new AdapterExecutionResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdapterName

`func (o *AdapterExecutionResults) GetAdapterName() string`

GetAdapterName returns the AdapterName field if non-nil, zero value otherwise.

### GetAdapterNameOk

`func (o *AdapterExecutionResults) GetAdapterNameOk() (*string, bool)`

GetAdapterNameOk returns a tuple with the AdapterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterName

`func (o *AdapterExecutionResults) SetAdapterName(v string)`

SetAdapterName sets AdapterName field to given value.

### HasAdapterName

`func (o *AdapterExecutionResults) HasAdapterName() bool`

HasAdapterName returns a boolean if a field has been set.

### GetErr

`func (o *AdapterExecutionResults) GetErr() map[string]interface{}`

GetErr returns the Err field if non-nil, zero value otherwise.

### GetErrOk

`func (o *AdapterExecutionResults) GetErrOk() (*map[string]interface{}, bool)`

GetErrOk returns a tuple with the Err field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErr

`func (o *AdapterExecutionResults) SetErr(v map[string]interface{})`

SetErr sets Err field to given value.

### HasErr

`func (o *AdapterExecutionResults) HasErr() bool`

HasErr returns a boolean if a field has been set.

### GetFinalState

`func (o *AdapterExecutionResults) GetFinalState() AdapterExecutionState`

GetFinalState returns the FinalState field if non-nil, zero value otherwise.

### GetFinalStateOk

`func (o *AdapterExecutionResults) GetFinalStateOk() (*AdapterExecutionState, bool)`

GetFinalStateOk returns a tuple with the FinalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalState

`func (o *AdapterExecutionResults) SetFinalState(v AdapterExecutionState)`

SetFinalState sets FinalState field to given value.

### HasFinalState

`func (o *AdapterExecutionResults) HasFinalState() bool`

HasFinalState returns a boolean if a field has been set.

### GetJobId

`func (o *AdapterExecutionResults) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *AdapterExecutionResults) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *AdapterExecutionResults) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *AdapterExecutionResults) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetResultFiles

`func (o *AdapterExecutionResults) GetResultFiles() []AdapterFileInfo`

GetResultFiles returns the ResultFiles field if non-nil, zero value otherwise.

### GetResultFilesOk

`func (o *AdapterExecutionResults) GetResultFilesOk() (*[]AdapterFileInfo, bool)`

GetResultFilesOk returns a tuple with the ResultFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFiles

`func (o *AdapterExecutionResults) SetResultFiles(v []AdapterFileInfo)`

SetResultFiles sets ResultFiles field to given value.

### HasResultFiles

`func (o *AdapterExecutionResults) HasResultFiles() bool`

HasResultFiles returns a boolean if a field has been set.

### GetResultTexts

`func (o *AdapterExecutionResults) GetResultTexts() []string`

GetResultTexts returns the ResultTexts field if non-nil, zero value otherwise.

### GetResultTextsOk

`func (o *AdapterExecutionResults) GetResultTextsOk() (*[]string, bool)`

GetResultTextsOk returns a tuple with the ResultTexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTexts

`func (o *AdapterExecutionResults) SetResultTexts(v []string)`

SetResultTexts sets ResultTexts field to given value.

### HasResultTexts

`func (o *AdapterExecutionResults) HasResultTexts() bool`

HasResultTexts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


