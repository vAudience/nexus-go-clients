# FileUploadCategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedMimeTypes** | Pointer to **[]string** | MIME types a file uploaded to this category may have. The type is detected from the file&#39;s content, and from its extension only where content detection is inconclusive (e.g. .docx, .md); the request&#39;s Content-Type is not used. Any other type is rejected with 400 invalid_file_upload_mime_type. | [optional] 
**Category** | Pointer to **string** | Name of the category, used as the {category} path segment of POST /v1/organizations/{org_id}/files/{category}. | [optional] 
**LlmInputType** | Pointer to **string** | How a chat passes a file of this category to the model: reference (the message carries only a link to the stored file), embedded (the file is converted to markdown and its content embedded in the message), or multimodal (the file is sent to the model as native multimodal input, such as an image or audio). | [optional] 
**MaxFileSize** | Pointer to **int64** | Largest accepted file size in bytes, inclusive; a larger file is rejected with 413 invalid_file_upload_size. | [optional] 
**MinFileSize** | Pointer to **int64** | Smallest accepted file size in bytes, inclusive; a smaller file is rejected with 413 invalid_file_upload_size. | [optional] 

## Methods

### NewFileUploadCategoryResponse

`func NewFileUploadCategoryResponse() *FileUploadCategoryResponse`

NewFileUploadCategoryResponse instantiates a new FileUploadCategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUploadCategoryResponseWithDefaults

`func NewFileUploadCategoryResponseWithDefaults() *FileUploadCategoryResponse`

NewFileUploadCategoryResponseWithDefaults instantiates a new FileUploadCategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedMimeTypes

`func (o *FileUploadCategoryResponse) GetAcceptedMimeTypes() []string`

GetAcceptedMimeTypes returns the AcceptedMimeTypes field if non-nil, zero value otherwise.

### GetAcceptedMimeTypesOk

`func (o *FileUploadCategoryResponse) GetAcceptedMimeTypesOk() (*[]string, bool)`

GetAcceptedMimeTypesOk returns a tuple with the AcceptedMimeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedMimeTypes

`func (o *FileUploadCategoryResponse) SetAcceptedMimeTypes(v []string)`

SetAcceptedMimeTypes sets AcceptedMimeTypes field to given value.

### HasAcceptedMimeTypes

`func (o *FileUploadCategoryResponse) HasAcceptedMimeTypes() bool`

HasAcceptedMimeTypes returns a boolean if a field has been set.

### GetCategory

`func (o *FileUploadCategoryResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FileUploadCategoryResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FileUploadCategoryResponse) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FileUploadCategoryResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLlmInputType

`func (o *FileUploadCategoryResponse) GetLlmInputType() string`

GetLlmInputType returns the LlmInputType field if non-nil, zero value otherwise.

### GetLlmInputTypeOk

`func (o *FileUploadCategoryResponse) GetLlmInputTypeOk() (*string, bool)`

GetLlmInputTypeOk returns a tuple with the LlmInputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmInputType

`func (o *FileUploadCategoryResponse) SetLlmInputType(v string)`

SetLlmInputType sets LlmInputType field to given value.

### HasLlmInputType

`func (o *FileUploadCategoryResponse) HasLlmInputType() bool`

HasLlmInputType returns a boolean if a field has been set.

### GetMaxFileSize

`func (o *FileUploadCategoryResponse) GetMaxFileSize() int64`

GetMaxFileSize returns the MaxFileSize field if non-nil, zero value otherwise.

### GetMaxFileSizeOk

`func (o *FileUploadCategoryResponse) GetMaxFileSizeOk() (*int64, bool)`

GetMaxFileSizeOk returns a tuple with the MaxFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSize

`func (o *FileUploadCategoryResponse) SetMaxFileSize(v int64)`

SetMaxFileSize sets MaxFileSize field to given value.

### HasMaxFileSize

`func (o *FileUploadCategoryResponse) HasMaxFileSize() bool`

HasMaxFileSize returns a boolean if a field has been set.

### GetMinFileSize

`func (o *FileUploadCategoryResponse) GetMinFileSize() int64`

GetMinFileSize returns the MinFileSize field if non-nil, zero value otherwise.

### GetMinFileSizeOk

`func (o *FileUploadCategoryResponse) GetMinFileSizeOk() (*int64, bool)`

GetMinFileSizeOk returns a tuple with the MinFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFileSize

`func (o *FileUploadCategoryResponse) SetMinFileSize(v int64)`

SetMinFileSize sets MinFileSize field to given value.

### HasMinFileSize

`func (o *FileUploadCategoryResponse) HasMinFileSize() bool`

HasMinFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


