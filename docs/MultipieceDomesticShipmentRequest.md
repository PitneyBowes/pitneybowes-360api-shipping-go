# MultipieceDomesticShipmentRequest

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
**MultiPieceParcels** | Pointer to [**[]MultipieceDomesticShipmentRequestMultiPieceParcelsInner**](MultipieceDomesticShipmentRequestMultiPieceParcelsInner.md) | description | [optional] 
**ToAddress** | Pointer to [**MultipieceDomesticShipmentRequestToAddress**](MultipieceDomesticShipmentRequestToAddress.md) |  | [optional] 

## Methods

### NewMultipieceDomesticShipmentRequest

`func NewMultipieceDomesticShipmentRequest() *MultipieceDomesticShipmentRequest`

NewMultipieceDomesticShipmentRequest instantiates a new MultipieceDomesticShipmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceDomesticShipmentRequestWithDefaults

`func NewMultipieceDomesticShipmentRequestWithDefaults() *MultipieceDomesticShipmentRequest`

NewMultipieceDomesticShipmentRequestWithDefaults instantiates a new MultipieceDomesticShipmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *MultipieceDomesticShipmentRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MultipieceDomesticShipmentRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MultipieceDomesticShipmentRequest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *MultipieceDomesticShipmentRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *MultipieceDomesticShipmentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultipieceDomesticShipmentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultipieceDomesticShipmentRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MultipieceDomesticShipmentRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFromAddress

`func (o *MultipieceDomesticShipmentRequest) GetFromAddress() MultipieceDomesticShipmentRequestFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MultipieceDomesticShipmentRequest) GetFromAddressOk() (*MultipieceDomesticShipmentRequestFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MultipieceDomesticShipmentRequest) SetFromAddress(v MultipieceDomesticShipmentRequestFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MultipieceDomesticShipmentRequest) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *MultipieceDomesticShipmentRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *MultipieceDomesticShipmentRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *MultipieceDomesticShipmentRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *MultipieceDomesticShipmentRequest) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetServiceId

`func (o *MultipieceDomesticShipmentRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MultipieceDomesticShipmentRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MultipieceDomesticShipmentRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MultipieceDomesticShipmentRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *MultipieceDomesticShipmentRequest) GetShipmentOptions() MultipieceDomesticShipmentRequestShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *MultipieceDomesticShipmentRequest) GetShipmentOptionsOk() (*MultipieceDomesticShipmentRequestShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *MultipieceDomesticShipmentRequest) SetShipmentOptions(v MultipieceDomesticShipmentRequestShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *MultipieceDomesticShipmentRequest) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetMetadata

`func (o *MultipieceDomesticShipmentRequest) GetMetadata() []MultipieceDomesticShipmentRequestMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MultipieceDomesticShipmentRequest) GetMetadataOk() (*[]MultipieceDomesticShipmentRequestMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MultipieceDomesticShipmentRequest) SetMetadata(v []MultipieceDomesticShipmentRequestMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MultipieceDomesticShipmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMultiPieceParcels

`func (o *MultipieceDomesticShipmentRequest) GetMultiPieceParcels() []MultipieceDomesticShipmentRequestMultiPieceParcelsInner`

GetMultiPieceParcels returns the MultiPieceParcels field if non-nil, zero value otherwise.

### GetMultiPieceParcelsOk

`func (o *MultipieceDomesticShipmentRequest) GetMultiPieceParcelsOk() (*[]MultipieceDomesticShipmentRequestMultiPieceParcelsInner, bool)`

GetMultiPieceParcelsOk returns a tuple with the MultiPieceParcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPieceParcels

`func (o *MultipieceDomesticShipmentRequest) SetMultiPieceParcels(v []MultipieceDomesticShipmentRequestMultiPieceParcelsInner)`

SetMultiPieceParcels sets MultiPieceParcels field to given value.

### HasMultiPieceParcels

`func (o *MultipieceDomesticShipmentRequest) HasMultiPieceParcels() bool`

HasMultiPieceParcels returns a boolean if a field has been set.

### GetToAddress

`func (o *MultipieceDomesticShipmentRequest) GetToAddress() MultipieceDomesticShipmentRequestToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *MultipieceDomesticShipmentRequest) GetToAddressOk() (*MultipieceDomesticShipmentRequestToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *MultipieceDomesticShipmentRequest) SetToAddress(v MultipieceDomesticShipmentRequestToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *MultipieceDomesticShipmentRequest) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


