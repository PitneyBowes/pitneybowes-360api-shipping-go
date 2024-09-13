# MultipieceRateShopRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | Pointer to [**MultipieceRatesRequestFromAddress**](MultipieceRatesRequestFromAddress.md) |  | [optional] 
**MultiPieceParcels** | Pointer to [**[]MultipieceRateShopRequestMultiPieceParcelsInner**](MultipieceRateShopRequestMultiPieceParcelsInner.md) | description | [optional] 
**CarrierAccounts** | Pointer to **[]string** | description | [optional] 
**ToAddress** | Pointer to [**MultipieceRatesRequestToAddress**](MultipieceRatesRequestToAddress.md) |  | [optional] 

## Methods

### NewMultipieceRateShopRequest

`func NewMultipieceRateShopRequest() *MultipieceRateShopRequest`

NewMultipieceRateShopRequest instantiates a new MultipieceRateShopRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceRateShopRequestWithDefaults

`func NewMultipieceRateShopRequestWithDefaults() *MultipieceRateShopRequest`

NewMultipieceRateShopRequestWithDefaults instantiates a new MultipieceRateShopRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *MultipieceRateShopRequest) GetFromAddress() MultipieceRatesRequestFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MultipieceRateShopRequest) GetFromAddressOk() (*MultipieceRatesRequestFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MultipieceRateShopRequest) SetFromAddress(v MultipieceRatesRequestFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MultipieceRateShopRequest) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetMultiPieceParcels

`func (o *MultipieceRateShopRequest) GetMultiPieceParcels() []MultipieceRateShopRequestMultiPieceParcelsInner`

GetMultiPieceParcels returns the MultiPieceParcels field if non-nil, zero value otherwise.

### GetMultiPieceParcelsOk

`func (o *MultipieceRateShopRequest) GetMultiPieceParcelsOk() (*[]MultipieceRateShopRequestMultiPieceParcelsInner, bool)`

GetMultiPieceParcelsOk returns a tuple with the MultiPieceParcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPieceParcels

`func (o *MultipieceRateShopRequest) SetMultiPieceParcels(v []MultipieceRateShopRequestMultiPieceParcelsInner)`

SetMultiPieceParcels sets MultiPieceParcels field to given value.

### HasMultiPieceParcels

`func (o *MultipieceRateShopRequest) HasMultiPieceParcels() bool`

HasMultiPieceParcels returns a boolean if a field has been set.

### GetCarrierAccounts

`func (o *MultipieceRateShopRequest) GetCarrierAccounts() []string`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *MultipieceRateShopRequest) GetCarrierAccountsOk() (*[]string, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *MultipieceRateShopRequest) SetCarrierAccounts(v []string)`

SetCarrierAccounts sets CarrierAccounts field to given value.

### HasCarrierAccounts

`func (o *MultipieceRateShopRequest) HasCarrierAccounts() bool`

HasCarrierAccounts returns a boolean if a field has been set.

### GetToAddress

`func (o *MultipieceRateShopRequest) GetToAddress() MultipieceRatesRequestToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *MultipieceRateShopRequest) GetToAddressOk() (*MultipieceRatesRequestToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *MultipieceRateShopRequest) SetToAddress(v MultipieceRatesRequestToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *MultipieceRateShopRequest) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


