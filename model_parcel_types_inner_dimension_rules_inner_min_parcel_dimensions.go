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

// checks if the ParcelTypesInnerDimensionRulesInnerMinParcelDimensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParcelTypesInnerDimensionRulesInnerMinParcelDimensions{}

// ParcelTypesInnerDimensionRulesInnerMinParcelDimensions struct for ParcelTypesInnerDimensionRulesInnerMinParcelDimensions
type ParcelTypesInnerDimensionRulesInnerMinParcelDimensions struct {
	// Height is a part of Dimension objet where it helps determine a parcel’s girth.
	Height *float32 `json:"height,omitempty"`
	// Length is a part of Dimension objet having highest numeric value out of three required parameters (length, width, and height) of Dimension. It helps determine a parcel’s girth.
	Length *float32 `json:"length,omitempty"`
	// DimUnit is a standard for measuring the physical quantities of specified dimension parameters.
	UnitOfMeasurement *string `json:"unitOfMeasurement,omitempty"`
	// Width is a part of Dimension object having lowest numeric value out of three required parameters of dimension (length, width, and height). This helps determine a parcel’s girth.
	Width *float32 `json:"width,omitempty"`
}

// NewParcelTypesInnerDimensionRulesInnerMinParcelDimensions instantiates a new ParcelTypesInnerDimensionRulesInnerMinParcelDimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelTypesInnerDimensionRulesInnerMinParcelDimensions() *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions {
	this := ParcelTypesInnerDimensionRulesInnerMinParcelDimensions{}
	return &this
}

// NewParcelTypesInnerDimensionRulesInnerMinParcelDimensionsWithDefaults instantiates a new ParcelTypesInnerDimensionRulesInnerMinParcelDimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelTypesInnerDimensionRulesInnerMinParcelDimensionsWithDefaults() *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions {
	this := ParcelTypesInnerDimensionRulesInnerMinParcelDimensions{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) SetHeight(v float32) {
	o.Height = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) GetLength() float32 {
	if o == nil || IsNil(o.Length) {
		var ret float32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) GetLengthOk() (*float32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given float32 and assigns it to the Length field.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) SetLength(v float32) {
	o.Length = &v
}

// GetUnitOfMeasurement returns the UnitOfMeasurement field value if set, zero value otherwise.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) GetUnitOfMeasurement() string {
	if o == nil || IsNil(o.UnitOfMeasurement) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasurement
}

// GetUnitOfMeasurementOk returns a tuple with the UnitOfMeasurement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) GetUnitOfMeasurementOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasurement) {
		return nil, false
	}
	return o.UnitOfMeasurement, true
}

// HasUnitOfMeasurement returns a boolean if a field has been set.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) HasUnitOfMeasurement() bool {
	if o != nil && !IsNil(o.UnitOfMeasurement) {
		return true
	}

	return false
}

// SetUnitOfMeasurement gets a reference to the given string and assigns it to the UnitOfMeasurement field.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) SetUnitOfMeasurement(v string) {
	o.UnitOfMeasurement = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) GetWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) SetWidth(v float32) {
	o.Width = &v
}

func (o ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.UnitOfMeasurement) {
		toSerialize["unitOfMeasurement"] = o.UnitOfMeasurement
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullableParcelTypesInnerDimensionRulesInnerMinParcelDimensions struct {
	value *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions
	isSet bool
}

func (v NullableParcelTypesInnerDimensionRulesInnerMinParcelDimensions) Get() *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions {
	return v.value
}

func (v *NullableParcelTypesInnerDimensionRulesInnerMinParcelDimensions) Set(val *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelTypesInnerDimensionRulesInnerMinParcelDimensions) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelTypesInnerDimensionRulesInnerMinParcelDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelTypesInnerDimensionRulesInnerMinParcelDimensions(val *ParcelTypesInnerDimensionRulesInnerMinParcelDimensions) *NullableParcelTypesInnerDimensionRulesInnerMinParcelDimensions {
	return &NullableParcelTypesInnerDimensionRulesInnerMinParcelDimensions{value: val, isSet: true}
}

func (v NullableParcelTypesInnerDimensionRulesInnerMinParcelDimensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelTypesInnerDimensionRulesInnerMinParcelDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


