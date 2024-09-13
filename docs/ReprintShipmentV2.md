# ReprintShipmentV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentId** | Pointer to **string** | The shipmentId, a unique identifier for an individual Shipment, which is used for Reprint or Cancel. | [optional] 
**ParcelTrackingNumber** | Pointer to **string** | The Tracking number given to the Parcel for tracking purpose. | [optional] 
**LabelLayout** | Pointer to [**[]ReprintShipmentV2LabelLayoutInner**](ReprintShipmentV2LabelLayoutInner.md) | It defines the layout of the shipping label. | [optional] 
**Parcel** | Pointer to [**ParcelV2**](ParcelV2.md) |  | [optional] 
**Rate** | Pointer to [**RateResponseV2**](RateResponseV2.md) |  | [optional] 
**References** | Pointer to [**CancelShipmentV2References**](CancelShipmentV2References.md) |  | [optional] 
**PrintStatus** | Pointer to **string** | Status of the Printed Label. | [optional] 

## Methods

### NewReprintShipmentV2

`func NewReprintShipmentV2() *ReprintShipmentV2`

NewReprintShipmentV2 instantiates a new ReprintShipmentV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReprintShipmentV2WithDefaults

`func NewReprintShipmentV2WithDefaults() *ReprintShipmentV2`

NewReprintShipmentV2WithDefaults instantiates a new ReprintShipmentV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipmentId

`func (o *ReprintShipmentV2) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *ReprintShipmentV2) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *ReprintShipmentV2) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *ReprintShipmentV2) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *ReprintShipmentV2) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *ReprintShipmentV2) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *ReprintShipmentV2) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *ReprintShipmentV2) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetLabelLayout

`func (o *ReprintShipmentV2) GetLabelLayout() []ReprintShipmentV2LabelLayoutInner`

GetLabelLayout returns the LabelLayout field if non-nil, zero value otherwise.

### GetLabelLayoutOk

`func (o *ReprintShipmentV2) GetLabelLayoutOk() (*[]ReprintShipmentV2LabelLayoutInner, bool)`

GetLabelLayoutOk returns a tuple with the LabelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelLayout

`func (o *ReprintShipmentV2) SetLabelLayout(v []ReprintShipmentV2LabelLayoutInner)`

SetLabelLayout sets LabelLayout field to given value.

### HasLabelLayout

`func (o *ReprintShipmentV2) HasLabelLayout() bool`

HasLabelLayout returns a boolean if a field has been set.

### GetParcel

`func (o *ReprintShipmentV2) GetParcel() ParcelV2`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *ReprintShipmentV2) GetParcelOk() (*ParcelV2, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *ReprintShipmentV2) SetParcel(v ParcelV2)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *ReprintShipmentV2) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetRate

`func (o *ReprintShipmentV2) GetRate() RateResponseV2`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ReprintShipmentV2) GetRateOk() (*RateResponseV2, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ReprintShipmentV2) SetRate(v RateResponseV2)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ReprintShipmentV2) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetReferences

`func (o *ReprintShipmentV2) GetReferences() CancelShipmentV2References`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ReprintShipmentV2) GetReferencesOk() (*CancelShipmentV2References, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ReprintShipmentV2) SetReferences(v CancelShipmentV2References)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ReprintShipmentV2) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetPrintStatus

`func (o *ReprintShipmentV2) GetPrintStatus() string`

GetPrintStatus returns the PrintStatus field if non-nil, zero value otherwise.

### GetPrintStatusOk

`func (o *ReprintShipmentV2) GetPrintStatusOk() (*string, bool)`

GetPrintStatusOk returns a tuple with the PrintStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintStatus

`func (o *ReprintShipmentV2) SetPrintStatus(v string)`

SetPrintStatus sets PrintStatus field to given value.

### HasPrintStatus

`func (o *ReprintShipmentV2) HasPrintStatus() bool`

HasPrintStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


