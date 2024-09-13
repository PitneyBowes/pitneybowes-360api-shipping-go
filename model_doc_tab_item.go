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

// checks if the DocTabItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocTabItem{}

// DocTabItem This field will either be part of the request or response payload or will be marked as a custom field. We need to pass this field only if we need to print it on the label.
type DocTabItem struct {
	// This is a mandatory field. It will be displayed on the label
	DisplayName *string `json:"displayName,omitempty"`
	// If the field is part of a request or response, the value will be picked up from there. In the case of custom fields, the user-provided value will be printed.
	Value *string `json:"value,omitempty"`
	// Row Position of the Item. The min value is 1.
	Row *int32 `json:"row,omitempty"`
	// Column Position of the Item. The min value is 1.
	Column *int32 `json:"column,omitempty"`
}

// NewDocTabItem instantiates a new DocTabItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocTabItem() *DocTabItem {
	this := DocTabItem{}
	return &this
}

// NewDocTabItemWithDefaults instantiates a new DocTabItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocTabItemWithDefaults() *DocTabItem {
	this := DocTabItem{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *DocTabItem) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocTabItem) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DocTabItem) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *DocTabItem) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DocTabItem) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocTabItem) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DocTabItem) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DocTabItem) SetValue(v string) {
	o.Value = &v
}

// GetRow returns the Row field value if set, zero value otherwise.
func (o *DocTabItem) GetRow() int32 {
	if o == nil || IsNil(o.Row) {
		var ret int32
		return ret
	}
	return *o.Row
}

// GetRowOk returns a tuple with the Row field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocTabItem) GetRowOk() (*int32, bool) {
	if o == nil || IsNil(o.Row) {
		return nil, false
	}
	return o.Row, true
}

// HasRow returns a boolean if a field has been set.
func (o *DocTabItem) HasRow() bool {
	if o != nil && !IsNil(o.Row) {
		return true
	}

	return false
}

// SetRow gets a reference to the given int32 and assigns it to the Row field.
func (o *DocTabItem) SetRow(v int32) {
	o.Row = &v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *DocTabItem) GetColumn() int32 {
	if o == nil || IsNil(o.Column) {
		var ret int32
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocTabItem) GetColumnOk() (*int32, bool) {
	if o == nil || IsNil(o.Column) {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *DocTabItem) HasColumn() bool {
	if o != nil && !IsNil(o.Column) {
		return true
	}

	return false
}

// SetColumn gets a reference to the given int32 and assigns it to the Column field.
func (o *DocTabItem) SetColumn(v int32) {
	o.Column = &v
}

func (o DocTabItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocTabItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Row) {
		toSerialize["row"] = o.Row
	}
	if !IsNil(o.Column) {
		toSerialize["column"] = o.Column
	}
	return toSerialize, nil
}

type NullableDocTabItem struct {
	value *DocTabItem
	isSet bool
}

func (v NullableDocTabItem) Get() *DocTabItem {
	return v.value
}

func (v *NullableDocTabItem) Set(val *DocTabItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDocTabItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDocTabItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocTabItem(val *DocTabItem) *NullableDocTabItem {
	return &NullableDocTabItem{value: val, isSet: true}
}

func (v NullableDocTabItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocTabItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


