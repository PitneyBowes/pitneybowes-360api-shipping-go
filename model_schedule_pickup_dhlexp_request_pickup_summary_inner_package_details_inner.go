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

// checks if the SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner{}

// SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner struct for SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner
type SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner struct {
	// It indicates the package width
	Width float32 `json:"width"`
	// It indicates the package height
	Height float32 `json:"height"`
	// It indicates the package length
	Length float32 `json:"length"`
	// It indicates the dimension unit, supported values are `IN` and `CM`
	DimUnit string `json:"dimUnit"`
	// It indicates the weight unit, supported values are `OZ` and `GM`
	WeightUnit string `json:"weightUnit"`
	// It indicates the package length
	Weight float32 `json:"weight"`
}

type _SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner

// NewSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner instantiates a new SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner(width float32, height float32, length float32, dimUnit string, weightUnit string, weight float32) *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner {
	this := SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner{}
	this.Width = width
	this.Height = height
	this.Length = length
	this.DimUnit = dimUnit
	this.WeightUnit = weightUnit
	this.Weight = weight
	return &this
}

// NewSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInnerWithDefaults instantiates a new SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInnerWithDefaults() *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner {
	this := SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner{}
	return &this
}

// GetWidth returns the Width field value
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) GetWidth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) SetWidth(v float32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) GetHeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) GetHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) SetHeight(v float32) {
	o.Height = v
}

// GetLength returns the Length field value
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) GetLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) GetLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) SetLength(v float32) {
	o.Length = v
}

// GetDimUnit returns the DimUnit field value
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) GetDimUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DimUnit
}

// GetDimUnitOk returns a tuple with the DimUnit field value
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) GetDimUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DimUnit, true
}

// SetDimUnit sets field value
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) SetDimUnit(v string) {
	o.DimUnit = v
}

// GetWeightUnit returns the WeightUnit field value
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) GetWeightUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) GetWeightUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WeightUnit, true
}

// SetWeightUnit sets field value
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) SetWeightUnit(v string) {
	o.WeightUnit = v
}

// GetWeight returns the Weight field value
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) GetWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) SetWeight(v float32) {
	o.Weight = v
}

func (o SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	toSerialize["length"] = o.Length
	toSerialize["dimUnit"] = o.DimUnit
	toSerialize["weightUnit"] = o.WeightUnit
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

func (o *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"width",
		"height",
		"length",
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

	varSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner := _SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner)

	if err != nil {
		return err
	}

	*o = SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner(varSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner)

	return err
}

type NullableSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner struct {
	value *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner
	isSet bool
}

func (v NullableSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) Get() *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner {
	return v.value
}

func (v *NullableSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) Set(val *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner(val *SchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) *NullableSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner {
	return &NullableSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner{value: val, isSet: true}
}

func (v NullableSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulePickupDHLEXPRequestPickupSummaryInnerPackageDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


