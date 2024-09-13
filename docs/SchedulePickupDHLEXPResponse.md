# SchedulePickupDHLEXPResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLocation** | Pointer to **string** | It specifies the location from where packages would be collected. | [optional] 
**PickupConfirmationNumber** | Pointer to **string** | It displays the unique confirmation number of the pickup | [optional] 
**PickupId** | Pointer to **string** | It displays the unique pickup Id which can be further used to get scheduled PDF and cancel pdf if required. | [optional] 
**Carrier** | Pointer to **string** | It dispays the carrier | [optional] 
**CarrierAccountId** | Pointer to **string** | It displays the carrier acount id which is used to create the shipment | [optional] 
**PickupAddress** | Pointer to [**SchedulePickupDHLEXPResponsePickupAddress**](SchedulePickupDHLEXPResponsePickupAddress.md) |  | [optional] 
**ShipmentIds** | Pointer to **[]string** | It indicates the shipmentIds for which pickup is scheduled. | [optional] 
**PickupSummary** | Pointer to [**[]SchedulePickupDHLEXPResponsePickupSummaryInner**](SchedulePickupDHLEXPResponsePickupSummaryInner.md) | It displays the package details provided in the request. | [optional] 
**Additionalnotes** | Pointer to **string** | It displays additional comments or remarks provided in the request, it would be printed on the scheduled pickup document | [optional] 
**Reference** | Pointer to **string** | It displays any reference information provided in the request. | [optional] 
**PickupDateTime** | Pointer to **string** | It displays the scheduled pickup date and time (in UTC) | [optional] 
**PickupTotalWeight** | Pointer to **float32** | It displays the total package weight. | [optional] 
**PickupTotalWeightUnit** | Pointer to **string** | It displays the weight unit. | [optional] 
**PickupOptions** | Pointer to [**SchedulePickupDHLEXPResponsePickupOptions**](SchedulePickupDHLEXPResponsePickupOptions.md) |  | [optional] 

## Methods

### NewSchedulePickupDHLEXPResponse

`func NewSchedulePickupDHLEXPResponse() *SchedulePickupDHLEXPResponse`

NewSchedulePickupDHLEXPResponse instantiates a new SchedulePickupDHLEXPResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupDHLEXPResponseWithDefaults

`func NewSchedulePickupDHLEXPResponseWithDefaults() *SchedulePickupDHLEXPResponse`

NewSchedulePickupDHLEXPResponseWithDefaults instantiates a new SchedulePickupDHLEXPResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupDHLEXPResponse) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupDHLEXPResponse) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupDHLEXPResponse) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *SchedulePickupDHLEXPResponse) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetPickupConfirmationNumber

`func (o *SchedulePickupDHLEXPResponse) GetPickupConfirmationNumber() string`

GetPickupConfirmationNumber returns the PickupConfirmationNumber field if non-nil, zero value otherwise.

### GetPickupConfirmationNumberOk

`func (o *SchedulePickupDHLEXPResponse) GetPickupConfirmationNumberOk() (*string, bool)`

GetPickupConfirmationNumberOk returns a tuple with the PickupConfirmationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupConfirmationNumber

`func (o *SchedulePickupDHLEXPResponse) SetPickupConfirmationNumber(v string)`

SetPickupConfirmationNumber sets PickupConfirmationNumber field to given value.

### HasPickupConfirmationNumber

`func (o *SchedulePickupDHLEXPResponse) HasPickupConfirmationNumber() bool`

HasPickupConfirmationNumber returns a boolean if a field has been set.

### GetPickupId

`func (o *SchedulePickupDHLEXPResponse) GetPickupId() string`

GetPickupId returns the PickupId field if non-nil, zero value otherwise.

### GetPickupIdOk

`func (o *SchedulePickupDHLEXPResponse) GetPickupIdOk() (*string, bool)`

GetPickupIdOk returns a tuple with the PickupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupId

`func (o *SchedulePickupDHLEXPResponse) SetPickupId(v string)`

SetPickupId sets PickupId field to given value.

### HasPickupId

`func (o *SchedulePickupDHLEXPResponse) HasPickupId() bool`

HasPickupId returns a boolean if a field has been set.

### GetCarrier

`func (o *SchedulePickupDHLEXPResponse) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *SchedulePickupDHLEXPResponse) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *SchedulePickupDHLEXPResponse) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *SchedulePickupDHLEXPResponse) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *SchedulePickupDHLEXPResponse) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupDHLEXPResponse) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupDHLEXPResponse) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *SchedulePickupDHLEXPResponse) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *SchedulePickupDHLEXPResponse) GetPickupAddress() SchedulePickupDHLEXPResponsePickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupDHLEXPResponse) GetPickupAddressOk() (*SchedulePickupDHLEXPResponsePickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupDHLEXPResponse) SetPickupAddress(v SchedulePickupDHLEXPResponsePickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *SchedulePickupDHLEXPResponse) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetShipmentIds

`func (o *SchedulePickupDHLEXPResponse) GetShipmentIds() []string`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *SchedulePickupDHLEXPResponse) GetShipmentIdsOk() (*[]string, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *SchedulePickupDHLEXPResponse) SetShipmentIds(v []string)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *SchedulePickupDHLEXPResponse) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupDHLEXPResponse) GetPickupSummary() []SchedulePickupDHLEXPResponsePickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupDHLEXPResponse) GetPickupSummaryOk() (*[]SchedulePickupDHLEXPResponsePickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupDHLEXPResponse) SetPickupSummary(v []SchedulePickupDHLEXPResponsePickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupDHLEXPResponse) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetAdditionalnotes

`func (o *SchedulePickupDHLEXPResponse) GetAdditionalnotes() string`

GetAdditionalnotes returns the Additionalnotes field if non-nil, zero value otherwise.

### GetAdditionalnotesOk

`func (o *SchedulePickupDHLEXPResponse) GetAdditionalnotesOk() (*string, bool)`

GetAdditionalnotesOk returns a tuple with the Additionalnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalnotes

`func (o *SchedulePickupDHLEXPResponse) SetAdditionalnotes(v string)`

SetAdditionalnotes sets Additionalnotes field to given value.

### HasAdditionalnotes

`func (o *SchedulePickupDHLEXPResponse) HasAdditionalnotes() bool`

HasAdditionalnotes returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickupDHLEXPResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickupDHLEXPResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickupDHLEXPResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickupDHLEXPResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPickupDateTime

`func (o *SchedulePickupDHLEXPResponse) GetPickupDateTime() string`

GetPickupDateTime returns the PickupDateTime field if non-nil, zero value otherwise.

### GetPickupDateTimeOk

`func (o *SchedulePickupDHLEXPResponse) GetPickupDateTimeOk() (*string, bool)`

GetPickupDateTimeOk returns a tuple with the PickupDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDateTime

`func (o *SchedulePickupDHLEXPResponse) SetPickupDateTime(v string)`

SetPickupDateTime sets PickupDateTime field to given value.

### HasPickupDateTime

`func (o *SchedulePickupDHLEXPResponse) HasPickupDateTime() bool`

HasPickupDateTime returns a boolean if a field has been set.

### GetPickupTotalWeight

`func (o *SchedulePickupDHLEXPResponse) GetPickupTotalWeight() float32`

GetPickupTotalWeight returns the PickupTotalWeight field if non-nil, zero value otherwise.

### GetPickupTotalWeightOk

`func (o *SchedulePickupDHLEXPResponse) GetPickupTotalWeightOk() (*float32, bool)`

GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeight

`func (o *SchedulePickupDHLEXPResponse) SetPickupTotalWeight(v float32)`

SetPickupTotalWeight sets PickupTotalWeight field to given value.

### HasPickupTotalWeight

`func (o *SchedulePickupDHLEXPResponse) HasPickupTotalWeight() bool`

HasPickupTotalWeight returns a boolean if a field has been set.

### GetPickupTotalWeightUnit

`func (o *SchedulePickupDHLEXPResponse) GetPickupTotalWeightUnit() string`

GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field if non-nil, zero value otherwise.

### GetPickupTotalWeightUnitOk

`func (o *SchedulePickupDHLEXPResponse) GetPickupTotalWeightUnitOk() (*string, bool)`

GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeightUnit

`func (o *SchedulePickupDHLEXPResponse) SetPickupTotalWeightUnit(v string)`

SetPickupTotalWeightUnit sets PickupTotalWeightUnit field to given value.

### HasPickupTotalWeightUnit

`func (o *SchedulePickupDHLEXPResponse) HasPickupTotalWeightUnit() bool`

HasPickupTotalWeightUnit returns a boolean if a field has been set.

### GetPickupOptions

`func (o *SchedulePickupDHLEXPResponse) GetPickupOptions() SchedulePickupDHLEXPResponsePickupOptions`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *SchedulePickupDHLEXPResponse) GetPickupOptionsOk() (*SchedulePickupDHLEXPResponsePickupOptions, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *SchedulePickupDHLEXPResponse) SetPickupOptions(v SchedulePickupDHLEXPResponsePickupOptions)`

SetPickupOptions sets PickupOptions field to given value.

### HasPickupOptions

`func (o *SchedulePickupDHLEXPResponse) HasPickupOptions() bool`

HasPickupOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


