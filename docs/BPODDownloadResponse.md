# BPODDownloadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileURL** | Pointer to **string** | This is the URL for the BPOD (Bulk Proof of Delivery) ZIP file. | [optional] 

## Methods

### NewBPODDownloadResponse

`func NewBPODDownloadResponse() *BPODDownloadResponse`

NewBPODDownloadResponse instantiates a new BPODDownloadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBPODDownloadResponseWithDefaults

`func NewBPODDownloadResponseWithDefaults() *BPODDownloadResponse`

NewBPODDownloadResponseWithDefaults instantiates a new BPODDownloadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileURL

`func (o *BPODDownloadResponse) GetFileURL() string`

GetFileURL returns the FileURL field if non-nil, zero value otherwise.

### GetFileURLOk

`func (o *BPODDownloadResponse) GetFileURLOk() (*string, bool)`

GetFileURLOk returns a tuple with the FileURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileURL

`func (o *BPODDownloadResponse) SetFileURL(v string)`

SetFileURL sets FileURL field to given value.

### HasFileURL

`func (o *BPODDownloadResponse) HasFileURL() bool`

HasFileURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


