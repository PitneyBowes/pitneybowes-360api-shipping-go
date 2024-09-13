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

// checks if the ReprintShipment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReprintShipment{}

// ReprintShipment struct for ReprintShipment
type ReprintShipment struct {
	// This is a GUID (globally unique identifier) that's automatically generated for every request that the webserver receives.
	CorrelationId *string `json:"correlationId,omitempty"`
	// This defines the label size of the Shipment, e.g., Shipping Label having Doc Size (8' X 11').
	Size *string `json:"size,omitempty"`
	// This defines the type of the Shipment, e.g., Shipping Label.
	Type *string `json:"type,omitempty"`
	// This defines the type of the shipment which is printed. For example Shipping label prints in PDF form.
	Format *string `json:"format,omitempty"`
	FromAddress *ReprintShipmentFromAddress `json:"fromAddress,omitempty"`
	Parcel *ReprintShipmentParcel `json:"parcel,omitempty"`
	// The Tracking number given to the Parcel for tracking purpose.
	ParcelTrackingNumber *string `json:"parcelTrackingNumber,omitempty"`
	Rate *ReprintShipmentRate `json:"rate,omitempty"`
	// A unique identifier associated with Shipment ID.
	ShipmentId *string `json:"shipmentId,omitempty"`
	ShipmentOptions *ShipmentOptions `json:"shipmentOptions,omitempty"`
	ToAddress *ReprintShipmentToAddress `json:"toAddress,omitempty"`
}

// NewReprintShipment instantiates a new ReprintShipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReprintShipment() *ReprintShipment {
	this := ReprintShipment{}
	return &this
}

// NewReprintShipmentWithDefaults instantiates a new ReprintShipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReprintShipmentWithDefaults() *ReprintShipment {
	this := ReprintShipment{}
	return &this
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *ReprintShipment) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipment) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *ReprintShipment) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *ReprintShipment) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ReprintShipment) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipment) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ReprintShipment) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *ReprintShipment) SetSize(v string) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReprintShipment) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipment) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReprintShipment) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReprintShipment) SetType(v string) {
	o.Type = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ReprintShipment) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipment) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ReprintShipment) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ReprintShipment) SetFormat(v string) {
	o.Format = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *ReprintShipment) GetFromAddress() ReprintShipmentFromAddress {
	if o == nil || IsNil(o.FromAddress) {
		var ret ReprintShipmentFromAddress
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipment) GetFromAddressOk() (*ReprintShipmentFromAddress, bool) {
	if o == nil || IsNil(o.FromAddress) {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *ReprintShipment) HasFromAddress() bool {
	if o != nil && !IsNil(o.FromAddress) {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given ReprintShipmentFromAddress and assigns it to the FromAddress field.
func (o *ReprintShipment) SetFromAddress(v ReprintShipmentFromAddress) {
	o.FromAddress = &v
}

// GetParcel returns the Parcel field value if set, zero value otherwise.
func (o *ReprintShipment) GetParcel() ReprintShipmentParcel {
	if o == nil || IsNil(o.Parcel) {
		var ret ReprintShipmentParcel
		return ret
	}
	return *o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipment) GetParcelOk() (*ReprintShipmentParcel, bool) {
	if o == nil || IsNil(o.Parcel) {
		return nil, false
	}
	return o.Parcel, true
}

// HasParcel returns a boolean if a field has been set.
func (o *ReprintShipment) HasParcel() bool {
	if o != nil && !IsNil(o.Parcel) {
		return true
	}

	return false
}

// SetParcel gets a reference to the given ReprintShipmentParcel and assigns it to the Parcel field.
func (o *ReprintShipment) SetParcel(v ReprintShipmentParcel) {
	o.Parcel = &v
}

// GetParcelTrackingNumber returns the ParcelTrackingNumber field value if set, zero value otherwise.
func (o *ReprintShipment) GetParcelTrackingNumber() string {
	if o == nil || IsNil(o.ParcelTrackingNumber) {
		var ret string
		return ret
	}
	return *o.ParcelTrackingNumber
}

// GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipment) GetParcelTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelTrackingNumber) {
		return nil, false
	}
	return o.ParcelTrackingNumber, true
}

// HasParcelTrackingNumber returns a boolean if a field has been set.
func (o *ReprintShipment) HasParcelTrackingNumber() bool {
	if o != nil && !IsNil(o.ParcelTrackingNumber) {
		return true
	}

	return false
}

// SetParcelTrackingNumber gets a reference to the given string and assigns it to the ParcelTrackingNumber field.
func (o *ReprintShipment) SetParcelTrackingNumber(v string) {
	o.ParcelTrackingNumber = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *ReprintShipment) GetRate() ReprintShipmentRate {
	if o == nil || IsNil(o.Rate) {
		var ret ReprintShipmentRate
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipment) GetRateOk() (*ReprintShipmentRate, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *ReprintShipment) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given ReprintShipmentRate and assigns it to the Rate field.
func (o *ReprintShipment) SetRate(v ReprintShipmentRate) {
	o.Rate = &v
}

// GetShipmentId returns the ShipmentId field value if set, zero value otherwise.
func (o *ReprintShipment) GetShipmentId() string {
	if o == nil || IsNil(o.ShipmentId) {
		var ret string
		return ret
	}
	return *o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipment) GetShipmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentId) {
		return nil, false
	}
	return o.ShipmentId, true
}

// HasShipmentId returns a boolean if a field has been set.
func (o *ReprintShipment) HasShipmentId() bool {
	if o != nil && !IsNil(o.ShipmentId) {
		return true
	}

	return false
}

// SetShipmentId gets a reference to the given string and assigns it to the ShipmentId field.
func (o *ReprintShipment) SetShipmentId(v string) {
	o.ShipmentId = &v
}

// GetShipmentOptions returns the ShipmentOptions field value if set, zero value otherwise.
func (o *ReprintShipment) GetShipmentOptions() ShipmentOptions {
	if o == nil || IsNil(o.ShipmentOptions) {
		var ret ShipmentOptions
		return ret
	}
	return *o.ShipmentOptions
}

// GetShipmentOptionsOk returns a tuple with the ShipmentOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipment) GetShipmentOptionsOk() (*ShipmentOptions, bool) {
	if o == nil || IsNil(o.ShipmentOptions) {
		return nil, false
	}
	return o.ShipmentOptions, true
}

// HasShipmentOptions returns a boolean if a field has been set.
func (o *ReprintShipment) HasShipmentOptions() bool {
	if o != nil && !IsNil(o.ShipmentOptions) {
		return true
	}

	return false
}

// SetShipmentOptions gets a reference to the given ShipmentOptions and assigns it to the ShipmentOptions field.
func (o *ReprintShipment) SetShipmentOptions(v ShipmentOptions) {
	o.ShipmentOptions = &v
}

// GetToAddress returns the ToAddress field value if set, zero value otherwise.
func (o *ReprintShipment) GetToAddress() ReprintShipmentToAddress {
	if o == nil || IsNil(o.ToAddress) {
		var ret ReprintShipmentToAddress
		return ret
	}
	return *o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprintShipment) GetToAddressOk() (*ReprintShipmentToAddress, bool) {
	if o == nil || IsNil(o.ToAddress) {
		return nil, false
	}
	return o.ToAddress, true
}

// HasToAddress returns a boolean if a field has been set.
func (o *ReprintShipment) HasToAddress() bool {
	if o != nil && !IsNil(o.ToAddress) {
		return true
	}

	return false
}

// SetToAddress gets a reference to the given ReprintShipmentToAddress and assigns it to the ToAddress field.
func (o *ReprintShipment) SetToAddress(v ReprintShipmentToAddress) {
	o.ToAddress = &v
}

func (o ReprintShipment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReprintShipment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorrelationId) {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.FromAddress) {
		toSerialize["fromAddress"] = o.FromAddress
	}
	if !IsNil(o.Parcel) {
		toSerialize["parcel"] = o.Parcel
	}
	if !IsNil(o.ParcelTrackingNumber) {
		toSerialize["parcelTrackingNumber"] = o.ParcelTrackingNumber
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.ShipmentId) {
		toSerialize["shipmentId"] = o.ShipmentId
	}
	if !IsNil(o.ShipmentOptions) {
		toSerialize["shipmentOptions"] = o.ShipmentOptions
	}
	if !IsNil(o.ToAddress) {
		toSerialize["toAddress"] = o.ToAddress
	}
	return toSerialize, nil
}

type NullableReprintShipment struct {
	value *ReprintShipment
	isSet bool
}

func (v NullableReprintShipment) Get() *ReprintShipment {
	return v.value
}

func (v *NullableReprintShipment) Set(val *ReprintShipment) {
	v.value = val
	v.isSet = true
}

func (v NullableReprintShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableReprintShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReprintShipment(val *ReprintShipment) *NullableReprintShipment {
	return &NullableReprintShipment{value: val, isSet: true}
}

func (v NullableReprintShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReprintShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


