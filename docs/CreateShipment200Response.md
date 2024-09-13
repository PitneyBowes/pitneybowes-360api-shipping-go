# CreateShipment200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | This is a GUID (globally unique identifier) that&#39;s automatically generated for every request that the webserver receives. | [optional] 
**LabelLayout** | Pointer to [**[]DomesticShipmentResponseLabelLayoutInner**](DomesticShipmentResponseLabelLayoutInner.md) | This indicates the label layout and generated label details | [optional] 
**FromAddress** | Pointer to [**ReprintShipmentFromAddress**](ReprintShipmentFromAddress.md) |  | [optional] 
**Parcel** | Pointer to [**ShipmentDomesticParcel**](ShipmentDomesticParcel.md) |  | [optional] 
**ParcelId** | Pointer to **string** | &gt;-Parcel Id is optional and would be visible in case when is present in the request. | [optional] 
**ParcelTrackingNumber** | Pointer to **string** | The Tracking number given to the Parcel for tracking purpose. | [optional] 
**Rate** | Pointer to [**InternationalShipmentResponseRate**](InternationalShipmentResponseRate.md) |  | [optional] 
**ShipmentId** | Pointer to **string** | A unique identifier associated with the Shipment. | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentOptions**](ShipmentOptions.md) |  | [optional] 
**ToAddress** | Pointer to [**ReprintShipmentToAddress**](ReprintShipmentToAddress.md) |  | [optional] 
**Customs** | Pointer to [**InternationalShipmentResponseCustoms**](InternationalShipmentResponseCustoms.md) |  | [optional] 

## Methods

### NewCreateShipment200Response

`func NewCreateShipment200Response() *CreateShipment200Response`

NewCreateShipment200Response instantiates a new CreateShipment200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShipment200ResponseWithDefaults

`func NewCreateShipment200ResponseWithDefaults() *CreateShipment200Response`

NewCreateShipment200ResponseWithDefaults instantiates a new CreateShipment200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *CreateShipment200Response) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *CreateShipment200Response) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *CreateShipment200Response) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *CreateShipment200Response) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetLabelLayout

`func (o *CreateShipment200Response) GetLabelLayout() []DomesticShipmentResponseLabelLayoutInner`

GetLabelLayout returns the LabelLayout field if non-nil, zero value otherwise.

### GetLabelLayoutOk

`func (o *CreateShipment200Response) GetLabelLayoutOk() (*[]DomesticShipmentResponseLabelLayoutInner, bool)`

GetLabelLayoutOk returns a tuple with the LabelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelLayout

`func (o *CreateShipment200Response) SetLabelLayout(v []DomesticShipmentResponseLabelLayoutInner)`

SetLabelLayout sets LabelLayout field to given value.

### HasLabelLayout

`func (o *CreateShipment200Response) HasLabelLayout() bool`

HasLabelLayout returns a boolean if a field has been set.

### GetFromAddress

`func (o *CreateShipment200Response) GetFromAddress() ReprintShipmentFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *CreateShipment200Response) GetFromAddressOk() (*ReprintShipmentFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *CreateShipment200Response) SetFromAddress(v ReprintShipmentFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *CreateShipment200Response) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcel

`func (o *CreateShipment200Response) GetParcel() ShipmentDomesticParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *CreateShipment200Response) GetParcelOk() (*ShipmentDomesticParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *CreateShipment200Response) SetParcel(v ShipmentDomesticParcel)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *CreateShipment200Response) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetParcelId

`func (o *CreateShipment200Response) GetParcelId() string`

GetParcelId returns the ParcelId field if non-nil, zero value otherwise.

### GetParcelIdOk

`func (o *CreateShipment200Response) GetParcelIdOk() (*string, bool)`

GetParcelIdOk returns a tuple with the ParcelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelId

`func (o *CreateShipment200Response) SetParcelId(v string)`

SetParcelId sets ParcelId field to given value.

### HasParcelId

`func (o *CreateShipment200Response) HasParcelId() bool`

HasParcelId returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *CreateShipment200Response) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *CreateShipment200Response) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *CreateShipment200Response) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *CreateShipment200Response) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetRate

`func (o *CreateShipment200Response) GetRate() InternationalShipmentResponseRate`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *CreateShipment200Response) GetRateOk() (*InternationalShipmentResponseRate, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *CreateShipment200Response) SetRate(v InternationalShipmentResponseRate)`

SetRate sets Rate field to given value.

### HasRate

`func (o *CreateShipment200Response) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetShipmentId

`func (o *CreateShipment200Response) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *CreateShipment200Response) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *CreateShipment200Response) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *CreateShipment200Response) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *CreateShipment200Response) GetShipmentOptions() ShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *CreateShipment200Response) GetShipmentOptionsOk() (*ShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *CreateShipment200Response) SetShipmentOptions(v ShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *CreateShipment200Response) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetToAddress

`func (o *CreateShipment200Response) GetToAddress() ReprintShipmentToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *CreateShipment200Response) GetToAddressOk() (*ReprintShipmentToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *CreateShipment200Response) SetToAddress(v ReprintShipmentToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *CreateShipment200Response) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetCustoms

`func (o *CreateShipment200Response) GetCustoms() InternationalShipmentResponseCustoms`

GetCustoms returns the Customs field if non-nil, zero value otherwise.

### GetCustomsOk

`func (o *CreateShipment200Response) GetCustomsOk() (*InternationalShipmentResponseCustoms, bool)`

GetCustomsOk returns a tuple with the Customs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustoms

`func (o *CreateShipment200Response) SetCustoms(v InternationalShipmentResponseCustoms)`

SetCustoms sets Customs field to given value.

### HasCustoms

`func (o *CreateShipment200Response) HasCustoms() bool`

HasCustoms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


