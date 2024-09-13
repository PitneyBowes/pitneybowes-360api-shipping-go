# ShipmentERRMetadataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The key of metadata is CostAccount. User needs to enter the name of the CostAccount, e.g. PBPUNE_COSTACC. | [optional] 
**Value** | Pointer to **string** | User needs to enter the CostAccount Code, e.g. CAC-00PBPUN. | [optional] 

## Methods

### NewShipmentERRMetadataInner

`func NewShipmentERRMetadataInner() *ShipmentERRMetadataInner`

NewShipmentERRMetadataInner instantiates a new ShipmentERRMetadataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentERRMetadataInnerWithDefaults

`func NewShipmentERRMetadataInnerWithDefaults() *ShipmentERRMetadataInner`

NewShipmentERRMetadataInnerWithDefaults instantiates a new ShipmentERRMetadataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ShipmentERRMetadataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipmentERRMetadataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipmentERRMetadataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipmentERRMetadataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ShipmentERRMetadataInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ShipmentERRMetadataInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ShipmentERRMetadataInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ShipmentERRMetadataInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


