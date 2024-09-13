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

// checks if the ShipmentCancelV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentCancelV2{}

// ShipmentCancelV2 struct for ShipmentCancelV2
type ShipmentCancelV2 struct {
	// The shipmentId is a unique identifier for an individual Shipment.
	ShipmentId string `json:"shipmentId"`
	// The tracking number associated with one parcel in a shipment. The parcel tracking number can be used to track one specific parcel.
	ParcelTrackingNumber *string `json:"parcelTrackingNumber,omitempty"`
	References *ShipmentReprintV2References `json:"references,omitempty"`
}

type _ShipmentCancelV2 ShipmentCancelV2

// NewShipmentCancelV2 instantiates a new ShipmentCancelV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentCancelV2(shipmentId string) *ShipmentCancelV2 {
	this := ShipmentCancelV2{}
	this.ShipmentId = shipmentId
	return &this
}

// NewShipmentCancelV2WithDefaults instantiates a new ShipmentCancelV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentCancelV2WithDefaults() *ShipmentCancelV2 {
	this := ShipmentCancelV2{}
	return &this
}

// GetShipmentId returns the ShipmentId field value
func (o *ShipmentCancelV2) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *ShipmentCancelV2) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *ShipmentCancelV2) SetShipmentId(v string) {
	o.ShipmentId = v
}

// GetParcelTrackingNumber returns the ParcelTrackingNumber field value if set, zero value otherwise.
func (o *ShipmentCancelV2) GetParcelTrackingNumber() string {
	if o == nil || IsNil(o.ParcelTrackingNumber) {
		var ret string
		return ret
	}
	return *o.ParcelTrackingNumber
}

// GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentCancelV2) GetParcelTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelTrackingNumber) {
		return nil, false
	}
	return o.ParcelTrackingNumber, true
}

// HasParcelTrackingNumber returns a boolean if a field has been set.
func (o *ShipmentCancelV2) HasParcelTrackingNumber() bool {
	if o != nil && !IsNil(o.ParcelTrackingNumber) {
		return true
	}

	return false
}

// SetParcelTrackingNumber gets a reference to the given string and assigns it to the ParcelTrackingNumber field.
func (o *ShipmentCancelV2) SetParcelTrackingNumber(v string) {
	o.ParcelTrackingNumber = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *ShipmentCancelV2) GetReferences() ShipmentReprintV2References {
	if o == nil || IsNil(o.References) {
		var ret ShipmentReprintV2References
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentCancelV2) GetReferencesOk() (*ShipmentReprintV2References, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *ShipmentCancelV2) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given ShipmentReprintV2References and assigns it to the References field.
func (o *ShipmentCancelV2) SetReferences(v ShipmentReprintV2References) {
	o.References = &v
}

func (o ShipmentCancelV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentCancelV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipmentId"] = o.ShipmentId
	if !IsNil(o.ParcelTrackingNumber) {
		toSerialize["parcelTrackingNumber"] = o.ParcelTrackingNumber
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	return toSerialize, nil
}

func (o *ShipmentCancelV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"shipmentId",
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

	varShipmentCancelV2 := _ShipmentCancelV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipmentCancelV2)

	if err != nil {
		return err
	}

	*o = ShipmentCancelV2(varShipmentCancelV2)

	return err
}

type NullableShipmentCancelV2 struct {
	value *ShipmentCancelV2
	isSet bool
}

func (v NullableShipmentCancelV2) Get() *ShipmentCancelV2 {
	return v.value
}

func (v *NullableShipmentCancelV2) Set(val *ShipmentCancelV2) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentCancelV2) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentCancelV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentCancelV2(val *ShipmentCancelV2) *NullableShipmentCancelV2 {
	return &NullableShipmentCancelV2{value: val, isSet: true}
}

func (v NullableShipmentCancelV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentCancelV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


