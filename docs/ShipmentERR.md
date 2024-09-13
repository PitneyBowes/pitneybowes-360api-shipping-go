# ShipmentERR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | This is a user-defined value provided by users just for their reference. This is for mapping purpose against each shipment. | [optional] 
**FromAddress** | [**FromAddress**](FromAddress.md) |  | 
**ToAddress** | [**ToAddress**](ToAddress.md) |  | 
**Parcel** | [**Parcel**](Parcel.md) |  | 
**CarrierAccountId** | Pointer to **string** | A unique identifier associated with the Carrier account used by client users during shipment process. | [optional] 
**ParcelType** | Pointer to **string** | &gt;-Packaging type varies as per selected carrier and its services, e.g., PKG, LGENV. | [optional] 
**ServiceId** | Pointer to **string** | &gt;-The unique identifier given to the carrier-specific service. ERR supports two services: First Class Mail (FCM) and Priority Mail (PM). | [optional] 
**DateOfShipment** | Pointer to **string** | Indicates the date when shipment is created. | [optional] 
**SpecialServices** | Pointer to [**[]SpecialServiceERRInner**](SpecialServiceERRInner.md) |  | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentOptionsERR**](ShipmentOptionsERR.md) |  | [optional] 
**Metadata** | Pointer to [**[]ShipmentERRMetadataInner**](ShipmentERRMetadataInner.md) | Additional metadata that needs to be stored for this shipment can be added here. For now, &#x60;costAccountName&#x60; is supported. | [optional] 

## Methods

### NewShipmentERR

`func NewShipmentERR(fromAddress FromAddress, toAddress ToAddress, parcel Parcel, ) *ShipmentERR`

NewShipmentERR instantiates a new ShipmentERR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentERRWithDefaults

`func NewShipmentERRWithDefaults() *ShipmentERR`

NewShipmentERRWithDefaults instantiates a new ShipmentERR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ShipmentERR) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ShipmentERR) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ShipmentERR) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ShipmentERR) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFromAddress

`func (o *ShipmentERR) GetFromAddress() FromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ShipmentERR) GetFromAddressOk() (*FromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ShipmentERR) SetFromAddress(v FromAddress)`

SetFromAddress sets FromAddress field to given value.


### GetToAddress

`func (o *ShipmentERR) GetToAddress() ToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *ShipmentERR) GetToAddressOk() (*ToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *ShipmentERR) SetToAddress(v ToAddress)`

SetToAddress sets ToAddress field to given value.


### GetParcel

`func (o *ShipmentERR) GetParcel() Parcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *ShipmentERR) GetParcelOk() (*Parcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *ShipmentERR) SetParcel(v Parcel)`

SetParcel sets Parcel field to given value.


### GetCarrierAccountId

`func (o *ShipmentERR) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *ShipmentERR) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *ShipmentERR) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *ShipmentERR) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetParcelType

`func (o *ShipmentERR) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *ShipmentERR) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *ShipmentERR) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *ShipmentERR) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetServiceId

`func (o *ShipmentERR) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ShipmentERR) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ShipmentERR) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ShipmentERR) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetDateOfShipment

`func (o *ShipmentERR) GetDateOfShipment() string`

GetDateOfShipment returns the DateOfShipment field if non-nil, zero value otherwise.

### GetDateOfShipmentOk

`func (o *ShipmentERR) GetDateOfShipmentOk() (*string, bool)`

GetDateOfShipmentOk returns a tuple with the DateOfShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfShipment

`func (o *ShipmentERR) SetDateOfShipment(v string)`

SetDateOfShipment sets DateOfShipment field to given value.

### HasDateOfShipment

`func (o *ShipmentERR) HasDateOfShipment() bool`

HasDateOfShipment returns a boolean if a field has been set.

### GetSpecialServices

`func (o *ShipmentERR) GetSpecialServices() []SpecialServiceERRInner`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *ShipmentERR) GetSpecialServicesOk() (*[]SpecialServiceERRInner, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *ShipmentERR) SetSpecialServices(v []SpecialServiceERRInner)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *ShipmentERR) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *ShipmentERR) GetShipmentOptions() ShipmentOptionsERR`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *ShipmentERR) GetShipmentOptionsOk() (*ShipmentOptionsERR, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *ShipmentERR) SetShipmentOptions(v ShipmentOptionsERR)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *ShipmentERR) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetMetadata

`func (o *ShipmentERR) GetMetadata() []ShipmentERRMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShipmentERR) GetMetadataOk() (*[]ShipmentERRMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShipmentERR) SetMetadata(v []ShipmentERRMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShipmentERR) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


