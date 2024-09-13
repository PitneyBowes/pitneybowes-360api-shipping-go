# SchedulePickupDHLEXPResponsePickupSummaryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** | The service id | [optional] 
**PackageCount** | Pointer to **float32** | The total number of packages | [optional] 
**TotalWeight** | Pointer to **float32** | The total weight of the packages | [optional] 
**WeightUnit** | Pointer to **string** | Weight Unit, supported values are &#x60;OZ&#x60; and &#x60;GM&#x60; | [optional] 
**CurrencyCode** | Pointer to **string** | Currency code | [optional] 
**TotalCustomsDeclaredValue** | Pointer to **string** | It indicates the custom declared value. It is required in case of international shipment. | [optional] 

## Methods

### NewSchedulePickupDHLEXPResponsePickupSummaryInner

`func NewSchedulePickupDHLEXPResponsePickupSummaryInner() *SchedulePickupDHLEXPResponsePickupSummaryInner`

NewSchedulePickupDHLEXPResponsePickupSummaryInner instantiates a new SchedulePickupDHLEXPResponsePickupSummaryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupDHLEXPResponsePickupSummaryInnerWithDefaults

`func NewSchedulePickupDHLEXPResponsePickupSummaryInnerWithDefaults() *SchedulePickupDHLEXPResponsePickupSummaryInner`

NewSchedulePickupDHLEXPResponsePickupSummaryInnerWithDefaults instantiates a new SchedulePickupDHLEXPResponsePickupSummaryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetPackageCount

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) GetPackageCount() float32`

GetPackageCount returns the PackageCount field if non-nil, zero value otherwise.

### GetPackageCountOk

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) GetPackageCountOk() (*float32, bool)`

GetPackageCountOk returns a tuple with the PackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCount

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) SetPackageCount(v float32)`

SetPackageCount sets PackageCount field to given value.

### HasPackageCount

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) HasPackageCount() bool`

HasPackageCount returns a boolean if a field has been set.

### GetTotalWeight

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) GetTotalWeight() float32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) GetTotalWeightOk() (*float32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) SetTotalWeight(v float32)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetTotalCustomsDeclaredValue

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) GetTotalCustomsDeclaredValue() string`

GetTotalCustomsDeclaredValue returns the TotalCustomsDeclaredValue field if non-nil, zero value otherwise.

### GetTotalCustomsDeclaredValueOk

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) GetTotalCustomsDeclaredValueOk() (*string, bool)`

GetTotalCustomsDeclaredValueOk returns a tuple with the TotalCustomsDeclaredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCustomsDeclaredValue

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) SetTotalCustomsDeclaredValue(v string)`

SetTotalCustomsDeclaredValue sets TotalCustomsDeclaredValue field to given value.

### HasTotalCustomsDeclaredValue

`func (o *SchedulePickupDHLEXPResponsePickupSummaryInner) HasTotalCustomsDeclaredValue() bool`

HasTotalCustomsDeclaredValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


