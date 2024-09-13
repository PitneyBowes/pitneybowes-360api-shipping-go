# GetAllPickupsDataInner

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
**PickupSummary** | Pointer to [**[]GetAllPickupsDataInnerPickupSummaryInner**](GetAllPickupsDataInnerPickupSummaryInner.md) | It displays the package details provided in the request. | [optional] 
**SpecialInstructions** | Pointer to **string** | It displays additional comments or remarks provided in the request, it would be printed on the scheduled pickup document | [optional] 
**Reference** | Pointer to **string** | It displays any reference information provided in the request. | [optional] 
**PickupDateTime** | Pointer to **string** | It displays the scheduled pickup date and time (in UTC) | [optional] 
**PickupTotalWeight** | Pointer to **float32** | It displays the total package weight. | [optional] 
**PickupTotalWeightUnit** | Pointer to **string** | It displays the weight unit. | [optional] 
**PickupOptions** | Pointer to [**GetAllPickupsDataInnerPickupOptions**](GetAllPickupsDataInnerPickupOptions.md) |  | [optional] 
**PickupRate** | Pointer to [**SchedulePickupUPSResponsePickupRate**](SchedulePickupUPSResponsePickupRate.md) |  | [optional] 

## Methods

### NewGetAllPickupsDataInner

`func NewGetAllPickupsDataInner() *GetAllPickupsDataInner`

NewGetAllPickupsDataInner instantiates a new GetAllPickupsDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllPickupsDataInnerWithDefaults

`func NewGetAllPickupsDataInnerWithDefaults() *GetAllPickupsDataInner`

NewGetAllPickupsDataInnerWithDefaults instantiates a new GetAllPickupsDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *GetAllPickupsDataInner) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *GetAllPickupsDataInner) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *GetAllPickupsDataInner) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.

### HasPackageLocation

`func (o *GetAllPickupsDataInner) HasPackageLocation() bool`

HasPackageLocation returns a boolean if a field has been set.

### GetPickupConfirmationNumber

`func (o *GetAllPickupsDataInner) GetPickupConfirmationNumber() string`

GetPickupConfirmationNumber returns the PickupConfirmationNumber field if non-nil, zero value otherwise.

### GetPickupConfirmationNumberOk

`func (o *GetAllPickupsDataInner) GetPickupConfirmationNumberOk() (*string, bool)`

GetPickupConfirmationNumberOk returns a tuple with the PickupConfirmationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupConfirmationNumber

`func (o *GetAllPickupsDataInner) SetPickupConfirmationNumber(v string)`

SetPickupConfirmationNumber sets PickupConfirmationNumber field to given value.

### HasPickupConfirmationNumber

`func (o *GetAllPickupsDataInner) HasPickupConfirmationNumber() bool`

HasPickupConfirmationNumber returns a boolean if a field has been set.

### GetPickupId

`func (o *GetAllPickupsDataInner) GetPickupId() string`

GetPickupId returns the PickupId field if non-nil, zero value otherwise.

### GetPickupIdOk

`func (o *GetAllPickupsDataInner) GetPickupIdOk() (*string, bool)`

GetPickupIdOk returns a tuple with the PickupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupId

`func (o *GetAllPickupsDataInner) SetPickupId(v string)`

SetPickupId sets PickupId field to given value.

### HasPickupId

`func (o *GetAllPickupsDataInner) HasPickupId() bool`

HasPickupId returns a boolean if a field has been set.

### GetCarrier

`func (o *GetAllPickupsDataInner) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *GetAllPickupsDataInner) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *GetAllPickupsDataInner) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *GetAllPickupsDataInner) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *GetAllPickupsDataInner) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *GetAllPickupsDataInner) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *GetAllPickupsDataInner) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *GetAllPickupsDataInner) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetPickupAddress

`func (o *GetAllPickupsDataInner) GetPickupAddress() SchedulePickupDHLEXPResponsePickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *GetAllPickupsDataInner) GetPickupAddressOk() (*SchedulePickupDHLEXPResponsePickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *GetAllPickupsDataInner) SetPickupAddress(v SchedulePickupDHLEXPResponsePickupAddress)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *GetAllPickupsDataInner) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetShipmentIds

`func (o *GetAllPickupsDataInner) GetShipmentIds() []string`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *GetAllPickupsDataInner) GetShipmentIdsOk() (*[]string, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *GetAllPickupsDataInner) SetShipmentIds(v []string)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *GetAllPickupsDataInner) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.

### GetPickupSummary

`func (o *GetAllPickupsDataInner) GetPickupSummary() []GetAllPickupsDataInnerPickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *GetAllPickupsDataInner) GetPickupSummaryOk() (*[]GetAllPickupsDataInnerPickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *GetAllPickupsDataInner) SetPickupSummary(v []GetAllPickupsDataInnerPickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *GetAllPickupsDataInner) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetSpecialInstructions

`func (o *GetAllPickupsDataInner) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *GetAllPickupsDataInner) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *GetAllPickupsDataInner) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *GetAllPickupsDataInner) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### GetReference

`func (o *GetAllPickupsDataInner) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GetAllPickupsDataInner) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GetAllPickupsDataInner) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GetAllPickupsDataInner) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPickupDateTime

`func (o *GetAllPickupsDataInner) GetPickupDateTime() string`

GetPickupDateTime returns the PickupDateTime field if non-nil, zero value otherwise.

### GetPickupDateTimeOk

`func (o *GetAllPickupsDataInner) GetPickupDateTimeOk() (*string, bool)`

GetPickupDateTimeOk returns a tuple with the PickupDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDateTime

`func (o *GetAllPickupsDataInner) SetPickupDateTime(v string)`

SetPickupDateTime sets PickupDateTime field to given value.

### HasPickupDateTime

`func (o *GetAllPickupsDataInner) HasPickupDateTime() bool`

HasPickupDateTime returns a boolean if a field has been set.

### GetPickupTotalWeight

`func (o *GetAllPickupsDataInner) GetPickupTotalWeight() float32`

GetPickupTotalWeight returns the PickupTotalWeight field if non-nil, zero value otherwise.

### GetPickupTotalWeightOk

`func (o *GetAllPickupsDataInner) GetPickupTotalWeightOk() (*float32, bool)`

GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeight

`func (o *GetAllPickupsDataInner) SetPickupTotalWeight(v float32)`

SetPickupTotalWeight sets PickupTotalWeight field to given value.

### HasPickupTotalWeight

`func (o *GetAllPickupsDataInner) HasPickupTotalWeight() bool`

HasPickupTotalWeight returns a boolean if a field has been set.

### GetPickupTotalWeightUnit

`func (o *GetAllPickupsDataInner) GetPickupTotalWeightUnit() string`

GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field if non-nil, zero value otherwise.

### GetPickupTotalWeightUnitOk

`func (o *GetAllPickupsDataInner) GetPickupTotalWeightUnitOk() (*string, bool)`

GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTotalWeightUnit

`func (o *GetAllPickupsDataInner) SetPickupTotalWeightUnit(v string)`

SetPickupTotalWeightUnit sets PickupTotalWeightUnit field to given value.

### HasPickupTotalWeightUnit

`func (o *GetAllPickupsDataInner) HasPickupTotalWeightUnit() bool`

HasPickupTotalWeightUnit returns a boolean if a field has been set.

### GetPickupOptions

`func (o *GetAllPickupsDataInner) GetPickupOptions() GetAllPickupsDataInnerPickupOptions`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *GetAllPickupsDataInner) GetPickupOptionsOk() (*GetAllPickupsDataInnerPickupOptions, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *GetAllPickupsDataInner) SetPickupOptions(v GetAllPickupsDataInnerPickupOptions)`

SetPickupOptions sets PickupOptions field to given value.

### HasPickupOptions

`func (o *GetAllPickupsDataInner) HasPickupOptions() bool`

HasPickupOptions returns a boolean if a field has been set.

### GetPickupRate

`func (o *GetAllPickupsDataInner) GetPickupRate() SchedulePickupUPSResponsePickupRate`

GetPickupRate returns the PickupRate field if non-nil, zero value otherwise.

### GetPickupRateOk

`func (o *GetAllPickupsDataInner) GetPickupRateOk() (*SchedulePickupUPSResponsePickupRate, bool)`

GetPickupRateOk returns a tuple with the PickupRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupRate

`func (o *GetAllPickupsDataInner) SetPickupRate(v SchedulePickupUPSResponsePickupRate)`

SetPickupRate sets PickupRate field to given value.

### HasPickupRate

`func (o *GetAllPickupsDataInner) HasPickupRate() bool`

HasPickupRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


