# MultipieceRatesResponseRatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCarrierCharge** | Pointer to **float32** | description | [optional] 
**Carrier** | Pointer to **string** | description | [optional] 
**CurrencyCode** | Pointer to **string** | description | [optional] 
**RateTypeId** | Pointer to **string** | description | [optional] 
**ServiceId** | Pointer to **string** | description | [optional] 
**MultiPieceParcels** | Pointer to [**[]MultipieceRatesResponseRatesInnerMultiPieceParcelsInner**](MultipieceRatesResponseRatesInnerMultiPieceParcelsInner.md) | description | [optional] 
**Surcharges** | Pointer to [**[]MultipieceRatesResponseRatesInnerSurchargesInner**](MultipieceRatesResponseRatesInnerSurchargesInner.md) | description | [optional] 

## Methods

### NewMultipieceRatesResponseRatesInner

`func NewMultipieceRatesResponseRatesInner() *MultipieceRatesResponseRatesInner`

NewMultipieceRatesResponseRatesInner instantiates a new MultipieceRatesResponseRatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceRatesResponseRatesInnerWithDefaults

`func NewMultipieceRatesResponseRatesInnerWithDefaults() *MultipieceRatesResponseRatesInner`

NewMultipieceRatesResponseRatesInnerWithDefaults instantiates a new MultipieceRatesResponseRatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCarrierCharge

`func (o *MultipieceRatesResponseRatesInner) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *MultipieceRatesResponseRatesInner) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *MultipieceRatesResponseRatesInner) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *MultipieceRatesResponseRatesInner) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.

### GetCarrier

`func (o *MultipieceRatesResponseRatesInner) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *MultipieceRatesResponseRatesInner) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *MultipieceRatesResponseRatesInner) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *MultipieceRatesResponseRatesInner) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *MultipieceRatesResponseRatesInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MultipieceRatesResponseRatesInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MultipieceRatesResponseRatesInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *MultipieceRatesResponseRatesInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetRateTypeId

`func (o *MultipieceRatesResponseRatesInner) GetRateTypeId() string`

GetRateTypeId returns the RateTypeId field if non-nil, zero value otherwise.

### GetRateTypeIdOk

`func (o *MultipieceRatesResponseRatesInner) GetRateTypeIdOk() (*string, bool)`

GetRateTypeIdOk returns a tuple with the RateTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTypeId

`func (o *MultipieceRatesResponseRatesInner) SetRateTypeId(v string)`

SetRateTypeId sets RateTypeId field to given value.

### HasRateTypeId

`func (o *MultipieceRatesResponseRatesInner) HasRateTypeId() bool`

HasRateTypeId returns a boolean if a field has been set.

### GetServiceId

`func (o *MultipieceRatesResponseRatesInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MultipieceRatesResponseRatesInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MultipieceRatesResponseRatesInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MultipieceRatesResponseRatesInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetMultiPieceParcels

`func (o *MultipieceRatesResponseRatesInner) GetMultiPieceParcels() []MultipieceRatesResponseRatesInnerMultiPieceParcelsInner`

GetMultiPieceParcels returns the MultiPieceParcels field if non-nil, zero value otherwise.

### GetMultiPieceParcelsOk

`func (o *MultipieceRatesResponseRatesInner) GetMultiPieceParcelsOk() (*[]MultipieceRatesResponseRatesInnerMultiPieceParcelsInner, bool)`

GetMultiPieceParcelsOk returns a tuple with the MultiPieceParcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPieceParcels

`func (o *MultipieceRatesResponseRatesInner) SetMultiPieceParcels(v []MultipieceRatesResponseRatesInnerMultiPieceParcelsInner)`

SetMultiPieceParcels sets MultiPieceParcels field to given value.

### HasMultiPieceParcels

`func (o *MultipieceRatesResponseRatesInner) HasMultiPieceParcels() bool`

HasMultiPieceParcels returns a boolean if a field has been set.

### GetSurcharges

`func (o *MultipieceRatesResponseRatesInner) GetSurcharges() []MultipieceRatesResponseRatesInnerSurchargesInner`

GetSurcharges returns the Surcharges field if non-nil, zero value otherwise.

### GetSurchargesOk

`func (o *MultipieceRatesResponseRatesInner) GetSurchargesOk() (*[]MultipieceRatesResponseRatesInnerSurchargesInner, bool)`

GetSurchargesOk returns a tuple with the Surcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharges

`func (o *MultipieceRatesResponseRatesInner) SetSurcharges(v []MultipieceRatesResponseRatesInnerSurchargesInner)`

SetSurcharges sets Surcharges field to given value.

### HasSurcharges

`func (o *MultipieceRatesResponseRatesInner) HasSurcharges() bool`

HasSurcharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


