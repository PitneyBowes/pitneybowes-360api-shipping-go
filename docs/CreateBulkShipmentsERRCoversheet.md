# CreateBulkShipmentsERRCoversheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the of ERR Batch which consists of multiple shipments (shipments in bulk) for Coversheet printing., e.g. ERR-Coversheet07. | 
**GroupName** | Pointer to **string** | Indicates the name of the group of batches, which consists of multiple Batch groups. | [optional] 
**Size** | **string** | This indicates the envelope size of the ERR Bulk Shipment, i.e., DocSize. We have indicated the example values in enum. | 
**Type** | **string** | This indicates the type of the ERR Batch Shipment, e.g., Coversheet. | 
**Format** | Pointer to **string** | This defines the format type of the shipment which is printed. For example Coversheet prints in PDF form. | [optional] 
**CarrierAccountId** | **string** | The unique identifier associated with the user&#39;s registered USPS Account which will be required for this batch. User can override this value by defining it at Shipment level. | 
**ParcelType** | **string** | Parcel Type is required for creating a shipment while rating a parcel. And it varies as per USPS selected services, e.g. LTR, LGENV. User can override this value by defining it at shipment level. | 
**ServiceId** | **string** | A unique identifier given  to the carrier-specific service which is used for this BulkShipment. User can override this value by defining it at shipment level. | 
**SpecialServices** | Pointer to [**[]SpecialServiceERRInner**](SpecialServiceERRInner.md) |  | [optional] 
**Shipments** | [**[]ShipmentERRCoversheet**](ShipmentERRCoversheet.md) |  | 

## Methods

### NewCreateBulkShipmentsERRCoversheet

`func NewCreateBulkShipmentsERRCoversheet(name string, size string, type_ string, carrierAccountId string, parcelType string, serviceId string, shipments []ShipmentERRCoversheet, ) *CreateBulkShipmentsERRCoversheet`

NewCreateBulkShipmentsERRCoversheet instantiates a new CreateBulkShipmentsERRCoversheet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBulkShipmentsERRCoversheetWithDefaults

`func NewCreateBulkShipmentsERRCoversheetWithDefaults() *CreateBulkShipmentsERRCoversheet`

NewCreateBulkShipmentsERRCoversheetWithDefaults instantiates a new CreateBulkShipmentsERRCoversheet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBulkShipmentsERRCoversheet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBulkShipmentsERRCoversheet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBulkShipmentsERRCoversheet) SetName(v string)`

SetName sets Name field to given value.


### GetGroupName

`func (o *CreateBulkShipmentsERRCoversheet) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CreateBulkShipmentsERRCoversheet) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CreateBulkShipmentsERRCoversheet) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CreateBulkShipmentsERRCoversheet) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSize

`func (o *CreateBulkShipmentsERRCoversheet) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateBulkShipmentsERRCoversheet) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateBulkShipmentsERRCoversheet) SetSize(v string)`

SetSize sets Size field to given value.


### GetType

`func (o *CreateBulkShipmentsERRCoversheet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateBulkShipmentsERRCoversheet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateBulkShipmentsERRCoversheet) SetType(v string)`

SetType sets Type field to given value.


### GetFormat

`func (o *CreateBulkShipmentsERRCoversheet) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateBulkShipmentsERRCoversheet) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateBulkShipmentsERRCoversheet) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateBulkShipmentsERRCoversheet) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *CreateBulkShipmentsERRCoversheet) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *CreateBulkShipmentsERRCoversheet) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *CreateBulkShipmentsERRCoversheet) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetParcelType

`func (o *CreateBulkShipmentsERRCoversheet) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *CreateBulkShipmentsERRCoversheet) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *CreateBulkShipmentsERRCoversheet) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.


### GetServiceId

`func (o *CreateBulkShipmentsERRCoversheet) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateBulkShipmentsERRCoversheet) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateBulkShipmentsERRCoversheet) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSpecialServices

`func (o *CreateBulkShipmentsERRCoversheet) GetSpecialServices() []SpecialServiceERRInner`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *CreateBulkShipmentsERRCoversheet) GetSpecialServicesOk() (*[]SpecialServiceERRInner, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *CreateBulkShipmentsERRCoversheet) SetSpecialServices(v []SpecialServiceERRInner)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *CreateBulkShipmentsERRCoversheet) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetShipments

`func (o *CreateBulkShipmentsERRCoversheet) GetShipments() []ShipmentERRCoversheet`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *CreateBulkShipmentsERRCoversheet) GetShipmentsOk() (*[]ShipmentERRCoversheet, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *CreateBulkShipmentsERRCoversheet) SetShipments(v []ShipmentERRCoversheet)`

SetShipments sets Shipments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


