# CreateShipmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **string** | This defines the label size of the Shipment, e.g., Shipping Label having Doc Size (8&#39; X 11&#39;). | 
**Type** | **string** | This defines the type of the Shipment, e.g., Shipping Label. | 
**Format** | Pointer to **string** | This defines the type of the shipment which is printed. For example Shipping label prints in PDF form. | [optional] 
**DateOfShipment** | Pointer to **string** | This defines the date of the Shipment in the format YYYY:MM:DD. | [optional] 
**FromAddress** | [**ShipmentInternationalFromAddress**](ShipmentInternationalFromAddress.md) |  | 
**Parcel** | [**ShipmentDomesticParcel**](ShipmentDomesticParcel.md) |  | 
**CarrierAccountId** | **string** | The unique identifier associated with the Carrier account used by client users during shipment process. | 
**ParcelType** | **string** | &gt;-Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection. It has categories like Package, Envelopes, Paks, Boxes, Tube, defined per specific carrier and used in abbreviated form, e.g., FRPKG, LGENV, TUBE,PKG. | 
**ParcelId** | Pointer to **string** | &gt;-Parcel Id is optional and required to be given in case of branded parcels which have brandedDimension and/or brandedWeight. If parcel has brandedDimension, in that case user need not to pass dimensionUnit and dimension details(length, width and height) in the parcel object. And if brandedWeight is also available for the parcel then in that case weightUnit and weight need not to be passed  in parcel object | [optional] 
**ServiceId** | **string** | &gt;-A unique identifier given to the carrier-specific service. This is required for creating a shipment, while it is optional for rating a parcel. | 
**SpecialServices** | Pointer to [**[]SpecialService**](SpecialService.md) | This provides a carrier-service based special or extra sevice. | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentOptions**](ShipmentOptions.md) |  | [optional] 
**Metadata** | Pointer to [**[]ShipmentInternationalMetadataInner**](ShipmentInternationalMetadataInner.md) | Additional metadata that needs to be stored for this shipment can be added here. For now, &#x60;costAccountName&#x60; is supported. | [optional] 
**ToAddress** | [**ShipmentInternationalToAddress**](ShipmentInternationalToAddress.md) |  | 
**Customs** | [**ShipmentInternationalCustoms**](ShipmentInternationalCustoms.md) |  | 

## Methods

### NewCreateShipmentRequest

`func NewCreateShipmentRequest(size string, type_ string, fromAddress ShipmentInternationalFromAddress, parcel ShipmentDomesticParcel, carrierAccountId string, parcelType string, serviceId string, toAddress ShipmentInternationalToAddress, customs ShipmentInternationalCustoms, ) *CreateShipmentRequest`

NewCreateShipmentRequest instantiates a new CreateShipmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShipmentRequestWithDefaults

`func NewCreateShipmentRequestWithDefaults() *CreateShipmentRequest`

NewCreateShipmentRequestWithDefaults instantiates a new CreateShipmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *CreateShipmentRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateShipmentRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateShipmentRequest) SetSize(v string)`

SetSize sets Size field to given value.


### GetType

`func (o *CreateShipmentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateShipmentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateShipmentRequest) SetType(v string)`

SetType sets Type field to given value.


### GetFormat

`func (o *CreateShipmentRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateShipmentRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateShipmentRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateShipmentRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetDateOfShipment

`func (o *CreateShipmentRequest) GetDateOfShipment() string`

GetDateOfShipment returns the DateOfShipment field if non-nil, zero value otherwise.

### GetDateOfShipmentOk

`func (o *CreateShipmentRequest) GetDateOfShipmentOk() (*string, bool)`

GetDateOfShipmentOk returns a tuple with the DateOfShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfShipment

`func (o *CreateShipmentRequest) SetDateOfShipment(v string)`

SetDateOfShipment sets DateOfShipment field to given value.

### HasDateOfShipment

`func (o *CreateShipmentRequest) HasDateOfShipment() bool`

HasDateOfShipment returns a boolean if a field has been set.

### GetFromAddress

`func (o *CreateShipmentRequest) GetFromAddress() ShipmentInternationalFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *CreateShipmentRequest) GetFromAddressOk() (*ShipmentInternationalFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *CreateShipmentRequest) SetFromAddress(v ShipmentInternationalFromAddress)`

SetFromAddress sets FromAddress field to given value.


### GetParcel

`func (o *CreateShipmentRequest) GetParcel() ShipmentDomesticParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *CreateShipmentRequest) GetParcelOk() (*ShipmentDomesticParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *CreateShipmentRequest) SetParcel(v ShipmentDomesticParcel)`

SetParcel sets Parcel field to given value.


### GetCarrierAccountId

`func (o *CreateShipmentRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *CreateShipmentRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *CreateShipmentRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetParcelType

`func (o *CreateShipmentRequest) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *CreateShipmentRequest) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *CreateShipmentRequest) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.


### GetParcelId

`func (o *CreateShipmentRequest) GetParcelId() string`

GetParcelId returns the ParcelId field if non-nil, zero value otherwise.

### GetParcelIdOk

`func (o *CreateShipmentRequest) GetParcelIdOk() (*string, bool)`

GetParcelIdOk returns a tuple with the ParcelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelId

`func (o *CreateShipmentRequest) SetParcelId(v string)`

SetParcelId sets ParcelId field to given value.

### HasParcelId

`func (o *CreateShipmentRequest) HasParcelId() bool`

HasParcelId returns a boolean if a field has been set.

### GetServiceId

`func (o *CreateShipmentRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateShipmentRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateShipmentRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSpecialServices

`func (o *CreateShipmentRequest) GetSpecialServices() []SpecialService`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *CreateShipmentRequest) GetSpecialServicesOk() (*[]SpecialService, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *CreateShipmentRequest) SetSpecialServices(v []SpecialService)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *CreateShipmentRequest) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *CreateShipmentRequest) GetShipmentOptions() ShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *CreateShipmentRequest) GetShipmentOptionsOk() (*ShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *CreateShipmentRequest) SetShipmentOptions(v ShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *CreateShipmentRequest) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateShipmentRequest) GetMetadata() []ShipmentInternationalMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateShipmentRequest) GetMetadataOk() (*[]ShipmentInternationalMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateShipmentRequest) SetMetadata(v []ShipmentInternationalMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateShipmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetToAddress

`func (o *CreateShipmentRequest) GetToAddress() ShipmentInternationalToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *CreateShipmentRequest) GetToAddressOk() (*ShipmentInternationalToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *CreateShipmentRequest) SetToAddress(v ShipmentInternationalToAddress)`

SetToAddress sets ToAddress field to given value.


### GetCustoms

`func (o *CreateShipmentRequest) GetCustoms() ShipmentInternationalCustoms`

GetCustoms returns the Customs field if non-nil, zero value otherwise.

### GetCustomsOk

`func (o *CreateShipmentRequest) GetCustomsOk() (*ShipmentInternationalCustoms, bool)`

GetCustomsOk returns a tuple with the Customs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustoms

`func (o *CreateShipmentRequest) SetCustoms(v ShipmentInternationalCustoms)`

SetCustoms sets Customs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


