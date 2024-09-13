# VoidBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Indicates the name of the Batch, i.e., a list of shipments, which need to be cancelled. | 
**Reason** | Pointer to **string** | Reason to cancel the batch of shipments. | [optional] 
**ShipmentIds** | Pointer to **[]string** | &gt;- Shipment ID is a unique identifier for an individual shipment. If ShipmentID(s) are passed in the array, then corresponding shipments will be cancelled. If ShipmentID(s) is/are not provided, then the entire shipments of the Batch will be cancelled. | [optional] 

## Methods

### NewVoidBatchRequest

`func NewVoidBatchRequest(name string, ) *VoidBatchRequest`

NewVoidBatchRequest instantiates a new VoidBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidBatchRequestWithDefaults

`func NewVoidBatchRequestWithDefaults() *VoidBatchRequest`

NewVoidBatchRequestWithDefaults instantiates a new VoidBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VoidBatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoidBatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoidBatchRequest) SetName(v string)`

SetName sets Name field to given value.


### GetReason

`func (o *VoidBatchRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *VoidBatchRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *VoidBatchRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *VoidBatchRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetShipmentIds

`func (o *VoidBatchRequest) GetShipmentIds() []string`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *VoidBatchRequest) GetShipmentIdsOk() (*[]string, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *VoidBatchRequest) SetShipmentIds(v []string)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *VoidBatchRequest) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


