# RateResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCharge** | Pointer to **float32** | The base service charge is payable to the carrier, excluding special service charges. | [optional] 
**Carrier** | Pointer to **string** | Carrier is a service used to transport the parcels or couriers from one place to another. | [optional] 
**CurrencyCode** | Pointer to **string** | A three-character (all uppercase letter) symbol of a currency according to the international ISO standard.&lt;br /&gt; As a rule, the first two letters denote the name of the country, and the third letter, the name of the currency thereof. For example, for US - the currency is Dollars and code is USD. Similarly for Canada, the currencycode is CAD, and for India, it is INR.  | [optional] 
**ParcelType** | Pointer to **string** | Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection.&lt;br /&gt; ParcelType have categories like Package, Envelopes, Paks, Boxes, Tube, etc. | [optional] 
**ServiceId** | Pointer to **string** | The unique identifier given to the carrier specific service. | [optional] 
**Surcharges** | Pointer to [**[]RateResponseV2SurchargesInner**](RateResponseV2SurchargesInner.md) | Additional fees or surcharges for the shipment. Each object in the array defines a surcharge and fee. | [optional] 
**TotalCarrierCharge** | Pointer to **float32** | The total amount payable to the carrier, including special service fees, surcharges, and any international taxes and duties, except as noted below: | [optional] 
**DeliveryCommitment** | Pointer to [**RateResponseV2DeliveryCommitment**](RateResponseV2DeliveryCommitment.md) |  | [optional] 

## Methods

### NewRateResponseV2

`func NewRateResponseV2() *RateResponseV2`

NewRateResponseV2 instantiates a new RateResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateResponseV2WithDefaults

`func NewRateResponseV2WithDefaults() *RateResponseV2`

NewRateResponseV2WithDefaults instantiates a new RateResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCharge

`func (o *RateResponseV2) GetBaseCharge() float32`

GetBaseCharge returns the BaseCharge field if non-nil, zero value otherwise.

### GetBaseChargeOk

`func (o *RateResponseV2) GetBaseChargeOk() (*float32, bool)`

GetBaseChargeOk returns a tuple with the BaseCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCharge

`func (o *RateResponseV2) SetBaseCharge(v float32)`

SetBaseCharge sets BaseCharge field to given value.

### HasBaseCharge

`func (o *RateResponseV2) HasBaseCharge() bool`

HasBaseCharge returns a boolean if a field has been set.

### GetCarrier

`func (o *RateResponseV2) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *RateResponseV2) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *RateResponseV2) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *RateResponseV2) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *RateResponseV2) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *RateResponseV2) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *RateResponseV2) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *RateResponseV2) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetParcelType

`func (o *RateResponseV2) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *RateResponseV2) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *RateResponseV2) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *RateResponseV2) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetServiceId

`func (o *RateResponseV2) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RateResponseV2) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RateResponseV2) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *RateResponseV2) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSurcharges

`func (o *RateResponseV2) GetSurcharges() []RateResponseV2SurchargesInner`

GetSurcharges returns the Surcharges field if non-nil, zero value otherwise.

### GetSurchargesOk

`func (o *RateResponseV2) GetSurchargesOk() (*[]RateResponseV2SurchargesInner, bool)`

GetSurchargesOk returns a tuple with the Surcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharges

`func (o *RateResponseV2) SetSurcharges(v []RateResponseV2SurchargesInner)`

SetSurcharges sets Surcharges field to given value.

### HasSurcharges

`func (o *RateResponseV2) HasSurcharges() bool`

HasSurcharges returns a boolean if a field has been set.

### GetTotalCarrierCharge

`func (o *RateResponseV2) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *RateResponseV2) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *RateResponseV2) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *RateResponseV2) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.

### GetDeliveryCommitment

`func (o *RateResponseV2) GetDeliveryCommitment() RateResponseV2DeliveryCommitment`

GetDeliveryCommitment returns the DeliveryCommitment field if non-nil, zero value otherwise.

### GetDeliveryCommitmentOk

`func (o *RateResponseV2) GetDeliveryCommitmentOk() (*RateResponseV2DeliveryCommitment, bool)`

GetDeliveryCommitmentOk returns a tuple with the DeliveryCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCommitment

`func (o *RateResponseV2) SetDeliveryCommitment(v RateResponseV2DeliveryCommitment)`

SetDeliveryCommitment sets DeliveryCommitment field to given value.

### HasDeliveryCommitment

`func (o *RateResponseV2) HasDeliveryCommitment() bool`

HasDeliveryCommitment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


