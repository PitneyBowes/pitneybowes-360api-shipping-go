# CreateBulkShipmentsERRLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the of ERR Batch which consists of multiple shipments (shipments in bulk) for Label printing, e.g. ERR-Bulk05. | 
**GroupName** | Pointer to **string** | Indicates the name of the group of batches, which consists of multiple Batch groups. | [optional] 
**Size** | **string** | This indicates the label size of the Bulk Shipment when it gets printed,i.e., DocSize. This has two options 8&#39; X 11&#39; or 4&#39; X 6&#39;. | 
**Type** | **string** | Indicates the type of the Batch Shipment, e.g., Shipping Label. | 
**Format** | Pointer to **string** | Defines the type of the shipment which is printed, e.g., Shipping label gets printed in PDF form. | [optional] 
**CarrierAccountId** | **string** | A unique identifier associated with the Carrier account used by client users during shipment process. Default CarrierAccountID for this batch will be user&#39;s registered USPS account. User can override this value by defining it at Shipment level. | 
**ServiceId** | **string** | A unique identifier given to the carrier-specific service. User can override this value by defining it at Shipment level. | 
**ParcelType** | **string** | Parcel Type is required for creating a shipment while rating a parcel. And it varies as per USPS selected services, e.g. FRPKG, LGENV, TUBE, and PKG. User can override this value by defining it at Shipment level. | 
**ParcelID** | Pointer to **string** | A unique identifier given to the parcel or package corresponding to the selected service. This is optional field, but is used in few cases. Examples include BLM10, B1095, MT1098, etc. User can override this value by defining it at Shipment level. | [optional] 
**SpecialServices** | Pointer to [**[]SpecialServiceERRInner**](SpecialServiceERRInner.md) |  | [optional] 
**Shipments** | [**[]ShipmentERR**](ShipmentERR.md) |  | 

## Methods

### NewCreateBulkShipmentsERRLabel

`func NewCreateBulkShipmentsERRLabel(name string, size string, type_ string, carrierAccountId string, serviceId string, parcelType string, shipments []ShipmentERR, ) *CreateBulkShipmentsERRLabel`

NewCreateBulkShipmentsERRLabel instantiates a new CreateBulkShipmentsERRLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBulkShipmentsERRLabelWithDefaults

`func NewCreateBulkShipmentsERRLabelWithDefaults() *CreateBulkShipmentsERRLabel`

NewCreateBulkShipmentsERRLabelWithDefaults instantiates a new CreateBulkShipmentsERRLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBulkShipmentsERRLabel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBulkShipmentsERRLabel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBulkShipmentsERRLabel) SetName(v string)`

SetName sets Name field to given value.


### GetGroupName

`func (o *CreateBulkShipmentsERRLabel) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CreateBulkShipmentsERRLabel) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CreateBulkShipmentsERRLabel) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CreateBulkShipmentsERRLabel) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSize

`func (o *CreateBulkShipmentsERRLabel) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateBulkShipmentsERRLabel) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateBulkShipmentsERRLabel) SetSize(v string)`

SetSize sets Size field to given value.


### GetType

`func (o *CreateBulkShipmentsERRLabel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateBulkShipmentsERRLabel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateBulkShipmentsERRLabel) SetType(v string)`

SetType sets Type field to given value.


### GetFormat

`func (o *CreateBulkShipmentsERRLabel) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateBulkShipmentsERRLabel) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateBulkShipmentsERRLabel) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateBulkShipmentsERRLabel) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *CreateBulkShipmentsERRLabel) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *CreateBulkShipmentsERRLabel) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *CreateBulkShipmentsERRLabel) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetServiceId

`func (o *CreateBulkShipmentsERRLabel) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateBulkShipmentsERRLabel) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateBulkShipmentsERRLabel) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetParcelType

`func (o *CreateBulkShipmentsERRLabel) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *CreateBulkShipmentsERRLabel) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *CreateBulkShipmentsERRLabel) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.


### GetParcelID

`func (o *CreateBulkShipmentsERRLabel) GetParcelID() string`

GetParcelID returns the ParcelID field if non-nil, zero value otherwise.

### GetParcelIDOk

`func (o *CreateBulkShipmentsERRLabel) GetParcelIDOk() (*string, bool)`

GetParcelIDOk returns a tuple with the ParcelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelID

`func (o *CreateBulkShipmentsERRLabel) SetParcelID(v string)`

SetParcelID sets ParcelID field to given value.

### HasParcelID

`func (o *CreateBulkShipmentsERRLabel) HasParcelID() bool`

HasParcelID returns a boolean if a field has been set.

### GetSpecialServices

`func (o *CreateBulkShipmentsERRLabel) GetSpecialServices() []SpecialServiceERRInner`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *CreateBulkShipmentsERRLabel) GetSpecialServicesOk() (*[]SpecialServiceERRInner, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *CreateBulkShipmentsERRLabel) SetSpecialServices(v []SpecialServiceERRInner)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *CreateBulkShipmentsERRLabel) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetShipments

`func (o *CreateBulkShipmentsERRLabel) GetShipments() []ShipmentERR`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *CreateBulkShipmentsERRLabel) GetShipmentsOk() (*[]ShipmentERR, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *CreateBulkShipmentsERRLabel) SetShipments(v []ShipmentERR)`

SetShipments sets Shipments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


