# GetSingleShipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | This is a GUID (globally unique identifier) that&#39;s automatically generated for every request that the webserver receives. | [optional] 
**FromAddress** | Pointer to [**GetSingleShipmentFromAddress**](GetSingleShipmentFromAddress.md) |  | [optional] 
**Parcel** | Pointer to [**ShipmentDomesticParcel**](ShipmentDomesticParcel.md) |  | [optional] 
**Metadata** | Pointer to [**[]GetAllShipmentsDataInnerMetadataInner**](GetAllShipmentsDataInnerMetadataInner.md) | Additional metadata that needs to be stored for this shipment can be added here. For now, &#x60;costAccountName&#x60; is supported. | [optional] 
**ParcelTrackingNumber** | Pointer to **string** | The Tracking number given to the Parcel for tracking purpose. | [optional] 
**Rate** | Pointer to [**GetSingleShipmentRate**](GetSingleShipmentRate.md) |  | [optional] 
**Service** | Pointer to **string** | This indicates the carrier based service that is used for shipment. | [optional] 
**ShipmentId** | Pointer to **string** | A unique identifier associated with Shipment ID. | [optional] 
**Status** | Pointer to **string** | The status of the Shipment. | [optional] 
**ToAddress** | Pointer to [**GetSingleShipmentToAddress**](GetSingleShipmentToAddress.md) |  | [optional] 

## Methods

### NewGetSingleShipment

`func NewGetSingleShipment() *GetSingleShipment`

NewGetSingleShipment instantiates a new GetSingleShipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleShipmentWithDefaults

`func NewGetSingleShipmentWithDefaults() *GetSingleShipment`

NewGetSingleShipmentWithDefaults instantiates a new GetSingleShipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *GetSingleShipment) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *GetSingleShipment) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *GetSingleShipment) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *GetSingleShipment) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetFromAddress

`func (o *GetSingleShipment) GetFromAddress() GetSingleShipmentFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *GetSingleShipment) GetFromAddressOk() (*GetSingleShipmentFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *GetSingleShipment) SetFromAddress(v GetSingleShipmentFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *GetSingleShipment) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcel

`func (o *GetSingleShipment) GetParcel() ShipmentDomesticParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *GetSingleShipment) GetParcelOk() (*ShipmentDomesticParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *GetSingleShipment) SetParcel(v ShipmentDomesticParcel)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *GetSingleShipment) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetMetadata

`func (o *GetSingleShipment) GetMetadata() []GetAllShipmentsDataInnerMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetSingleShipment) GetMetadataOk() (*[]GetAllShipmentsDataInnerMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetSingleShipment) SetMetadata(v []GetAllShipmentsDataInnerMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetSingleShipment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *GetSingleShipment) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *GetSingleShipment) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *GetSingleShipment) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *GetSingleShipment) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetRate

`func (o *GetSingleShipment) GetRate() GetSingleShipmentRate`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetSingleShipment) GetRateOk() (*GetSingleShipmentRate, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetSingleShipment) SetRate(v GetSingleShipmentRate)`

SetRate sets Rate field to given value.

### HasRate

`func (o *GetSingleShipment) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetService

`func (o *GetSingleShipment) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetSingleShipment) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetSingleShipment) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetSingleShipment) HasService() bool`

HasService returns a boolean if a field has been set.

### GetShipmentId

`func (o *GetSingleShipment) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *GetSingleShipment) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *GetSingleShipment) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *GetSingleShipment) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetStatus

`func (o *GetSingleShipment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSingleShipment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSingleShipment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSingleShipment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToAddress

`func (o *GetSingleShipment) GetToAddress() GetSingleShipmentToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *GetSingleShipment) GetToAddressOk() (*GetSingleShipmentToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *GetSingleShipment) SetToAddress(v GetSingleShipmentToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *GetSingleShipment) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


