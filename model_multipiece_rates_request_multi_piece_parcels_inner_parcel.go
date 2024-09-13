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

// checks if the MultipieceRatesRequestMultiPieceParcelsInnerParcel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipieceRatesRequestMultiPieceParcelsInnerParcel{}

// MultipieceRatesRequestMultiPieceParcelsInnerParcel description
type MultipieceRatesRequestMultiPieceParcelsInnerParcel struct {
	// description
	WeightUnit *string `json:"weightUnit,omitempty"`
	// description
	Weight *int32 `json:"weight,omitempty"`
	// description
	DimUnit *string `json:"dimUnit,omitempty"`
	// description
	Length *int32 `json:"length,omitempty"`
	// description
	Width *int32 `json:"width,omitempty"`
	// description
	Height *float32 `json:"height,omitempty"`
}

// NewMultipieceRatesRequestMultiPieceParcelsInnerParcel instantiates a new MultipieceRatesRequestMultiPieceParcelsInnerParcel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipieceRatesRequestMultiPieceParcelsInnerParcel() *MultipieceRatesRequestMultiPieceParcelsInnerParcel {
	this := MultipieceRatesRequestMultiPieceParcelsInnerParcel{}
	return &this
}

// NewMultipieceRatesRequestMultiPieceParcelsInnerParcelWithDefaults instantiates a new MultipieceRatesRequestMultiPieceParcelsInnerParcel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipieceRatesRequestMultiPieceParcelsInnerParcelWithDefaults() *MultipieceRatesRequestMultiPieceParcelsInnerParcel {
	this := MultipieceRatesRequestMultiPieceParcelsInnerParcel{}
	return &this
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) GetWeightUnit() string {
	if o == nil || IsNil(o.WeightUnit) {
		var ret string
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) GetWeightUnitOk() (*string, bool) {
	if o == nil || IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) HasWeightUnit() bool {
	if o != nil && !IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given string and assigns it to the WeightUnit field.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) SetWeightUnit(v string) {
	o.WeightUnit = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) SetWeight(v int32) {
	o.Weight = &v
}

// GetDimUnit returns the DimUnit field value if set, zero value otherwise.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) GetDimUnit() string {
	if o == nil || IsNil(o.DimUnit) {
		var ret string
		return ret
	}
	return *o.DimUnit
}

// GetDimUnitOk returns a tuple with the DimUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) GetDimUnitOk() (*string, bool) {
	if o == nil || IsNil(o.DimUnit) {
		return nil, false
	}
	return o.DimUnit, true
}

// HasDimUnit returns a boolean if a field has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) HasDimUnit() bool {
	if o != nil && !IsNil(o.DimUnit) {
		return true
	}

	return false
}

// SetDimUnit gets a reference to the given string and assigns it to the DimUnit field.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) SetDimUnit(v string) {
	o.DimUnit = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) GetLength() int32 {
	if o == nil || IsNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) GetLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) SetLength(v int32) {
	o.Length = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *MultipieceRatesRequestMultiPieceParcelsInnerParcel) SetHeight(v float32) {
	o.Height = &v
}

func (o MultipieceRatesRequestMultiPieceParcelsInnerParcel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipieceRatesRequestMultiPieceParcelsInnerParcel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WeightUnit) {
		toSerialize["weightUnit"] = o.WeightUnit
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.DimUnit) {
		toSerialize["dimUnit"] = o.DimUnit
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableMultipieceRatesRequestMultiPieceParcelsInnerParcel struct {
	value *MultipieceRatesRequestMultiPieceParcelsInnerParcel
	isSet bool
}

func (v NullableMultipieceRatesRequestMultiPieceParcelsInnerParcel) Get() *MultipieceRatesRequestMultiPieceParcelsInnerParcel {
	return v.value
}

func (v *NullableMultipieceRatesRequestMultiPieceParcelsInnerParcel) Set(val *MultipieceRatesRequestMultiPieceParcelsInnerParcel) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipieceRatesRequestMultiPieceParcelsInnerParcel) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipieceRatesRequestMultiPieceParcelsInnerParcel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipieceRatesRequestMultiPieceParcelsInnerParcel(val *MultipieceRatesRequestMultiPieceParcelsInnerParcel) *NullableMultipieceRatesRequestMultiPieceParcelsInnerParcel {
	return &NullableMultipieceRatesRequestMultiPieceParcelsInnerParcel{value: val, isSet: true}
}

func (v NullableMultipieceRatesRequestMultiPieceParcelsInnerParcel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipieceRatesRequestMultiPieceParcelsInnerParcel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


