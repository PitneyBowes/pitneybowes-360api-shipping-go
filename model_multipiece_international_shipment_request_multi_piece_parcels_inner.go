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

// checks if the MultipieceInternationalShipmentRequestMultiPieceParcelsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipieceInternationalShipmentRequestMultiPieceParcelsInner{}

// MultipieceInternationalShipmentRequestMultiPieceParcelsInner struct for MultipieceInternationalShipmentRequestMultiPieceParcelsInner
type MultipieceInternationalShipmentRequestMultiPieceParcelsInner struct {
	// Description
	ParcelType *string `json:"parcelType,omitempty"`
	// Description
	ParcelId *string `json:"parcelId,omitempty"`
	Parcel *MultipieceInternationalShipmentRequestMultiPieceParcelsInnerParcel `json:"parcel,omitempty"`
}

// NewMultipieceInternationalShipmentRequestMultiPieceParcelsInner instantiates a new MultipieceInternationalShipmentRequestMultiPieceParcelsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipieceInternationalShipmentRequestMultiPieceParcelsInner() *MultipieceInternationalShipmentRequestMultiPieceParcelsInner {
	this := MultipieceInternationalShipmentRequestMultiPieceParcelsInner{}
	return &this
}

// NewMultipieceInternationalShipmentRequestMultiPieceParcelsInnerWithDefaults instantiates a new MultipieceInternationalShipmentRequestMultiPieceParcelsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipieceInternationalShipmentRequestMultiPieceParcelsInnerWithDefaults() *MultipieceInternationalShipmentRequestMultiPieceParcelsInner {
	this := MultipieceInternationalShipmentRequestMultiPieceParcelsInner{}
	return &this
}

// GetParcelType returns the ParcelType field value if set, zero value otherwise.
func (o *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) GetParcelType() string {
	if o == nil || IsNil(o.ParcelType) {
		var ret string
		return ret
	}
	return *o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) GetParcelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelType) {
		return nil, false
	}
	return o.ParcelType, true
}

// HasParcelType returns a boolean if a field has been set.
func (o *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) HasParcelType() bool {
	if o != nil && !IsNil(o.ParcelType) {
		return true
	}

	return false
}

// SetParcelType gets a reference to the given string and assigns it to the ParcelType field.
func (o *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) SetParcelType(v string) {
	o.ParcelType = &v
}

// GetParcelId returns the ParcelId field value if set, zero value otherwise.
func (o *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) GetParcelId() string {
	if o == nil || IsNil(o.ParcelId) {
		var ret string
		return ret
	}
	return *o.ParcelId
}

// GetParcelIdOk returns a tuple with the ParcelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) GetParcelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelId) {
		return nil, false
	}
	return o.ParcelId, true
}

// HasParcelId returns a boolean if a field has been set.
func (o *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) HasParcelId() bool {
	if o != nil && !IsNil(o.ParcelId) {
		return true
	}

	return false
}

// SetParcelId gets a reference to the given string and assigns it to the ParcelId field.
func (o *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) SetParcelId(v string) {
	o.ParcelId = &v
}

// GetParcel returns the Parcel field value if set, zero value otherwise.
func (o *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) GetParcel() MultipieceInternationalShipmentRequestMultiPieceParcelsInnerParcel {
	if o == nil || IsNil(o.Parcel) {
		var ret MultipieceInternationalShipmentRequestMultiPieceParcelsInnerParcel
		return ret
	}
	return *o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) GetParcelOk() (*MultipieceInternationalShipmentRequestMultiPieceParcelsInnerParcel, bool) {
	if o == nil || IsNil(o.Parcel) {
		return nil, false
	}
	return o.Parcel, true
}

// HasParcel returns a boolean if a field has been set.
func (o *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) HasParcel() bool {
	if o != nil && !IsNil(o.Parcel) {
		return true
	}

	return false
}

// SetParcel gets a reference to the given MultipieceInternationalShipmentRequestMultiPieceParcelsInnerParcel and assigns it to the Parcel field.
func (o *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) SetParcel(v MultipieceInternationalShipmentRequestMultiPieceParcelsInnerParcel) {
	o.Parcel = &v
}

func (o MultipieceInternationalShipmentRequestMultiPieceParcelsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipieceInternationalShipmentRequestMultiPieceParcelsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParcelType) {
		toSerialize["parcelType"] = o.ParcelType
	}
	if !IsNil(o.ParcelId) {
		toSerialize["parcelId"] = o.ParcelId
	}
	if !IsNil(o.Parcel) {
		toSerialize["parcel"] = o.Parcel
	}
	return toSerialize, nil
}

type NullableMultipieceInternationalShipmentRequestMultiPieceParcelsInner struct {
	value *MultipieceInternationalShipmentRequestMultiPieceParcelsInner
	isSet bool
}

func (v NullableMultipieceInternationalShipmentRequestMultiPieceParcelsInner) Get() *MultipieceInternationalShipmentRequestMultiPieceParcelsInner {
	return v.value
}

func (v *NullableMultipieceInternationalShipmentRequestMultiPieceParcelsInner) Set(val *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipieceInternationalShipmentRequestMultiPieceParcelsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipieceInternationalShipmentRequestMultiPieceParcelsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipieceInternationalShipmentRequestMultiPieceParcelsInner(val *MultipieceInternationalShipmentRequestMultiPieceParcelsInner) *NullableMultipieceInternationalShipmentRequestMultiPieceParcelsInner {
	return &NullableMultipieceInternationalShipmentRequestMultiPieceParcelsInner{value: val, isSet: true}
}

func (v NullableMultipieceInternationalShipmentRequestMultiPieceParcelsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipieceInternationalShipmentRequestMultiPieceParcelsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


