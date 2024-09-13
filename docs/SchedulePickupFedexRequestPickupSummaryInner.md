# SchedulePickupFedexRequestPickupSummaryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | The service id | 
**PackageCount** | **float32** | The total number of packages | 
**TotalWeight** | **float32** | The total weight of the packages | 
**WeightUnit** | **string** | Weight Unit, supported values are &#x60;OZ&#x60; and &#x60;GM&#x60; | 

## Methods

### NewSchedulePickupFedexRequestPickupSummaryInner

`func NewSchedulePickupFedexRequestPickupSummaryInner(serviceId string, packageCount float32, totalWeight float32, weightUnit string, ) *SchedulePickupFedexRequestPickupSummaryInner`

NewSchedulePickupFedexRequestPickupSummaryInner instantiates a new SchedulePickupFedexRequestPickupSummaryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupFedexRequestPickupSummaryInnerWithDefaults

`func NewSchedulePickupFedexRequestPickupSummaryInnerWithDefaults() *SchedulePickupFedexRequestPickupSummaryInner`

NewSchedulePickupFedexRequestPickupSummaryInnerWithDefaults instantiates a new SchedulePickupFedexRequestPickupSummaryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *SchedulePickupFedexRequestPickupSummaryInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SchedulePickupFedexRequestPickupSummaryInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SchedulePickupFedexRequestPickupSummaryInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetPackageCount

`func (o *SchedulePickupFedexRequestPickupSummaryInner) GetPackageCount() float32`

GetPackageCount returns the PackageCount field if non-nil, zero value otherwise.

### GetPackageCountOk

`func (o *SchedulePickupFedexRequestPickupSummaryInner) GetPackageCountOk() (*float32, bool)`

GetPackageCountOk returns a tuple with the PackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCount

`func (o *SchedulePickupFedexRequestPickupSummaryInner) SetPackageCount(v float32)`

SetPackageCount sets PackageCount field to given value.


### GetTotalWeight

`func (o *SchedulePickupFedexRequestPickupSummaryInner) GetTotalWeight() float32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *SchedulePickupFedexRequestPickupSummaryInner) GetTotalWeightOk() (*float32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *SchedulePickupFedexRequestPickupSummaryInner) SetTotalWeight(v float32)`

SetTotalWeight sets TotalWeight field to given value.


### GetWeightUnit

`func (o *SchedulePickupFedexRequestPickupSummaryInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *SchedulePickupFedexRequestPickupSummaryInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *SchedulePickupFedexRequestPickupSummaryInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


