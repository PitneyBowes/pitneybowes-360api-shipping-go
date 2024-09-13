# SchedulePickupUPSRequestPickupSummaryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | The service id | 
**PackageCount** | **float32** | The total number of packages | 
**TotalWeight** | **float32** | The total weight of the packages | 
**WeightUnit** | **string** | Weight Unit, supported values are &#x60;OZ&#x60; and &#x60;GM&#x60; | 
**ParcelType** | **string** | It indicates the parcel type, applicable values are- &#x60;PKG&#x60;,&#x60;LTR&#x60;,&#x60;TUBE&#x60;,&#x60;PACK&#x60;,&#x60;BOX&#x60;,&#x60;25KG&#x60;,&#x60;10KG&#x60;,&#x60;SMALL_EXP_BOX&#x60;,&#x60;MED_EXP_BOX&#x60;,&#x60;LG_EXP_BOX&#x60; | 
**ToAddressCountryCode** | **string** | It indicates the 2 characters- ISO country code of recipient of the shipment. | 

## Methods

### NewSchedulePickupUPSRequestPickupSummaryInner

`func NewSchedulePickupUPSRequestPickupSummaryInner(serviceId string, packageCount float32, totalWeight float32, weightUnit string, parcelType string, toAddressCountryCode string, ) *SchedulePickupUPSRequestPickupSummaryInner`

NewSchedulePickupUPSRequestPickupSummaryInner instantiates a new SchedulePickupUPSRequestPickupSummaryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUPSRequestPickupSummaryInnerWithDefaults

`func NewSchedulePickupUPSRequestPickupSummaryInnerWithDefaults() *SchedulePickupUPSRequestPickupSummaryInner`

NewSchedulePickupUPSRequestPickupSummaryInnerWithDefaults instantiates a new SchedulePickupUPSRequestPickupSummaryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetPackageCount

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetPackageCount() float32`

GetPackageCount returns the PackageCount field if non-nil, zero value otherwise.

### GetPackageCountOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetPackageCountOk() (*float32, bool)`

GetPackageCountOk returns a tuple with the PackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCount

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetPackageCount(v float32)`

SetPackageCount sets PackageCount field to given value.


### GetTotalWeight

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetTotalWeight() float32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetTotalWeightOk() (*float32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetTotalWeight(v float32)`

SetTotalWeight sets TotalWeight field to given value.


### GetWeightUnit

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.


### GetParcelType

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.


### GetToAddressCountryCode

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetToAddressCountryCode() string`

GetToAddressCountryCode returns the ToAddressCountryCode field if non-nil, zero value otherwise.

### GetToAddressCountryCodeOk

`func (o *SchedulePickupUPSRequestPickupSummaryInner) GetToAddressCountryCodeOk() (*string, bool)`

GetToAddressCountryCodeOk returns a tuple with the ToAddressCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddressCountryCode

`func (o *SchedulePickupUPSRequestPickupSummaryInner) SetToAddressCountryCode(v string)`

SetToAddressCountryCode sets ToAddressCountryCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


