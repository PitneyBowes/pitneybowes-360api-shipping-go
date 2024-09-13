/*
Shipping APIs

### Introduction  The Shipping APIs include a variety of operations that allow users to manage and track their shipping requests.   Some of the key API operations available in the Shipping API includes: ### Shipment API  | Operation      | Description | | ----------- | ----------- |  | Get Carriers    | This operation fetches all onboarded carriers. Typically, user will use this service to get list of onboarded carriers and supported properties for those carriers.  |  | Get Countries | This operation fetches list of supported destination countries for a provided carrier and origin country.  |  | Get Services | This operation fetches a list of supported services for a carrier with respect to specific origin and destination country. |  | Get ParcelTypes| This operation fetches ParcelTypes based on carrier, origin and destination country. |  | Get Special Services| This operation fetches Special Services for a given carrier, service, origin and destination country. |  | Get Carrier Accounts| This operation retrieves onboarded Carriers with their Carrier Account Ids which uniquely identify multiple accounts of same carrier.  |  | Rate Shop and Get Single Rate| This API contains 2 operations, rate shop and single rate. Rate shop will fetch rates for all carrier services based on the given addresses (From and To), weight, and dimension for given parcelType. Single rate will get rate for specific service and special service (if requested) based on the given addresses (From and To), weight, and dimension, parcelType and serviceId with or without specialServices. Single rate will be used mainly to a rate a shipment before creating shipment.  |  | Create Shipment| This operation creates a new Shipment or Shipment Label. This is for both Domestic and International. | | Get All Shipments| This operation fetches all created Shipments. |  | Get Shipment by Id| Retrieves single shipment using Shipment Id. |  | Reprint Shipment| This operation reprints Shipment by the shipmentId. It retrieves an existing shipping label to reprint. The API sends the shipmentId returned by the original Created Shipment request. Use this only if the shipping label in the Create Shipment response was spoilt or lost. |  | Cancel Shipment| This operation cancels previously created shipment. |  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipping

import (
	"encoding/json"
)

// checks if the SchedulePickupDHLEXPResponsePickupOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulePickupDHLEXPResponsePickupOptions{}

// SchedulePickupDHLEXPResponsePickupOptions It is used to provide any pickup options while scheduling pickups
type SchedulePickupDHLEXPResponsePickupOptions struct {
	// It specifies the earliest date and time that your parcels will be available for pickup (UTC time)
	PickupStartDateTime *string `json:"pickupStartDateTime,omitempty"`
	// It specifies the latest date and time that your parcels will be available for pickup (UTC time)
	PickupEndDateTime *string `json:"pickupEndDateTime,omitempty"`
}

// NewSchedulePickupDHLEXPResponsePickupOptions instantiates a new SchedulePickupDHLEXPResponsePickupOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulePickupDHLEXPResponsePickupOptions() *SchedulePickupDHLEXPResponsePickupOptions {
	this := SchedulePickupDHLEXPResponsePickupOptions{}
	return &this
}

// NewSchedulePickupDHLEXPResponsePickupOptionsWithDefaults instantiates a new SchedulePickupDHLEXPResponsePickupOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulePickupDHLEXPResponsePickupOptionsWithDefaults() *SchedulePickupDHLEXPResponsePickupOptions {
	this := SchedulePickupDHLEXPResponsePickupOptions{}
	return &this
}

// GetPickupStartDateTime returns the PickupStartDateTime field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponsePickupOptions) GetPickupStartDateTime() string {
	if o == nil || IsNil(o.PickupStartDateTime) {
		var ret string
		return ret
	}
	return *o.PickupStartDateTime
}

// GetPickupStartDateTimeOk returns a tuple with the PickupStartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponsePickupOptions) GetPickupStartDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PickupStartDateTime) {
		return nil, false
	}
	return o.PickupStartDateTime, true
}

// HasPickupStartDateTime returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponsePickupOptions) HasPickupStartDateTime() bool {
	if o != nil && !IsNil(o.PickupStartDateTime) {
		return true
	}

	return false
}

// SetPickupStartDateTime gets a reference to the given string and assigns it to the PickupStartDateTime field.
func (o *SchedulePickupDHLEXPResponsePickupOptions) SetPickupStartDateTime(v string) {
	o.PickupStartDateTime = &v
}

// GetPickupEndDateTime returns the PickupEndDateTime field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponsePickupOptions) GetPickupEndDateTime() string {
	if o == nil || IsNil(o.PickupEndDateTime) {
		var ret string
		return ret
	}
	return *o.PickupEndDateTime
}

// GetPickupEndDateTimeOk returns a tuple with the PickupEndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponsePickupOptions) GetPickupEndDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PickupEndDateTime) {
		return nil, false
	}
	return o.PickupEndDateTime, true
}

// HasPickupEndDateTime returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponsePickupOptions) HasPickupEndDateTime() bool {
	if o != nil && !IsNil(o.PickupEndDateTime) {
		return true
	}

	return false
}

// SetPickupEndDateTime gets a reference to the given string and assigns it to the PickupEndDateTime field.
func (o *SchedulePickupDHLEXPResponsePickupOptions) SetPickupEndDateTime(v string) {
	o.PickupEndDateTime = &v
}

func (o SchedulePickupDHLEXPResponsePickupOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulePickupDHLEXPResponsePickupOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PickupStartDateTime) {
		toSerialize["pickupStartDateTime"] = o.PickupStartDateTime
	}
	if !IsNil(o.PickupEndDateTime) {
		toSerialize["pickupEndDateTime"] = o.PickupEndDateTime
	}
	return toSerialize, nil
}

type NullableSchedulePickupDHLEXPResponsePickupOptions struct {
	value *SchedulePickupDHLEXPResponsePickupOptions
	isSet bool
}

func (v NullableSchedulePickupDHLEXPResponsePickupOptions) Get() *SchedulePickupDHLEXPResponsePickupOptions {
	return v.value
}

func (v *NullableSchedulePickupDHLEXPResponsePickupOptions) Set(val *SchedulePickupDHLEXPResponsePickupOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulePickupDHLEXPResponsePickupOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulePickupDHLEXPResponsePickupOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulePickupDHLEXPResponsePickupOptions(val *SchedulePickupDHLEXPResponsePickupOptions) *NullableSchedulePickupDHLEXPResponsePickupOptions {
	return &NullableSchedulePickupDHLEXPResponsePickupOptions{value: val, isSet: true}
}

func (v NullableSchedulePickupDHLEXPResponsePickupOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulePickupDHLEXPResponsePickupOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


