# GetShipmentsForBatchDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** | This is a system-generated unique identifier assigned to the Batch while it is processed. | [optional] 
**CarrierAccountId** | Pointer to **string** | A unique identifier associated with the Carrier account used by client users during shipment process. | [optional] 
**ExternalId** | Pointer to **string** | This is a user-defined value provided by users just for their reference. This is for mapping purpose against each shipment. | [optional] 
**FromAddress** | Pointer to [**FromAddress**](FromAddress.md) |  | [optional] 
**LabelLayout** | Pointer to [**GetShipmentsForBatchDataInnerLabelLayout**](GetShipmentsForBatchDataInnerLabelLayout.md) |  | [optional] 
**Metadata** | Pointer to [**[]GetShipmentsForBatchDataInnerMetadataInner**](GetShipmentsForBatchDataInnerMetadataInner.md) | Additional metadata that needs to be stored for this shipment can be added here. For now, &#x60;costAccountName&#x60; is supported. | [optional] 
**Parcel** | Pointer to [**Parcel**](Parcel.md) |  | [optional] 
**ParcelType** | Pointer to **string** | Parcel Type is required for creating a shipment while rating a parcel. And it varies as per carrier selection and corresponding services. | [optional] 
**ServiceId** | Pointer to **string** | A unique identifier given to the carrier-specific service. User can override this value by defining it at Shipment level. | [optional] 
**ShipmentId** | Pointer to **string** | Shipment ID is a unique identifier for an individual shipment | [optional] 
**ShipmentIdentifier** | Pointer to **string** | Unique identifier generated for each shipment, it can be either success or failed. | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentOptions**](ShipmentOptions.md) |  | [optional] 
**SpecialServices** | Pointer to [**[]GetShipmentsForBatchDataInnerSpecialServicesInner**](GetShipmentsForBatchDataInnerSpecialServicesInner.md) | Special services used to create shipment | [optional] 
**StepStatus** | Pointer to [**GetShipmentsForBatchDataInnerStepStatus**](GetShipmentsForBatchDataInnerStepStatus.md) |  | [optional] 
**ToAddress** | Pointer to [**ToAddress**](ToAddress.md) |  | [optional] 

## Methods

### NewGetShipmentsForBatchDataInner

`func NewGetShipmentsForBatchDataInner() *GetShipmentsForBatchDataInner`

NewGetShipmentsForBatchDataInner instantiates a new GetShipmentsForBatchDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetShipmentsForBatchDataInnerWithDefaults

`func NewGetShipmentsForBatchDataInnerWithDefaults() *GetShipmentsForBatchDataInner`

NewGetShipmentsForBatchDataInnerWithDefaults instantiates a new GetShipmentsForBatchDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *GetShipmentsForBatchDataInner) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *GetShipmentsForBatchDataInner) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *GetShipmentsForBatchDataInner) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *GetShipmentsForBatchDataInner) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetCarrierAccountId

`func (o *GetShipmentsForBatchDataInner) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *GetShipmentsForBatchDataInner) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *GetShipmentsForBatchDataInner) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *GetShipmentsForBatchDataInner) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetExternalId

`func (o *GetShipmentsForBatchDataInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetShipmentsForBatchDataInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetShipmentsForBatchDataInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetShipmentsForBatchDataInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFromAddress

`func (o *GetShipmentsForBatchDataInner) GetFromAddress() FromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *GetShipmentsForBatchDataInner) GetFromAddressOk() (*FromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *GetShipmentsForBatchDataInner) SetFromAddress(v FromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *GetShipmentsForBatchDataInner) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetLabelLayout

`func (o *GetShipmentsForBatchDataInner) GetLabelLayout() GetShipmentsForBatchDataInnerLabelLayout`

GetLabelLayout returns the LabelLayout field if non-nil, zero value otherwise.

### GetLabelLayoutOk

`func (o *GetShipmentsForBatchDataInner) GetLabelLayoutOk() (*GetShipmentsForBatchDataInnerLabelLayout, bool)`

GetLabelLayoutOk returns a tuple with the LabelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelLayout

`func (o *GetShipmentsForBatchDataInner) SetLabelLayout(v GetShipmentsForBatchDataInnerLabelLayout)`

SetLabelLayout sets LabelLayout field to given value.

### HasLabelLayout

`func (o *GetShipmentsForBatchDataInner) HasLabelLayout() bool`

HasLabelLayout returns a boolean if a field has been set.

### GetMetadata

`func (o *GetShipmentsForBatchDataInner) GetMetadata() []GetShipmentsForBatchDataInnerMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetShipmentsForBatchDataInner) GetMetadataOk() (*[]GetShipmentsForBatchDataInnerMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetShipmentsForBatchDataInner) SetMetadata(v []GetShipmentsForBatchDataInnerMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetShipmentsForBatchDataInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetParcel

`func (o *GetShipmentsForBatchDataInner) GetParcel() Parcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *GetShipmentsForBatchDataInner) GetParcelOk() (*Parcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *GetShipmentsForBatchDataInner) SetParcel(v Parcel)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *GetShipmentsForBatchDataInner) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetParcelType

`func (o *GetShipmentsForBatchDataInner) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *GetShipmentsForBatchDataInner) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *GetShipmentsForBatchDataInner) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *GetShipmentsForBatchDataInner) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetServiceId

`func (o *GetShipmentsForBatchDataInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GetShipmentsForBatchDataInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GetShipmentsForBatchDataInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *GetShipmentsForBatchDataInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetShipmentId

`func (o *GetShipmentsForBatchDataInner) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *GetShipmentsForBatchDataInner) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *GetShipmentsForBatchDataInner) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *GetShipmentsForBatchDataInner) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetShipmentIdentifier

`func (o *GetShipmentsForBatchDataInner) GetShipmentIdentifier() string`

GetShipmentIdentifier returns the ShipmentIdentifier field if non-nil, zero value otherwise.

### GetShipmentIdentifierOk

`func (o *GetShipmentsForBatchDataInner) GetShipmentIdentifierOk() (*string, bool)`

GetShipmentIdentifierOk returns a tuple with the ShipmentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIdentifier

`func (o *GetShipmentsForBatchDataInner) SetShipmentIdentifier(v string)`

SetShipmentIdentifier sets ShipmentIdentifier field to given value.

### HasShipmentIdentifier

`func (o *GetShipmentsForBatchDataInner) HasShipmentIdentifier() bool`

HasShipmentIdentifier returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *GetShipmentsForBatchDataInner) GetShipmentOptions() ShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *GetShipmentsForBatchDataInner) GetShipmentOptionsOk() (*ShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *GetShipmentsForBatchDataInner) SetShipmentOptions(v ShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *GetShipmentsForBatchDataInner) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetSpecialServices

`func (o *GetShipmentsForBatchDataInner) GetSpecialServices() []GetShipmentsForBatchDataInnerSpecialServicesInner`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *GetShipmentsForBatchDataInner) GetSpecialServicesOk() (*[]GetShipmentsForBatchDataInnerSpecialServicesInner, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *GetShipmentsForBatchDataInner) SetSpecialServices(v []GetShipmentsForBatchDataInnerSpecialServicesInner)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *GetShipmentsForBatchDataInner) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetStepStatus

`func (o *GetShipmentsForBatchDataInner) GetStepStatus() GetShipmentsForBatchDataInnerStepStatus`

GetStepStatus returns the StepStatus field if non-nil, zero value otherwise.

### GetStepStatusOk

`func (o *GetShipmentsForBatchDataInner) GetStepStatusOk() (*GetShipmentsForBatchDataInnerStepStatus, bool)`

GetStepStatusOk returns a tuple with the StepStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepStatus

`func (o *GetShipmentsForBatchDataInner) SetStepStatus(v GetShipmentsForBatchDataInnerStepStatus)`

SetStepStatus sets StepStatus field to given value.

### HasStepStatus

`func (o *GetShipmentsForBatchDataInner) HasStepStatus() bool`

HasStepStatus returns a boolean if a field has been set.

### GetToAddress

`func (o *GetShipmentsForBatchDataInner) GetToAddress() ToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *GetShipmentsForBatchDataInner) GetToAddressOk() (*ToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *GetShipmentsForBatchDataInner) SetToAddress(v ToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *GetShipmentsForBatchDataInner) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


