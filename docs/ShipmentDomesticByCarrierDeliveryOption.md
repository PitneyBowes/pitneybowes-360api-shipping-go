# ShipmentDomesticByCarrierDeliveryOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliverBy** | Pointer to **string** | Delivery date is the data when shipment is to be delivered, which is scheduled by sender. The format of the Date is YYYY-MM-DD. &lt;br /&gt; This field will be mandatory to provide, if the customer chooses ruleType is deliverBy. | [optional] 
**UseBestNextDate** | Pointer to **bool** | When this is set to true, if the scheduled delivery date falls on a Holiday, then the next business day will be considered to deliver the shipment. | [optional] 

## Methods

### NewShipmentDomesticByCarrierDeliveryOption

`func NewShipmentDomesticByCarrierDeliveryOption() *ShipmentDomesticByCarrierDeliveryOption`

NewShipmentDomesticByCarrierDeliveryOption instantiates a new ShipmentDomesticByCarrierDeliveryOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentDomesticByCarrierDeliveryOptionWithDefaults

`func NewShipmentDomesticByCarrierDeliveryOptionWithDefaults() *ShipmentDomesticByCarrierDeliveryOption`

NewShipmentDomesticByCarrierDeliveryOptionWithDefaults instantiates a new ShipmentDomesticByCarrierDeliveryOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliverBy

`func (o *ShipmentDomesticByCarrierDeliveryOption) GetDeliverBy() string`

GetDeliverBy returns the DeliverBy field if non-nil, zero value otherwise.

### GetDeliverByOk

`func (o *ShipmentDomesticByCarrierDeliveryOption) GetDeliverByOk() (*string, bool)`

GetDeliverByOk returns a tuple with the DeliverBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverBy

`func (o *ShipmentDomesticByCarrierDeliveryOption) SetDeliverBy(v string)`

SetDeliverBy sets DeliverBy field to given value.

### HasDeliverBy

`func (o *ShipmentDomesticByCarrierDeliveryOption) HasDeliverBy() bool`

HasDeliverBy returns a boolean if a field has been set.

### GetUseBestNextDate

`func (o *ShipmentDomesticByCarrierDeliveryOption) GetUseBestNextDate() bool`

GetUseBestNextDate returns the UseBestNextDate field if non-nil, zero value otherwise.

### GetUseBestNextDateOk

`func (o *ShipmentDomesticByCarrierDeliveryOption) GetUseBestNextDateOk() (*bool, bool)`

GetUseBestNextDateOk returns a tuple with the UseBestNextDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBestNextDate

`func (o *ShipmentDomesticByCarrierDeliveryOption) SetUseBestNextDate(v bool)`

SetUseBestNextDate sets UseBestNextDate field to given value.

### HasUseBestNextDate

`func (o *ShipmentDomesticByCarrierDeliveryOption) HasUseBestNextDate() bool`

HasUseBestNextDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


