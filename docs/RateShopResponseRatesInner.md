# RateShopResponseRatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCharge** | Pointer to **float32** | The base service charge is payable to the carrier, excluding special service charges. | [optional] 
**Carrier** | Pointer to **string** | Carrier is a service used to transport the parcels or couriers from one place to another. | [optional] 
**CarrierAccount** | Pointer to **string** | Carrier Account is the account associated to specific carrier service. | [optional] 
**DeliveryCommitment** | Pointer to [**RateShopResponseRatesInnerDeliveryCommitment**](RateShopResponseRatesInnerDeliveryCommitment.md) |  | [optional] 
**CurrencyCode** | Pointer to **string** | A three-character (all uppercase letter) symbol of a currency according to the international ISO standard. As a rule, the first two letters denote the name of the country, and the third letter, the name of the currency thereof.For example, for US - the currency is Dollars and code is USD. Similarly for Canada, the currencycode is CAD, and for India, it is INR.  | [optional] 
**ParcelType** | Pointer to **string** | Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection. ParcelType have categories like Package, Envelopes, Paks, Boxes, Tube, etc.  | [optional] 
**ParcelId** | Pointer to **string** | It indicates parcelId chosen in the request | [optional] 
**RateTypeId** | Pointer to **string** | Its value can be CONTRACT_RATES, COMMERCIAL or COMMERCIAL_BASE for USPS and COMMERCIAL for other carriers depending on the Pitney Bowes contract/subscription | [optional] 
**ServiceId** | Pointer to **string** | The unique identifier given to the carrier specific service. | [optional] 
**Surcharges** | Pointer to [**[]RateShopResponseRatesInnerSurchargesInner**](RateShopResponseRatesInnerSurchargesInner.md) |  Additional fees or surcharges for the shipment. Each object in the array defines a surcharge and fee. | [optional] 
**TotalCarrierCharge** | Pointer to **float32** | The total amount payable to the carrier, including special service fees, surcharges, and any international taxes and duties | [optional] 

## Methods

### NewRateShopResponseRatesInner

`func NewRateShopResponseRatesInner() *RateShopResponseRatesInner`

NewRateShopResponseRatesInner instantiates a new RateShopResponseRatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateShopResponseRatesInnerWithDefaults

`func NewRateShopResponseRatesInnerWithDefaults() *RateShopResponseRatesInner`

NewRateShopResponseRatesInnerWithDefaults instantiates a new RateShopResponseRatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCharge

`func (o *RateShopResponseRatesInner) GetBaseCharge() float32`

GetBaseCharge returns the BaseCharge field if non-nil, zero value otherwise.

### GetBaseChargeOk

`func (o *RateShopResponseRatesInner) GetBaseChargeOk() (*float32, bool)`

GetBaseChargeOk returns a tuple with the BaseCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCharge

`func (o *RateShopResponseRatesInner) SetBaseCharge(v float32)`

SetBaseCharge sets BaseCharge field to given value.

### HasBaseCharge

`func (o *RateShopResponseRatesInner) HasBaseCharge() bool`

HasBaseCharge returns a boolean if a field has been set.

### GetCarrier

`func (o *RateShopResponseRatesInner) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *RateShopResponseRatesInner) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *RateShopResponseRatesInner) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *RateShopResponseRatesInner) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierAccount

`func (o *RateShopResponseRatesInner) GetCarrierAccount() string`

GetCarrierAccount returns the CarrierAccount field if non-nil, zero value otherwise.

### GetCarrierAccountOk

`func (o *RateShopResponseRatesInner) GetCarrierAccountOk() (*string, bool)`

GetCarrierAccountOk returns a tuple with the CarrierAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccount

`func (o *RateShopResponseRatesInner) SetCarrierAccount(v string)`

SetCarrierAccount sets CarrierAccount field to given value.

### HasCarrierAccount

`func (o *RateShopResponseRatesInner) HasCarrierAccount() bool`

HasCarrierAccount returns a boolean if a field has been set.

### GetDeliveryCommitment

`func (o *RateShopResponseRatesInner) GetDeliveryCommitment() RateShopResponseRatesInnerDeliveryCommitment`

GetDeliveryCommitment returns the DeliveryCommitment field if non-nil, zero value otherwise.

### GetDeliveryCommitmentOk

`func (o *RateShopResponseRatesInner) GetDeliveryCommitmentOk() (*RateShopResponseRatesInnerDeliveryCommitment, bool)`

GetDeliveryCommitmentOk returns a tuple with the DeliveryCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCommitment

`func (o *RateShopResponseRatesInner) SetDeliveryCommitment(v RateShopResponseRatesInnerDeliveryCommitment)`

SetDeliveryCommitment sets DeliveryCommitment field to given value.

### HasDeliveryCommitment

`func (o *RateShopResponseRatesInner) HasDeliveryCommitment() bool`

HasDeliveryCommitment returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *RateShopResponseRatesInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *RateShopResponseRatesInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *RateShopResponseRatesInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *RateShopResponseRatesInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetParcelType

`func (o *RateShopResponseRatesInner) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *RateShopResponseRatesInner) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *RateShopResponseRatesInner) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *RateShopResponseRatesInner) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetParcelId

`func (o *RateShopResponseRatesInner) GetParcelId() string`

GetParcelId returns the ParcelId field if non-nil, zero value otherwise.

### GetParcelIdOk

`func (o *RateShopResponseRatesInner) GetParcelIdOk() (*string, bool)`

GetParcelIdOk returns a tuple with the ParcelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelId

`func (o *RateShopResponseRatesInner) SetParcelId(v string)`

SetParcelId sets ParcelId field to given value.

### HasParcelId

`func (o *RateShopResponseRatesInner) HasParcelId() bool`

HasParcelId returns a boolean if a field has been set.

### GetRateTypeId

`func (o *RateShopResponseRatesInner) GetRateTypeId() string`

GetRateTypeId returns the RateTypeId field if non-nil, zero value otherwise.

### GetRateTypeIdOk

`func (o *RateShopResponseRatesInner) GetRateTypeIdOk() (*string, bool)`

GetRateTypeIdOk returns a tuple with the RateTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTypeId

`func (o *RateShopResponseRatesInner) SetRateTypeId(v string)`

SetRateTypeId sets RateTypeId field to given value.

### HasRateTypeId

`func (o *RateShopResponseRatesInner) HasRateTypeId() bool`

HasRateTypeId returns a boolean if a field has been set.

### GetServiceId

`func (o *RateShopResponseRatesInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RateShopResponseRatesInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RateShopResponseRatesInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *RateShopResponseRatesInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSurcharges

`func (o *RateShopResponseRatesInner) GetSurcharges() []RateShopResponseRatesInnerSurchargesInner`

GetSurcharges returns the Surcharges field if non-nil, zero value otherwise.

### GetSurchargesOk

`func (o *RateShopResponseRatesInner) GetSurchargesOk() (*[]RateShopResponseRatesInnerSurchargesInner, bool)`

GetSurchargesOk returns a tuple with the Surcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharges

`func (o *RateShopResponseRatesInner) SetSurcharges(v []RateShopResponseRatesInnerSurchargesInner)`

SetSurcharges sets Surcharges field to given value.

### HasSurcharges

`func (o *RateShopResponseRatesInner) HasSurcharges() bool`

HasSurcharges returns a boolean if a field has been set.

### GetTotalCarrierCharge

`func (o *RateShopResponseRatesInner) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *RateShopResponseRatesInner) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *RateShopResponseRatesInner) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *RateShopResponseRatesInner) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


