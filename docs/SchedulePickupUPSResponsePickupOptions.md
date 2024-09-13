# SchedulePickupUPSResponsePickupOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PickupStartDateTime** | Pointer to **string** | It specifies the earliest date and time that your parcels will be available for pickup (UTC time) | [optional] 
**PickupEndDateTime** | Pointer to **string** | It specifies the latest date and time that your parcels will be available for pickup (UTC time) | [optional] 
**Overweight** | Pointer to **float32** | It specifies the count of packages which are overwight (weight&gt; 70 lbs) | [optional] 

## Methods

### NewSchedulePickupUPSResponsePickupOptions

`func NewSchedulePickupUPSResponsePickupOptions() *SchedulePickupUPSResponsePickupOptions`

NewSchedulePickupUPSResponsePickupOptions instantiates a new SchedulePickupUPSResponsePickupOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUPSResponsePickupOptionsWithDefaults

`func NewSchedulePickupUPSResponsePickupOptionsWithDefaults() *SchedulePickupUPSResponsePickupOptions`

NewSchedulePickupUPSResponsePickupOptionsWithDefaults instantiates a new SchedulePickupUPSResponsePickupOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPickupStartDateTime

`func (o *SchedulePickupUPSResponsePickupOptions) GetPickupStartDateTime() string`

GetPickupStartDateTime returns the PickupStartDateTime field if non-nil, zero value otherwise.

### GetPickupStartDateTimeOk

`func (o *SchedulePickupUPSResponsePickupOptions) GetPickupStartDateTimeOk() (*string, bool)`

GetPickupStartDateTimeOk returns a tuple with the PickupStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStartDateTime

`func (o *SchedulePickupUPSResponsePickupOptions) SetPickupStartDateTime(v string)`

SetPickupStartDateTime sets PickupStartDateTime field to given value.

### HasPickupStartDateTime

`func (o *SchedulePickupUPSResponsePickupOptions) HasPickupStartDateTime() bool`

HasPickupStartDateTime returns a boolean if a field has been set.

### GetPickupEndDateTime

`func (o *SchedulePickupUPSResponsePickupOptions) GetPickupEndDateTime() string`

GetPickupEndDateTime returns the PickupEndDateTime field if non-nil, zero value otherwise.

### GetPickupEndDateTimeOk

`func (o *SchedulePickupUPSResponsePickupOptions) GetPickupEndDateTimeOk() (*string, bool)`

GetPickupEndDateTimeOk returns a tuple with the PickupEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupEndDateTime

`func (o *SchedulePickupUPSResponsePickupOptions) SetPickupEndDateTime(v string)`

SetPickupEndDateTime sets PickupEndDateTime field to given value.

### HasPickupEndDateTime

`func (o *SchedulePickupUPSResponsePickupOptions) HasPickupEndDateTime() bool`

HasPickupEndDateTime returns a boolean if a field has been set.

### GetOverweight

`func (o *SchedulePickupUPSResponsePickupOptions) GetOverweight() float32`

GetOverweight returns the Overweight field if non-nil, zero value otherwise.

### GetOverweightOk

`func (o *SchedulePickupUPSResponsePickupOptions) GetOverweightOk() (*float32, bool)`

GetOverweightOk returns a tuple with the Overweight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverweight

`func (o *SchedulePickupUPSResponsePickupOptions) SetOverweight(v float32)`

SetOverweight sets Overweight field to given value.

### HasOverweight

`func (o *SchedulePickupUPSResponsePickupOptions) HasOverweight() bool`

HasOverweight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


