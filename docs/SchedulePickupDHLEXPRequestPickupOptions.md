# SchedulePickupDHLEXPRequestPickupOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PickupStartDateTime** | **string** | It specifies the earliest date and time that your parcels will be available for pickup (UTC time) | 
**PickupEndDateTime** | **string** | It specifies the latest date and time that your parcels will be available for pickup (UTC time) | 

## Methods

### NewSchedulePickupDHLEXPRequestPickupOptions

`func NewSchedulePickupDHLEXPRequestPickupOptions(pickupStartDateTime string, pickupEndDateTime string, ) *SchedulePickupDHLEXPRequestPickupOptions`

NewSchedulePickupDHLEXPRequestPickupOptions instantiates a new SchedulePickupDHLEXPRequestPickupOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupDHLEXPRequestPickupOptionsWithDefaults

`func NewSchedulePickupDHLEXPRequestPickupOptionsWithDefaults() *SchedulePickupDHLEXPRequestPickupOptions`

NewSchedulePickupDHLEXPRequestPickupOptionsWithDefaults instantiates a new SchedulePickupDHLEXPRequestPickupOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPickupStartDateTime

`func (o *SchedulePickupDHLEXPRequestPickupOptions) GetPickupStartDateTime() string`

GetPickupStartDateTime returns the PickupStartDateTime field if non-nil, zero value otherwise.

### GetPickupStartDateTimeOk

`func (o *SchedulePickupDHLEXPRequestPickupOptions) GetPickupStartDateTimeOk() (*string, bool)`

GetPickupStartDateTimeOk returns a tuple with the PickupStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStartDateTime

`func (o *SchedulePickupDHLEXPRequestPickupOptions) SetPickupStartDateTime(v string)`

SetPickupStartDateTime sets PickupStartDateTime field to given value.


### GetPickupEndDateTime

`func (o *SchedulePickupDHLEXPRequestPickupOptions) GetPickupEndDateTime() string`

GetPickupEndDateTime returns the PickupEndDateTime field if non-nil, zero value otherwise.

### GetPickupEndDateTimeOk

`func (o *SchedulePickupDHLEXPRequestPickupOptions) GetPickupEndDateTimeOk() (*string, bool)`

GetPickupEndDateTimeOk returns a tuple with the PickupEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupEndDateTime

`func (o *SchedulePickupDHLEXPRequestPickupOptions) SetPickupEndDateTime(v string)`

SetPickupEndDateTime sets PickupEndDateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


