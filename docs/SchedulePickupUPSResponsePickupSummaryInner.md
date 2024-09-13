# SchedulePickupUPSResponsePickupSummaryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** | The service id | [optional] 
**PackageCount** | Pointer to **float32** | The total number of packages | [optional] 
**TotalWeight** | Pointer to **float32** | The total weight of the packages | [optional] 
**WeightUnit** | Pointer to **string** | Weight Unit, supported values are &#x60;OZ&#x60; and &#x60;GM&#x60; | [optional] 
**ParcelType** | Pointer to **string** | It indicates the parcel type, applicable values are- &#x60;PKG&#x60;,&#x60;LTR&#x60;,&#x60;TUBE&#x60;,&#x60;PACK&#x60;,&#x60;BOX&#x60;,&#x60;25KG&#x60;,&#x60;10KG&#x60;,&#x60;SMALL_EXP_BOX&#x60;,&#x60;MED_EXP_BOX&#x60;,&#x60;LG_EXP_BOX&#x60; | [optional] 
**ToAddressCountryCode** | Pointer to **string** | It indicates the 2 characters- ISO country code of recipient of the shipment. | [optional] 

## Methods

### NewSchedulePickupUPSResponsePickupSummaryInner

`func NewSchedulePickupUPSResponsePickupSummaryInner() *SchedulePickupUPSResponsePickupSummaryInner`

NewSchedulePickupUPSResponsePickupSummaryInner instantiates a new SchedulePickupUPSResponsePickupSummaryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUPSResponsePickupSummaryInnerWithDefaults

`func NewSchedulePickupUPSResponsePickupSummaryInnerWithDefaults() *SchedulePickupUPSResponsePickupSummaryInner`

NewSchedulePickupUPSResponsePickupSummaryInnerWithDefaults instantiates a new SchedulePickupUPSResponsePickupSummaryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetPackageCount

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetPackageCount() float32`

GetPackageCount returns the PackageCount field if non-nil, zero value otherwise.

### GetPackageCountOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetPackageCountOk() (*float32, bool)`

GetPackageCountOk returns a tuple with the PackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCount

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetPackageCount(v float32)`

SetPackageCount sets PackageCount field to given value.

### HasPackageCount

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasPackageCount() bool`

HasPackageCount returns a boolean if a field has been set.

### GetTotalWeight

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetTotalWeight() float32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetTotalWeightOk() (*float32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetTotalWeight(v float32)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetParcelType

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetToAddressCountryCode

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetToAddressCountryCode() string`

GetToAddressCountryCode returns the ToAddressCountryCode field if non-nil, zero value otherwise.

### GetToAddressCountryCodeOk

`func (o *SchedulePickupUPSResponsePickupSummaryInner) GetToAddressCountryCodeOk() (*string, bool)`

GetToAddressCountryCodeOk returns a tuple with the ToAddressCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddressCountryCode

`func (o *SchedulePickupUPSResponsePickupSummaryInner) SetToAddressCountryCode(v string)`

SetToAddressCountryCode sets ToAddressCountryCode field to given value.

### HasToAddressCountryCode

`func (o *SchedulePickupUPSResponsePickupSummaryInner) HasToAddressCountryCode() bool`

HasToAddressCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


