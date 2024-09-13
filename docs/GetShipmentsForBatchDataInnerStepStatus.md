# GetShipmentsForBatchDataInnerStepStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Import** | Pointer to **string** | This is only visible for batch submitted during Import via .CSV. | [optional] 
**AddressValidation** | Pointer to **string** | It indicates the status of shipments at addressValidation stage. | [optional] 
**Rating** | Pointer to **string** | It indicates status of shipments at rating level. | [optional] 
**LabelGeneration** | Pointer to **string** | It indicates status of shipments at labelGeneration step. | [optional] 
**VoidLabel** | Pointer to **string** | It will only be visible when batch labels are cancelled. | [optional] 

## Methods

### NewGetShipmentsForBatchDataInnerStepStatus

`func NewGetShipmentsForBatchDataInnerStepStatus() *GetShipmentsForBatchDataInnerStepStatus`

NewGetShipmentsForBatchDataInnerStepStatus instantiates a new GetShipmentsForBatchDataInnerStepStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetShipmentsForBatchDataInnerStepStatusWithDefaults

`func NewGetShipmentsForBatchDataInnerStepStatusWithDefaults() *GetShipmentsForBatchDataInnerStepStatus`

NewGetShipmentsForBatchDataInnerStepStatusWithDefaults instantiates a new GetShipmentsForBatchDataInnerStepStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImport

`func (o *GetShipmentsForBatchDataInnerStepStatus) GetImport() string`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *GetShipmentsForBatchDataInnerStepStatus) GetImportOk() (*string, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *GetShipmentsForBatchDataInnerStepStatus) SetImport(v string)`

SetImport sets Import field to given value.

### HasImport

`func (o *GetShipmentsForBatchDataInnerStepStatus) HasImport() bool`

HasImport returns a boolean if a field has been set.

### GetAddressValidation

`func (o *GetShipmentsForBatchDataInnerStepStatus) GetAddressValidation() string`

GetAddressValidation returns the AddressValidation field if non-nil, zero value otherwise.

### GetAddressValidationOk

`func (o *GetShipmentsForBatchDataInnerStepStatus) GetAddressValidationOk() (*string, bool)`

GetAddressValidationOk returns a tuple with the AddressValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressValidation

`func (o *GetShipmentsForBatchDataInnerStepStatus) SetAddressValidation(v string)`

SetAddressValidation sets AddressValidation field to given value.

### HasAddressValidation

`func (o *GetShipmentsForBatchDataInnerStepStatus) HasAddressValidation() bool`

HasAddressValidation returns a boolean if a field has been set.

### GetRating

`func (o *GetShipmentsForBatchDataInnerStepStatus) GetRating() string`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GetShipmentsForBatchDataInnerStepStatus) GetRatingOk() (*string, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GetShipmentsForBatchDataInnerStepStatus) SetRating(v string)`

SetRating sets Rating field to given value.

### HasRating

`func (o *GetShipmentsForBatchDataInnerStepStatus) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetLabelGeneration

`func (o *GetShipmentsForBatchDataInnerStepStatus) GetLabelGeneration() string`

GetLabelGeneration returns the LabelGeneration field if non-nil, zero value otherwise.

### GetLabelGenerationOk

`func (o *GetShipmentsForBatchDataInnerStepStatus) GetLabelGenerationOk() (*string, bool)`

GetLabelGenerationOk returns a tuple with the LabelGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelGeneration

`func (o *GetShipmentsForBatchDataInnerStepStatus) SetLabelGeneration(v string)`

SetLabelGeneration sets LabelGeneration field to given value.

### HasLabelGeneration

`func (o *GetShipmentsForBatchDataInnerStepStatus) HasLabelGeneration() bool`

HasLabelGeneration returns a boolean if a field has been set.

### GetVoidLabel

`func (o *GetShipmentsForBatchDataInnerStepStatus) GetVoidLabel() string`

GetVoidLabel returns the VoidLabel field if non-nil, zero value otherwise.

### GetVoidLabelOk

`func (o *GetShipmentsForBatchDataInnerStepStatus) GetVoidLabelOk() (*string, bool)`

GetVoidLabelOk returns a tuple with the VoidLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidLabel

`func (o *GetShipmentsForBatchDataInnerStepStatus) SetVoidLabel(v string)`

SetVoidLabel sets VoidLabel field to given value.

### HasVoidLabel

`func (o *GetShipmentsForBatchDataInnerStepStatus) HasVoidLabel() bool`

HasVoidLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


