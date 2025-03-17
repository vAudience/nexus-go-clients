# AIgencyMessageContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**AIgencyMessageFile**](AIgencyMessageFile.md) |  | [optional] 
**Text** | **string** |  | 
**Type** | [**AIgencyMessageContentType**](AIgencyMessageContentType.md) |  | 

## Methods

### NewAIgencyMessageContent

`func NewAIgencyMessageContent(text string, type_ AIgencyMessageContentType, ) *AIgencyMessageContent`

NewAIgencyMessageContent instantiates a new AIgencyMessageContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIgencyMessageContentWithDefaults

`func NewAIgencyMessageContentWithDefaults() *AIgencyMessageContent`

NewAIgencyMessageContentWithDefaults instantiates a new AIgencyMessageContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *AIgencyMessageContent) GetFile() AIgencyMessageFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AIgencyMessageContent) GetFileOk() (*AIgencyMessageFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AIgencyMessageContent) SetFile(v AIgencyMessageFile)`

SetFile sets File field to given value.

### HasFile

`func (o *AIgencyMessageContent) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetText

`func (o *AIgencyMessageContent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AIgencyMessageContent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AIgencyMessageContent) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *AIgencyMessageContent) GetType() AIgencyMessageContentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AIgencyMessageContent) GetTypeOk() (*AIgencyMessageContentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AIgencyMessageContent) SetType(v AIgencyMessageContentType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


