# ShipmentBatchResponseERR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** | This is a system-generated unique identifier assigned to the ERR Batch while it is processed. | [optional] 
**Name** | Pointer to **string** | Name of the of ERR Batch which consists of multiple shipments (shipments in bulk). | [optional] 
**GroupName** | Pointer to **string** |  Indicates the name of the group of batches, which consists of multiple Batch groups. | [optional] 
**Status** | Pointer to [**JobStatus**](JobStatus.md) |  | [optional] 
**UploadURL** | Pointer to **string** | For the stored Batch file under S3, this is the S3 returned URL. The URL is uploaded along with the .CSV file to get the BatchID, which is used to track the import status. | [optional] 

## Methods

### NewShipmentBatchResponseERR

`func NewShipmentBatchResponseERR() *ShipmentBatchResponseERR`

NewShipmentBatchResponseERR instantiates a new ShipmentBatchResponseERR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentBatchResponseERRWithDefaults

`func NewShipmentBatchResponseERRWithDefaults() *ShipmentBatchResponseERR`

NewShipmentBatchResponseERRWithDefaults instantiates a new ShipmentBatchResponseERR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *ShipmentBatchResponseERR) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *ShipmentBatchResponseERR) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *ShipmentBatchResponseERR) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *ShipmentBatchResponseERR) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetName

`func (o *ShipmentBatchResponseERR) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipmentBatchResponseERR) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipmentBatchResponseERR) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipmentBatchResponseERR) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroupName

`func (o *ShipmentBatchResponseERR) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ShipmentBatchResponseERR) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ShipmentBatchResponseERR) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ShipmentBatchResponseERR) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetStatus

`func (o *ShipmentBatchResponseERR) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipmentBatchResponseERR) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipmentBatchResponseERR) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShipmentBatchResponseERR) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUploadURL

`func (o *ShipmentBatchResponseERR) GetUploadURL() string`

GetUploadURL returns the UploadURL field if non-nil, zero value otherwise.

### GetUploadURLOk

`func (o *ShipmentBatchResponseERR) GetUploadURLOk() (*string, bool)`

GetUploadURLOk returns a tuple with the UploadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadURL

`func (o *ShipmentBatchResponseERR) SetUploadURL(v string)`

SetUploadURL sets UploadURL field to given value.

### HasUploadURL

`func (o *ShipmentBatchResponseERR) HasUploadURL() bool`

HasUploadURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


