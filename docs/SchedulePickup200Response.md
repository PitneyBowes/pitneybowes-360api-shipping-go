# SchedulePickup200Response

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
**PickupRate** | Pointer to [**SchedulePickupUPSResponsePickupRate**](SchedulePickupUPSResponsePickupRate.md) |  | [optional] 

## Methods

### NewSchedulePickup200Response

`func NewSchedulePickup200Response() *SchedulePickup200Response`

NewSchedulePickup200Response instantiates a new SchedulePickup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickup200ResponseWithDefaults

`func NewSchedulePickup200ResponseWithDefaults() *SchedulePickup200Response`

NewSchedulePickup200ResponseWithDefaults instantiates a new SchedulePickup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickup200Response) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickup200Response) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickup200Response) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *SchedulePickup200Response) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetPickupConfirmationNumber

`func (o *SchedulePickup200Response) GetPickupConfirmationNumber() string`

GetPickupConfirmationNumber returns the PickupConfirmationNumber field if non-nil, zero value otherwise.

### GetPickupConfirmationNumberOk

`func (o *SchedulePickup200Response) GetPickupConfirmationNumberOk() (*string, bool)`

GetPickupConfirmationNumberOk returns a tuple with the PickupConfirmationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupConfirmationNumber

`func (o *SchedulePickup200Response) SetPickupConfirmationNumber(v string)`

SetPickupConfirmationNumber sets PickupConfirmationNumber field to given value.

### HasPickupConfirmationNumber

`func (o *SchedulePickup200Response) HasPickupConfirmationNumber() bool`

HasPickupConfirmationNumber returns a boolean if a field has been set.

### GetPickupId

`func (o *SchedulePickup200Response) GetPickupId() string`

GetPickupId returns the PickupId field if non-nil, zero value otherwise.

### GetPickupIdOk

`func (o *SchedulePickup200Response) GetPickupIdOk() (*string, bool)`

GetPickupIdOk returns a tuple with the PickupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupId

`func (o *SchedulePickup200Response) SetPickupId(v string)`

SetPickupId sets PickupId field to given value.

### HasPickupId

`func (o *SchedulePickup200Response) HasPickupId() bool`

HasPickupId returns a boolean if a field has been set.

### GetCarrier

`func (o *SchedulePickup200Response) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *SchedulePickup200Response) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *SchedulePickup200Response) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *SchedulePickup200Response) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *SchedulePickup200Response) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickup200Response) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickup200Response) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *SchedulePickup200Response) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *SchedulePickup200Response) GetPickupAddress() SchedulePickupDHLEXPResponsePickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickup200Response) GetPickupAddressOk() (*SchedulePickupDHLEXPResponsePickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickup200Response) SetPickupAddress(v SchedulePickupDHLEXPResponsePickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *SchedulePickup200Response) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetShipmentIds

`func (o *SchedulePickup200Response) GetShipmentIds() []string`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *SchedulePickup200Response) GetShipmentIdsOk() (*[]string, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *SchedulePickup200Response) SetShipmentIds(v []string)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *SchedulePickup200Response) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickup200Response) GetPickupSummary() []SchedulePickupUSPSResponsePickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickup200Response) GetPickupSummaryOk() (*[]SchedulePickupUSPSResponsePickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickup200Response) SetPickupSummary(v []SchedulePickupUSPSResponsePickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickup200Response) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetAdditionalnotes

`func (o *SchedulePickup200Response) GetAdditionalnotes() string`

GetAdditionalnotes returns the Additionalnotes field if non-nil, zero value otherwise.

### GetAdditionalnotesOk

`func (o *SchedulePickup200Response) GetAdditionalnotesOk() (*string, bool)`

GetAdditionalnotesOk returns a tuple with the Additionalnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalnotes

`func (o *SchedulePickup200Response) SetAdditionalnotes(v string)`

SetAdditionalnotes sets Additionalnotes field to given value.

### HasAdditionalnotes

`func (o *SchedulePickup200Response) HasAdditionalnotes() bool`

HasAdditionalnotes returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickup200Response) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickup200Response) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickup200Response) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickup200Response) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPickupDateTime

`func (o *SchedulePickup200Response) GetPickupDateTime() string`

GetPickupDateTime returns the PickupDateTime field if non-nil, zero value otherwise.

### GetPickupDateTimeOk

`func (o *SchedulePickup200Response) GetPickupDateTimeOk() (*string, bool)`

GetPickupDateTimeOk returns a tuple with the PickupDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDateTime

`func (o *SchedulePickup200Response) SetPickupDateTime(v string)`

SetPickupDateTime sets PickupDateTime field to given value.

### HasPickupDateTime

`func (o *SchedulePickup200Response) HasPickupDateTime() bool`

HasPickupDateTime returns a boolean if a field has been set.

### GetPickupTotalWeight

`func (o *SchedulePickup200Response) GetPickupTotalWeight() float32`

GetPickupTotalWeight returns the PickupTotalWeight field if non-nil, zero value otherwise.

### GetPickupTotalWeightOk

`func (o *SchedulePickup200Response) GetPickupTotalWeightOk() (*float32, bool)`

GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeight

`func (o *SchedulePickup200Response) SetPickupTotalWeight(v float32)`

SetPickupTotalWeight sets PickupTotalWeight field to given value.

### HasPickupTotalWeight

`func (o *SchedulePickup200Response) HasPickupTotalWeight() bool`

HasPickupTotalWeight returns a boolean if a field has been set.

### GetPickupTotalWeightUnit

`func (o *SchedulePickup200Response) GetPickupTotalWeightUnit() string`

GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field if non-nil, zero value otherwise.

### GetPickupTotalWeightUnitOk

`func (o *SchedulePickup200Response) GetPickupTotalWeightUnitOk() (*string, bool)`

GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeightUnit

`func (o *SchedulePickup200Response) SetPickupTotalWeightUnit(v string)`

SetPickupTotalWeightUnit sets PickupTotalWeightUnit field to given value.

### HasPickupTotalWeightUnit

`func (o *SchedulePickup200Response) HasPickupTotalWeightUnit() bool`

HasPickupTotalWeightUnit returns a boolean if a field has been set.

### GetPickupOptions

`func (o *SchedulePickup200Response) GetPickupOptions() GetAllPickupsDataInnerPickupOptions`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *SchedulePickup200Response) GetPickupOptionsOk() (*GetAllPickupsDataInnerPickupOptions, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *SchedulePickup200Response) SetPickupOptions(v GetAllPickupsDataInnerPickupOptions)`

SetPickupOptions sets PickupOptions field to given value.

### HasPickupOptions

`func (o *SchedulePickup200Response) HasPickupOptions() bool`

HasPickupOptions returns a boolean if a field has been set.

### GetPickupRate

`func (o *SchedulePickup200Response) GetPickupRate() SchedulePickupUPSResponsePickupRate`

GetPickupRate returns the PickupRate field if non-nil, zero value otherwise.

### GetPickupRateOk

`func (o *SchedulePickup200Response) GetPickupRateOk() (*SchedulePickupUPSResponsePickupRate, bool)`

GetPickupRateOk returns a tuple with the PickupRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupRate

`func (o *SchedulePickup200Response) SetPickupRate(v SchedulePickupUPSResponsePickupRate)`

SetPickupRate sets PickupRate field to given value.

### HasPickupRate

`func (o *SchedulePickup200Response) HasPickupRate() bool`

HasPickupRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


