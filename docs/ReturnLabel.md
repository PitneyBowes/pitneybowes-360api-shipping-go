# ReturnLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **string** | This defines the label size of the Shipment, e.g., Shipping Label having Doc Size (8&#39; X 11&#39;). | [optional] 
**Type** | Pointer to **string** | This defines the type of the Shipment, e.g., Shipping Label. | [optional] 
**FromAddress** | Pointer to [**ReturnLabelFromAddress**](ReturnLabelFromAddress.md) |  | [optional] 
**Parcel** | Pointer to [**ShipmentDomesticParcel**](ShipmentDomesticParcel.md) |  | [optional] 
**CarrierAccountId** | Pointer to **string** | The unique identifier associated with the Carrier account used by client users during shipment process. | [optional] 
**ParcelType** | Pointer to **string** | &gt;-Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection. It has categories like Package, Envelopes, Paks, Boxes, Tube, defined per specific carrier and used in abbreviated form, e.g., FRPKG, LGENV, TUBE,PKG. | [optional] 
**ParcelId** | Pointer to **string** | A unique identifier associated with the Parcel. | [optional] 
**ServiceId** | Pointer to **string** | &gt;-A unique identifier given to the carrier-specific service. This is required for creating a shipment, while it is optional for rating a parcel. | [optional] 
**SpecialServices** | Pointer to [**[]ReturnLabelSpecialServicesInner**](ReturnLabelSpecialServicesInner.md) |  | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentOptionsV2**](ShipmentOptionsV2.md) |  | [optional] 
**Metadata** | Pointer to [**[]ShipmentInternationalMetadataInner**](ShipmentInternationalMetadataInner.md) | Additional metadata that needs to be stored for this shipment can be added here. For now, &#x60;costAccountName&#x60; is supported. | [optional] 
**ToAddress** | Pointer to [**ReturnLabelToAddress**](ReturnLabelToAddress.md) |  | [optional] 

## Methods

### NewReturnLabel

`func NewReturnLabel() *ReturnLabel`

NewReturnLabel instantiates a new ReturnLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnLabelWithDefaults

`func NewReturnLabelWithDefaults() *ReturnLabel`

NewReturnLabelWithDefaults instantiates a new ReturnLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *ReturnLabel) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ReturnLabel) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ReturnLabel) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ReturnLabel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *ReturnLabel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReturnLabel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReturnLabel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReturnLabel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFromAddress

`func (o *ReturnLabel) GetFromAddress() ReturnLabelFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ReturnLabel) GetFromAddressOk() (*ReturnLabelFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ReturnLabel) SetFromAddress(v ReturnLabelFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *ReturnLabel) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcel

`func (o *ReturnLabel) GetParcel() ShipmentDomesticParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *ReturnLabel) GetParcelOk() (*ShipmentDomesticParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *ReturnLabel) SetParcel(v ShipmentDomesticParcel)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *ReturnLabel) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *ReturnLabel) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *ReturnLabel) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *ReturnLabel) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *ReturnLabel) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetParcelType

`func (o *ReturnLabel) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *ReturnLabel) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *ReturnLabel) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *ReturnLabel) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetParcelId

`func (o *ReturnLabel) GetParcelId() string`

GetParcelId returns the ParcelId field if non-nil, zero value otherwise.

### GetParcelIdOk

`func (o *ReturnLabel) GetParcelIdOk() (*string, bool)`

GetParcelIdOk returns a tuple with the ParcelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelId

`func (o *ReturnLabel) SetParcelId(v string)`

SetParcelId sets ParcelId field to given value.

### HasParcelId

`func (o *ReturnLabel) HasParcelId() bool`

HasParcelId returns a boolean if a field has been set.

### GetServiceId

`func (o *ReturnLabel) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ReturnLabel) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ReturnLabel) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ReturnLabel) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSpecialServices

`func (o *ReturnLabel) GetSpecialServices() []ReturnLabelSpecialServicesInner`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *ReturnLabel) GetSpecialServicesOk() (*[]ReturnLabelSpecialServicesInner, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *ReturnLabel) SetSpecialServices(v []ReturnLabelSpecialServicesInner)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *ReturnLabel) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *ReturnLabel) GetShipmentOptions() ShipmentOptionsV2`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *ReturnLabel) GetShipmentOptionsOk() (*ShipmentOptionsV2, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *ReturnLabel) SetShipmentOptions(v ShipmentOptionsV2)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *ReturnLabel) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetMetadata

`func (o *ReturnLabel) GetMetadata() []ShipmentInternationalMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReturnLabel) GetMetadataOk() (*[]ShipmentInternationalMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReturnLabel) SetMetadata(v []ShipmentInternationalMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ReturnLabel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetToAddress

`func (o *ReturnLabel) GetToAddress() ReturnLabelToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *ReturnLabel) GetToAddressOk() (*ReturnLabelToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *ReturnLabel) SetToAddress(v ReturnLabelToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *ReturnLabel) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


