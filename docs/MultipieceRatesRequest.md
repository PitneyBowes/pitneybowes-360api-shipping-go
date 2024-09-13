# MultipieceRatesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | Pointer to [**MultipieceRatesRequestFromAddress**](MultipieceRatesRequestFromAddress.md) |  | [optional] 
**MultiPieceParcels** | Pointer to [**[]MultipieceRatesRequestMultiPieceParcelsInner**](MultipieceRatesRequestMultiPieceParcelsInner.md) | description | [optional] 
**CarrierAccounts** | Pointer to **[]string** | description | [optional] 
**ToAddress** | Pointer to [**MultipieceRatesRequestToAddress**](MultipieceRatesRequestToAddress.md) |  | [optional] 
**ServiceId** | Pointer to **string** | description | [optional] 

## Methods

### NewMultipieceRatesRequest

`func NewMultipieceRatesRequest() *MultipieceRatesRequest`

NewMultipieceRatesRequest instantiates a new MultipieceRatesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceRatesRequestWithDefaults

`func NewMultipieceRatesRequestWithDefaults() *MultipieceRatesRequest`

NewMultipieceRatesRequestWithDefaults instantiates a new MultipieceRatesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *MultipieceRatesRequest) GetFromAddress() MultipieceRatesRequestFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MultipieceRatesRequest) GetFromAddressOk() (*MultipieceRatesRequestFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MultipieceRatesRequest) SetFromAddress(v MultipieceRatesRequestFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MultipieceRatesRequest) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetMultiPieceParcels

`func (o *MultipieceRatesRequest) GetMultiPieceParcels() []MultipieceRatesRequestMultiPieceParcelsInner`

GetMultiPieceParcels returns the MultiPieceParcels field if non-nil, zero value otherwise.

### GetMultiPieceParcelsOk

`func (o *MultipieceRatesRequest) GetMultiPieceParcelsOk() (*[]MultipieceRatesRequestMultiPieceParcelsInner, bool)`

GetMultiPieceParcelsOk returns a tuple with the MultiPieceParcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPieceParcels

`func (o *MultipieceRatesRequest) SetMultiPieceParcels(v []MultipieceRatesRequestMultiPieceParcelsInner)`

SetMultiPieceParcels sets MultiPieceParcels field to given value.

### HasMultiPieceParcels

`func (o *MultipieceRatesRequest) HasMultiPieceParcels() bool`

HasMultiPieceParcels returns a boolean if a field has been set.

### GetCarrierAccounts

`func (o *MultipieceRatesRequest) GetCarrierAccounts() []string`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *MultipieceRatesRequest) GetCarrierAccountsOk() (*[]string, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *MultipieceRatesRequest) SetCarrierAccounts(v []string)`

SetCarrierAccounts sets CarrierAccounts field to given value.

### HasCarrierAccounts

`func (o *MultipieceRatesRequest) HasCarrierAccounts() bool`

HasCarrierAccounts returns a boolean if a field has been set.

### GetToAddress

`func (o *MultipieceRatesRequest) GetToAddress() MultipieceRatesRequestToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *MultipieceRatesRequest) GetToAddressOk() (*MultipieceRatesRequestToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *MultipieceRatesRequest) SetToAddress(v MultipieceRatesRequestToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *MultipieceRatesRequest) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetServiceId

`func (o *MultipieceRatesRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MultipieceRatesRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MultipieceRatesRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MultipieceRatesRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


