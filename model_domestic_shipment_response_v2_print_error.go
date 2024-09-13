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

// checks if the DomesticShipmentResponseV2PrintError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomesticShipmentResponseV2PrintError{}

// DomesticShipmentResponseV2PrintError struct for DomesticShipmentResponseV2PrintError
type DomesticShipmentResponseV2PrintError struct {
	// indicates error code of print
	Code *string `json:"code,omitempty"`
	// Error message if print failed 
	Message *string `json:"message,omitempty"`
}

// NewDomesticShipmentResponseV2PrintError instantiates a new DomesticShipmentResponseV2PrintError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomesticShipmentResponseV2PrintError() *DomesticShipmentResponseV2PrintError {
	this := DomesticShipmentResponseV2PrintError{}
	return &this
}

// NewDomesticShipmentResponseV2PrintErrorWithDefaults instantiates a new DomesticShipmentResponseV2PrintError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomesticShipmentResponseV2PrintErrorWithDefaults() *DomesticShipmentResponseV2PrintError {
	this := DomesticShipmentResponseV2PrintError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DomesticShipmentResponseV2PrintError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseV2PrintError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DomesticShipmentResponseV2PrintError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *DomesticShipmentResponseV2PrintError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DomesticShipmentResponseV2PrintError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomesticShipmentResponseV2PrintError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DomesticShipmentResponseV2PrintError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DomesticShipmentResponseV2PrintError) SetMessage(v string) {
	o.Message = &v
}

func (o DomesticShipmentResponseV2PrintError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomesticShipmentResponseV2PrintError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableDomesticShipmentResponseV2PrintError struct {
	value *DomesticShipmentResponseV2PrintError
	isSet bool
}

func (v NullableDomesticShipmentResponseV2PrintError) Get() *DomesticShipmentResponseV2PrintError {
	return v.value
}

func (v *NullableDomesticShipmentResponseV2PrintError) Set(val *DomesticShipmentResponseV2PrintError) {
	v.value = val
	v.isSet = true
}

func (v NullableDomesticShipmentResponseV2PrintError) IsSet() bool {
	return v.isSet
}

func (v *NullableDomesticShipmentResponseV2PrintError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomesticShipmentResponseV2PrintError(val *DomesticShipmentResponseV2PrintError) *NullableDomesticShipmentResponseV2PrintError {
	return &NullableDomesticShipmentResponseV2PrintError{value: val, isSet: true}
}

func (v NullableDomesticShipmentResponseV2PrintError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomesticShipmentResponseV2PrintError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


