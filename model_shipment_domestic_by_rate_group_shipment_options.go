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

// checks if the ShipmentDomesticByRateGroupShipmentOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentDomesticByRateGroupShipmentOptions{}

// ShipmentDomesticByRateGroupShipmentOptions Shipment Options have an option of Manifest.<br /> With Manifest, the Mail Center agent can print the Manifest (End of day records of all created shipment) of selected carrier.
type ShipmentDomesticByRateGroupShipmentOptions struct {
	// This option asks if the shipment is to be added for Manifest, so that the shipment will reflect in the Manifest Form while compilation. <br /> The value can be 'true' or 'false'.
	AddToManifest *bool `json:"addToManifest,omitempty"`
}

// NewShipmentDomesticByRateGroupShipmentOptions instantiates a new ShipmentDomesticByRateGroupShipmentOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDomesticByRateGroupShipmentOptions() *ShipmentDomesticByRateGroupShipmentOptions {
	this := ShipmentDomesticByRateGroupShipmentOptions{}
	return &this
}

// NewShipmentDomesticByRateGroupShipmentOptionsWithDefaults instantiates a new ShipmentDomesticByRateGroupShipmentOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDomesticByRateGroupShipmentOptionsWithDefaults() *ShipmentDomesticByRateGroupShipmentOptions {
	this := ShipmentDomesticByRateGroupShipmentOptions{}
	return &this
}

// GetAddToManifest returns the AddToManifest field value if set, zero value otherwise.
func (o *ShipmentDomesticByRateGroupShipmentOptions) GetAddToManifest() bool {
	if o == nil || IsNil(o.AddToManifest) {
		var ret bool
		return ret
	}
	return *o.AddToManifest
}

// GetAddToManifestOk returns a tuple with the AddToManifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomesticByRateGroupShipmentOptions) GetAddToManifestOk() (*bool, bool) {
	if o == nil || IsNil(o.AddToManifest) {
		return nil, false
	}
	return o.AddToManifest, true
}

// HasAddToManifest returns a boolean if a field has been set.
func (o *ShipmentDomesticByRateGroupShipmentOptions) HasAddToManifest() bool {
	if o != nil && !IsNil(o.AddToManifest) {
		return true
	}

	return false
}

// SetAddToManifest gets a reference to the given bool and assigns it to the AddToManifest field.
func (o *ShipmentDomesticByRateGroupShipmentOptions) SetAddToManifest(v bool) {
	o.AddToManifest = &v
}

func (o ShipmentDomesticByRateGroupShipmentOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentDomesticByRateGroupShipmentOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddToManifest) {
		toSerialize["addToManifest"] = o.AddToManifest
	}
	return toSerialize, nil
}

type NullableShipmentDomesticByRateGroupShipmentOptions struct {
	value *ShipmentDomesticByRateGroupShipmentOptions
	isSet bool
}

func (v NullableShipmentDomesticByRateGroupShipmentOptions) Get() *ShipmentDomesticByRateGroupShipmentOptions {
	return v.value
}

func (v *NullableShipmentDomesticByRateGroupShipmentOptions) Set(val *ShipmentDomesticByRateGroupShipmentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDomesticByRateGroupShipmentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDomesticByRateGroupShipmentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDomesticByRateGroupShipmentOptions(val *ShipmentDomesticByRateGroupShipmentOptions) *NullableShipmentDomesticByRateGroupShipmentOptions {
	return &NullableShipmentDomesticByRateGroupShipmentOptions{value: val, isSet: true}
}

func (v NullableShipmentDomesticByRateGroupShipmentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDomesticByRateGroupShipmentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


