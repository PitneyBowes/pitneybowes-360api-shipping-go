# SchedulePickupDHLEXPRequestPickupSummaryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | The service id | 
**PackageCount** | **float32** | The total number of packages | 
**TotalWeight** | **float32** | The total weight of the packages | 
**WeightUnit** | **string** | Weight Unit, supported values are &#x60;OZ&#x60; and &#x60;GM&#x60; | 
**CurrencyCode** | **string** | Currency code | 
**TotalCustomsDeclaredValue** | Pointer to **float32** | It indicates the custom declared value. It is required in case of international shipment. | [optional] 
**PackageDetails** | [**[]SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner**](SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner.md) | It descibes the individual package details | 

## Methods

### NewSchedulePickupDHLEXPRequestPickupSummaryInner

`func NewSchedulePickupDHLEXPRequestPickupSummaryInner(serviceId string, packageCount float32, totalWeight float32, weightUnit string, currencyCode string, packageDetails []SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner, ) *SchedulePickupDHLEXPRequestPickupSummaryInner`

NewSchedulePickupDHLEXPRequestPickupSummaryInner instantiates a new SchedulePickupDHLEXPRequestPickupSummaryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupDHLEXPRequestPickupSummaryInnerWithDefaults

`func NewSchedulePickupDHLEXPRequestPickupSummaryInnerWithDefaults() *SchedulePickupDHLEXPRequestPickupSummaryInner`

NewSchedulePickupDHLEXPRequestPickupSummaryInnerWithDefaults instantiates a new SchedulePickupDHLEXPRequestPickupSummaryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetPackageCount

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetPackageCount() float32`

GetPackageCount returns the PackageCount field if non-nil, zero value otherwise.

### GetPackageCountOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetPackageCountOk() (*float32, bool)`

GetPackageCountOk returns a tuple with the PackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCount

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetPackageCount(v float32)`

SetPackageCount sets PackageCount field to given value.


### GetTotalWeight

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetTotalWeight() float32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetTotalWeightOk() (*float32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetTotalWeight(v float32)`

SetTotalWeight sets TotalWeight field to given value.


### GetWeightUnit

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.


### GetCurrencyCode

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetTotalCustomsDeclaredValue

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetTotalCustomsDeclaredValue() float32`

GetTotalCustomsDeclaredValue returns the TotalCustomsDeclaredValue field if non-nil, zero value otherwise.

### GetTotalCustomsDeclaredValueOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetTotalCustomsDeclaredValueOk() (*float32, bool)`

GetTotalCustomsDeclaredValueOk returns a tuple with the TotalCustomsDeclaredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCustomsDeclaredValue

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetTotalCustomsDeclaredValue(v float32)`

SetTotalCustomsDeclaredValue sets TotalCustomsDeclaredValue field to given value.

### HasTotalCustomsDeclaredValue

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) HasTotalCustomsDeclaredValue() bool`

HasTotalCustomsDeclaredValue returns a boolean if a field has been set.

### GetPackageDetails

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetPackageDetails() []SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner`

GetPackageDetails returns the PackageDetails field if non-nil, zero value otherwise.

### GetPackageDetailsOk

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) GetPackageDetailsOk() (*[]SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner, bool)`

GetPackageDetailsOk returns a tuple with the PackageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDetails

`func (o *SchedulePickupDHLEXPRequestPickupSummaryInner) SetPackageDetails(v []SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner)`

SetPackageDetails sets PackageDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


