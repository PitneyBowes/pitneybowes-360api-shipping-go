# ImportCountStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **float32** | The number of shipments which have been successfully imported. | [optional] 
**Failed** | Pointer to **float32** | The number of shipments failed during Import. | [optional] 
**Pending** | Pointer to **float32** | The number of shipments which are pending and in-queue to be imported. | [optional] 

## Methods

### NewImportCountStatus

`func NewImportCountStatus() *ImportCountStatus`

NewImportCountStatus instantiates a new ImportCountStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportCountStatusWithDefaults

`func NewImportCountStatusWithDefaults() *ImportCountStatus`

NewImportCountStatusWithDefaults instantiates a new ImportCountStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ImportCountStatus) GetSuccess() float32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ImportCountStatus) GetSuccessOk() (*float32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ImportCountStatus) SetSuccess(v float32)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ImportCountStatus) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailed

`func (o *ImportCountStatus) GetFailed() float32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ImportCountStatus) GetFailedOk() (*float32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ImportCountStatus) SetFailed(v float32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ImportCountStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetPending

`func (o *ImportCountStatus) GetPending() float32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ImportCountStatus) GetPendingOk() (*float32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ImportCountStatus) SetPending(v float32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *ImportCountStatus) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


