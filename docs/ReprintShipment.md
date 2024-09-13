# ReprintShipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | This is a GUID (globally unique identifier) that&#39;s automatically generated for every request that the webserver receives. | [optional] 
**Size** | Pointer to **string** | This defines the label size of the Shipment, e.g., Shipping Label having Doc Size (8&#39; X 11&#39;). | [optional] 
**Type** | Pointer to **string** | This defines the type of the Shipment, e.g., Shipping Label. | [optional] 
**Format** | Pointer to **string** | This defines the type of the shipment which is printed. For example Shipping label prints in PDF form. | [optional] 
**FromAddress** | Pointer to [**ReprintShipmentFromAddress**](ReprintShipmentFromAddress.md) |  | [optional] 
**Parcel** | Pointer to [**ReprintShipmentParcel**](ReprintShipmentParcel.md) |  | [optional] 
**ParcelTrackingNumber** | Pointer to **string** | The Tracking number given to the Parcel for tracking purpose. | [optional] 
**Rate** | Pointer to [**ReprintShipmentRate**](ReprintShipmentRate.md) |  | [optional] 
**ShipmentId** | Pointer to **string** | A unique identifier associated with Shipment ID. | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentOptions**](ShipmentOptions.md) |  | [optional] 
**ToAddress** | Pointer to [**ReprintShipmentToAddress**](ReprintShipmentToAddress.md) |  | [optional] 

## Methods

### NewReprintShipment

`func NewReprintShipment() *ReprintShipment`

NewReprintShipment instantiates a new ReprintShipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReprintShipmentWithDefaults

`func NewReprintShipmentWithDefaults() *ReprintShipment`

NewReprintShipmentWithDefaults instantiates a new ReprintShipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *ReprintShipment) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *ReprintShipment) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *ReprintShipment) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *ReprintShipment) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetSize

`func (o *ReprintShipment) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ReprintShipment) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ReprintShipment) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ReprintShipment) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *ReprintShipment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReprintShipment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReprintShipment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReprintShipment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFormat

`func (o *ReprintShipment) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ReprintShipment) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ReprintShipment) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ReprintShipment) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFromAddress

`func (o *ReprintShipment) GetFromAddress() ReprintShipmentFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ReprintShipment) GetFromAddressOk() (*ReprintShipmentFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ReprintShipment) SetFromAddress(v ReprintShipmentFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *ReprintShipment) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcel

`func (o *ReprintShipment) GetParcel() ReprintShipmentParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *ReprintShipment) GetParcelOk() (*ReprintShipmentParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *ReprintShipment) SetParcel(v ReprintShipmentParcel)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *ReprintShipment) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *ReprintShipment) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *ReprintShipment) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *ReprintShipment) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *ReprintShipment) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetRate

`func (o *ReprintShipment) GetRate() ReprintShipmentRate`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ReprintShipment) GetRateOk() (*ReprintShipmentRate, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ReprintShipment) SetRate(v ReprintShipmentRate)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ReprintShipment) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetShipmentId

`func (o *ReprintShipment) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *ReprintShipment) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *ReprintShipment) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *ReprintShipment) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *ReprintShipment) GetShipmentOptions() ShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *ReprintShipment) GetShipmentOptionsOk() (*ShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *ReprintShipment) SetShipmentOptions(v ShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *ReprintShipment) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetToAddress

`func (o *ReprintShipment) GetToAddress() ReprintShipmentToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *ReprintShipment) GetToAddressOk() (*ReprintShipmentToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *ReprintShipment) SetToAddress(v ReprintShipmentToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *ReprintShipment) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


