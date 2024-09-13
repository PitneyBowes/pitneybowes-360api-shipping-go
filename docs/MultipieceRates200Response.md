# MultipieceRates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | Pointer to [**MultipieceRatesRequestFromAddress**](MultipieceRatesRequestFromAddress.md) |  | [optional] 
**ToAddress** | Pointer to [**MultipieceRatesRequestToAddress**](MultipieceRatesRequestToAddress.md) |  | [optional] 
**ServiceId** | Pointer to **string** | description | [optional] 
**Rates** | Pointer to [**[]MultipieceRatesResponseRatesInner**](MultipieceRatesResponseRatesInner.md) | description | [optional] 
**Errors** | Pointer to [**[]MultipieceRateShopResponseErrorsInner**](MultipieceRateShopResponseErrorsInner.md) | It display any error while getting rates | [optional] 

## Methods

### NewMultipieceRates200Response

`func NewMultipieceRates200Response() *MultipieceRates200Response`

NewMultipieceRates200Response instantiates a new MultipieceRates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceRates200ResponseWithDefaults

`func NewMultipieceRates200ResponseWithDefaults() *MultipieceRates200Response`

NewMultipieceRates200ResponseWithDefaults instantiates a new MultipieceRates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *MultipieceRates200Response) GetFromAddress() MultipieceRatesRequestFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MultipieceRates200Response) GetFromAddressOk() (*MultipieceRatesRequestFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MultipieceRates200Response) SetFromAddress(v MultipieceRatesRequestFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MultipieceRates200Response) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetToAddress

`func (o *MultipieceRates200Response) GetToAddress() MultipieceRatesRequestToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *MultipieceRates200Response) GetToAddressOk() (*MultipieceRatesRequestToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *MultipieceRates200Response) SetToAddress(v MultipieceRatesRequestToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *MultipieceRates200Response) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetServiceId

`func (o *MultipieceRates200Response) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MultipieceRates200Response) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MultipieceRates200Response) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MultipieceRates200Response) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetRates

`func (o *MultipieceRates200Response) GetRates() []MultipieceRatesResponseRatesInner`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *MultipieceRates200Response) GetRatesOk() (*[]MultipieceRatesResponseRatesInner, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *MultipieceRates200Response) SetRates(v []MultipieceRatesResponseRatesInner)`

SetRates sets Rates field to given value.

### HasRates

`func (o *MultipieceRates200Response) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetErrors

`func (o *MultipieceRates200Response) GetErrors() []MultipieceRateShopResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MultipieceRates200Response) GetErrorsOk() (*[]MultipieceRateShopResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MultipieceRates200Response) SetErrors(v []MultipieceRateShopResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MultipieceRates200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


