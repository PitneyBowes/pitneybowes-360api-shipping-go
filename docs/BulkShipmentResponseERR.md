# BulkShipmentResponseERR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** |  This is a system-generated unique identifier assigned to the Batch while it is processed. | [optional] 
**Name** | Pointer to **string** |  Name of the of ERR Batch which consists of multiple shipments (shipments in bulk). | [optional] 
**GroupName** | Pointer to **string** | Indicates the name of the group of batches, which consists of multiple Batch groups. | [optional] 
**Status** | Pointer to [**JobStatus**](JobStatus.md) |  | [optional] 
**AddressValidation** | Pointer to [**AddressCountStatus**](AddressCountStatus.md) |  | [optional] 
**Rating** | Pointer to [**RatingCountStatusERR**](RatingCountStatusERR.md) |  | [optional] 
**LabelGeneration** | Pointer to [**LabelGenerationCountStatus**](LabelGenerationCountStatus.md) |  | [optional] 
**LabelDetails** | Pointer to [**BulkShipmentResponseERRLabelDetails**](BulkShipmentResponseERRLabelDetails.md) |  | [optional] 

## Methods

### NewBulkShipmentResponseERR

`func NewBulkShipmentResponseERR() *BulkShipmentResponseERR`

NewBulkShipmentResponseERR instantiates a new BulkShipmentResponseERR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkShipmentResponseERRWithDefaults

`func NewBulkShipmentResponseERRWithDefaults() *BulkShipmentResponseERR`

NewBulkShipmentResponseERRWithDefaults instantiates a new BulkShipmentResponseERR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *BulkShipmentResponseERR) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *BulkShipmentResponseERR) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *BulkShipmentResponseERR) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *BulkShipmentResponseERR) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetName

`func (o *BulkShipmentResponseERR) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkShipmentResponseERR) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkShipmentResponseERR) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkShipmentResponseERR) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroupName

`func (o *BulkShipmentResponseERR) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *BulkShipmentResponseERR) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *BulkShipmentResponseERR) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *BulkShipmentResponseERR) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetStatus

`func (o *BulkShipmentResponseERR) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkShipmentResponseERR) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkShipmentResponseERR) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkShipmentResponseERR) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAddressValidation

`func (o *BulkShipmentResponseERR) GetAddressValidation() AddressCountStatus`

GetAddressValidation returns the AddressValidation field if non-nil, zero value otherwise.

### GetAddressValidationOk

`func (o *BulkShipmentResponseERR) GetAddressValidationOk() (*AddressCountStatus, bool)`

GetAddressValidationOk returns a tuple with the AddressValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressValidation

`func (o *BulkShipmentResponseERR) SetAddressValidation(v AddressCountStatus)`

SetAddressValidation sets AddressValidation field to given value.

### HasAddressValidation

`func (o *BulkShipmentResponseERR) HasAddressValidation() bool`

HasAddressValidation returns a boolean if a field has been set.

### GetRating

`func (o *BulkShipmentResponseERR) GetRating() RatingCountStatusERR`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *BulkShipmentResponseERR) GetRatingOk() (*RatingCountStatusERR, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *BulkShipmentResponseERR) SetRating(v RatingCountStatusERR)`

SetRating sets Rating field to given value.

### HasRating

`func (o *BulkShipmentResponseERR) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetLabelGeneration

`func (o *BulkShipmentResponseERR) GetLabelGeneration() LabelGenerationCountStatus`

GetLabelGeneration returns the LabelGeneration field if non-nil, zero value otherwise.

### GetLabelGenerationOk

`func (o *BulkShipmentResponseERR) GetLabelGenerationOk() (*LabelGenerationCountStatus, bool)`

GetLabelGenerationOk returns a tuple with the LabelGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelGeneration

`func (o *BulkShipmentResponseERR) SetLabelGeneration(v LabelGenerationCountStatus)`

SetLabelGeneration sets LabelGeneration field to given value.

### HasLabelGeneration

`func (o *BulkShipmentResponseERR) HasLabelGeneration() bool`

HasLabelGeneration returns a boolean if a field has been set.

### GetLabelDetails

`func (o *BulkShipmentResponseERR) GetLabelDetails() BulkShipmentResponseERRLabelDetails`

GetLabelDetails returns the LabelDetails field if non-nil, zero value otherwise.

### GetLabelDetailsOk

`func (o *BulkShipmentResponseERR) GetLabelDetailsOk() (*BulkShipmentResponseERRLabelDetails, bool)`

GetLabelDetailsOk returns a tuple with the LabelDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelDetails

`func (o *BulkShipmentResponseERR) SetLabelDetails(v BulkShipmentResponseERRLabelDetails)`

SetLabelDetails sets LabelDetails field to given value.

### HasLabelDetails

`func (o *BulkShipmentResponseERR) HasLabelDetails() bool`

HasLabelDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


