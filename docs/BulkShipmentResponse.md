# BulkShipmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** |  This is the system generated Batch ID. | [optional] 
**GroupName** | Pointer to **string** | This indicates group name for the Batch. | [optional] 
**Name** | Pointer to **string** |  The name of the Batch. | [optional] 
**Status** | Pointer to [**JobStatus**](JobStatus.md) |  | [optional] 
**AddressValidation** | Pointer to [**CounterStatus**](CounterStatus.md) |  | [optional] 
**Rating** | Pointer to [**CounterStatus**](CounterStatus.md) |  | [optional] 
**LabelGeneration** | Pointer to [**CounterStatus**](CounterStatus.md) |  | [optional] 

## Methods

### NewBulkShipmentResponse

`func NewBulkShipmentResponse() *BulkShipmentResponse`

NewBulkShipmentResponse instantiates a new BulkShipmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkShipmentResponseWithDefaults

`func NewBulkShipmentResponseWithDefaults() *BulkShipmentResponse`

NewBulkShipmentResponseWithDefaults instantiates a new BulkShipmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *BulkShipmentResponse) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *BulkShipmentResponse) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *BulkShipmentResponse) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *BulkShipmentResponse) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetGroupName

`func (o *BulkShipmentResponse) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *BulkShipmentResponse) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *BulkShipmentResponse) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *BulkShipmentResponse) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetName

`func (o *BulkShipmentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkShipmentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkShipmentResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkShipmentResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *BulkShipmentResponse) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkShipmentResponse) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkShipmentResponse) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkShipmentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAddressValidation

`func (o *BulkShipmentResponse) GetAddressValidation() CounterStatus`

GetAddressValidation returns the AddressValidation field if non-nil, zero value otherwise.

### GetAddressValidationOk

`func (o *BulkShipmentResponse) GetAddressValidationOk() (*CounterStatus, bool)`

GetAddressValidationOk returns a tuple with the AddressValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressValidation

`func (o *BulkShipmentResponse) SetAddressValidation(v CounterStatus)`

SetAddressValidation sets AddressValidation field to given value.

### HasAddressValidation

`func (o *BulkShipmentResponse) HasAddressValidation() bool`

HasAddressValidation returns a boolean if a field has been set.

### GetRating

`func (o *BulkShipmentResponse) GetRating() CounterStatus`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *BulkShipmentResponse) GetRatingOk() (*CounterStatus, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *BulkShipmentResponse) SetRating(v CounterStatus)`

SetRating sets Rating field to given value.

### HasRating

`func (o *BulkShipmentResponse) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetLabelGeneration

`func (o *BulkShipmentResponse) GetLabelGeneration() CounterStatus`

GetLabelGeneration returns the LabelGeneration field if non-nil, zero value otherwise.

### GetLabelGenerationOk

`func (o *BulkShipmentResponse) GetLabelGenerationOk() (*CounterStatus, bool)`

GetLabelGenerationOk returns a tuple with the LabelGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelGeneration

`func (o *BulkShipmentResponse) SetLabelGeneration(v CounterStatus)`

SetLabelGeneration sets LabelGeneration field to given value.

### HasLabelGeneration

`func (o *BulkShipmentResponse) HasLabelGeneration() bool`

HasLabelGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


