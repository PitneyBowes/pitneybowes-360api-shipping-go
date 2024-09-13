# CancelShipmentV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Carrier** | Pointer to **string** | The name of the Carrier. | [optional] 
**TotalCarrierCharge** | Pointer to **float32** | The total amount payable to the carrier, including special service fees, surcharges, and any international taxes and duties. | [optional] 
**ParcelTrackingNumber** | Pointer to **string** | The Tracking number given to the Parcel for tracking purpose. | [optional] 
**Status** | Pointer to **string** | Status of the Shipment. | [optional] 
**References** | Pointer to [**CancelShipmentV2References**](CancelShipmentV2References.md) |  | [optional] 

## Methods

### NewCancelShipmentV2

`func NewCancelShipmentV2() *CancelShipmentV2`

NewCancelShipmentV2 instantiates a new CancelShipmentV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelShipmentV2WithDefaults

`func NewCancelShipmentV2WithDefaults() *CancelShipmentV2`

NewCancelShipmentV2WithDefaults instantiates a new CancelShipmentV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrier

`func (o *CancelShipmentV2) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *CancelShipmentV2) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *CancelShipmentV2) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *CancelShipmentV2) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetTotalCarrierCharge

`func (o *CancelShipmentV2) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *CancelShipmentV2) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *CancelShipmentV2) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *CancelShipmentV2) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *CancelShipmentV2) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *CancelShipmentV2) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *CancelShipmentV2) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *CancelShipmentV2) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetStatus

`func (o *CancelShipmentV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CancelShipmentV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CancelShipmentV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CancelShipmentV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReferences

`func (o *CancelShipmentV2) GetReferences() CancelShipmentV2References`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CancelShipmentV2) GetReferencesOk() (*CancelShipmentV2References, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CancelShipmentV2) SetReferences(v CancelShipmentV2References)`

SetReferences sets References field to given value.

### HasReferences

`func (o *CancelShipmentV2) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


