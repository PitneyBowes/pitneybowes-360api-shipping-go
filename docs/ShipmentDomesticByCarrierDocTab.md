# ShipmentDomesticByCarrierDocTab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateName** | Pointer to **string** | This is an optional field and will be used when the client has multiple doctab options. | [optional] 
**ToAddressName** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**ParcelTrackingNumber** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**Carrier** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**ServiceId** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**DateOfShipment** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**PieceNumber** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**PackageTotalCarrierCharge** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**TotalCarrierCharge** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**PackageWeight** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**TotalWeight** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**CustomField1** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**CustomField2** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**CustomField3** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 
**CustomField4** | Pointer to [**DocTabItem**](DocTabItem.md) |  | [optional] 

## Methods

### NewShipmentDomesticByCarrierDocTab

`func NewShipmentDomesticByCarrierDocTab() *ShipmentDomesticByCarrierDocTab`

NewShipmentDomesticByCarrierDocTab instantiates a new ShipmentDomesticByCarrierDocTab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentDomesticByCarrierDocTabWithDefaults

`func NewShipmentDomesticByCarrierDocTabWithDefaults() *ShipmentDomesticByCarrierDocTab`

NewShipmentDomesticByCarrierDocTabWithDefaults instantiates a new ShipmentDomesticByCarrierDocTab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateName

`func (o *ShipmentDomesticByCarrierDocTab) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ShipmentDomesticByCarrierDocTab) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ShipmentDomesticByCarrierDocTab) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *ShipmentDomesticByCarrierDocTab) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetToAddressName

`func (o *ShipmentDomesticByCarrierDocTab) GetToAddressName() DocTabItem`

GetToAddressName returns the ToAddressName field if non-nil, zero value otherwise.

### GetToAddressNameOk

`func (o *ShipmentDomesticByCarrierDocTab) GetToAddressNameOk() (*DocTabItem, bool)`

GetToAddressNameOk returns a tuple with the ToAddressName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddressName

`func (o *ShipmentDomesticByCarrierDocTab) SetToAddressName(v DocTabItem)`

SetToAddressName sets ToAddressName field to given value.

### HasToAddressName

`func (o *ShipmentDomesticByCarrierDocTab) HasToAddressName() bool`

HasToAddressName returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *ShipmentDomesticByCarrierDocTab) GetParcelTrackingNumber() DocTabItem`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *ShipmentDomesticByCarrierDocTab) GetParcelTrackingNumberOk() (*DocTabItem, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *ShipmentDomesticByCarrierDocTab) SetParcelTrackingNumber(v DocTabItem)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *ShipmentDomesticByCarrierDocTab) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetCarrier

`func (o *ShipmentDomesticByCarrierDocTab) GetCarrier() DocTabItem`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *ShipmentDomesticByCarrierDocTab) GetCarrierOk() (*DocTabItem, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *ShipmentDomesticByCarrierDocTab) SetCarrier(v DocTabItem)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *ShipmentDomesticByCarrierDocTab) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetServiceId

`func (o *ShipmentDomesticByCarrierDocTab) GetServiceId() DocTabItem`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ShipmentDomesticByCarrierDocTab) GetServiceIdOk() (*DocTabItem, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ShipmentDomesticByCarrierDocTab) SetServiceId(v DocTabItem)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ShipmentDomesticByCarrierDocTab) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetDateOfShipment

`func (o *ShipmentDomesticByCarrierDocTab) GetDateOfShipment() DocTabItem`

GetDateOfShipment returns the DateOfShipment field if non-nil, zero value otherwise.

### GetDateOfShipmentOk

`func (o *ShipmentDomesticByCarrierDocTab) GetDateOfShipmentOk() (*DocTabItem, bool)`

GetDateOfShipmentOk returns a tuple with the DateOfShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfShipment

`func (o *ShipmentDomesticByCarrierDocTab) SetDateOfShipment(v DocTabItem)`

SetDateOfShipment sets DateOfShipment field to given value.

### HasDateOfShipment

`func (o *ShipmentDomesticByCarrierDocTab) HasDateOfShipment() bool`

HasDateOfShipment returns a boolean if a field has been set.

### GetPieceNumber

`func (o *ShipmentDomesticByCarrierDocTab) GetPieceNumber() DocTabItem`

GetPieceNumber returns the PieceNumber field if non-nil, zero value otherwise.

### GetPieceNumberOk

`func (o *ShipmentDomesticByCarrierDocTab) GetPieceNumberOk() (*DocTabItem, bool)`

GetPieceNumberOk returns a tuple with the PieceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceNumber

`func (o *ShipmentDomesticByCarrierDocTab) SetPieceNumber(v DocTabItem)`

SetPieceNumber sets PieceNumber field to given value.

### HasPieceNumber

`func (o *ShipmentDomesticByCarrierDocTab) HasPieceNumber() bool`

HasPieceNumber returns a boolean if a field has been set.

### GetPackageTotalCarrierCharge

`func (o *ShipmentDomesticByCarrierDocTab) GetPackageTotalCarrierCharge() DocTabItem`

GetPackageTotalCarrierCharge returns the PackageTotalCarrierCharge field if non-nil, zero value otherwise.

### GetPackageTotalCarrierChargeOk

`func (o *ShipmentDomesticByCarrierDocTab) GetPackageTotalCarrierChargeOk() (*DocTabItem, bool)`

GetPackageTotalCarrierChargeOk returns a tuple with the PackageTotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageTotalCarrierCharge

`func (o *ShipmentDomesticByCarrierDocTab) SetPackageTotalCarrierCharge(v DocTabItem)`

SetPackageTotalCarrierCharge sets PackageTotalCarrierCharge field to given value.

### HasPackageTotalCarrierCharge

`func (o *ShipmentDomesticByCarrierDocTab) HasPackageTotalCarrierCharge() bool`

HasPackageTotalCarrierCharge returns a boolean if a field has been set.

### GetTotalCarrierCharge

`func (o *ShipmentDomesticByCarrierDocTab) GetTotalCarrierCharge() DocTabItem`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *ShipmentDomesticByCarrierDocTab) GetTotalCarrierChargeOk() (*DocTabItem, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *ShipmentDomesticByCarrierDocTab) SetTotalCarrierCharge(v DocTabItem)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *ShipmentDomesticByCarrierDocTab) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.

### GetPackageWeight

`func (o *ShipmentDomesticByCarrierDocTab) GetPackageWeight() DocTabItem`

GetPackageWeight returns the PackageWeight field if non-nil, zero value otherwise.

### GetPackageWeightOk

`func (o *ShipmentDomesticByCarrierDocTab) GetPackageWeightOk() (*DocTabItem, bool)`

GetPackageWeightOk returns a tuple with the PackageWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageWeight

`func (o *ShipmentDomesticByCarrierDocTab) SetPackageWeight(v DocTabItem)`

SetPackageWeight sets PackageWeight field to given value.

### HasPackageWeight

`func (o *ShipmentDomesticByCarrierDocTab) HasPackageWeight() bool`

HasPackageWeight returns a boolean if a field has been set.

### GetTotalWeight

`func (o *ShipmentDomesticByCarrierDocTab) GetTotalWeight() DocTabItem`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *ShipmentDomesticByCarrierDocTab) GetTotalWeightOk() (*DocTabItem, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *ShipmentDomesticByCarrierDocTab) SetTotalWeight(v DocTabItem)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *ShipmentDomesticByCarrierDocTab) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetCustomField1

`func (o *ShipmentDomesticByCarrierDocTab) GetCustomField1() DocTabItem`

GetCustomField1 returns the CustomField1 field if non-nil, zero value otherwise.

### GetCustomField1Ok

`func (o *ShipmentDomesticByCarrierDocTab) GetCustomField1Ok() (*DocTabItem, bool)`

GetCustomField1Ok returns a tuple with the CustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField1

`func (o *ShipmentDomesticByCarrierDocTab) SetCustomField1(v DocTabItem)`

SetCustomField1 sets CustomField1 field to given value.

### HasCustomField1

`func (o *ShipmentDomesticByCarrierDocTab) HasCustomField1() bool`

HasCustomField1 returns a boolean if a field has been set.

### GetCustomField2

`func (o *ShipmentDomesticByCarrierDocTab) GetCustomField2() DocTabItem`

GetCustomField2 returns the CustomField2 field if non-nil, zero value otherwise.

### GetCustomField2Ok

`func (o *ShipmentDomesticByCarrierDocTab) GetCustomField2Ok() (*DocTabItem, bool)`

GetCustomField2Ok returns a tuple with the CustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField2

`func (o *ShipmentDomesticByCarrierDocTab) SetCustomField2(v DocTabItem)`

SetCustomField2 sets CustomField2 field to given value.

### HasCustomField2

`func (o *ShipmentDomesticByCarrierDocTab) HasCustomField2() bool`

HasCustomField2 returns a boolean if a field has been set.

### GetCustomField3

`func (o *ShipmentDomesticByCarrierDocTab) GetCustomField3() DocTabItem`

GetCustomField3 returns the CustomField3 field if non-nil, zero value otherwise.

### GetCustomField3Ok

`func (o *ShipmentDomesticByCarrierDocTab) GetCustomField3Ok() (*DocTabItem, bool)`

GetCustomField3Ok returns a tuple with the CustomField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField3

`func (o *ShipmentDomesticByCarrierDocTab) SetCustomField3(v DocTabItem)`

SetCustomField3 sets CustomField3 field to given value.

### HasCustomField3

`func (o *ShipmentDomesticByCarrierDocTab) HasCustomField3() bool`

HasCustomField3 returns a boolean if a field has been set.

### GetCustomField4

`func (o *ShipmentDomesticByCarrierDocTab) GetCustomField4() DocTabItem`

GetCustomField4 returns the CustomField4 field if non-nil, zero value otherwise.

### GetCustomField4Ok

`func (o *ShipmentDomesticByCarrierDocTab) GetCustomField4Ok() (*DocTabItem, bool)`

GetCustomField4Ok returns a tuple with the CustomField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField4

`func (o *ShipmentDomesticByCarrierDocTab) SetCustomField4(v DocTabItem)`

SetCustomField4 sets CustomField4 field to given value.

### HasCustomField4

`func (o *ShipmentDomesticByCarrierDocTab) HasCustomField4() bool`

HasCustomField4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


