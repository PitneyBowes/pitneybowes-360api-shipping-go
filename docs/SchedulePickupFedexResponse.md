# SchedulePickupFedexResponse

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
**PickupDateTime** | Pointer to **string** | It displays the scheduled pickup date and time (in UTC) | [optional] 
**PickupTotalWeight** | Pointer to **float32** | It displays the total package weight. | [optional] 
**PickupTotalWeightUnit** | Pointer to **string** | It displays the weight unit. | [optional] 
**PickupOptions** | Pointer to [**GetAllPickupsDataInnerPickupOptions**](GetAllPickupsDataInnerPickupOptions.md) |  | [optional] 

## Methods

### NewSchedulePickupFedexResponse

`func NewSchedulePickupFedexResponse() *SchedulePickupFedexResponse`

NewSchedulePickupFedexResponse instantiates a new SchedulePickupFedexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupFedexResponseWithDefaults

`func NewSchedulePickupFedexResponseWithDefaults() *SchedulePickupFedexResponse`

NewSchedulePickupFedexResponseWithDefaults instantiates a new SchedulePickupFedexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupFedexResponse) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupFedexResponse) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupFedexResponse) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *SchedulePickupFedexResponse) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetPickupConfirmationNumber

`func (o *SchedulePickupFedexResponse) GetPickupConfirmationNumber() string`

GetPickupConfirmationNumber returns the PickupConfirmationNumber field if non-nil, zero value otherwise.

### GetPickupConfirmationNumberOk

`func (o *SchedulePickupFedexResponse) GetPickupConfirmationNumberOk() (*string, bool)`

GetPickupConfirmationNumberOk returns a tuple with the PickupConfirmationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupConfirmationNumber

`func (o *SchedulePickupFedexResponse) SetPickupConfirmationNumber(v string)`

SetPickupConfirmationNumber sets PickupConfirmationNumber field to given value.

### HasPickupConfirmationNumber

`func (o *SchedulePickupFedexResponse) HasPickupConfirmationNumber() bool`

HasPickupConfirmationNumber returns a boolean if a field has been set.

### GetPickupId

`func (o *SchedulePickupFedexResponse) GetPickupId() string`

GetPickupId returns the PickupId field if non-nil, zero value otherwise.

### GetPickupIdOk

`func (o *SchedulePickupFedexResponse) GetPickupIdOk() (*string, bool)`

GetPickupIdOk returns a tuple with the PickupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupId

`func (o *SchedulePickupFedexResponse) SetPickupId(v string)`

SetPickupId sets PickupId field to given value.

### HasPickupId

`func (o *SchedulePickupFedexResponse) HasPickupId() bool`

HasPickupId returns a boolean if a field has been set.

### GetCarrier

`func (o *SchedulePickupFedexResponse) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *SchedulePickupFedexResponse) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *SchedulePickupFedexResponse) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *SchedulePickupFedexResponse) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *SchedulePickupFedexResponse) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupFedexResponse) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupFedexResponse) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *SchedulePickupFedexResponse) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *SchedulePickupFedexResponse) GetPickupAddress() SchedulePickupDHLEXPResponsePickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupFedexResponse) GetPickupAddressOk() (*SchedulePickupDHLEXPResponsePickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupFedexResponse) SetPickupAddress(v SchedulePickupDHLEXPResponsePickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *SchedulePickupFedexResponse) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetShipmentIds

`func (o *SchedulePickupFedexResponse) GetShipmentIds() []string`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *SchedulePickupFedexResponse) GetShipmentIdsOk() (*[]string, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *SchedulePickupFedexResponse) SetShipmentIds(v []string)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *SchedulePickupFedexResponse) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupFedexResponse) GetPickupSummary() []SchedulePickupUSPSResponsePickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupFedexResponse) GetPickupSummaryOk() (*[]SchedulePickupUSPSResponsePickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupFedexResponse) SetPickupSummary(v []SchedulePickupUSPSResponsePickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupFedexResponse) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetAdditionalnotes

`func (o *SchedulePickupFedexResponse) GetAdditionalnotes() string`

GetAdditionalnotes returns the Additionalnotes field if non-nil, zero value otherwise.

### GetAdditionalnotesOk

`func (o *SchedulePickupFedexResponse) GetAdditionalnotesOk() (*string, bool)`

GetAdditionalnotesOk returns a tuple with the Additionalnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalnotes

`func (o *SchedulePickupFedexResponse) SetAdditionalnotes(v string)`

SetAdditionalnotes sets Additionalnotes field to given value.

### HasAdditionalnotes

`func (o *SchedulePickupFedexResponse) HasAdditionalnotes() bool`

HasAdditionalnotes returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickupFedexResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickupFedexResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickupFedexResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickupFedexResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPickupDateTime

`func (o *SchedulePickupFedexResponse) GetPickupDateTime() string`

GetPickupDateTime returns the PickupDateTime field if non-nil, zero value otherwise.

### GetPickupDateTimeOk

`func (o *SchedulePickupFedexResponse) GetPickupDateTimeOk() (*string, bool)`

GetPickupDateTimeOk returns a tuple with the PickupDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDateTime

`func (o *SchedulePickupFedexResponse) SetPickupDateTime(v string)`

SetPickupDateTime sets PickupDateTime field to given value.

### HasPickupDateTime

`func (o *SchedulePickupFedexResponse) HasPickupDateTime() bool`

HasPickupDateTime returns a boolean if a field has been set.

### GetPickupTotalWeight

`func (o *SchedulePickupFedexResponse) GetPickupTotalWeight() float32`

GetPickupTotalWeight returns the PickupTotalWeight field if non-nil, zero value otherwise.

### GetPickupTotalWeightOk

`func (o *SchedulePickupFedexResponse) GetPickupTotalWeightOk() (*float32, bool)`

GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeight

`func (o *SchedulePickupFedexResponse) SetPickupTotalWeight(v float32)`

SetPickupTotalWeight sets PickupTotalWeight field to given value.

### HasPickupTotalWeight

`func (o *SchedulePickupFedexResponse) HasPickupTotalWeight() bool`

HasPickupTotalWeight returns a boolean if a field has been set.

### GetPickupTotalWeightUnit

`func (o *SchedulePickupFedexResponse) GetPickupTotalWeightUnit() string`

GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field if non-nil, zero value otherwise.

### GetPickupTotalWeightUnitOk

`func (o *SchedulePickupFedexResponse) GetPickupTotalWeightUnitOk() (*string, bool)`

GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeightUnit

`func (o *SchedulePickupFedexResponse) SetPickupTotalWeightUnit(v string)`

SetPickupTotalWeightUnit sets PickupTotalWeightUnit field to given value.

### HasPickupTotalWeightUnit

`func (o *SchedulePickupFedexResponse) HasPickupTotalWeightUnit() bool`

HasPickupTotalWeightUnit returns a boolean if a field has been set.

### GetPickupOptions

`func (o *SchedulePickupFedexResponse) GetPickupOptions() GetAllPickupsDataInnerPickupOptions`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *SchedulePickupFedexResponse) GetPickupOptionsOk() (*GetAllPickupsDataInnerPickupOptions, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *SchedulePickupFedexResponse) SetPickupOptions(v GetAllPickupsDataInnerPickupOptions)`

SetPickupOptions sets PickupOptions field to given value.

### HasPickupOptions

`func (o *SchedulePickupFedexResponse) HasPickupOptions() bool`

HasPickupOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


