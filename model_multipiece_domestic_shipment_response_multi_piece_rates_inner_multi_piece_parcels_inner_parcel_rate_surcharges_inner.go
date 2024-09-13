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

// checks if the MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner{}

// MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner description
type MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner struct {
	// description
	Fee *float32 `json:"fee,omitempty"`
	// description
	Name *string `json:"name,omitempty"`
}

// NewMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner instantiates a new MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner() *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner {
	this := MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner{}
	return &this
}

// NewMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInnerWithDefaults instantiates a new MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInnerWithDefaults() *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner {
	this := MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner{}
	return &this
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) GetFee() float32 {
	if o == nil || IsNil(o.Fee) {
		var ret float32
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) GetFeeOk() (*float32, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given float32 and assigns it to the Fee field.
func (o *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) SetFee(v float32) {
	o.Fee = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) SetName(v string) {
	o.Name = &v
}

func (o MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner struct {
	value *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner
	isSet bool
}

func (v NullableMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) Get() *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner {
	return v.value
}

func (v *NullableMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) Set(val *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner(val *MultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) *NullableMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner {
	return &NullableMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner{value: val, isSet: true}
}

func (v NullableMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipieceDomesticShipmentResponseMultiPieceRatesInnerMultiPieceParcelsInnerParcelRateSurchargesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


