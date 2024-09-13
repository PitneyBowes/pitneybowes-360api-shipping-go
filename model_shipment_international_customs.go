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

// checks if the ShipmentInternationalCustoms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentInternationalCustoms{}

// ShipmentInternationalCustoms struct for ShipmentInternationalCustoms
type ShipmentInternationalCustoms struct {
	CustomsItems []ShipmentInternationalCustomsCustomsItemsInner `json:"customsItems,omitempty"`
	CustomsInfo ShipmentInternationalCustomsCustomsInfo `json:"customsInfo"`
}

type _ShipmentInternationalCustoms ShipmentInternationalCustoms

// NewShipmentInternationalCustoms instantiates a new ShipmentInternationalCustoms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentInternationalCustoms(customsInfo ShipmentInternationalCustomsCustomsInfo) *ShipmentInternationalCustoms {
	this := ShipmentInternationalCustoms{}
	this.CustomsInfo = customsInfo
	return &this
}

// NewShipmentInternationalCustomsWithDefaults instantiates a new ShipmentInternationalCustoms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentInternationalCustomsWithDefaults() *ShipmentInternationalCustoms {
	this := ShipmentInternationalCustoms{}
	return &this
}

// GetCustomsItems returns the CustomsItems field value if set, zero value otherwise.
func (o *ShipmentInternationalCustoms) GetCustomsItems() []ShipmentInternationalCustomsCustomsItemsInner {
	if o == nil || IsNil(o.CustomsItems) {
		var ret []ShipmentInternationalCustomsCustomsItemsInner
		return ret
	}
	return o.CustomsItems
}

// GetCustomsItemsOk returns a tuple with the CustomsItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustoms) GetCustomsItemsOk() ([]ShipmentInternationalCustomsCustomsItemsInner, bool) {
	if o == nil || IsNil(o.CustomsItems) {
		return nil, false
	}
	return o.CustomsItems, true
}

// HasCustomsItems returns a boolean if a field has been set.
func (o *ShipmentInternationalCustoms) HasCustomsItems() bool {
	if o != nil && !IsNil(o.CustomsItems) {
		return true
	}

	return false
}

// SetCustomsItems gets a reference to the given []ShipmentInternationalCustomsCustomsItemsInner and assigns it to the CustomsItems field.
func (o *ShipmentInternationalCustoms) SetCustomsItems(v []ShipmentInternationalCustomsCustomsItemsInner) {
	o.CustomsItems = v
}

// GetCustomsInfo returns the CustomsInfo field value
func (o *ShipmentInternationalCustoms) GetCustomsInfo() ShipmentInternationalCustomsCustomsInfo {
	if o == nil {
		var ret ShipmentInternationalCustomsCustomsInfo
		return ret
	}

	return o.CustomsInfo
}

// GetCustomsInfoOk returns a tuple with the CustomsInfo field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustoms) GetCustomsInfoOk() (*ShipmentInternationalCustomsCustomsInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomsInfo, true
}

// SetCustomsInfo sets field value
func (o *ShipmentInternationalCustoms) SetCustomsInfo(v ShipmentInternationalCustomsCustomsInfo) {
	o.CustomsInfo = v
}

func (o ShipmentInternationalCustoms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentInternationalCustoms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomsItems) {
		toSerialize["customsItems"] = o.CustomsItems
	}
	toSerialize["customsInfo"] = o.CustomsInfo
	return toSerialize, nil
}

func (o *ShipmentInternationalCustoms) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customsInfo",
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

	varShipmentInternationalCustoms := _ShipmentInternationalCustoms{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipmentInternationalCustoms)

	if err != nil {
		return err
	}

	*o = ShipmentInternationalCustoms(varShipmentInternationalCustoms)

	return err
}

type NullableShipmentInternationalCustoms struct {
	value *ShipmentInternationalCustoms
	isSet bool
}

func (v NullableShipmentInternationalCustoms) Get() *ShipmentInternationalCustoms {
	return v.value
}

func (v *NullableShipmentInternationalCustoms) Set(val *ShipmentInternationalCustoms) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentInternationalCustoms) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentInternationalCustoms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentInternationalCustoms(val *ShipmentInternationalCustoms) *NullableShipmentInternationalCustoms {
	return &NullableShipmentInternationalCustoms{value: val, isSet: true}
}

func (v NullableShipmentInternationalCustoms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentInternationalCustoms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


