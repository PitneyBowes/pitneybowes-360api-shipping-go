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

// checks if the DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner{}

// DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner struct for DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner
type DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner struct {
	// Name or ID of the Input Parameters.
	Name *string `json:"name,omitempty"`
	// Value of the Input Parameters.
	Value *string `json:"value,omitempty"`
}

// NewDomesticShipmentResponseRateSpecialServicesInnerInputParametersInner instantiates a new DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomesticShipmentResponseRateSpecialServicesInnerInputParametersInner() *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner {
	this := DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner{}
	return &this
}

// NewDomesticShipmentResponseRateSpecialServicesInnerInputParametersInnerWithDefaults instantiates a new DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomesticShipmentResponseRateSpecialServicesInnerInputParametersInnerWithDefaults() *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner {
	this := DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) SetValue(v string) {
	o.Value = &v
}

func (o DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableDomesticShipmentResponseRateSpecialServicesInnerInputParametersInner struct {
	value *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner
	isSet bool
}

func (v NullableDomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) Get() *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner {
	return v.value
}

func (v *NullableDomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) Set(val *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomesticShipmentResponseRateSpecialServicesInnerInputParametersInner(val *DomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) *NullableDomesticShipmentResponseRateSpecialServicesInnerInputParametersInner {
	return &NullableDomesticShipmentResponseRateSpecialServicesInnerInputParametersInner{value: val, isSet: true}
}

func (v NullableDomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomesticShipmentResponseRateSpecialServicesInnerInputParametersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


