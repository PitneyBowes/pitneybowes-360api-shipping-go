# SingleRateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | Pointer to [**SingleRateResponseFromAddress**](SingleRateResponseFromAddress.md) |  | [optional] 
**Parcel** | Pointer to [**SingleRateResponseParcel**](SingleRateResponseParcel.md) |  | [optional] 
**Rates** | Pointer to [**[]SingleRateResponseRatesInner**](SingleRateResponseRatesInner.md) | It displays the rate for the required carrier-service and/or special service combinaion | [optional] 
**ToAddress** | Pointer to [**SingleRateResponseToAddress**](SingleRateResponseToAddress.md) |  | [optional] 

## Methods

### NewSingleRateResponse

`func NewSingleRateResponse() *SingleRateResponse`

NewSingleRateResponse instantiates a new SingleRateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleRateResponseWithDefaults

`func NewSingleRateResponseWithDefaults() *SingleRateResponse`

NewSingleRateResponseWithDefaults instantiates a new SingleRateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *SingleRateResponse) GetFromAddress() SingleRateResponseFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *SingleRateResponse) GetFromAddressOk() (*SingleRateResponseFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *SingleRateResponse) SetFromAddress(v SingleRateResponseFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *SingleRateResponse) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcel

`func (o *SingleRateResponse) GetParcel() SingleRateResponseParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *SingleRateResponse) GetParcelOk() (*SingleRateResponseParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *SingleRateResponse) SetParcel(v SingleRateResponseParcel)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *SingleRateResponse) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetRates

`func (o *SingleRateResponse) GetRates() []SingleRateResponseRatesInner`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *SingleRateResponse) GetRatesOk() (*[]SingleRateResponseRatesInner, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *SingleRateResponse) SetRates(v []SingleRateResponseRatesInner)`

SetRates sets Rates field to given value.

### HasRates

`func (o *SingleRateResponse) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetToAddress

`func (o *SingleRateResponse) GetToAddress() SingleRateResponseToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *SingleRateResponse) GetToAddressOk() (*SingleRateResponseToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *SingleRateResponse) SetToAddress(v SingleRateResponseToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *SingleRateResponse) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


