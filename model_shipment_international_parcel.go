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

// checks if the ShipmentInternationalParcel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentInternationalParcel{}

// ShipmentInternationalParcel struct for ShipmentInternationalParcel
type ShipmentInternationalParcel struct {
	Height float32 `json:"height"`
	Length float32 `json:"length"`
	Width float32 `json:"width"`
	DimUnit string `json:"dimUnit"`
	WeightUnit string `json:"weightUnit"`
	Weight float32 `json:"weight"`
}

type _ShipmentInternationalParcel ShipmentInternationalParcel

// NewShipmentInternationalParcel instantiates a new ShipmentInternationalParcel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentInternationalParcel(height float32, length float32, width float32, dimUnit string, weightUnit string, weight float32) *ShipmentInternationalParcel {
	this := ShipmentInternationalParcel{}
	this.Height = height
	this.Length = length
	this.Width = width
	this.DimUnit = dimUnit
	this.WeightUnit = weightUnit
	this.Weight = weight
	return &this
}

// NewShipmentInternationalParcelWithDefaults instantiates a new ShipmentInternationalParcel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentInternationalParcelWithDefaults() *ShipmentInternationalParcel {
	this := ShipmentInternationalParcel{}
	return &this
}

// GetHeight returns the Height field value
func (o *ShipmentInternationalParcel) GetHeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalParcel) GetHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *ShipmentInternationalParcel) SetHeight(v float32) {
	o.Height = v
}

// GetLength returns the Length field value
func (o *ShipmentInternationalParcel) GetLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalParcel) GetLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *ShipmentInternationalParcel) SetLength(v float32) {
	o.Length = v
}

// GetWidth returns the Width field value
func (o *ShipmentInternationalParcel) GetWidth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalParcel) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *ShipmentInternationalParcel) SetWidth(v float32) {
	o.Width = v
}

// GetDimUnit returns the DimUnit field value
func (o *ShipmentInternationalParcel) GetDimUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DimUnit
}

// GetDimUnitOk returns a tuple with the DimUnit field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalParcel) GetDimUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DimUnit, true
}

// SetDimUnit sets field value
func (o *ShipmentInternationalParcel) SetDimUnit(v string) {
	o.DimUnit = v
}

// GetWeightUnit returns the WeightUnit field value
func (o *ShipmentInternationalParcel) GetWeightUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalParcel) GetWeightUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WeightUnit, true
}

// SetWeightUnit sets field value
func (o *ShipmentInternationalParcel) SetWeightUnit(v string) {
	o.WeightUnit = v
}

// GetWeight returns the Weight field value
func (o *ShipmentInternationalParcel) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalParcel) GetWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *ShipmentInternationalParcel) SetWeight(v float32) {
	o.Weight = v
}

func (o ShipmentInternationalParcel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentInternationalParcel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["height"] = o.Height
	toSerialize["length"] = o.Length
	toSerialize["width"] = o.Width
	toSerialize["dimUnit"] = o.DimUnit
	toSerialize["weightUnit"] = o.WeightUnit
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

func (o *ShipmentInternationalParcel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"height",
		"length",
		"width",
		"dimUnit",
		"weightUnit",
		"weight",
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

	varShipmentInternationalParcel := _ShipmentInternationalParcel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipmentInternationalParcel)

	if err != nil {
		return err
	}

	*o = ShipmentInternationalParcel(varShipmentInternationalParcel)

	return err
}

type NullableShipmentInternationalParcel struct {
	value *ShipmentInternationalParcel
	isSet bool
}

func (v NullableShipmentInternationalParcel) Get() *ShipmentInternationalParcel {
	return v.value
}

func (v *NullableShipmentInternationalParcel) Set(val *ShipmentInternationalParcel) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentInternationalParcel) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentInternationalParcel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentInternationalParcel(val *ShipmentInternationalParcel) *NullableShipmentInternationalParcel {
	return &NullableShipmentInternationalParcel{value: val, isSet: true}
}

func (v NullableShipmentInternationalParcel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentInternationalParcel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


