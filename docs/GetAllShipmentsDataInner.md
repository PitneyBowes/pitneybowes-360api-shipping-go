# GetAllShipmentsDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | This is a GUID (globally unique identifier) that&#39;s automatically generated for every request that the webserver receives. | [optional] 
**FromAddress** | Pointer to [**GetAllShipmentsDataInnerFromAddress**](GetAllShipmentsDataInnerFromAddress.md) |  | [optional] 
**Parcel** | Pointer to [**GetAllShipmentsDataInnerParcel**](GetAllShipmentsDataInnerParcel.md) |  | [optional] 
**Metadata** | Pointer to [**[]GetAllShipmentsDataInnerMetadataInner**](GetAllShipmentsDataInnerMetadataInner.md) | Additional metadata that needs to be stored for this shipment can be added here. For now, &#x60;costAccountName&#x60; is supported. | [optional] 
**ParcelTrackingNumber** | Pointer to **string** | The Tracking number given to the Parcel for tracking purpose. | [optional] 
**Rate** | Pointer to [**GetAllShipmentsDataInnerRate**](GetAllShipmentsDataInnerRate.md) |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**ShipmentId** | Pointer to **string** | A unique identifier associated with the Shipment. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ToAddress** | Pointer to [**GetAllShipmentsDataInnerToAddress**](GetAllShipmentsDataInnerToAddress.md) |  | [optional] 

## Methods

### NewGetAllShipmentsDataInner

`func NewGetAllShipmentsDataInner() *GetAllShipmentsDataInner`

NewGetAllShipmentsDataInner instantiates a new GetAllShipmentsDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllShipmentsDataInnerWithDefaults

`func NewGetAllShipmentsDataInnerWithDefaults() *GetAllShipmentsDataInner`

NewGetAllShipmentsDataInnerWithDefaults instantiates a new GetAllShipmentsDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *GetAllShipmentsDataInner) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *GetAllShipmentsDataInner) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *GetAllShipmentsDataInner) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *GetAllShipmentsDataInner) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetFromAddress

`func (o *GetAllShipmentsDataInner) GetFromAddress() GetAllShipmentsDataInnerFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *GetAllShipmentsDataInner) GetFromAddressOk() (*GetAllShipmentsDataInnerFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *GetAllShipmentsDataInner) SetFromAddress(v GetAllShipmentsDataInnerFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *GetAllShipmentsDataInner) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcel

`func (o *GetAllShipmentsDataInner) GetParcel() GetAllShipmentsDataInnerParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *GetAllShipmentsDataInner) GetParcelOk() (*GetAllShipmentsDataInnerParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *GetAllShipmentsDataInner) SetParcel(v GetAllShipmentsDataInnerParcel)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *GetAllShipmentsDataInner) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetMetadata

`func (o *GetAllShipmentsDataInner) GetMetadata() []GetAllShipmentsDataInnerMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetAllShipmentsDataInner) GetMetadataOk() (*[]GetAllShipmentsDataInnerMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetAllShipmentsDataInner) SetMetadata(v []GetAllShipmentsDataInnerMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetAllShipmentsDataInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *GetAllShipmentsDataInner) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *GetAllShipmentsDataInner) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *GetAllShipmentsDataInner) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *GetAllShipmentsDataInner) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetRate

`func (o *GetAllShipmentsDataInner) GetRate() GetAllShipmentsDataInnerRate`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetAllShipmentsDataInner) GetRateOk() (*GetAllShipmentsDataInnerRate, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetAllShipmentsDataInner) SetRate(v GetAllShipmentsDataInnerRate)`

SetRate sets Rate field to given value.

### HasRate

`func (o *GetAllShipmentsDataInner) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetService

`func (o *GetAllShipmentsDataInner) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetAllShipmentsDataInner) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetAllShipmentsDataInner) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetAllShipmentsDataInner) HasService() bool`

HasService returns a boolean if a field has been set.

### GetShipmentId

`func (o *GetAllShipmentsDataInner) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *GetAllShipmentsDataInner) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *GetAllShipmentsDataInner) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *GetAllShipmentsDataInner) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetStatus

`func (o *GetAllShipmentsDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAllShipmentsDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAllShipmentsDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAllShipmentsDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToAddress

`func (o *GetAllShipmentsDataInner) GetToAddress() GetAllShipmentsDataInnerToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *GetAllShipmentsDataInner) GetToAddressOk() (*GetAllShipmentsDataInnerToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *GetAllShipmentsDataInner) SetToAddress(v GetAllShipmentsDataInnerToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *GetAllShipmentsDataInner) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


