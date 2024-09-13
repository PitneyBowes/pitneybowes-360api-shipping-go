# ShipmentBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** | This is the system generated Batch ID. | [optional] 
**Name** | Pointer to **string** | Name of the Batch. | [optional] 
**GroupName** | Pointer to **string** |  The group name of the Batch. | [optional] 
**Status** | Pointer to [**JobStatus**](JobStatus.md) |  | [optional] 
**UploadURL** | Pointer to **string** | If this is an import batch, this is the URL to upload the .CSV file. | [optional] 

## Methods

### NewShipmentBatch

`func NewShipmentBatch() *ShipmentBatch`

NewShipmentBatch instantiates a new ShipmentBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentBatchWithDefaults

`func NewShipmentBatchWithDefaults() *ShipmentBatch`

NewShipmentBatchWithDefaults instantiates a new ShipmentBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *ShipmentBatch) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *ShipmentBatch) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *ShipmentBatch) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *ShipmentBatch) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetName

`func (o *ShipmentBatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipmentBatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipmentBatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipmentBatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroupName

`func (o *ShipmentBatch) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ShipmentBatch) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ShipmentBatch) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ShipmentBatch) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetStatus

`func (o *ShipmentBatch) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipmentBatch) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipmentBatch) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShipmentBatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUploadURL

`func (o *ShipmentBatch) GetUploadURL() string`

GetUploadURL returns the UploadURL field if non-nil, zero value otherwise.

### GetUploadURLOk

`func (o *ShipmentBatch) GetUploadURLOk() (*string, bool)`

GetUploadURLOk returns a tuple with the UploadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadURL

`func (o *ShipmentBatch) SetUploadURL(v string)`

SetUploadURL sets UploadURL field to given value.

### HasUploadURL

`func (o *ShipmentBatch) HasUploadURL() bool`

HasUploadURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


