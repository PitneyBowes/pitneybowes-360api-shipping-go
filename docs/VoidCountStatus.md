# VoidCountStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **float32** | The number of shipment labels which have been successfully cancelled. | [optional] 
**Failed** | Pointer to **float32** | The number of shipment labels which failed cancelling due to some validation issue. | [optional] 
**Pending** | Pointer to **float32** | The number of shipment labels which are pending and in-queue to be cancelled. | [optional] 

## Methods

### NewVoidCountStatus

`func NewVoidCountStatus() *VoidCountStatus`

NewVoidCountStatus instantiates a new VoidCountStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidCountStatusWithDefaults

`func NewVoidCountStatusWithDefaults() *VoidCountStatus`

NewVoidCountStatusWithDefaults instantiates a new VoidCountStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *VoidCountStatus) GetSuccess() float32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *VoidCountStatus) GetSuccessOk() (*float32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *VoidCountStatus) SetSuccess(v float32)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *VoidCountStatus) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailed

`func (o *VoidCountStatus) GetFailed() float32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *VoidCountStatus) GetFailedOk() (*float32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *VoidCountStatus) SetFailed(v float32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *VoidCountStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetPending

`func (o *VoidCountStatus) GetPending() float32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *VoidCountStatus) GetPendingOk() (*float32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *VoidCountStatus) SetPending(v float32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *VoidCountStatus) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


