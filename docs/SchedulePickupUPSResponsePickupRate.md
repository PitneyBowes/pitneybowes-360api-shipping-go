# SchedulePickupUPSResponsePickupRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCharge** | Pointer to **float32** | It displays the base charge for this pickup | [optional] 
**Surcharges** | Pointer to [**[]SchedulePickupUPSResponsePickupRateSurchargesInner**](SchedulePickupUPSResponsePickupRateSurchargesInner.md) | It displays the surcharges if any for the pickup | [optional] 
**TotalCarrierCharge** | Pointer to **float32** | It displays the total charge for the pickup including surcharges. | [optional] 

## Methods

### NewSchedulePickupUPSResponsePickupRate

`func NewSchedulePickupUPSResponsePickupRate() *SchedulePickupUPSResponsePickupRate`

NewSchedulePickupUPSResponsePickupRate instantiates a new SchedulePickupUPSResponsePickupRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUPSResponsePickupRateWithDefaults

`func NewSchedulePickupUPSResponsePickupRateWithDefaults() *SchedulePickupUPSResponsePickupRate`

NewSchedulePickupUPSResponsePickupRateWithDefaults instantiates a new SchedulePickupUPSResponsePickupRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCharge

`func (o *SchedulePickupUPSResponsePickupRate) GetBaseCharge() float32`

GetBaseCharge returns the BaseCharge field if non-nil, zero value otherwise.

### GetBaseChargeOk

`func (o *SchedulePickupUPSResponsePickupRate) GetBaseChargeOk() (*float32, bool)`

GetBaseChargeOk returns a tuple with the BaseCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCharge

`func (o *SchedulePickupUPSResponsePickupRate) SetBaseCharge(v float32)`

SetBaseCharge sets BaseCharge field to given value.

### HasBaseCharge

`func (o *SchedulePickupUPSResponsePickupRate) HasBaseCharge() bool`

HasBaseCharge returns a boolean if a field has been set.

### GetSurcharges

`func (o *SchedulePickupUPSResponsePickupRate) GetSurcharges() []SchedulePickupUPSResponsePickupRateSurchargesInner`

GetSurcharges returns the Surcharges field if non-nil, zero value otherwise.

### GetSurchargesOk

`func (o *SchedulePickupUPSResponsePickupRate) GetSurchargesOk() (*[]SchedulePickupUPSResponsePickupRateSurchargesInner, bool)`

GetSurchargesOk returns a tuple with the Surcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharges

`func (o *SchedulePickupUPSResponsePickupRate) SetSurcharges(v []SchedulePickupUPSResponsePickupRateSurchargesInner)`

SetSurcharges sets Surcharges field to given value.

### HasSurcharges

`func (o *SchedulePickupUPSResponsePickupRate) HasSurcharges() bool`

HasSurcharges returns a boolean if a field has been set.

### GetTotalCarrierCharge

`func (o *SchedulePickupUPSResponsePickupRate) GetTotalCarrierCharge() float32`

GetTotalCarrierCharge returns the TotalCarrierCharge field if non-nil, zero value otherwise.

### GetTotalCarrierChargeOk

`func (o *SchedulePickupUPSResponsePickupRate) GetTotalCarrierChargeOk() (*float32, bool)`

GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCarrierCharge

`func (o *SchedulePickupUPSResponsePickupRate) SetTotalCarrierCharge(v float32)`

SetTotalCarrierCharge sets TotalCarrierCharge field to given value.

### HasTotalCarrierCharge

`func (o *SchedulePickupUPSResponsePickupRate) HasTotalCarrierCharge() bool`

HasTotalCarrierCharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


