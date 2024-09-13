# MultipieceRateShopResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | Pointer to [**MultipieceRatesRequestFromAddress**](MultipieceRatesRequestFromAddress.md) |  | [optional] 
**ToAddress** | Pointer to [**MultipieceRatesRequestToAddress**](MultipieceRatesRequestToAddress.md) |  | [optional] 
**ServiceId** | Pointer to **string** | description | [optional] 
**Rates** | Pointer to [**[]MultipieceRateShopResponseRatesInner**](MultipieceRateShopResponseRatesInner.md) | description | [optional] 
**Errors** | Pointer to [**[]MultipieceRateShopResponseErrorsInner**](MultipieceRateShopResponseErrorsInner.md) | It display any error while getting rates | [optional] 

## Methods

### NewMultipieceRateShopResponse

`func NewMultipieceRateShopResponse() *MultipieceRateShopResponse`

NewMultipieceRateShopResponse instantiates a new MultipieceRateShopResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceRateShopResponseWithDefaults

`func NewMultipieceRateShopResponseWithDefaults() *MultipieceRateShopResponse`

NewMultipieceRateShopResponseWithDefaults instantiates a new MultipieceRateShopResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *MultipieceRateShopResponse) GetFromAddress() MultipieceRatesRequestFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MultipieceRateShopResponse) GetFromAddressOk() (*MultipieceRatesRequestFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MultipieceRateShopResponse) SetFromAddress(v MultipieceRatesRequestFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MultipieceRateShopResponse) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetToAddress

`func (o *MultipieceRateShopResponse) GetToAddress() MultipieceRatesRequestToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *MultipieceRateShopResponse) GetToAddressOk() (*MultipieceRatesRequestToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *MultipieceRateShopResponse) SetToAddress(v MultipieceRatesRequestToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *MultipieceRateShopResponse) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetServiceId

`func (o *MultipieceRateShopResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MultipieceRateShopResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MultipieceRateShopResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MultipieceRateShopResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetRates

`func (o *MultipieceRateShopResponse) GetRates() []MultipieceRateShopResponseRatesInner`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *MultipieceRateShopResponse) GetRatesOk() (*[]MultipieceRateShopResponseRatesInner, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *MultipieceRateShopResponse) SetRates(v []MultipieceRateShopResponseRatesInner)`

SetRates sets Rates field to given value.

### HasRates

`func (o *MultipieceRateShopResponse) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetErrors

`func (o *MultipieceRateShopResponse) GetErrors() []MultipieceRateShopResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MultipieceRateShopResponse) GetErrorsOk() (*[]MultipieceRateShopResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MultipieceRateShopResponse) SetErrors(v []MultipieceRateShopResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MultipieceRateShopResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


