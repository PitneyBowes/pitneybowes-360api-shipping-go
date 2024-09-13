# CreateBatchRequestERR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the of ERR Batch which is imported, e.g. ERR-Import-05. | [optional] 
**GroupName** | Pointer to **string** | Indicates the name of the group of batches, which consists of multiple Batch groups. | [optional] 
**Size** | **string** | This indicates the label size of the Bulk Shipment when it gets printed,i.e., DocSize. | 
**Type** | **string** | This indicates the type of the Batch Shipment, e.g., Shipping Label and Coversheet. | 
**Format** | Pointer to **string** | This defines the type of the shipment which is printed, e.g., Shipping label gets printed in PDF form. | [optional] 
**CarrierAccountId** | **string** | A unique identifier associated with the carrier account used by client users during shipment process. Default CarrierAccountID for this batch will be user&#39;s registered USPS account. | 
**ServiceId** | **string** | A unique identifier given to the carrier-specific service. User can override this value by defining it at Shipment level. | 
**ParcelType** | **string** | Parcel Type is required for creating a shipment while rating a parcel. And it varies as per USPS selected services, e.g. FRPKG, LGENV, TUBE, and PKG. | 
**SpecialServices** | Pointer to [**[]SpecialServiceBatchERR**](SpecialServiceBatchERR.md) |  | [optional] 

## Methods

### NewCreateBatchRequestERR

`func NewCreateBatchRequestERR(size string, type_ string, carrierAccountId string, serviceId string, parcelType string, ) *CreateBatchRequestERR`

NewCreateBatchRequestERR instantiates a new CreateBatchRequestERR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchRequestERRWithDefaults

`func NewCreateBatchRequestERRWithDefaults() *CreateBatchRequestERR`

NewCreateBatchRequestERRWithDefaults instantiates a new CreateBatchRequestERR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBatchRequestERR) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBatchRequestERR) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBatchRequestERR) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateBatchRequestERR) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroupName

`func (o *CreateBatchRequestERR) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CreateBatchRequestERR) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CreateBatchRequestERR) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CreateBatchRequestERR) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSize

`func (o *CreateBatchRequestERR) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateBatchRequestERR) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateBatchRequestERR) SetSize(v string)`

SetSize sets Size field to given value.


### GetType

`func (o *CreateBatchRequestERR) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateBatchRequestERR) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateBatchRequestERR) SetType(v string)`

SetType sets Type field to given value.


### GetFormat

`func (o *CreateBatchRequestERR) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateBatchRequestERR) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateBatchRequestERR) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateBatchRequestERR) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *CreateBatchRequestERR) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *CreateBatchRequestERR) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *CreateBatchRequestERR) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetServiceId

`func (o *CreateBatchRequestERR) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateBatchRequestERR) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateBatchRequestERR) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetParcelType

`func (o *CreateBatchRequestERR) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *CreateBatchRequestERR) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *CreateBatchRequestERR) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.


### GetSpecialServices

`func (o *CreateBatchRequestERR) GetSpecialServices() []SpecialServiceBatchERR`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *CreateBatchRequestERR) GetSpecialServicesOk() (*[]SpecialServiceBatchERR, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *CreateBatchRequestERR) SetSpecialServices(v []SpecialServiceBatchERR)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *CreateBatchRequestERR) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


