# SchedulePickupUPSRequestPickupOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PickupStartDateTime** | **string** | It specifies the earliest date and time that your parcels will be available for pickup (UTC time) | 
**PickupEndDateTime** | **string** | It specifies the latest date and time that your parcels will be available for pickup (UTC time) | 
**Overweight** | Pointer to **float32** | It specifies the count of packages which are overwight (weight&gt; 70 lbs) | [optional] 

## Methods

### NewSchedulePickupUPSRequestPickupOptions

`func NewSchedulePickupUPSRequestPickupOptions(pickupStartDateTime string, pickupEndDateTime string, ) *SchedulePickupUPSRequestPickupOptions`

NewSchedulePickupUPSRequestPickupOptions instantiates a new SchedulePickupUPSRequestPickupOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUPSRequestPickupOptionsWithDefaults

`func NewSchedulePickupUPSRequestPickupOptionsWithDefaults() *SchedulePickupUPSRequestPickupOptions`

NewSchedulePickupUPSRequestPickupOptionsWithDefaults instantiates a new SchedulePickupUPSRequestPickupOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPickupStartDateTime

`func (o *SchedulePickupUPSRequestPickupOptions) GetPickupStartDateTime() string`

GetPickupStartDateTime returns the PickupStartDateTime field if non-nil, zero value otherwise.

### GetPickupStartDateTimeOk

`func (o *SchedulePickupUPSRequestPickupOptions) GetPickupStartDateTimeOk() (*string, bool)`

GetPickupStartDateTimeOk returns a tuple with the PickupStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStartDateTime

`func (o *SchedulePickupUPSRequestPickupOptions) SetPickupStartDateTime(v string)`

SetPickupStartDateTime sets PickupStartDateTime field to given value.


### GetPickupEndDateTime

`func (o *SchedulePickupUPSRequestPickupOptions) GetPickupEndDateTime() string`

GetPickupEndDateTime returns the PickupEndDateTime field if non-nil, zero value otherwise.

### GetPickupEndDateTimeOk

`func (o *SchedulePickupUPSRequestPickupOptions) GetPickupEndDateTimeOk() (*string, bool)`

GetPickupEndDateTimeOk returns a tuple with the PickupEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupEndDateTime

`func (o *SchedulePickupUPSRequestPickupOptions) SetPickupEndDateTime(v string)`

SetPickupEndDateTime sets PickupEndDateTime field to given value.


### GetOverweight

`func (o *SchedulePickupUPSRequestPickupOptions) GetOverweight() float32`

GetOverweight returns the Overweight field if non-nil, zero value otherwise.

### GetOverweightOk

`func (o *SchedulePickupUPSRequestPickupOptions) GetOverweightOk() (*float32, bool)`

GetOverweightOk returns a tuple with the Overweight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverweight

`func (o *SchedulePickupUPSRequestPickupOptions) SetOverweight(v float32)`

SetOverweight sets Overweight field to given value.

### HasOverweight

`func (o *SchedulePickupUPSRequestPickupOptions) HasOverweight() bool`

HasOverweight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


