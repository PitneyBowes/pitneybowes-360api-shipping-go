# ShipmentCancelV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentId** | **string** | The shipmentId is a unique identifier for an individual Shipment. | 
**ParcelTrackingNumber** | Pointer to **string** | The tracking number associated with one parcel in a shipment. The parcel tracking number can be used to track one specific parcel. | [optional] 
**References** | Pointer to [**ShipmentReprintV2References**](ShipmentReprintV2References.md) |  | [optional] 

## Methods

### NewShipmentCancelV2

`func NewShipmentCancelV2(shipmentId string, ) *ShipmentCancelV2`

NewShipmentCancelV2 instantiates a new ShipmentCancelV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentCancelV2WithDefaults

`func NewShipmentCancelV2WithDefaults() *ShipmentCancelV2`

NewShipmentCancelV2WithDefaults instantiates a new ShipmentCancelV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipmentId

`func (o *ShipmentCancelV2) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *ShipmentCancelV2) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *ShipmentCancelV2) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.


### GetParcelTrackingNumber

`func (o *ShipmentCancelV2) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *ShipmentCancelV2) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *ShipmentCancelV2) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *ShipmentCancelV2) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetReferences

`func (o *ShipmentCancelV2) GetReferences() ShipmentReprintV2References`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ShipmentCancelV2) GetReferencesOk() (*ShipmentReprintV2References, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ShipmentCancelV2) SetReferences(v ShipmentReprintV2References)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ShipmentCancelV2) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


