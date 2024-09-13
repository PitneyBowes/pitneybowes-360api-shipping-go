# ReturnLabelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | Pointer to [**ReturnLabelFromAddress**](ReturnLabelFromAddress.md) |  | [optional] 
**Parcel** | Pointer to [**ShipmentDomesticParcel**](ShipmentDomesticParcel.md) |  | [optional] 
**CorrelationId** | Pointer to **string** | correction id | [optional] 
**ServiceId** | Pointer to **string** | &gt;-A unique identifier given to the carrier-specific service. This is required for creating a shipment, while it is optional for rating a parcel. | [optional] 
**ParcelTrackingNumber** | Pointer to **string** | &gt;-A unique identifier parcel tracking number | [optional] 
**ShipmentId** | Pointer to **string** | &gt;-A unique identifier shipment tracking number | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentOptionsV2**](ShipmentOptionsV2.md) |  | [optional] 
**LabelLayout** | Pointer to [**[]ReturnLabelResponseLabelLayoutInner**](ReturnLabelResponseLabelLayoutInner.md) | labelLayout details | [optional] 
**ToAddress** | Pointer to [**ReturnLabelResponseToAddress**](ReturnLabelResponseToAddress.md) |  | [optional] 
**Rate** | Pointer to [**ReturnLabelResponseRate**](ReturnLabelResponseRate.md) |  | [optional] 

## Methods

### NewReturnLabelResponse

`func NewReturnLabelResponse() *ReturnLabelResponse`

NewReturnLabelResponse instantiates a new ReturnLabelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnLabelResponseWithDefaults

`func NewReturnLabelResponseWithDefaults() *ReturnLabelResponse`

NewReturnLabelResponseWithDefaults instantiates a new ReturnLabelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *ReturnLabelResponse) GetFromAddress() ReturnLabelFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ReturnLabelResponse) GetFromAddressOk() (*ReturnLabelFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ReturnLabelResponse) SetFromAddress(v ReturnLabelFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *ReturnLabelResponse) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcel

`func (o *ReturnLabelResponse) GetParcel() ShipmentDomesticParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *ReturnLabelResponse) GetParcelOk() (*ShipmentDomesticParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *ReturnLabelResponse) SetParcel(v ShipmentDomesticParcel)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *ReturnLabelResponse) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetCorrelationId

`func (o *ReturnLabelResponse) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *ReturnLabelResponse) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *ReturnLabelResponse) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *ReturnLabelResponse) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetServiceId

`func (o *ReturnLabelResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ReturnLabelResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ReturnLabelResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ReturnLabelResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *ReturnLabelResponse) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *ReturnLabelResponse) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *ReturnLabelResponse) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *ReturnLabelResponse) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetShipmentId

`func (o *ReturnLabelResponse) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *ReturnLabelResponse) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *ReturnLabelResponse) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *ReturnLabelResponse) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *ReturnLabelResponse) GetShipmentOptions() ShipmentOptionsV2`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *ReturnLabelResponse) GetShipmentOptionsOk() (*ShipmentOptionsV2, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *ReturnLabelResponse) SetShipmentOptions(v ShipmentOptionsV2)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *ReturnLabelResponse) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetLabelLayout

`func (o *ReturnLabelResponse) GetLabelLayout() []ReturnLabelResponseLabelLayoutInner`

GetLabelLayout returns the LabelLayout field if non-nil, zero value otherwise.

### GetLabelLayoutOk

`func (o *ReturnLabelResponse) GetLabelLayoutOk() (*[]ReturnLabelResponseLabelLayoutInner, bool)`

GetLabelLayoutOk returns a tuple with the LabelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelLayout

`func (o *ReturnLabelResponse) SetLabelLayout(v []ReturnLabelResponseLabelLayoutInner)`

SetLabelLayout sets LabelLayout field to given value.

### HasLabelLayout

`func (o *ReturnLabelResponse) HasLabelLayout() bool`

HasLabelLayout returns a boolean if a field has been set.

### GetToAddress

`func (o *ReturnLabelResponse) GetToAddress() ReturnLabelResponseToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *ReturnLabelResponse) GetToAddressOk() (*ReturnLabelResponseToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *ReturnLabelResponse) SetToAddress(v ReturnLabelResponseToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *ReturnLabelResponse) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetRate

`func (o *ReturnLabelResponse) GetRate() ReturnLabelResponseRate`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ReturnLabelResponse) GetRateOk() (*ReturnLabelResponseRate, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ReturnLabelResponse) SetRate(v ReturnLabelResponseRate)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ReturnLabelResponse) HasRate() bool`

HasRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


