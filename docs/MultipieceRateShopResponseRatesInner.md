# MultipieceRateShopResponseRatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCarrierCharge** | Pointer to **float32** | description | [optional] 
**Carrier** | Pointer to **string** | description | [optional] 
**CurrencyCode** | Pointer to **string** | description | [optional] 
**RateTypeId** | Pointer to **string** | description | [optional] 
**ServiceId** | Pointer to **string** | description | [optional] 
**DeliveryCommitment** | Pointer to [**RateShopResponseRatesInnerDeliveryCommitment**](RateShopResponseRatesInnerDeliveryCommitment.md) |  | [optional] 
**MultiPieceParcels** | Pointer to [**[]MultipieceRateShopResponseRatesInnerMultiPieceParcelsInner**](MultipieceRateShopResponseRatesInnerMultiPieceParcelsInner.md) | description | [optional] 
**Surcharges** | Pointer to [**[]MultipieceRatesResponseRatesInnerSurchargesInner**](MultipieceRatesResponseRatesInnerSurchargesInner.md) | description | [optional] 

## Methods

### NewMultipieceRateShopResponseRatesInner

`func NewMultipieceRateShopResponseRatesInner() *MultipieceRateShopResponseRatesInner`

NewMultipieceRateShopResponseRatesInner instantiates a new MultipieceRateShopResponseRatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceRateShopResponseRatesInnerWithDefaults

`func NewMultipieceRateShopResponseRatesInnerWithDefaults() *MultipieceRateShopResponseRatesInner`

NewMultipieceRateShopResponseRatesInnerWithDefaults instantiates a new MultipieceRateShopResponseRatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCarrierCharge

`func (o *MultipieceRateShopResponseRatesInner) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *MultipieceRateShopResponseRatesInner) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *MultipieceRateShopResponseRatesInner) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *MultipieceRateShopResponseRatesInner) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.

### GetCarrier

`func (o *MultipieceRateShopResponseRatesInner) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *MultipieceRateShopResponseRatesInner) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *MultipieceRateShopResponseRatesInner) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *MultipieceRateShopResponseRatesInner) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *MultipieceRateShopResponseRatesInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MultipieceRateShopResponseRatesInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MultipieceRateShopResponseRatesInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *MultipieceRateShopResponseRatesInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetRateTypeId

`func (o *MultipieceRateShopResponseRatesInner) GetRateTypeId() string`

GetRateTypeId returns the RateTypeId field if non-nil, zero value otherwise.

### GetRateTypeIdOk

`func (o *MultipieceRateShopResponseRatesInner) GetRateTypeIdOk() (*string, bool)`

GetRateTypeIdOk returns a tuple with the RateTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTypeId

`func (o *MultipieceRateShopResponseRatesInner) SetRateTypeId(v string)`

SetRateTypeId sets RateTypeId field to given value.

### HasRateTypeId

`func (o *MultipieceRateShopResponseRatesInner) HasRateTypeId() bool`

HasRateTypeId returns a boolean if a field has been set.

### GetServiceId

`func (o *MultipieceRateShopResponseRatesInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MultipieceRateShopResponseRatesInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MultipieceRateShopResponseRatesInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MultipieceRateShopResponseRatesInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetDeliveryCommitment

`func (o *MultipieceRateShopResponseRatesInner) GetDeliveryCommitment() RateShopResponseRatesInnerDeliveryCommitment`

GetDeliveryCommitment returns the DeliveryCommitment field if non-nil, zero value otherwise.

### GetDeliveryCommitmentOk

`func (o *MultipieceRateShopResponseRatesInner) GetDeliveryCommitmentOk() (*RateShopResponseRatesInnerDeliveryCommitment, bool)`

GetDeliveryCommitmentOk returns a tuple with the DeliveryCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCommitment

`func (o *MultipieceRateShopResponseRatesInner) SetDeliveryCommitment(v RateShopResponseRatesInnerDeliveryCommitment)`

SetDeliveryCommitment sets DeliveryCommitment field to given value.

### HasDeliveryCommitment

`func (o *MultipieceRateShopResponseRatesInner) HasDeliveryCommitment() bool`

HasDeliveryCommitment returns a boolean if a field has been set.

### GetMultiPieceParcels

`func (o *MultipieceRateShopResponseRatesInner) GetMultiPieceParcels() []MultipieceRateShopResponseRatesInnerMultiPieceParcelsInner`

GetMultiPieceParcels returns the MultiPieceParcels field if non-nil, zero value otherwise.

### GetMultiPieceParcelsOk

`func (o *MultipieceRateShopResponseRatesInner) GetMultiPieceParcelsOk() (*[]MultipieceRateShopResponseRatesInnerMultiPieceParcelsInner, bool)`

GetMultiPieceParcelsOk returns a tuple with the MultiPieceParcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPieceParcels

`func (o *MultipieceRateShopResponseRatesInner) SetMultiPieceParcels(v []MultipieceRateShopResponseRatesInnerMultiPieceParcelsInner)`

SetMultiPieceParcels sets MultiPieceParcels field to given value.

### HasMultiPieceParcels

`func (o *MultipieceRateShopResponseRatesInner) HasMultiPieceParcels() bool`

HasMultiPieceParcels returns a boolean if a field has been set.

### GetSurcharges

`func (o *MultipieceRateShopResponseRatesInner) GetSurcharges() []MultipieceRatesResponseRatesInnerSurchargesInner`

GetSurcharges returns the Surcharges field if non-nil, zero value otherwise.

### GetSurchargesOk

`func (o *MultipieceRateShopResponseRatesInner) GetSurchargesOk() (*[]MultipieceRatesResponseRatesInnerSurchargesInner, bool)`

GetSurchargesOk returns a tuple with the Surcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharges

`func (o *MultipieceRateShopResponseRatesInner) SetSurcharges(v []MultipieceRatesResponseRatesInnerSurchargesInner)`

SetSurcharges sets Surcharges field to given value.

### HasSurcharges

`func (o *MultipieceRateShopResponseRatesInner) HasSurcharges() bool`

HasSurcharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


