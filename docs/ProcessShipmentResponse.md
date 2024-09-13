# ProcessShipmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** |  This is a system-generated unique identifier assigned to the Batch while it is processed. | [optional] 
**Name** | Pointer to **string** | Name of the of Batch which consists of multiple shipments (shipments in bulk). | [optional] 
**GroupName** | Pointer to **string** | Indicates the name of the group of batches, which consists of multiple Batch groups. | [optional] 
**Status** | Pointer to **string** |  The batch will show the status &#x60;SUBMITTED&#x60; on successful execution of Batch. | [optional] 

## Methods

### NewProcessShipmentResponse

`func NewProcessShipmentResponse() *ProcessShipmentResponse`

NewProcessShipmentResponse instantiates a new ProcessShipmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessShipmentResponseWithDefaults

`func NewProcessShipmentResponseWithDefaults() *ProcessShipmentResponse`

NewProcessShipmentResponseWithDefaults instantiates a new ProcessShipmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *ProcessShipmentResponse) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *ProcessShipmentResponse) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *ProcessShipmentResponse) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *ProcessShipmentResponse) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetName

`func (o *ProcessShipmentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessShipmentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessShipmentResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProcessShipmentResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroupName

`func (o *ProcessShipmentResponse) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ProcessShipmentResponse) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ProcessShipmentResponse) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ProcessShipmentResponse) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetStatus

`func (o *ProcessShipmentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProcessShipmentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProcessShipmentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProcessShipmentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


