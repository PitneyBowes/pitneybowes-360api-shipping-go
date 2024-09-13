# MultipieceRatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | Pointer to [**MultipieceRatesRequestFromAddress**](MultipieceRatesRequestFromAddress.md) |  | [optional] 
**ToAddress** | Pointer to [**MultipieceRatesRequestToAddress**](MultipieceRatesRequestToAddress.md) |  | [optional] 
**ServiceId** | Pointer to **string** | description | [optional] 
**Rates** | Pointer to [**[]MultipieceRatesResponseRatesInner**](MultipieceRatesResponseRatesInner.md) | description | [optional] 

## Methods

### NewMultipieceRatesResponse

`func NewMultipieceRatesResponse() *MultipieceRatesResponse`

NewMultipieceRatesResponse instantiates a new MultipieceRatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceRatesResponseWithDefaults

`func NewMultipieceRatesResponseWithDefaults() *MultipieceRatesResponse`

NewMultipieceRatesResponseWithDefaults instantiates a new MultipieceRatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *MultipieceRatesResponse) GetFromAddress() MultipieceRatesRequestFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MultipieceRatesResponse) GetFromAddressOk() (*MultipieceRatesRequestFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MultipieceRatesResponse) SetFromAddress(v MultipieceRatesRequestFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MultipieceRatesResponse) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetToAddress

`func (o *MultipieceRatesResponse) GetToAddress() MultipieceRatesRequestToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *MultipieceRatesResponse) GetToAddressOk() (*MultipieceRatesRequestToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *MultipieceRatesResponse) SetToAddress(v MultipieceRatesRequestToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *MultipieceRatesResponse) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetServiceId

`func (o *MultipieceRatesResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MultipieceRatesResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MultipieceRatesResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MultipieceRatesResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetRates

`func (o *MultipieceRatesResponse) GetRates() []MultipieceRatesResponseRatesInner`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *MultipieceRatesResponse) GetRatesOk() (*[]MultipieceRatesResponseRatesInner, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *MultipieceRatesResponse) SetRates(v []MultipieceRatesResponseRatesInner)`

SetRates sets Rates field to given value.

### HasRates

`func (o *MultipieceRatesResponse) HasRates() bool`

HasRates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


