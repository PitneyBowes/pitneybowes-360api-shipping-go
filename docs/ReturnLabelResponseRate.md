# ReturnLabelResponseRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCharge** | Pointer to **float32** | Description | [optional] 
**Carrier** | Pointer to **string** | Description | [optional] 
**CurrencyCode** | Pointer to **string** | Description | [optional] 
**ParcelType** | Pointer to **string** | Description | [optional] 
**RateTypeId** | Pointer to **string** | Description | [optional] 
**ServiceId** | Pointer to **string** | Description | [optional] 
**SpecialServices** | Pointer to [**[]ReturnLabelResponseRateSpecialServicesInner**](ReturnLabelResponseRateSpecialServicesInner.md) |  | [optional] 
**Surcharges** | Pointer to [**[]ReturnLabelResponseRateSurchargesInner**](ReturnLabelResponseRateSurchargesInner.md) |  | [optional] 
**TotalCarrierCharge** | Pointer to **float32** | Description | [optional] 

## Methods

### NewReturnLabelResponseRate

`func NewReturnLabelResponseRate() *ReturnLabelResponseRate`

NewReturnLabelResponseRate instantiates a new ReturnLabelResponseRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnLabelResponseRateWithDefaults

`func NewReturnLabelResponseRateWithDefaults() *ReturnLabelResponseRate`

NewReturnLabelResponseRateWithDefaults instantiates a new ReturnLabelResponseRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCharge

`func (o *ReturnLabelResponseRate) GetBaseCharge() float32`

GetBaseCharge returns the BaseCharge field if non-nil, zero value otherwise.

### GetBaseChargeOk

`func (o *ReturnLabelResponseRate) GetBaseChargeOk() (*float32, bool)`

GetBaseChargeOk returns a tuple with the BaseCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCharge

`func (o *ReturnLabelResponseRate) SetBaseCharge(v float32)`

SetBaseCharge sets BaseCharge field to given value.

### HasBaseCharge

`func (o *ReturnLabelResponseRate) HasBaseCharge() bool`

HasBaseCharge returns a boolean if a field has been set.

### GetCarrier

`func (o *ReturnLabelResponseRate) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *ReturnLabelResponseRate) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *ReturnLabelResponseRate) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *ReturnLabelResponseRate) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ReturnLabelResponseRate) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ReturnLabelResponseRate) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ReturnLabelResponseRate) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ReturnLabelResponseRate) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetParcelType

`func (o *ReturnLabelResponseRate) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *ReturnLabelResponseRate) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *ReturnLabelResponseRate) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *ReturnLabelResponseRate) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetRateTypeId

`func (o *ReturnLabelResponseRate) GetRateTypeId() string`

GetRateTypeId returns the RateTypeId field if non-nil, zero value otherwise.

### GetRateTypeIdOk

`func (o *ReturnLabelResponseRate) GetRateTypeIdOk() (*string, bool)`

GetRateTypeIdOk returns a tuple with the RateTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTypeId

`func (o *ReturnLabelResponseRate) SetRateTypeId(v string)`

SetRateTypeId sets RateTypeId field to given value.

### HasRateTypeId

`func (o *ReturnLabelResponseRate) HasRateTypeId() bool`

HasRateTypeId returns a boolean if a field has been set.

### GetServiceId

`func (o *ReturnLabelResponseRate) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ReturnLabelResponseRate) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ReturnLabelResponseRate) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ReturnLabelResponseRate) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSpecialServices

`func (o *ReturnLabelResponseRate) GetSpecialServices() []ReturnLabelResponseRateSpecialServicesInner`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *ReturnLabelResponseRate) GetSpecialServicesOk() (*[]ReturnLabelResponseRateSpecialServicesInner, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *ReturnLabelResponseRate) SetSpecialServices(v []ReturnLabelResponseRateSpecialServicesInner)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *ReturnLabelResponseRate) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetSurcharges

`func (o *ReturnLabelResponseRate) GetSurcharges() []ReturnLabelResponseRateSurchargesInner`

GetSurcharges returns the Surcharges field if non-nil, zero value otherwise.

### GetSurchargesOk

`func (o *ReturnLabelResponseRate) GetSurchargesOk() (*[]ReturnLabelResponseRateSurchargesInner, bool)`

GetSurchargesOk returns a tuple with the Surcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharges

`func (o *ReturnLabelResponseRate) SetSurcharges(v []ReturnLabelResponseRateSurchargesInner)`

SetSurcharges sets Surcharges field to given value.

### HasSurcharges

`func (o *ReturnLabelResponseRate) HasSurcharges() bool`

HasSurcharges returns a boolean if a field has been set.

### GetTotalCarrierCharge

`func (o *ReturnLabelResponseRate) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *ReturnLabelResponseRate) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *ReturnLabelResponseRate) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *ReturnLabelResponseRate) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


