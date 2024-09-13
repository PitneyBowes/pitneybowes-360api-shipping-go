# ShipmentInternationalCustomsCustomsItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A detailed description of the commodity, up to 255 characters. The description will appear on the customs form. If the shipment has multiple types of items, create a separate customsItems object for each. Each description will appear on the form. | 
**HSTariffCode** | Pointer to **string** | The destination country’s tariff-classification number (HS code) for the commodity. Most countries use the six-digit Harmonized System (HS) as the basis for their tariff classifications and then add digits for more detail. The maximum length for an HS code is 14 characters. The HS code will appear on the customs form. If the shipment has multiple types of items, create a separate customsItems object for each. | [optional] 
**OriginCountryCode** | Pointer to **string** | The two-character ISO country code of the shipment’s origin country. Use ISO 3166-1 alpha-2 standard values. | [optional] 
**Quantity** | **float32** | Enter the total number of items of this type of commodity. | 
**UnitPrice** | **float32** | The price of one item of this type of commodity. | 
**WeightUnit** | Pointer to **string** | The unit of measurement. This field is required by the unitWeight object. | [optional] 
**Weight** | Pointer to **float32** | The weight of the item. | [optional] 

## Methods

### NewShipmentInternationalCustomsCustomsItemsInner

`func NewShipmentInternationalCustomsCustomsItemsInner(description string, quantity float32, unitPrice float32, ) *ShipmentInternationalCustomsCustomsItemsInner`

NewShipmentInternationalCustomsCustomsItemsInner instantiates a new ShipmentInternationalCustomsCustomsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentInternationalCustomsCustomsItemsInnerWithDefaults

`func NewShipmentInternationalCustomsCustomsItemsInnerWithDefaults() *ShipmentInternationalCustomsCustomsItemsInner`

NewShipmentInternationalCustomsCustomsItemsInnerWithDefaults instantiates a new ShipmentInternationalCustomsCustomsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShipmentInternationalCustomsCustomsItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHSTariffCode

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetHSTariffCode() string`

GetHSTariffCode returns the HSTariffCode field if non-nil, zero value otherwise.

### GetHSTariffCodeOk

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetHSTariffCodeOk() (*string, bool)`

GetHSTariffCodeOk returns a tuple with the HSTariffCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSTariffCode

`func (o *ShipmentInternationalCustomsCustomsItemsInner) SetHSTariffCode(v string)`

SetHSTariffCode sets HSTariffCode field to given value.

### HasHSTariffCode

`func (o *ShipmentInternationalCustomsCustomsItemsInner) HasHSTariffCode() bool`

HasHSTariffCode returns a boolean if a field has been set.

### GetOriginCountryCode

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetOriginCountryCode() string`

GetOriginCountryCode returns the OriginCountryCode field if non-nil, zero value otherwise.

### GetOriginCountryCodeOk

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetOriginCountryCodeOk() (*string, bool)`

GetOriginCountryCodeOk returns a tuple with the OriginCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountryCode

`func (o *ShipmentInternationalCustomsCustomsItemsInner) SetOriginCountryCode(v string)`

SetOriginCountryCode sets OriginCountryCode field to given value.

### HasOriginCountryCode

`func (o *ShipmentInternationalCustomsCustomsItemsInner) HasOriginCountryCode() bool`

HasOriginCountryCode returns a boolean if a field has been set.

### GetQuantity

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ShipmentInternationalCustomsCustomsItemsInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitPrice

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *ShipmentInternationalCustomsCustomsItemsInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.


### GetWeightUnit

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ShipmentInternationalCustomsCustomsItemsInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *ShipmentInternationalCustomsCustomsItemsInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetWeight

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ShipmentInternationalCustomsCustomsItemsInner) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ShipmentInternationalCustomsCustomsItemsInner) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ShipmentInternationalCustomsCustomsItemsInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


