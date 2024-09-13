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

// checks if the ShipmentInternationalCustomsCustomsItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentInternationalCustomsCustomsItemsInner{}

// ShipmentInternationalCustomsCustomsItemsInner struct for ShipmentInternationalCustomsCustomsItemsInner
type ShipmentInternationalCustomsCustomsItemsInner struct {
	// A detailed description of the commodity, up to 255 characters. The description will appear on the customs form. If the shipment has multiple types of items, create a separate customsItems object for each. Each description will appear on the form.
	Description string `json:"description"`
	// The destination country’s tariff-classification number (HS code) for the commodity. Most countries use the six-digit Harmonized System (HS) as the basis for their tariff classifications and then add digits for more detail. The maximum length for an HS code is 14 characters. The HS code will appear on the customs form. If the shipment has multiple types of items, create a separate customsItems object for each.
	HSTariffCode *string `json:"hSTariffCode,omitempty"`
	// The two-character ISO country code of the shipment’s origin country. Use ISO 3166-1 alpha-2 standard values.
	OriginCountryCode *string `json:"originCountryCode,omitempty"`
	// Enter the total number of items of this type of commodity.
	Quantity float32 `json:"quantity"`
	// The price of one item of this type of commodity.
	UnitPrice float32 `json:"unitPrice"`
	// The unit of measurement. This field is required by the unitWeight object.
	WeightUnit *string `json:"weightUnit,omitempty"`
	// The weight of the item.
	Weight *float32 `json:"weight,omitempty"`
}

type _ShipmentInternationalCustomsCustomsItemsInner ShipmentInternationalCustomsCustomsItemsInner

// NewShipmentInternationalCustomsCustomsItemsInner instantiates a new ShipmentInternationalCustomsCustomsItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentInternationalCustomsCustomsItemsInner(description string, quantity float32, unitPrice float32) *ShipmentInternationalCustomsCustomsItemsInner {
	this := ShipmentInternationalCustomsCustomsItemsInner{}
	this.Description = description
	this.Quantity = quantity
	this.UnitPrice = unitPrice
	return &this
}

// NewShipmentInternationalCustomsCustomsItemsInnerWithDefaults instantiates a new ShipmentInternationalCustomsCustomsItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentInternationalCustomsCustomsItemsInnerWithDefaults() *ShipmentInternationalCustomsCustomsItemsInner {
	this := ShipmentInternationalCustomsCustomsItemsInner{}
	return &this
}

// GetDescription returns the Description field value
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ShipmentInternationalCustomsCustomsItemsInner) SetDescription(v string) {
	o.Description = v
}

// GetHSTariffCode returns the HSTariffCode field value if set, zero value otherwise.
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetHSTariffCode() string {
	if o == nil || IsNil(o.HSTariffCode) {
		var ret string
		return ret
	}
	return *o.HSTariffCode
}

// GetHSTariffCodeOk returns a tuple with the HSTariffCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetHSTariffCodeOk() (*string, bool) {
	if o == nil || IsNil(o.HSTariffCode) {
		return nil, false
	}
	return o.HSTariffCode, true
}

// HasHSTariffCode returns a boolean if a field has been set.
func (o *ShipmentInternationalCustomsCustomsItemsInner) HasHSTariffCode() bool {
	if o != nil && !IsNil(o.HSTariffCode) {
		return true
	}

	return false
}

// SetHSTariffCode gets a reference to the given string and assigns it to the HSTariffCode field.
func (o *ShipmentInternationalCustomsCustomsItemsInner) SetHSTariffCode(v string) {
	o.HSTariffCode = &v
}

// GetOriginCountryCode returns the OriginCountryCode field value if set, zero value otherwise.
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetOriginCountryCode() string {
	if o == nil || IsNil(o.OriginCountryCode) {
		var ret string
		return ret
	}
	return *o.OriginCountryCode
}

// GetOriginCountryCodeOk returns a tuple with the OriginCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetOriginCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OriginCountryCode) {
		return nil, false
	}
	return o.OriginCountryCode, true
}

// HasOriginCountryCode returns a boolean if a field has been set.
func (o *ShipmentInternationalCustomsCustomsItemsInner) HasOriginCountryCode() bool {
	if o != nil && !IsNil(o.OriginCountryCode) {
		return true
	}

	return false
}

// SetOriginCountryCode gets a reference to the given string and assigns it to the OriginCountryCode field.
func (o *ShipmentInternationalCustomsCustomsItemsInner) SetOriginCountryCode(v string) {
	o.OriginCountryCode = &v
}

// GetQuantity returns the Quantity field value
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ShipmentInternationalCustomsCustomsItemsInner) SetQuantity(v float32) {
	o.Quantity = v
}

// GetUnitPrice returns the UnitPrice field value
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetUnitPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetUnitPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPrice, true
}

// SetUnitPrice sets field value
func (o *ShipmentInternationalCustomsCustomsItemsInner) SetUnitPrice(v float32) {
	o.UnitPrice = v
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetWeightUnit() string {
	if o == nil || IsNil(o.WeightUnit) {
		var ret string
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetWeightUnitOk() (*string, bool) {
	if o == nil || IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *ShipmentInternationalCustomsCustomsItemsInner) HasWeightUnit() bool {
	if o != nil && !IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given string and assigns it to the WeightUnit field.
func (o *ShipmentInternationalCustomsCustomsItemsInner) SetWeightUnit(v string) {
	o.WeightUnit = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetWeight() float32 {
	if o == nil || IsNil(o.Weight) {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentInternationalCustomsCustomsItemsInner) GetWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ShipmentInternationalCustomsCustomsItemsInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *ShipmentInternationalCustomsCustomsItemsInner) SetWeight(v float32) {
	o.Weight = &v
}

func (o ShipmentInternationalCustomsCustomsItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentInternationalCustomsCustomsItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	if !IsNil(o.HSTariffCode) {
		toSerialize["hSTariffCode"] = o.HSTariffCode
	}
	if !IsNil(o.OriginCountryCode) {
		toSerialize["originCountryCode"] = o.OriginCountryCode
	}
	toSerialize["quantity"] = o.Quantity
	toSerialize["unitPrice"] = o.UnitPrice
	if !IsNil(o.WeightUnit) {
		toSerialize["weightUnit"] = o.WeightUnit
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

func (o *ShipmentInternationalCustomsCustomsItemsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"quantity",
		"unitPrice",
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

	varShipmentInternationalCustomsCustomsItemsInner := _ShipmentInternationalCustomsCustomsItemsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipmentInternationalCustomsCustomsItemsInner)

	if err != nil {
		return err
	}

	*o = ShipmentInternationalCustomsCustomsItemsInner(varShipmentInternationalCustomsCustomsItemsInner)

	return err
}

type NullableShipmentInternationalCustomsCustomsItemsInner struct {
	value *ShipmentInternationalCustomsCustomsItemsInner
	isSet bool
}

func (v NullableShipmentInternationalCustomsCustomsItemsInner) Get() *ShipmentInternationalCustomsCustomsItemsInner {
	return v.value
}

func (v *NullableShipmentInternationalCustomsCustomsItemsInner) Set(val *ShipmentInternationalCustomsCustomsItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentInternationalCustomsCustomsItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentInternationalCustomsCustomsItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentInternationalCustomsCustomsItemsInner(val *ShipmentInternationalCustomsCustomsItemsInner) *NullableShipmentInternationalCustomsCustomsItemsInner {
	return &NullableShipmentInternationalCustomsCustomsItemsInner{value: val, isSet: true}
}

func (v NullableShipmentInternationalCustomsCustomsItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentInternationalCustomsCustomsItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


