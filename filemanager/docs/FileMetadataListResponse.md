# FileMetadataListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | [**[]FileMetadataResponse**](FileMetadataResponse.md) |  | 
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewFileMetadataListResponse

`func NewFileMetadataListResponse(files []FileMetadataResponse, limit int32, offset int32, total int32, ) *FileMetadataListResponse`

NewFileMetadataListResponse instantiates a new FileMetadataListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileMetadataListResponseWithDefaults

`func NewFileMetadataListResponseWithDefaults() *FileMetadataListResponse`

NewFileMetadataListResponseWithDefaults instantiates a new FileMetadataListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *FileMetadataListResponse) GetFiles() []FileMetadataResponse`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FileMetadataListResponse) GetFilesOk() (*[]FileMetadataResponse, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FileMetadataListResponse) SetFiles(v []FileMetadataResponse)`

SetFiles sets Files field to given value.


### GetLimit

`func (o *FileMetadataListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FileMetadataListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FileMetadataListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *FileMetadataListResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *FileMetadataListResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *FileMetadataListResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *FileMetadataListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FileMetadataListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FileMetadataListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


