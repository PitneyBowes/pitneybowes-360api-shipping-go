# VoidBatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchID** | Pointer to **string** | This is a system-generated unique identifier assigned to the Batch while it is processed. | [optional] 
**Status** | Pointer to **string** | Indicates the status of the Batch while executing &#x60;voidShippingLabel&#x60;. | [optional] 

## Methods

### NewVoidBatchResponse

`func NewVoidBatchResponse() *VoidBatchResponse`

NewVoidBatchResponse instantiates a new VoidBatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidBatchResponseWithDefaults

`func NewVoidBatchResponseWithDefaults() *VoidBatchResponse`

NewVoidBatchResponseWithDefaults instantiates a new VoidBatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchID

`func (o *VoidBatchResponse) GetBatchID() string`

GetBatchID returns the BatchID field if non-nil, zero value otherwise.

### GetBatchIDOk

`func (o *VoidBatchResponse) GetBatchIDOk() (*string, bool)`

GetBatchIDOk returns a tuple with the BatchID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchID

`func (o *VoidBatchResponse) SetBatchID(v string)`

SetBatchID sets BatchID field to given value.

### HasBatchID

`func (o *VoidBatchResponse) HasBatchID() bool`

HasBatchID returns a boolean if a field has been set.

### GetStatus

`func (o *VoidBatchResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VoidBatchResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VoidBatchResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VoidBatchResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


