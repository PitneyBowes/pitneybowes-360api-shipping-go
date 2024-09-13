# ShipmentDomesticByRateGroupByRateGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | **string** | The Rule Type is a condition applied to RateGroup by Product side as per the customer requirement, which can have following options: Cheapest, Fastest, and deliverBy. &lt;br /&gt; If ruleType is deliverBy, then deliverBy date under deliveryOption will be mandatory to provide. | 
**RateGroupId** | **string** | This is a unique identifier assigned to the created RateGroup, which is used in the shipment creation. | 

## Methods

### NewShipmentDomesticByRateGroupByRateGroup

`func NewShipmentDomesticByRateGroupByRateGroup(ruleType string, rateGroupId string, ) *ShipmentDomesticByRateGroupByRateGroup`

NewShipmentDomesticByRateGroupByRateGroup instantiates a new ShipmentDomesticByRateGroupByRateGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentDomesticByRateGroupByRateGroupWithDefaults

`func NewShipmentDomesticByRateGroupByRateGroupWithDefaults() *ShipmentDomesticByRateGroupByRateGroup`

NewShipmentDomesticByRateGroupByRateGroupWithDefaults instantiates a new ShipmentDomesticByRateGroupByRateGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleType

`func (o *ShipmentDomesticByRateGroupByRateGroup) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *ShipmentDomesticByRateGroupByRateGroup) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *ShipmentDomesticByRateGroupByRateGroup) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetRateGroupId

`func (o *ShipmentDomesticByRateGroupByRateGroup) GetRateGroupId() string`

GetRateGroupId returns the RateGroupId field if non-nil, zero value otherwise.

### GetRateGroupIdOk

`func (o *ShipmentDomesticByRateGroupByRateGroup) GetRateGroupIdOk() (*string, bool)`

GetRateGroupIdOk returns a tuple with the RateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateGroupId

`func (o *ShipmentDomesticByRateGroupByRateGroup) SetRateGroupId(v string)`

SetRateGroupId sets RateGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


