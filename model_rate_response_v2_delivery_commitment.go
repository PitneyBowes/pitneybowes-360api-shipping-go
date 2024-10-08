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

// checks if the RateResponseV2DeliveryCommitment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateResponseV2DeliveryCommitment{}

// RateResponseV2DeliveryCommitment Check for estimated delivery date, guarantee (if any), and number of days for shipment to be delivered.
type RateResponseV2DeliveryCommitment struct {
	// Estimated Delivery Date.
	EstimatedDeliveryDateTime *string `json:"estimatedDeliveryDateTime,omitempty"`
	// Max days to deliver shipment.
	MaxEstimatedNumberOfDays *string `json:"maxEstimatedNumberOfDays,omitempty"`
	// Checks if there is any guarantee or committment for shipment delivery.
	Guarantee *string `json:"guarantee,omitempty"`
}

// NewRateResponseV2DeliveryCommitment instantiates a new RateResponseV2DeliveryCommitment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateResponseV2DeliveryCommitment() *RateResponseV2DeliveryCommitment {
	this := RateResponseV2DeliveryCommitment{}
	return &this
}

// NewRateResponseV2DeliveryCommitmentWithDefaults instantiates a new RateResponseV2DeliveryCommitment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateResponseV2DeliveryCommitmentWithDefaults() *RateResponseV2DeliveryCommitment {
	this := RateResponseV2DeliveryCommitment{}
	return &this
}

// GetEstimatedDeliveryDateTime returns the EstimatedDeliveryDateTime field value if set, zero value otherwise.
func (o *RateResponseV2DeliveryCommitment) GetEstimatedDeliveryDateTime() string {
	if o == nil || IsNil(o.EstimatedDeliveryDateTime) {
		var ret string
		return ret
	}
	return *o.EstimatedDeliveryDateTime
}

// GetEstimatedDeliveryDateTimeOk returns a tuple with the EstimatedDeliveryDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateResponseV2DeliveryCommitment) GetEstimatedDeliveryDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedDeliveryDateTime) {
		return nil, false
	}
	return o.EstimatedDeliveryDateTime, true
}

// HasEstimatedDeliveryDateTime returns a boolean if a field has been set.
func (o *RateResponseV2DeliveryCommitment) HasEstimatedDeliveryDateTime() bool {
	if o != nil && !IsNil(o.EstimatedDeliveryDateTime) {
		return true
	}

	return false
}

// SetEstimatedDeliveryDateTime gets a reference to the given string and assigns it to the EstimatedDeliveryDateTime field.
func (o *RateResponseV2DeliveryCommitment) SetEstimatedDeliveryDateTime(v string) {
	o.EstimatedDeliveryDateTime = &v
}

// GetMaxEstimatedNumberOfDays returns the MaxEstimatedNumberOfDays field value if set, zero value otherwise.
func (o *RateResponseV2DeliveryCommitment) GetMaxEstimatedNumberOfDays() string {
	if o == nil || IsNil(o.MaxEstimatedNumberOfDays) {
		var ret string
		return ret
	}
	return *o.MaxEstimatedNumberOfDays
}

// GetMaxEstimatedNumberOfDaysOk returns a tuple with the MaxEstimatedNumberOfDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateResponseV2DeliveryCommitment) GetMaxEstimatedNumberOfDaysOk() (*string, bool) {
	if o == nil || IsNil(o.MaxEstimatedNumberOfDays) {
		return nil, false
	}
	return o.MaxEstimatedNumberOfDays, true
}

// HasMaxEstimatedNumberOfDays returns a boolean if a field has been set.
func (o *RateResponseV2DeliveryCommitment) HasMaxEstimatedNumberOfDays() bool {
	if o != nil && !IsNil(o.MaxEstimatedNumberOfDays) {
		return true
	}

	return false
}

// SetMaxEstimatedNumberOfDays gets a reference to the given string and assigns it to the MaxEstimatedNumberOfDays field.
func (o *RateResponseV2DeliveryCommitment) SetMaxEstimatedNumberOfDays(v string) {
	o.MaxEstimatedNumberOfDays = &v
}

// GetGuarantee returns the Guarantee field value if set, zero value otherwise.
func (o *RateResponseV2DeliveryCommitment) GetGuarantee() string {
	if o == nil || IsNil(o.Guarantee) {
		var ret string
		return ret
	}
	return *o.Guarantee
}

// GetGuaranteeOk returns a tuple with the Guarantee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateResponseV2DeliveryCommitment) GetGuaranteeOk() (*string, bool) {
	if o == nil || IsNil(o.Guarantee) {
		return nil, false
	}
	return o.Guarantee, true
}

// HasGuarantee returns a boolean if a field has been set.
func (o *RateResponseV2DeliveryCommitment) HasGuarantee() bool {
	if o != nil && !IsNil(o.Guarantee) {
		return true
	}

	return false
}

// SetGuarantee gets a reference to the given string and assigns it to the Guarantee field.
func (o *RateResponseV2DeliveryCommitment) SetGuarantee(v string) {
	o.Guarantee = &v
}

func (o RateResponseV2DeliveryCommitment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateResponseV2DeliveryCommitment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EstimatedDeliveryDateTime) {
		toSerialize["estimatedDeliveryDateTime"] = o.EstimatedDeliveryDateTime
	}
	if !IsNil(o.MaxEstimatedNumberOfDays) {
		toSerialize["maxEstimatedNumberOfDays"] = o.MaxEstimatedNumberOfDays
	}
	if !IsNil(o.Guarantee) {
		toSerialize["guarantee"] = o.Guarantee
	}
	return toSerialize, nil
}

type NullableRateResponseV2DeliveryCommitment struct {
	value *RateResponseV2DeliveryCommitment
	isSet bool
}

func (v NullableRateResponseV2DeliveryCommitment) Get() *RateResponseV2DeliveryCommitment {
	return v.value
}

func (v *NullableRateResponseV2DeliveryCommitment) Set(val *RateResponseV2DeliveryCommitment) {
	v.value = val
	v.isSet = true
}

func (v NullableRateResponseV2DeliveryCommitment) IsSet() bool {
	return v.isSet
}

func (v *NullableRateResponseV2DeliveryCommitment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateResponseV2DeliveryCommitment(val *RateResponseV2DeliveryCommitment) *NullableRateResponseV2DeliveryCommitment {
	return &NullableRateResponseV2DeliveryCommitment{value: val, isSet: true}
}

func (v NullableRateResponseV2DeliveryCommitment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateResponseV2DeliveryCommitment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


