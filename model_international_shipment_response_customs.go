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

// checks if the InternationalShipmentResponseCustoms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternationalShipmentResponseCustoms{}

// InternationalShipmentResponseCustoms struct for InternationalShipmentResponseCustoms
type InternationalShipmentResponseCustoms struct {
	CustomsInfo *InternationalShipmentResponseCustomsCustomsInfo `json:"customsInfo,omitempty"`
	CustomsItems []InternationalShipmentResponseCustomsCustomsItemsInner `json:"customsItems,omitempty"`
}

// NewInternationalShipmentResponseCustoms instantiates a new InternationalShipmentResponseCustoms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternationalShipmentResponseCustoms() *InternationalShipmentResponseCustoms {
	this := InternationalShipmentResponseCustoms{}
	return &this
}

// NewInternationalShipmentResponseCustomsWithDefaults instantiates a new InternationalShipmentResponseCustoms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternationalShipmentResponseCustomsWithDefaults() *InternationalShipmentResponseCustoms {
	this := InternationalShipmentResponseCustoms{}
	return &this
}

// GetCustomsInfo returns the CustomsInfo field value if set, zero value otherwise.
func (o *InternationalShipmentResponseCustoms) GetCustomsInfo() InternationalShipmentResponseCustomsCustomsInfo {
	if o == nil || IsNil(o.CustomsInfo) {
		var ret InternationalShipmentResponseCustomsCustomsInfo
		return ret
	}
	return *o.CustomsInfo
}

// GetCustomsInfoOk returns a tuple with the CustomsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalShipmentResponseCustoms) GetCustomsInfoOk() (*InternationalShipmentResponseCustomsCustomsInfo, bool) {
	if o == nil || IsNil(o.CustomsInfo) {
		return nil, false
	}
	return o.CustomsInfo, true
}

// HasCustomsInfo returns a boolean if a field has been set.
func (o *InternationalShipmentResponseCustoms) HasCustomsInfo() bool {
	if o != nil && !IsNil(o.CustomsInfo) {
		return true
	}

	return false
}

// SetCustomsInfo gets a reference to the given InternationalShipmentResponseCustomsCustomsInfo and assigns it to the CustomsInfo field.
func (o *InternationalShipmentResponseCustoms) SetCustomsInfo(v InternationalShipmentResponseCustomsCustomsInfo) {
	o.CustomsInfo = &v
}

// GetCustomsItems returns the CustomsItems field value if set, zero value otherwise.
func (o *InternationalShipmentResponseCustoms) GetCustomsItems() []InternationalShipmentResponseCustomsCustomsItemsInner {
	if o == nil || IsNil(o.CustomsItems) {
		var ret []InternationalShipmentResponseCustomsCustomsItemsInner
		return ret
	}
	return o.CustomsItems
}

// GetCustomsItemsOk returns a tuple with the CustomsItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalShipmentResponseCustoms) GetCustomsItemsOk() ([]InternationalShipmentResponseCustomsCustomsItemsInner, bool) {
	if o == nil || IsNil(o.CustomsItems) {
		return nil, false
	}
	return o.CustomsItems, true
}

// HasCustomsItems returns a boolean if a field has been set.
func (o *InternationalShipmentResponseCustoms) HasCustomsItems() bool {
	if o != nil && !IsNil(o.CustomsItems) {
		return true
	}

	return false
}

// SetCustomsItems gets a reference to the given []InternationalShipmentResponseCustomsCustomsItemsInner and assigns it to the CustomsItems field.
func (o *InternationalShipmentResponseCustoms) SetCustomsItems(v []InternationalShipmentResponseCustomsCustomsItemsInner) {
	o.CustomsItems = v
}

func (o InternationalShipmentResponseCustoms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternationalShipmentResponseCustoms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomsInfo) {
		toSerialize["customsInfo"] = o.CustomsInfo
	}
	if !IsNil(o.CustomsItems) {
		toSerialize["customsItems"] = o.CustomsItems
	}
	return toSerialize, nil
}

type NullableInternationalShipmentResponseCustoms struct {
	value *InternationalShipmentResponseCustoms
	isSet bool
}

func (v NullableInternationalShipmentResponseCustoms) Get() *InternationalShipmentResponseCustoms {
	return v.value
}

func (v *NullableInternationalShipmentResponseCustoms) Set(val *InternationalShipmentResponseCustoms) {
	v.value = val
	v.isSet = true
}

func (v NullableInternationalShipmentResponseCustoms) IsSet() bool {
	return v.isSet
}

func (v *NullableInternationalShipmentResponseCustoms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternationalShipmentResponseCustoms(val *InternationalShipmentResponseCustoms) *NullableInternationalShipmentResponseCustoms {
	return &NullableInternationalShipmentResponseCustoms{value: val, isSet: true}
}

func (v NullableInternationalShipmentResponseCustoms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternationalShipmentResponseCustoms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


