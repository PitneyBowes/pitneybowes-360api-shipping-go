/*
Shipping APIs

### Introduction  The Shipping APIs include a variety of operations that allow users to manage and track their shipping requests.   Some of the key API operations available in the Shipping API includes: ### Shipment API  | Operation      | Description | | ----------- | ----------- |  | Get Carriers    | This operation fetches all onboarded carriers. Typically, user will use this service to get list of onboarded carriers and supported properties for those carriers.  |  | Get Countries | This operation fetches list of supported destination countries for a provided carrier and origin country.  |  | Get Services | This operation fetches a list of supported services for a carrier with respect to specific origin and destination country. |  | Get ParcelTypes| This operation fetches ParcelTypes based on carrier, origin and destination country. |  | Get Special Services| This operation fetches Special Services for a given carrier, service, origin and destination country. |  | Get Carrier Accounts| This operation retrieves onboarded Carriers with their Carrier Account Ids which uniquely identify multiple accounts of same carrier.  |  | Rate Shop and Get Single Rate| This API contains 2 operations, rate shop and single rate. Rate shop will fetch rates for all carrier services based on the given addresses (From and To), weight, and dimension for given parcelType. Single rate will get rate for specific service and special service (if requested) based on the given addresses (From and To), weight, and dimension, parcelType and serviceId with or without specialServices. Single rate will be used mainly to a rate a shipment before creating shipment.  |  | Create Shipment| This operation creates a new Shipment or Shipment Label. This is for both Domestic and International. | | Get All Shipments| This operation fetches all created Shipments. |  | Get Shipment by Id| Retrieves single shipment using Shipment Id. |  | Reprint Shipment| This operation reprints Shipment by the shipmentId. It retrieves an existing shipping label to reprint. The API sends the shipmentId returned by the original Created Shipment request. Use this only if the shipping label in the Create Shipment response was spoilt or lost. |  | Cancel Shipment| This operation cancels previously created shipment. |  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipping

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SchedulePickupCancelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulePickupCancelRequest{}

// SchedulePickupCancelRequest struct for SchedulePickupCancelRequest
type SchedulePickupCancelRequest struct {
	// It specifies the pickup Ids for which you would like to cancel the request. Only pickupIds of the same carrier should be provided in the array.
	PickupIds []string `json:"pickupIds"`
	// It is required to be provided for DHL Express pickup cancellation. Both `REQUESTOR_NAME` and `REASON_FOR_CANCEL` are required for DHL Express.
	Options []SchedulePickupCancelRequestOptionsInner `json:"options,omitempty"`
}

type _SchedulePickupCancelRequest SchedulePickupCancelRequest

// NewSchedulePickupCancelRequest instantiates a new SchedulePickupCancelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulePickupCancelRequest(pickupIds []string) *SchedulePickupCancelRequest {
	this := SchedulePickupCancelRequest{}
	this.PickupIds = pickupIds
	return &this
}

// NewSchedulePickupCancelRequestWithDefaults instantiates a new SchedulePickupCancelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulePickupCancelRequestWithDefaults() *SchedulePickupCancelRequest {
	this := SchedulePickupCancelRequest{}
	return &this
}

// GetPickupIds returns the PickupIds field value
func (o *SchedulePickupCancelRequest) GetPickupIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PickupIds
}

// GetPickupIdsOk returns a tuple with the PickupIds field value
// and a boolean to check if the value has been set.
func (o *SchedulePickupCancelRequest) GetPickupIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PickupIds, true
}

// SetPickupIds sets field value
func (o *SchedulePickupCancelRequest) SetPickupIds(v []string) {
	o.PickupIds = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SchedulePickupCancelRequest) GetOptions() []SchedulePickupCancelRequestOptionsInner {
	if o == nil || IsNil(o.Options) {
		var ret []SchedulePickupCancelRequestOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupCancelRequest) GetOptionsOk() ([]SchedulePickupCancelRequestOptionsInner, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SchedulePickupCancelRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []SchedulePickupCancelRequestOptionsInner and assigns it to the Options field.
func (o *SchedulePickupCancelRequest) SetOptions(v []SchedulePickupCancelRequestOptionsInner) {
	o.Options = v
}

func (o SchedulePickupCancelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulePickupCancelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pickupIds"] = o.PickupIds
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

func (o *SchedulePickupCancelRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pickupIds",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSchedulePickupCancelRequest := _SchedulePickupCancelRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchedulePickupCancelRequest)

	if err != nil {
		return err
	}

	*o = SchedulePickupCancelRequest(varSchedulePickupCancelRequest)

	return err
}

type NullableSchedulePickupCancelRequest struct {
	value *SchedulePickupCancelRequest
	isSet bool
}

func (v NullableSchedulePickupCancelRequest) Get() *SchedulePickupCancelRequest {
	return v.value
}

func (v *NullableSchedulePickupCancelRequest) Set(val *SchedulePickupCancelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulePickupCancelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulePickupCancelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulePickupCancelRequest(val *SchedulePickupCancelRequest) *NullableSchedulePickupCancelRequest {
	return &NullableSchedulePickupCancelRequest{value: val, isSet: true}
}

func (v NullableSchedulePickupCancelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulePickupCancelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


