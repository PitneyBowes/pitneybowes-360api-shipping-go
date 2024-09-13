# SchedulePickupUSPSResponse

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
**PickupSummary** | Pointer to [**[]SchedulePickupUSPSResponsePickupSummaryInner**](SchedulePickupUSPSResponsePickupSummaryInner.md) | It displays the package details provided in the request. | [optional] 
**Additionalnotes** | Pointer to **string** | It displays additional comments or remarks provided in the request, it would be printed on the scheduled pickup document | [optional] 
**Reference** | Pointer to **string** | It displays any reference information provided in the request. | [optional] 
**PickupDateTime** | Pointer to **string** | It displays the scheduled pickup date | [optional] 
**PickupTotalWeight** | Pointer to **float32** | It displays the total package weight. | [optional] 
**PickupTotalWeightUnit** | Pointer to **string** | It displays the weight unit. | [optional] 

## Methods

### NewSchedulePickupUSPSResponse

`func NewSchedulePickupUSPSResponse() *SchedulePickupUSPSResponse`

NewSchedulePickupUSPSResponse instantiates a new SchedulePickupUSPSResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUSPSResponseWithDefaults

`func NewSchedulePickupUSPSResponseWithDefaults() *SchedulePickupUSPSResponse`

NewSchedulePickupUSPSResponseWithDefaults instantiates a new SchedulePickupUSPSResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupUSPSResponse) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupUSPSResponse) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupUSPSResponse) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *SchedulePickupUSPSResponse) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetPickupConfirmationNumber

`func (o *SchedulePickupUSPSResponse) GetPickupConfirmationNumber() string`

GetPickupConfirmationNumber returns the PickupConfirmationNumber field if non-nil, zero value otherwise.

### GetPickupConfirmationNumberOk

`func (o *SchedulePickupUSPSResponse) GetPickupConfirmationNumberOk() (*string, bool)`

GetPickupConfirmationNumberOk returns a tuple with the PickupConfirmationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupConfirmationNumber

`func (o *SchedulePickupUSPSResponse) SetPickupConfirmationNumber(v string)`

SetPickupConfirmationNumber sets PickupConfirmationNumber field to given value.

### HasPickupConfirmationNumber

`func (o *SchedulePickupUSPSResponse) HasPickupConfirmationNumber() bool`

HasPickupConfirmationNumber returns a boolean if a field has been set.

### GetPickupId

`func (o *SchedulePickupUSPSResponse) GetPickupId() string`

GetPickupId returns the PickupId field if non-nil, zero value otherwise.

### GetPickupIdOk

`func (o *SchedulePickupUSPSResponse) GetPickupIdOk() (*string, bool)`

GetPickupIdOk returns a tuple with the PickupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupId

`func (o *SchedulePickupUSPSResponse) SetPickupId(v string)`

SetPickupId sets PickupId field to given value.

### HasPickupId

`func (o *SchedulePickupUSPSResponse) HasPickupId() bool`

HasPickupId returns a boolean if a field has been set.

### GetCarrier

`func (o *SchedulePickupUSPSResponse) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *SchedulePickupUSPSResponse) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *SchedulePickupUSPSResponse) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *SchedulePickupUSPSResponse) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *SchedulePickupUSPSResponse) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupUSPSResponse) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupUSPSResponse) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *SchedulePickupUSPSResponse) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *SchedulePickupUSPSResponse) GetPickupAddress() SchedulePickupDHLEXPResponsePickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupUSPSResponse) GetPickupAddressOk() (*SchedulePickupDHLEXPResponsePickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupUSPSResponse) SetPickupAddress(v SchedulePickupDHLEXPResponsePickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *SchedulePickupUSPSResponse) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetShipmentIds

`func (o *SchedulePickupUSPSResponse) GetShipmentIds() []string`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *SchedulePickupUSPSResponse) GetShipmentIdsOk() (*[]string, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *SchedulePickupUSPSResponse) SetShipmentIds(v []string)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *SchedulePickupUSPSResponse) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupUSPSResponse) GetPickupSummary() []SchedulePickupUSPSResponsePickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupUSPSResponse) GetPickupSummaryOk() (*[]SchedulePickupUSPSResponsePickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupUSPSResponse) SetPickupSummary(v []SchedulePickupUSPSResponsePickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupUSPSResponse) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetAdditionalnotes

`func (o *SchedulePickupUSPSResponse) GetAdditionalnotes() string`

GetAdditionalnotes returns the Additionalnotes field if non-nil, zero value otherwise.

### GetAdditionalnotesOk

`func (o *SchedulePickupUSPSResponse) GetAdditionalnotesOk() (*string, bool)`

GetAdditionalnotesOk returns a tuple with the Additionalnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalnotes

`func (o *SchedulePickupUSPSResponse) SetAdditionalnotes(v string)`

SetAdditionalnotes sets Additionalnotes field to given value.

### HasAdditionalnotes

`func (o *SchedulePickupUSPSResponse) HasAdditionalnotes() bool`

HasAdditionalnotes returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickupUSPSResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickupUSPSResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickupUSPSResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickupUSPSResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPickupDateTime

`func (o *SchedulePickupUSPSResponse) GetPickupDateTime() string`

GetPickupDateTime returns the PickupDateTime field if non-nil, zero value otherwise.

### GetPickupDateTimeOk

`func (o *SchedulePickupUSPSResponse) GetPickupDateTimeOk() (*string, bool)`

GetPickupDateTimeOk returns a tuple with the PickupDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDateTime

`func (o *SchedulePickupUSPSResponse) SetPickupDateTime(v string)`

SetPickupDateTime sets PickupDateTime field to given value.

### HasPickupDateTime

`func (o *SchedulePickupUSPSResponse) HasPickupDateTime() bool`

HasPickupDateTime returns a boolean if a field has been set.

### GetPickupTotalWeight

`func (o *SchedulePickupUSPSResponse) GetPickupTotalWeight() float32`

GetPickupTotalWeight returns the PickupTotalWeight field if non-nil, zero value otherwise.

### GetPickupTotalWeightOk

`func (o *SchedulePickupUSPSResponse) GetPickupTotalWeightOk() (*float32, bool)`

GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeight

`func (o *SchedulePickupUSPSResponse) SetPickupTotalWeight(v float32)`

SetPickupTotalWeight sets PickupTotalWeight field to given value.

### HasPickupTotalWeight

`func (o *SchedulePickupUSPSResponse) HasPickupTotalWeight() bool`

HasPickupTotalWeight returns a boolean if a field has been set.

### GetPickupTotalWeightUnit

`func (o *SchedulePickupUSPSResponse) GetPickupTotalWeightUnit() string`

GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field if non-nil, zero value otherwise.

### GetPickupTotalWeightUnitOk

`func (o *SchedulePickupUSPSResponse) GetPickupTotalWeightUnitOk() (*string, bool)`

GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeightUnit

`func (o *SchedulePickupUSPSResponse) SetPickupTotalWeightUnit(v string)`

SetPickupTotalWeightUnit sets PickupTotalWeightUnit field to given value.

### HasPickupTotalWeightUnit

`func (o *SchedulePickupUSPSResponse) HasPickupTotalWeightUnit() bool`

HasPickupTotalWeightUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


