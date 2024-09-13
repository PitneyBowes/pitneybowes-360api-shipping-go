# RateShopResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | Pointer to [**RateShopResponseFromAddress**](RateShopResponseFromAddress.md) |  | [optional] 
**Parcel** | Pointer to [**RateShopResponseParcel**](RateShopResponseParcel.md) |  | [optional] 
**Rates** | Pointer to [**[]RateShopResponseRatesInner**](RateShopResponseRatesInner.md) | It displays all available rates for each services | [optional] 
**ToAddress** | Pointer to [**RateShopResponseToAddress**](RateShopResponseToAddress.md) |  | [optional] 
**IsHazmat** | Pointer to **bool** | isHazmat if set to true,only services which support Hazardous material shipment would be visible in the response | [optional] 
**Errors** | Pointer to [**[]RateShopResponseErrorsInner**](RateShopResponseErrorsInner.md) | It display any error while getting rates | [optional] 

## Methods

### NewRateShopResponse

`func NewRateShopResponse() *RateShopResponse`

NewRateShopResponse instantiates a new RateShopResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateShopResponseWithDefaults

`func NewRateShopResponseWithDefaults() *RateShopResponse`

NewRateShopResponseWithDefaults instantiates a new RateShopResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *RateShopResponse) GetFromAddress() RateShopResponseFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *RateShopResponse) GetFromAddressOk() (*RateShopResponseFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *RateShopResponse) SetFromAddress(v RateShopResponseFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *RateShopResponse) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcel

`func (o *RateShopResponse) GetParcel() RateShopResponseParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *RateShopResponse) GetParcelOk() (*RateShopResponseParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *RateShopResponse) SetParcel(v RateShopResponseParcel)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *RateShopResponse) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetRates

`func (o *RateShopResponse) GetRates() []RateShopResponseRatesInner`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *RateShopResponse) GetRatesOk() (*[]RateShopResponseRatesInner, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *RateShopResponse) SetRates(v []RateShopResponseRatesInner)`

SetRates sets Rates field to given value.

### HasRates

`func (o *RateShopResponse) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetToAddress

`func (o *RateShopResponse) GetToAddress() RateShopResponseToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *RateShopResponse) GetToAddressOk() (*RateShopResponseToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *RateShopResponse) SetToAddress(v RateShopResponseToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *RateShopResponse) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetIsHazmat

`func (o *RateShopResponse) GetIsHazmat() bool`

GetIsHazmat returns the IsHazmat field if non-nil, zero value otherwise.

### GetIsHazmatOk

`func (o *RateShopResponse) GetIsHazmatOk() (*bool, bool)`

GetIsHazmatOk returns a tuple with the IsHazmat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHazmat

`func (o *RateShopResponse) SetIsHazmat(v bool)`

SetIsHazmat sets IsHazmat field to given value.

### HasIsHazmat

`func (o *RateShopResponse) HasIsHazmat() bool`

HasIsHazmat returns a boolean if a field has been set.

### GetErrors

`func (o *RateShopResponse) GetErrors() []RateShopResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RateShopResponse) GetErrorsOk() (*[]RateShopResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RateShopResponse) SetErrors(v []RateShopResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *RateShopResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


