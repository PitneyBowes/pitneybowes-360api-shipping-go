# ShipmentReprintV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentId** | **string** | The shipmentId is a unique identifier for an individual Shipment. | 
**PrinterAliasName** | Pointer to **string** | Refers to a printer connected (directly or via network) to a computer. &#x60;Max length &#x3D; 60&#x60; | [optional] 
**References** | Pointer to [**ShipmentReprintV2References**](ShipmentReprintV2References.md) |  | [optional] 

## Methods

### NewShipmentReprintV2

`func NewShipmentReprintV2(shipmentId string, ) *ShipmentReprintV2`

NewShipmentReprintV2 instantiates a new ShipmentReprintV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentReprintV2WithDefaults

`func NewShipmentReprintV2WithDefaults() *ShipmentReprintV2`

NewShipmentReprintV2WithDefaults instantiates a new ShipmentReprintV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipmentId

`func (o *ShipmentReprintV2) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *ShipmentReprintV2) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *ShipmentReprintV2) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.


### GetPrinterAliasName

`func (o *ShipmentReprintV2) GetPrinterAliasName() string`

GetPrinterAliasName returns the PrinterAliasName field if non-nil, zero value otherwise.

### GetPrinterAliasNameOk

`func (o *ShipmentReprintV2) GetPrinterAliasNameOk() (*string, bool)`

GetPrinterAliasNameOk returns a tuple with the PrinterAliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinterAliasName

`func (o *ShipmentReprintV2) SetPrinterAliasName(v string)`

SetPrinterAliasName sets PrinterAliasName field to given value.

### HasPrinterAliasName

`func (o *ShipmentReprintV2) HasPrinterAliasName() bool`

HasPrinterAliasName returns a boolean if a field has been set.

### GetReferences

`func (o *ShipmentReprintV2) GetReferences() ShipmentReprintV2References`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ShipmentReprintV2) GetReferencesOk() (*ShipmentReprintV2References, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ShipmentReprintV2) SetReferences(v ShipmentReprintV2References)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ShipmentReprintV2) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


