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

// checks if the SingleRateResponseRatesInnerSpecialServicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleRateResponseRatesInnerSpecialServicesInner{}

// SingleRateResponseRatesInnerSpecialServicesInner struct for SingleRateResponseRatesInnerSpecialServicesInner
type SingleRateResponseRatesInnerSpecialServicesInner struct {
	// The amount of the specialSevice.
	Fee *float32 `json:"fee,omitempty"`
	// This is the unique identifier given to various specialservice, which is used while Rating.
	SpecialserviceId *string `json:"specialserviceId,omitempty"`
}

// NewSingleRateResponseRatesInnerSpecialServicesInner instantiates a new SingleRateResponseRatesInnerSpecialServicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleRateResponseRatesInnerSpecialServicesInner() *SingleRateResponseRatesInnerSpecialServicesInner {
	this := SingleRateResponseRatesInnerSpecialServicesInner{}
	return &this
}

// NewSingleRateResponseRatesInnerSpecialServicesInnerWithDefaults instantiates a new SingleRateResponseRatesInnerSpecialServicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleRateResponseRatesInnerSpecialServicesInnerWithDefaults() *SingleRateResponseRatesInnerSpecialServicesInner {
	this := SingleRateResponseRatesInnerSpecialServicesInner{}
	return &this
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *SingleRateResponseRatesInnerSpecialServicesInner) GetFee() float32 {
	if o == nil || IsNil(o.Fee) {
		var ret float32
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponseRatesInnerSpecialServicesInner) GetFeeOk() (*float32, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *SingleRateResponseRatesInnerSpecialServicesInner) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given float32 and assigns it to the Fee field.
func (o *SingleRateResponseRatesInnerSpecialServicesInner) SetFee(v float32) {
	o.Fee = &v
}

// GetSpecialserviceId returns the SpecialserviceId field value if set, zero value otherwise.
func (o *SingleRateResponseRatesInnerSpecialServicesInner) GetSpecialserviceId() string {
	if o == nil || IsNil(o.SpecialserviceId) {
		var ret string
		return ret
	}
	return *o.SpecialserviceId
}

// GetSpecialserviceIdOk returns a tuple with the SpecialserviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRateResponseRatesInnerSpecialServicesInner) GetSpecialserviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialserviceId) {
		return nil, false
	}
	return o.SpecialserviceId, true
}

// HasSpecialserviceId returns a boolean if a field has been set.
func (o *SingleRateResponseRatesInnerSpecialServicesInner) HasSpecialserviceId() bool {
	if o != nil && !IsNil(o.SpecialserviceId) {
		return true
	}

	return false
}

// SetSpecialserviceId gets a reference to the given string and assigns it to the SpecialserviceId field.
func (o *SingleRateResponseRatesInnerSpecialServicesInner) SetSpecialserviceId(v string) {
	o.SpecialserviceId = &v
}

func (o SingleRateResponseRatesInnerSpecialServicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleRateResponseRatesInnerSpecialServicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.SpecialserviceId) {
		toSerialize["specialserviceId"] = o.SpecialserviceId
	}
	return toSerialize, nil
}

type NullableSingleRateResponseRatesInnerSpecialServicesInner struct {
	value *SingleRateResponseRatesInnerSpecialServicesInner
	isSet bool
}

func (v NullableSingleRateResponseRatesInnerSpecialServicesInner) Get() *SingleRateResponseRatesInnerSpecialServicesInner {
	return v.value
}

func (v *NullableSingleRateResponseRatesInnerSpecialServicesInner) Set(val *SingleRateResponseRatesInnerSpecialServicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleRateResponseRatesInnerSpecialServicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleRateResponseRatesInnerSpecialServicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleRateResponseRatesInnerSpecialServicesInner(val *SingleRateResponseRatesInnerSpecialServicesInner) *NullableSingleRateResponseRatesInnerSpecialServicesInner {
	return &NullableSingleRateResponseRatesInnerSpecialServicesInner{value: val, isSet: true}
}

func (v NullableSingleRateResponseRatesInnerSpecialServicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleRateResponseRatesInnerSpecialServicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


