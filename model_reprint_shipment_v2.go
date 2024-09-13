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

// checks if the ReprintShipmentV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReprintShipmentV2{}

// ReprintShipmentV2 struct for ReprintShipmentV2
type ReprintShipmentV2 struct {
	// The shipmentId, a unique identifier for an individual Shipment, which is used for Reprint or Cancel.
	ShipmentId *string `json:"shipmentId,omitempty"`
	// The Tracking number given to the Parcel for tracking purpose.
	ParcelTrackingNumber *string `json:"parcelTrackingNumber,omitempty"`
	// It defines the layout of the shipping label.
	LabelLayout []ReprintShipmentV2LabelLayoutInner `json:"labelLayout,omitempty"`
	Parcel *ParcelV2 `json:"parcel,omitempty"`
	Rate *RateResponseV2 `json:"rate,omitempty"`
	References *CancelShipmentV2References `json:"references,omitempty"`
	// Status of the Printed Label.
	PrintStatus *string `json:"printStatus,omitempty"`
}

// NewReprintShipmentV2 instantiates a new ReprintShipmentV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReprintShipmentV2() *ReprintShipmentV2 {
	this := ReprintShipmentV2{}
	return &this
}

// NewReprintShipmentV2WithDefaults instantiates a new ReprintShipmentV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReprintShipmentV2WithDefaults() *ReprintShipmentV2 {
	this := ReprintShipmentV2{}
	return &this
}

// GetShipmentId returns the ShipmentId field value if set, zero value otherwise.
func (o *ReprintShipmentV2) GetShipmentId() string {
	if o == nil || IsNil(o.ShipmentId) {
		var ret string
		return ret
	}
	return *o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipmentV2) GetShipmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentId) {
		return nil, false
	}
	return o.ShipmentId, true
}

// HasShipmentId returns a boolean if a field has been set.
func (o *ReprintShipmentV2) HasShipmentId() bool {
	if o != nil && !IsNil(o.ShipmentId) {
		return true
	}

	return false
}

// SetShipmentId gets a reference to the given string and assigns it to the ShipmentId field.
func (o *ReprintShipmentV2) SetShipmentId(v string) {
	o.ShipmentId = &v
}

// GetParcelTrackingNumber returns the ParcelTrackingNumber field value if set, zero value otherwise.
func (o *ReprintShipmentV2) GetParcelTrackingNumber() string {
	if o == nil || IsNil(o.ParcelTrackingNumber) {
		var ret string
		return ret
	}
	return *o.ParcelTrackingNumber
}

// GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipmentV2) GetParcelTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelTrackingNumber) {
		return nil, false
	}
	return o.ParcelTrackingNumber, true
}

// HasParcelTrackingNumber returns a boolean if a field has been set.
func (o *ReprintShipmentV2) HasParcelTrackingNumber() bool {
	if o != nil && !IsNil(o.ParcelTrackingNumber) {
		return true
	}

	return false
}

// SetParcelTrackingNumber gets a reference to the given string and assigns it to the ParcelTrackingNumber field.
func (o *ReprintShipmentV2) SetParcelTrackingNumber(v string) {
	o.ParcelTrackingNumber = &v
}

// GetLabelLayout returns the LabelLayout field value if set, zero value otherwise.
func (o *ReprintShipmentV2) GetLabelLayout() []ReprintShipmentV2LabelLayoutInner {
	if o == nil || IsNil(o.LabelLayout) {
		var ret []ReprintShipmentV2LabelLayoutInner
		return ret
	}
	return o.LabelLayout
}

// GetLabelLayoutOk returns a tuple with the LabelLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipmentV2) GetLabelLayoutOk() ([]ReprintShipmentV2LabelLayoutInner, bool) {
	if o == nil || IsNil(o.LabelLayout) {
		return nil, false
	}
	return o.LabelLayout, true
}

// HasLabelLayout returns a boolean if a field has been set.
func (o *ReprintShipmentV2) HasLabelLayout() bool {
	if o != nil && !IsNil(o.LabelLayout) {
		return true
	}

	return false
}

// SetLabelLayout gets a reference to the given []ReprintShipmentV2LabelLayoutInner and assigns it to the LabelLayout field.
func (o *ReprintShipmentV2) SetLabelLayout(v []ReprintShipmentV2LabelLayoutInner) {
	o.LabelLayout = v
}

// GetParcel returns the Parcel field value if set, zero value otherwise.
func (o *ReprintShipmentV2) GetParcel() ParcelV2 {
	if o == nil || IsNil(o.Parcel) {
		var ret ParcelV2
		return ret
	}
	return *o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipmentV2) GetParcelOk() (*ParcelV2, bool) {
	if o == nil || IsNil(o.Parcel) {
		return nil, false
	}
	return o.Parcel, true
}

// HasParcel returns a boolean if a field has been set.
func (o *ReprintShipmentV2) HasParcel() bool {
	if o != nil && !IsNil(o.Parcel) {
		return true
	}

	return false
}

// SetParcel gets a reference to the given ParcelV2 and assigns it to the Parcel field.
func (o *ReprintShipmentV2) SetParcel(v ParcelV2) {
	o.Parcel = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *ReprintShipmentV2) GetRate() RateResponseV2 {
	if o == nil || IsNil(o.Rate) {
		var ret RateResponseV2
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipmentV2) GetRateOk() (*RateResponseV2, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *ReprintShipmentV2) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given RateResponseV2 and assigns it to the Rate field.
func (o *ReprintShipmentV2) SetRate(v RateResponseV2) {
	o.Rate = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *ReprintShipmentV2) GetReferences() CancelShipmentV2References {
	if o == nil || IsNil(o.References) {
		var ret CancelShipmentV2References
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipmentV2) GetReferencesOk() (*CancelShipmentV2References, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *ReprintShipmentV2) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given CancelShipmentV2References and assigns it to the References field.
func (o *ReprintShipmentV2) SetReferences(v CancelShipmentV2References) {
	o.References = &v
}

// GetPrintStatus returns the PrintStatus field value if set, zero value otherwise.
func (o *ReprintShipmentV2) GetPrintStatus() string {
	if o == nil || IsNil(o.PrintStatus) {
		var ret string
		return ret
	}
	return *o.PrintStatus
}

// GetPrintStatusOk returns a tuple with the PrintStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipmentV2) GetPrintStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PrintStatus) {
		return nil, false
	}
	return o.PrintStatus, true
}

// HasPrintStatus returns a boolean if a field has been set.
func (o *ReprintShipmentV2) HasPrintStatus() bool {
	if o != nil && !IsNil(o.PrintStatus) {
		return true
	}

	return false
}

// SetPrintStatus gets a reference to the given string and assigns it to the PrintStatus field.
func (o *ReprintShipmentV2) SetPrintStatus(v string) {
	o.PrintStatus = &v
}

func (o ReprintShipmentV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReprintShipmentV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShipmentId) {
		toSerialize["shipmentId"] = o.ShipmentId
	}
	if !IsNil(o.ParcelTrackingNumber) {
		toSerialize["parcelTrackingNumber"] = o.ParcelTrackingNumber
	}
	if !IsNil(o.LabelLayout) {
		toSerialize["labelLayout"] = o.LabelLayout
	}
	if !IsNil(o.Parcel) {
		toSerialize["parcel"] = o.Parcel
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.PrintStatus) {
		toSerialize["printStatus"] = o.PrintStatus
	}
	return toSerialize, nil
}

type NullableReprintShipmentV2 struct {
	value *ReprintShipmentV2
	isSet bool
}

func (v NullableReprintShipmentV2) Get() *ReprintShipmentV2 {
	return v.value
}

func (v *NullableReprintShipmentV2) Set(val *ReprintShipmentV2) {
	v.value = val
	v.isSet = true
}

func (v NullableReprintShipmentV2) IsSet() bool {
	return v.isSet
}

func (v *NullableReprintShipmentV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReprintShipmentV2(val *ReprintShipmentV2) *NullableReprintShipmentV2 {
	return &NullableReprintShipmentV2{value: val, isSet: true}
}

func (v NullableReprintShipmentV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReprintShipmentV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


