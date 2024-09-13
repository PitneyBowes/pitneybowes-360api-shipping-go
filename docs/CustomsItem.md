# CustomsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A detailed description of the commodity. | 
**HSTariffCode** | Pointer to **string** | &gt;- The destination country&#39;s tariff-classification number for the commodity. Most countries use the six-digit Harmonized System (HS) as the basis for their tariff classifications and add additional digits for more detail. The maximum length is 14 characters.If you are issuing the HS Code API, you can use this field to help with the HS code prediction by entering the commodity&#39;s HS code from another country, such as from the origin country. Enter the country that the code comes from in the hSTariffCodeCountry field. | [optional] 
**Quantity** | **int32** | Enter the total number of items of this type of commodity. | 
**UnitPrice** | **float64** | &gt;- The price of one item of this type of commodity. Currency should be declared under customsInfo. | 
**WeightUnit** | **string** | The unit of measurement. This field is required by the unitWeight object. | 
**Weight** | **float32** | The weight of the item. | 

## Methods

### NewCustomsItem

`func NewCustomsItem(description string, quantity int32, unitPrice float64, weightUnit string, weight float32, ) *CustomsItem`

NewCustomsItem instantiates a new CustomsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomsItemWithDefaults

`func NewCustomsItemWithDefaults() *CustomsItem`

NewCustomsItemWithDefaults instantiates a new CustomsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CustomsItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomsItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomsItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHSTariffCode

`func (o *CustomsItem) GetHSTariffCode() string`

GetHSTariffCode returns the HSTariffCode field if non-nil, zero value otherwise.

### GetHSTariffCodeOk

`func (o *CustomsItem) GetHSTariffCodeOk() (*string, bool)`

GetHSTariffCodeOk returns a tuple with the HSTariffCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSTariffCode

`func (o *CustomsItem) SetHSTariffCode(v string)`

SetHSTariffCode sets HSTariffCode field to given value.

### HasHSTariffCode

`func (o *CustomsItem) HasHSTariffCode() bool`

HasHSTariffCode returns a boolean if a field has been set.

### GetQuantity

`func (o *CustomsItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CustomsItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CustomsItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitPrice

`func (o *CustomsItem) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *CustomsItem) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *CustomsItem) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.


### GetWeightUnit

`func (o *CustomsItem) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *CustomsItem) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *CustomsItem) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.


### GetWeight

`func (o *CustomsItem) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CustomsItem) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CustomsItem) SetWeight(v float32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


