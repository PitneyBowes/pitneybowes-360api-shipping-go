# LabelGenerationCountStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **float32** | The number of shipment labels which have been successfully generated. | [optional] 
**Failed** | Pointer to **float32** | The number of shipment labels which failed processing and could not be generated due to some errors with shipment. | [optional] 
**Pending** | Pointer to **float32** | The number of shipment labels which are pending and in-queue to be processed. | [optional] 

## Methods

### NewLabelGenerationCountStatus

`func NewLabelGenerationCountStatus() *LabelGenerationCountStatus`

NewLabelGenerationCountStatus instantiates a new LabelGenerationCountStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelGenerationCountStatusWithDefaults

`func NewLabelGenerationCountStatusWithDefaults() *LabelGenerationCountStatus`

NewLabelGenerationCountStatusWithDefaults instantiates a new LabelGenerationCountStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *LabelGenerationCountStatus) GetSuccess() float32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *LabelGenerationCountStatus) GetSuccessOk() (*float32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *LabelGenerationCountStatus) SetSuccess(v float32)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *LabelGenerationCountStatus) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailed

`func (o *LabelGenerationCountStatus) GetFailed() float32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *LabelGenerationCountStatus) GetFailedOk() (*float32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *LabelGenerationCountStatus) SetFailed(v float32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *LabelGenerationCountStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetPending

`func (o *LabelGenerationCountStatus) GetPending() float32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *LabelGenerationCountStatus) GetPendingOk() (*float32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *LabelGenerationCountStatus) SetPending(v float32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *LabelGenerationCountStatus) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


