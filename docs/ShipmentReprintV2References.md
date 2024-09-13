# ShipmentReprintV2References

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalReference1** | Pointer to **string** | Additional Reference is hardly used, but sender can mention anything as per requirement, just for Recipient&#39;s information. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 
**AdditionalReference2** | Pointer to **string** | Any tags or information that to be shown to Recipient, can be mentioned by Sender, which is not indicated on AdditionalReference1 field, e.g. PO No, Order No. etc.&lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 

## Methods

### NewShipmentReprintV2References

`func NewShipmentReprintV2References() *ShipmentReprintV2References`

NewShipmentReprintV2References instantiates a new ShipmentReprintV2References object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentReprintV2ReferencesWithDefaults

`func NewShipmentReprintV2ReferencesWithDefaults() *ShipmentReprintV2References`

NewShipmentReprintV2ReferencesWithDefaults instantiates a new ShipmentReprintV2References object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalReference1

`func (o *ShipmentReprintV2References) GetAdditionalReference1() string`

GetAdditionalReference1 returns the AdditionalReference1 field if non-nil, zero value otherwise.

### GetAdditionalReference1Ok

`func (o *ShipmentReprintV2References) GetAdditionalReference1Ok() (*string, bool)`

GetAdditionalReference1Ok returns a tuple with the AdditionalReference1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalReference1

`func (o *ShipmentReprintV2References) SetAdditionalReference1(v string)`

SetAdditionalReference1 sets AdditionalReference1 field to given value.

### HasAdditionalReference1

`func (o *ShipmentReprintV2References) HasAdditionalReference1() bool`

HasAdditionalReference1 returns a boolean if a field has been set.

### GetAdditionalReference2

`func (o *ShipmentReprintV2References) GetAdditionalReference2() string`

GetAdditionalReference2 returns the AdditionalReference2 field if non-nil, zero value otherwise.

### GetAdditionalReference2Ok

`func (o *ShipmentReprintV2References) GetAdditionalReference2Ok() (*string, bool)`

GetAdditionalReference2Ok returns a tuple with the AdditionalReference2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalReference2

`func (o *ShipmentReprintV2References) SetAdditionalReference2(v string)`

SetAdditionalReference2 sets AdditionalReference2 field to given value.

### HasAdditionalReference2

`func (o *ShipmentReprintV2References) HasAdditionalReference2() bool`

HasAdditionalReference2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


