# ProcessingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** | ErrorCode is the structured terminal error code for a failed file, drawn from the same vocabulary as the &#x60;code&#x60; of an error response (e.g. \&quot;corpus_token_quota_exceeded\&quot;), so the same condition reads identically whether it surfaced as an HTTP error or as a status on a file. Two values occur only here -- \&quot;corpus_processing_failed\&quot; and \&quot;upstream_processing_failed\&quot; -- for failures the synchronous path can only report as a generic 500; see corpusStatusOnlyCodes in corpus_errors.go. Surfaced alongside Error so a token-quota failure is distinguishable from a generic one. Omitted when absent. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ProcessedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TokenCount** | Pointer to **int64** | TokenCount is the deepr-reported document token count for the file. Omitted while the count has not been fetched yet or the file failed, and present only once deepr reports one (typically on completion) — so a completed file may still omit it. Distinct from the collection-level token_count on CollectionResponse (the corpus-usage total, a different number). | [optional] 

## Methods

### NewProcessingStatus

`func NewProcessingStatus() *ProcessingStatus`

NewProcessingStatus instantiates a new ProcessingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessingStatusWithDefaults

`func NewProcessingStatusWithDefaults() *ProcessingStatus`

NewProcessingStatusWithDefaults instantiates a new ProcessingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ProcessingStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ProcessingStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ProcessingStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ProcessingStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorCode

`func (o *ProcessingStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ProcessingStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ProcessingStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ProcessingStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMetadata

`func (o *ProcessingStatus) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProcessingStatus) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProcessingStatus) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProcessingStatus) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProcessedAt

`func (o *ProcessingStatus) GetProcessedAt() string`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *ProcessingStatus) GetProcessedAtOk() (*string, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *ProcessingStatus) SetProcessedAt(v string)`

SetProcessedAt sets ProcessedAt field to given value.

### HasProcessedAt

`func (o *ProcessingStatus) HasProcessedAt() bool`

HasProcessedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ProcessingStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProcessingStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProcessingStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProcessingStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTokenCount

`func (o *ProcessingStatus) GetTokenCount() int64`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *ProcessingStatus) GetTokenCountOk() (*int64, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *ProcessingStatus) SetTokenCount(v int64)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *ProcessingStatus) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


