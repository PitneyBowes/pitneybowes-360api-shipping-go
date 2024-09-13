# GetAllShipmentsDataInnerRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCharge** | Pointer to **float32** | The base service charge is payable to the carrier, excluding special service charges. | [optional] 
**Carrier** | Pointer to **string** | Carrier is a service used to transport the parcels or couriers from one place to another. | [optional] 
**CurrencyCode** | Pointer to **string** | A three-character (all uppercase letter) symbol of a currency according to the international ISO standard. As a rule, the first two letters denote the name of the country, and the third letter, the name of the currency thereof.For example, for US - the currency is Dollars and code is USD. Similarly for Canada, the currencycode is CAD, and for India, it is INR.  | [optional] 
**InductionPostalCode** | Pointer to **string** | The postal code where a shipment or shipments are tendered to a carrier. This can be different from the Sender’s address. | [optional] 
**ParcelType** | Pointer to **string** | Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection. ParcelType have categories like Package, Envelopes, Paks, Boxes, Tube, etc. | [optional] 
**RateTypeId** | Pointer to **string** | Its value can be CONTRACT_RATES, COMMERCIAL or COMMERCIAL_BASE for USPS and COMMERCIAL for other carriers depending on the Pitney Bowes contract/subscription | [optional] 
**ServiceId** | Pointer to **string** | The unique identifier given to the carrier specific service. | [optional] 
**SpecialServices** | Pointer to [**[]GetAllShipmentsDataInnerRateSpecialServicesInner**](GetAllShipmentsDataInnerRateSpecialServicesInner.md) | This provides a carrier-service based special or extra sevice. | [optional] 
**TotalCarrierCharge** | Pointer to **float32** | The total amount payable to the carrier, including special service fees, surcharges, and any international taxes and duties, except as noted below: | [optional] 

## Methods

### NewGetAllShipmentsDataInnerRate

`func NewGetAllShipmentsDataInnerRate() *GetAllShipmentsDataInnerRate`

NewGetAllShipmentsDataInnerRate instantiates a new GetAllShipmentsDataInnerRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllShipmentsDataInnerRateWithDefaults

`func NewGetAllShipmentsDataInnerRateWithDefaults() *GetAllShipmentsDataInnerRate`

NewGetAllShipmentsDataInnerRateWithDefaults instantiates a new GetAllShipmentsDataInnerRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCharge

`func (o *GetAllShipmentsDataInnerRate) GetBaseCharge() float32`

GetBaseCharge returns the BaseCharge field if non-nil, zero value otherwise.

### GetBaseChargeOk

`func (o *GetAllShipmentsDataInnerRate) GetBaseChargeOk() (*float32, bool)`

GetBaseChargeOk returns a tuple with the BaseCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCharge

`func (o *GetAllShipmentsDataInnerRate) SetBaseCharge(v float32)`

SetBaseCharge sets BaseCharge field to given value.

### HasBaseCharge

`func (o *GetAllShipmentsDataInnerRate) HasBaseCharge() bool`

HasBaseCharge returns a boolean if a field has been set.

### GetCarrier

`func (o *GetAllShipmentsDataInnerRate) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *GetAllShipmentsDataInnerRate) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *GetAllShipmentsDataInnerRate) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *GetAllShipmentsDataInnerRate) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GetAllShipmentsDataInnerRate) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GetAllShipmentsDataInnerRate) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GetAllShipmentsDataInnerRate) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GetAllShipmentsDataInnerRate) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetInductionPostalCode

`func (o *GetAllShipmentsDataInnerRate) GetInductionPostalCode() string`

GetInductionPostalCode returns the InductionPostalCode field if non-nil, zero value otherwise.

### GetInductionPostalCodeOk

`func (o *GetAllShipmentsDataInnerRate) GetInductionPostalCodeOk() (*string, bool)`

GetInductionPostalCodeOk returns a tuple with the InductionPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInductionPostalCode

`func (o *GetAllShipmentsDataInnerRate) SetInductionPostalCode(v string)`

SetInductionPostalCode sets InductionPostalCode field to given value.

### HasInductionPostalCode

`func (o *GetAllShipmentsDataInnerRate) HasInductionPostalCode() bool`

HasInductionPostalCode returns a boolean if a field has been set.

### GetParcelType

`func (o *GetAllShipmentsDataInnerRate) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *GetAllShipmentsDataInnerRate) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *GetAllShipmentsDataInnerRate) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *GetAllShipmentsDataInnerRate) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetRateTypeId

`func (o *GetAllShipmentsDataInnerRate) GetRateTypeId() string`

GetRateTypeId returns the RateTypeId field if non-nil, zero value otherwise.

### GetRateTypeIdOk

`func (o *GetAllShipmentsDataInnerRate) GetRateTypeIdOk() (*string, bool)`

GetRateTypeIdOk returns a tuple with the RateTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTypeId

`func (o *GetAllShipmentsDataInnerRate) SetRateTypeId(v string)`

SetRateTypeId sets RateTypeId field to given value.

### HasRateTypeId

`func (o *GetAllShipmentsDataInnerRate) HasRateTypeId() bool`

HasRateTypeId returns a boolean if a field has been set.

### GetServiceId

`func (o *GetAllShipmentsDataInnerRate) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GetAllShipmentsDataInnerRate) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GetAllShipmentsDataInnerRate) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *GetAllShipmentsDataInnerRate) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSpecialServices

`func (o *GetAllShipmentsDataInnerRate) GetSpecialServices() []GetAllShipmentsDataInnerRateSpecialServicesInner`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *GetAllShipmentsDataInnerRate) GetSpecialServicesOk() (*[]GetAllShipmentsDataInnerRateSpecialServicesInner, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *GetAllShipmentsDataInnerRate) SetSpecialServices(v []GetAllShipmentsDataInnerRateSpecialServicesInner)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *GetAllShipmentsDataInnerRate) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetTotalCarrierCharge

`func (o *GetAllShipmentsDataInnerRate) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *GetAllShipmentsDataInnerRate) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *GetAllShipmentsDataInnerRate) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *GetAllShipmentsDataInnerRate) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


