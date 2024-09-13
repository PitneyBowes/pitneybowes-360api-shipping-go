# DomesticShipmentResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | Key assigned by the shipping system to the transaction. | [optional] 
**ShipmentId** | Pointer to **string** | The shipmentId, a unique identifier for an individual Shipment. | [optional] 
**ParcelTrackingNumber** | Pointer to **string** | The Tracking number given to the Parcel for tracking purpose. | [optional] 
**LabelLayout** | Pointer to [**[]DomesticShipmentResponseV2LabelLayoutInner**](DomesticShipmentResponseV2LabelLayoutInner.md) | It defines the layout of the shipping label. | [optional] 
**Parcel** | Pointer to [**ParcelV2**](ParcelV2.md) |  | [optional] 
**Rate** | Pointer to [**RateResponseV2**](RateResponseV2.md) |  | [optional] 
**References** | Pointer to [**ReferenceV2**](ReferenceV2.md) |  | [optional] 
**PrintStatus** | Pointer to **string** | Status of the Printed Label. | [optional] 
**PrintError** | Pointer to [**DomesticShipmentResponseV2PrintError**](DomesticShipmentResponseV2PrintError.md) |  | [optional] 
**FromAddress** | Pointer to [**FromAddressV2Response**](FromAddressV2Response.md) |  | [optional] 
**ToAddress** | Pointer to [**ToAddressV2Response**](ToAddressV2Response.md) |  | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentOptions**](ShipmentOptions.md) |  | [optional] 

## Methods

### NewDomesticShipmentResponseV2

`func NewDomesticShipmentResponseV2() *DomesticShipmentResponseV2`

NewDomesticShipmentResponseV2 instantiates a new DomesticShipmentResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomesticShipmentResponseV2WithDefaults

`func NewDomesticShipmentResponseV2WithDefaults() *DomesticShipmentResponseV2`

NewDomesticShipmentResponseV2WithDefaults instantiates a new DomesticShipmentResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *DomesticShipmentResponseV2) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *DomesticShipmentResponseV2) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *DomesticShipmentResponseV2) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *DomesticShipmentResponseV2) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetShipmentId

`func (o *DomesticShipmentResponseV2) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *DomesticShipmentResponseV2) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *DomesticShipmentResponseV2) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *DomesticShipmentResponseV2) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *DomesticShipmentResponseV2) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *DomesticShipmentResponseV2) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *DomesticShipmentResponseV2) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *DomesticShipmentResponseV2) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetLabelLayout

`func (o *DomesticShipmentResponseV2) GetLabelLayout() []DomesticShipmentResponseV2LabelLayoutInner`

GetLabelLayout returns the LabelLayout field if non-nil, zero value otherwise.

### GetLabelLayoutOk

`func (o *DomesticShipmentResponseV2) GetLabelLayoutOk() (*[]DomesticShipmentResponseV2LabelLayoutInner, bool)`

GetLabelLayoutOk returns a tuple with the LabelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelLayout

`func (o *DomesticShipmentResponseV2) SetLabelLayout(v []DomesticShipmentResponseV2LabelLayoutInner)`

SetLabelLayout sets LabelLayout field to given value.

### HasLabelLayout

`func (o *DomesticShipmentResponseV2) HasLabelLayout() bool`

HasLabelLayout returns a boolean if a field has been set.

### GetParcel

`func (o *DomesticShipmentResponseV2) GetParcel() ParcelV2`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *DomesticShipmentResponseV2) GetParcelOk() (*ParcelV2, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *DomesticShipmentResponseV2) SetParcel(v ParcelV2)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *DomesticShipmentResponseV2) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetRate

`func (o *DomesticShipmentResponseV2) GetRate() RateResponseV2`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *DomesticShipmentResponseV2) GetRateOk() (*RateResponseV2, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *DomesticShipmentResponseV2) SetRate(v RateResponseV2)`

SetRate sets Rate field to given value.

### HasRate

`func (o *DomesticShipmentResponseV2) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetReferences

`func (o *DomesticShipmentResponseV2) GetReferences() ReferenceV2`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *DomesticShipmentResponseV2) GetReferencesOk() (*ReferenceV2, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *DomesticShipmentResponseV2) SetReferences(v ReferenceV2)`

SetReferences sets References field to given value.

### HasReferences

`func (o *DomesticShipmentResponseV2) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetPrintStatus

`func (o *DomesticShipmentResponseV2) GetPrintStatus() string`

GetPrintStatus returns the PrintStatus field if non-nil, zero value otherwise.

### GetPrintStatusOk

`func (o *DomesticShipmentResponseV2) GetPrintStatusOk() (*string, bool)`

GetPrintStatusOk returns a tuple with the PrintStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintStatus

`func (o *DomesticShipmentResponseV2) SetPrintStatus(v string)`

SetPrintStatus sets PrintStatus field to given value.

### HasPrintStatus

`func (o *DomesticShipmentResponseV2) HasPrintStatus() bool`

HasPrintStatus returns a boolean if a field has been set.

### GetPrintError

`func (o *DomesticShipmentResponseV2) GetPrintError() DomesticShipmentResponseV2PrintError`

GetPrintError returns the PrintError field if non-nil, zero value otherwise.

### GetPrintErrorOk

`func (o *DomesticShipmentResponseV2) GetPrintErrorOk() (*DomesticShipmentResponseV2PrintError, bool)`

GetPrintErrorOk returns a tuple with the PrintError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintError

`func (o *DomesticShipmentResponseV2) SetPrintError(v DomesticShipmentResponseV2PrintError)`

SetPrintError sets PrintError field to given value.

### HasPrintError

`func (o *DomesticShipmentResponseV2) HasPrintError() bool`

HasPrintError returns a boolean if a field has been set.

### GetFromAddress

`func (o *DomesticShipmentResponseV2) GetFromAddress() FromAddressV2Response`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *DomesticShipmentResponseV2) GetFromAddressOk() (*FromAddressV2Response, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *DomesticShipmentResponseV2) SetFromAddress(v FromAddressV2Response)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *DomesticShipmentResponseV2) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetToAddress

`func (o *DomesticShipmentResponseV2) GetToAddress() ToAddressV2Response`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *DomesticShipmentResponseV2) GetToAddressOk() (*ToAddressV2Response, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *DomesticShipmentResponseV2) SetToAddress(v ToAddressV2Response)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *DomesticShipmentResponseV2) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *DomesticShipmentResponseV2) GetShipmentOptions() ShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *DomesticShipmentResponseV2) GetShipmentOptionsOk() (*ShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *DomesticShipmentResponseV2) SetShipmentOptions(v ShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *DomesticShipmentResponseV2) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


