# MultipieceInternationalShipmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **string** | description | [optional] 
**Type** | Pointer to **string** | description | [optional] 
**FromAddress** | Pointer to [**MultipieceDomesticShipmentRequestFromAddress**](MultipieceDomesticShipmentRequestFromAddress.md) |  | [optional] 
**CarrierAccountId** | Pointer to **string** | description | [optional] 
**ServiceId** | Pointer to **string** | description | [optional] 
**ShipmentOptions** | Pointer to [**MultipieceDomesticShipmentRequestShipmentOptions**](MultipieceDomesticShipmentRequestShipmentOptions.md) |  | [optional] 
**Metadata** | Pointer to [**[]MultipieceDomesticShipmentRequestMetadataInner**](MultipieceDomesticShipmentRequestMetadataInner.md) | description | [optional] 
**MultiPieceParcels** | Pointer to [**[]MultipieceInternationalShipmentRequestMultiPieceParcelsInner**](MultipieceInternationalShipmentRequestMultiPieceParcelsInner.md) | description | [optional] 
**ToAddress** | Pointer to [**MultipieceInternationalShipmentRequestToAddress**](MultipieceInternationalShipmentRequestToAddress.md) |  | [optional] 
**Customs** | Pointer to [**MultipieceInternationalShipmentRequestCustoms**](MultipieceInternationalShipmentRequestCustoms.md) |  | [optional] 

## Methods

### NewMultipieceInternationalShipmentRequest

`func NewMultipieceInternationalShipmentRequest() *MultipieceInternationalShipmentRequest`

NewMultipieceInternationalShipmentRequest instantiates a new MultipieceInternationalShipmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceInternationalShipmentRequestWithDefaults

`func NewMultipieceInternationalShipmentRequestWithDefaults() *MultipieceInternationalShipmentRequest`

NewMultipieceInternationalShipmentRequestWithDefaults instantiates a new MultipieceInternationalShipmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *MultipieceInternationalShipmentRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MultipieceInternationalShipmentRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MultipieceInternationalShipmentRequest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *MultipieceInternationalShipmentRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *MultipieceInternationalShipmentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultipieceInternationalShipmentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultipieceInternationalShipmentRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MultipieceInternationalShipmentRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFromAddress

`func (o *MultipieceInternationalShipmentRequest) GetFromAddress() MultipieceDomesticShipmentRequestFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MultipieceInternationalShipmentRequest) GetFromAddressOk() (*MultipieceDomesticShipmentRequestFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MultipieceInternationalShipmentRequest) SetFromAddress(v MultipieceDomesticShipmentRequestFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MultipieceInternationalShipmentRequest) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *MultipieceInternationalShipmentRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *MultipieceInternationalShipmentRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *MultipieceInternationalShipmentRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *MultipieceInternationalShipmentRequest) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetServiceId

`func (o *MultipieceInternationalShipmentRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MultipieceInternationalShipmentRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MultipieceInternationalShipmentRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MultipieceInternationalShipmentRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *MultipieceInternationalShipmentRequest) GetShipmentOptions() MultipieceDomesticShipmentRequestShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *MultipieceInternationalShipmentRequest) GetShipmentOptionsOk() (*MultipieceDomesticShipmentRequestShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *MultipieceInternationalShipmentRequest) SetShipmentOptions(v MultipieceDomesticShipmentRequestShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *MultipieceInternationalShipmentRequest) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetMetadata

`func (o *MultipieceInternationalShipmentRequest) GetMetadata() []MultipieceDomesticShipmentRequestMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MultipieceInternationalShipmentRequest) GetMetadataOk() (*[]MultipieceDomesticShipmentRequestMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MultipieceInternationalShipmentRequest) SetMetadata(v []MultipieceDomesticShipmentRequestMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MultipieceInternationalShipmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMultiPieceParcels

`func (o *MultipieceInternationalShipmentRequest) GetMultiPieceParcels() []MultipieceInternationalShipmentRequestMultiPieceParcelsInner`

GetMultiPieceParcels returns the MultiPieceParcels field if non-nil, zero value otherwise.

### GetMultiPieceParcelsOk

`func (o *MultipieceInternationalShipmentRequest) GetMultiPieceParcelsOk() (*[]MultipieceInternationalShipmentRequestMultiPieceParcelsInner, bool)`

GetMultiPieceParcelsOk returns a tuple with the MultiPieceParcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPieceParcels

`func (o *MultipieceInternationalShipmentRequest) SetMultiPieceParcels(v []MultipieceInternationalShipmentRequestMultiPieceParcelsInner)`

SetMultiPieceParcels sets MultiPieceParcels field to given value.

### HasMultiPieceParcels

`func (o *MultipieceInternationalShipmentRequest) HasMultiPieceParcels() bool`

HasMultiPieceParcels returns a boolean if a field has been set.

### GetToAddress

`func (o *MultipieceInternationalShipmentRequest) GetToAddress() MultipieceInternationalShipmentRequestToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *MultipieceInternationalShipmentRequest) GetToAddressOk() (*MultipieceInternationalShipmentRequestToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *MultipieceInternationalShipmentRequest) SetToAddress(v MultipieceInternationalShipmentRequestToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *MultipieceInternationalShipmentRequest) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetCustoms

`func (o *MultipieceInternationalShipmentRequest) GetCustoms() MultipieceInternationalShipmentRequestCustoms`

GetCustoms returns the Customs field if non-nil, zero value otherwise.

### GetCustomsOk

`func (o *MultipieceInternationalShipmentRequest) GetCustomsOk() (*MultipieceInternationalShipmentRequestCustoms, bool)`

GetCustomsOk returns a tuple with the Customs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustoms

`func (o *MultipieceInternationalShipmentRequest) SetCustoms(v MultipieceInternationalShipmentRequestCustoms)`

SetCustoms sets Customs field to given value.

### HasCustoms

`func (o *MultipieceInternationalShipmentRequest) HasCustoms() bool`

HasCustoms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


