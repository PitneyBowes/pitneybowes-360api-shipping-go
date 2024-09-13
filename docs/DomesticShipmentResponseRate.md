# DomesticShipmentResponseRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCharge** | Pointer to **float32** | The base service charge is payable to the carrier, excluding special service charges. | [optional] 
**Carrier** | Pointer to **string** | Carrier is a service used to transport the parcels or couriers from one place to another. | [optional] 
**CurrencyCode** | Pointer to **string** | A three-character (all uppercase letter) symbol of a currency according to the international ISO standard. As a rule, the first two letters denote the name of the country, and the third letter, the name of the currency thereof.For example, for US - the currency is Dollars and code is USD. Similarly for Canada, the currencycode is CAD, and for India, it is INR.  | [optional] 
**ParcelType** | Pointer to **string** | Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection. ParcelType have categories like Package, Envelopes, Paks, Boxes, Tube, etc.  | [optional] 
**RateTypeId** | Pointer to **string** | Its value can be CONTRACT_RATES, COMMERCIAL or COMMERCIAL_BASE for USPS and COMMERCIAL for other carriers depending on the Pitney Bowes contract/subscription | [optional] 
**ServiceId** | Pointer to **string** | The unique identifier given to the carrier specific service. | [optional] 
**SpecialServices** | Pointer to [**[]DomesticShipmentResponseRateSpecialServicesInner**](DomesticShipmentResponseRateSpecialServicesInner.md) | This provides a carrier-service based special or extra sevice. | [optional] 
**TotalCarrierCharge** | Pointer to **float32** | The total amount payable to the carrier, including special service fees, surcharges, and any international taxes and duties, except as noted below: | [optional] 

## Methods

### NewDomesticShipmentResponseRate

`func NewDomesticShipmentResponseRate() *DomesticShipmentResponseRate`

NewDomesticShipmentResponseRate instantiates a new DomesticShipmentResponseRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomesticShipmentResponseRateWithDefaults

`func NewDomesticShipmentResponseRateWithDefaults() *DomesticShipmentResponseRate`

NewDomesticShipmentResponseRateWithDefaults instantiates a new DomesticShipmentResponseRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCharge

`func (o *DomesticShipmentResponseRate) GetBaseCharge() float32`

GetBaseCharge returns the BaseCharge field if non-nil, zero value otherwise.

### GetBaseChargeOk

`func (o *DomesticShipmentResponseRate) GetBaseChargeOk() (*float32, bool)`

GetBaseChargeOk returns a tuple with the BaseCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCharge

`func (o *DomesticShipmentResponseRate) SetBaseCharge(v float32)`

SetBaseCharge sets BaseCharge field to given value.

### HasBaseCharge

`func (o *DomesticShipmentResponseRate) HasBaseCharge() bool`

HasBaseCharge returns a boolean if a field has been set.

### GetCarrier

`func (o *DomesticShipmentResponseRate) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *DomesticShipmentResponseRate) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *DomesticShipmentResponseRate) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *DomesticShipmentResponseRate) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *DomesticShipmentResponseRate) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *DomesticShipmentResponseRate) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *DomesticShipmentResponseRate) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *DomesticShipmentResponseRate) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetParcelType

`func (o *DomesticShipmentResponseRate) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *DomesticShipmentResponseRate) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *DomesticShipmentResponseRate) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *DomesticShipmentResponseRate) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetRateTypeId

`func (o *DomesticShipmentResponseRate) GetRateTypeId() string`

GetRateTypeId returns the RateTypeId field if non-nil, zero value otherwise.

### GetRateTypeIdOk

`func (o *DomesticShipmentResponseRate) GetRateTypeIdOk() (*string, bool)`

GetRateTypeIdOk returns a tuple with the RateTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTypeId

`func (o *DomesticShipmentResponseRate) SetRateTypeId(v string)`

SetRateTypeId sets RateTypeId field to given value.

### HasRateTypeId

`func (o *DomesticShipmentResponseRate) HasRateTypeId() bool`

HasRateTypeId returns a boolean if a field has been set.

### GetServiceId

`func (o *DomesticShipmentResponseRate) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DomesticShipmentResponseRate) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DomesticShipmentResponseRate) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DomesticShipmentResponseRate) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSpecialServices

`func (o *DomesticShipmentResponseRate) GetSpecialServices() []DomesticShipmentResponseRateSpecialServicesInner`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *DomesticShipmentResponseRate) GetSpecialServicesOk() (*[]DomesticShipmentResponseRateSpecialServicesInner, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *DomesticShipmentResponseRate) SetSpecialServices(v []DomesticShipmentResponseRateSpecialServicesInner)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *DomesticShipmentResponseRate) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetTotalCarrierCharge

`func (o *DomesticShipmentResponseRate) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *DomesticShipmentResponseRate) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *DomesticShipmentResponseRate) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *DomesticShipmentResponseRate) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


