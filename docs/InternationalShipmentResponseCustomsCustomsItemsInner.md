# InternationalShipmentResponseCustomsCustomsItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A detailed description of the commodity, up to 255 characters. The description will appear on the customs form. If the shipment has multiple types of items, create a separate customsItems object for each. Each description will appear on the form. | [optional] 
**HSTariffCode** | Pointer to **string** | The destination country’s tariff-classification number (HS code) for the commodity. Most countries use the six-digit Harmonized System (HS) as the basis for their tariff classifications and then add digits for more detail. The maximum length for an HS code is 14 characters. The HS code will appear on the customs form. If the shipment has multiple types of items, create a separate customsItems object for each. | [optional] 
**OriginCountryCode** | Pointer to **string** | The two-character ISO country code of the shipment’s origin country. Use ISO 3166-1 alpha-2 standard values. | [optional] 
**Quantity** | Pointer to **float32** | Enter the total number of items of this type of commodity. | [optional] 
**UnitPrice** | Pointer to **float32** | The price of one item of this type of commodity. | [optional] 
**WeightUnit** | Pointer to **string** | WeightUnit is a standard for measuring the physical quantities of specified weight. &lt;br&gt; The valid values are:  &lt;br&gt;- OZ: Ounces &lt;br&gt;- GM: Grams &lt;br&gt;- LB: Pounds &lt;br&gt;- KG: Kilograms For USPS shipments, set this to OZ. | [optional] 
**Weight** | Pointer to **float32** | Weight is the measure of how heavy an object is. | [optional] 

## Methods

### NewInternationalShipmentResponseCustomsCustomsItemsInner

`func NewInternationalShipmentResponseCustomsCustomsItemsInner() *InternationalShipmentResponseCustomsCustomsItemsInner`

NewInternationalShipmentResponseCustomsCustomsItemsInner instantiates a new InternationalShipmentResponseCustomsCustomsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalShipmentResponseCustomsCustomsItemsInnerWithDefaults

`func NewInternationalShipmentResponseCustomsCustomsItemsInnerWithDefaults() *InternationalShipmentResponseCustomsCustomsItemsInner`

NewInternationalShipmentResponseCustomsCustomsItemsInnerWithDefaults instantiates a new InternationalShipmentResponseCustomsCustomsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHSTariffCode

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetHSTariffCode() string`

GetHSTariffCode returns the HSTariffCode field if non-nil, zero value otherwise.

### GetHSTariffCodeOk

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetHSTariffCodeOk() (*string, bool)`

GetHSTariffCodeOk returns a tuple with the HSTariffCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSTariffCode

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) SetHSTariffCode(v string)`

SetHSTariffCode sets HSTariffCode field to given value.

### HasHSTariffCode

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) HasHSTariffCode() bool`

HasHSTariffCode returns a boolean if a field has been set.

### GetOriginCountryCode

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetOriginCountryCode() string`

GetOriginCountryCode returns the OriginCountryCode field if non-nil, zero value otherwise.

### GetOriginCountryCodeOk

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetOriginCountryCodeOk() (*string, bool)`

GetOriginCountryCodeOk returns a tuple with the OriginCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountryCode

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) SetOriginCountryCode(v string)`

SetOriginCountryCode sets OriginCountryCode field to given value.

### HasOriginCountryCode

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) HasOriginCountryCode() bool`

HasOriginCountryCode returns a boolean if a field has been set.

### GetQuantity

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitPrice

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetWeightUnit

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetWeight

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *InternationalShipmentResponseCustomsCustomsItemsInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


