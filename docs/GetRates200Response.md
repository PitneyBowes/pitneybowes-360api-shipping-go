# GetRates200Response

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

### NewGetRates200Response

`func NewGetRates200Response() *GetRates200Response`

NewGetRates200Response instantiates a new GetRates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRates200ResponseWithDefaults

`func NewGetRates200ResponseWithDefaults() *GetRates200Response`

NewGetRates200ResponseWithDefaults instantiates a new GetRates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *GetRates200Response) GetFromAddress() RateShopResponseFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *GetRates200Response) GetFromAddressOk() (*RateShopResponseFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *GetRates200Response) SetFromAddress(v RateShopResponseFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *GetRates200Response) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcel

`func (o *GetRates200Response) GetParcel() RateShopResponseParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *GetRates200Response) GetParcelOk() (*RateShopResponseParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *GetRates200Response) SetParcel(v RateShopResponseParcel)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *GetRates200Response) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetRates

`func (o *GetRates200Response) GetRates() []RateShopResponseRatesInner`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *GetRates200Response) GetRatesOk() (*[]RateShopResponseRatesInner, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *GetRates200Response) SetRates(v []RateShopResponseRatesInner)`

SetRates sets Rates field to given value.

### HasRates

`func (o *GetRates200Response) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetToAddress

`func (o *GetRates200Response) GetToAddress() RateShopResponseToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *GetRates200Response) GetToAddressOk() (*RateShopResponseToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *GetRates200Response) SetToAddress(v RateShopResponseToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *GetRates200Response) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetIsHazmat

`func (o *GetRates200Response) GetIsHazmat() bool`

GetIsHazmat returns the IsHazmat field if non-nil, zero value otherwise.

### GetIsHazmatOk

`func (o *GetRates200Response) GetIsHazmatOk() (*bool, bool)`

GetIsHazmatOk returns a tuple with the IsHazmat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHazmat

`func (o *GetRates200Response) SetIsHazmat(v bool)`

SetIsHazmat sets IsHazmat field to given value.

### HasIsHazmat

`func (o *GetRates200Response) HasIsHazmat() bool`

HasIsHazmat returns a boolean if a field has been set.

### GetErrors

`func (o *GetRates200Response) GetErrors() []RateShopResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetRates200Response) GetErrorsOk() (*[]RateShopResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetRates200Response) SetErrors(v []RateShopResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetRates200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


