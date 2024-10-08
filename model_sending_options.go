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

// checks if the SendingOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendingOptions{}

// SendingOptions Sending Options will include carrier and its account details, sender details, and parcel details.
type SendingOptions struct {
	// Name of the carrier. E.g., FedEx.
	Carrier *string `json:"carrier,omitempty"`
	// Name of the carrier-based service. E.g., 2DA.
	Service *string `json:"service,omitempty"`
	CarrierAccounts *SendingOptionsCarrierAccounts `json:"carrierAccounts,omitempty"`
	// Size of the label, e.g., DOC_4X6.
	LabelSize *string `json:"labelSize,omitempty"`
	// Type of the Label, e.g., Shipping_Label.
	LabelType *string `json:"labelType,omitempty"`
	// Format of the Label, e.g., PDF.
	LabelFormat *string `json:"labelFormat,omitempty"`
	FromAddress *FromAddressV2 `json:"fromAddress,omitempty"`
	Parcel *ParcelV2 `json:"parcel,omitempty"`
}

// NewSendingOptions instantiates a new SendingOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendingOptions() *SendingOptions {
	this := SendingOptions{}
	return &this
}

// NewSendingOptionsWithDefaults instantiates a new SendingOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendingOptionsWithDefaults() *SendingOptions {
	this := SendingOptions{}
	return &this
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *SendingOptions) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendingOptions) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *SendingOptions) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *SendingOptions) SetCarrier(v string) {
	o.Carrier = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SendingOptions) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendingOptions) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SendingOptions) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SendingOptions) SetService(v string) {
	o.Service = &v
}

// GetCarrierAccounts returns the CarrierAccounts field value if set, zero value otherwise.
func (o *SendingOptions) GetCarrierAccounts() SendingOptionsCarrierAccounts {
	if o == nil || IsNil(o.CarrierAccounts) {
		var ret SendingOptionsCarrierAccounts
		return ret
	}
	return *o.CarrierAccounts
}

// GetCarrierAccountsOk returns a tuple with the CarrierAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendingOptions) GetCarrierAccountsOk() (*SendingOptionsCarrierAccounts, bool) {
	if o == nil || IsNil(o.CarrierAccounts) {
		return nil, false
	}
	return o.CarrierAccounts, true
}

// HasCarrierAccounts returns a boolean if a field has been set.
func (o *SendingOptions) HasCarrierAccounts() bool {
	if o != nil && !IsNil(o.CarrierAccounts) {
		return true
	}

	return false
}

// SetCarrierAccounts gets a reference to the given SendingOptionsCarrierAccounts and assigns it to the CarrierAccounts field.
func (o *SendingOptions) SetCarrierAccounts(v SendingOptionsCarrierAccounts) {
	o.CarrierAccounts = &v
}

// GetLabelSize returns the LabelSize field value if set, zero value otherwise.
func (o *SendingOptions) GetLabelSize() string {
	if o == nil || IsNil(o.LabelSize) {
		var ret string
		return ret
	}
	return *o.LabelSize
}

// GetLabelSizeOk returns a tuple with the LabelSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendingOptions) GetLabelSizeOk() (*string, bool) {
	if o == nil || IsNil(o.LabelSize) {
		return nil, false
	}
	return o.LabelSize, true
}

// HasLabelSize returns a boolean if a field has been set.
func (o *SendingOptions) HasLabelSize() bool {
	if o != nil && !IsNil(o.LabelSize) {
		return true
	}

	return false
}

// SetLabelSize gets a reference to the given string and assigns it to the LabelSize field.
func (o *SendingOptions) SetLabelSize(v string) {
	o.LabelSize = &v
}

// GetLabelType returns the LabelType field value if set, zero value otherwise.
func (o *SendingOptions) GetLabelType() string {
	if o == nil || IsNil(o.LabelType) {
		var ret string
		return ret
	}
	return *o.LabelType
}

// GetLabelTypeOk returns a tuple with the LabelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendingOptions) GetLabelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LabelType) {
		return nil, false
	}
	return o.LabelType, true
}

// HasLabelType returns a boolean if a field has been set.
func (o *SendingOptions) HasLabelType() bool {
	if o != nil && !IsNil(o.LabelType) {
		return true
	}

	return false
}

// SetLabelType gets a reference to the given string and assigns it to the LabelType field.
func (o *SendingOptions) SetLabelType(v string) {
	o.LabelType = &v
}

// GetLabelFormat returns the LabelFormat field value if set, zero value otherwise.
func (o *SendingOptions) GetLabelFormat() string {
	if o == nil || IsNil(o.LabelFormat) {
		var ret string
		return ret
	}
	return *o.LabelFormat
}

// GetLabelFormatOk returns a tuple with the LabelFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendingOptions) GetLabelFormatOk() (*string, bool) {
	if o == nil || IsNil(o.LabelFormat) {
		return nil, false
	}
	return o.LabelFormat, true
}

// HasLabelFormat returns a boolean if a field has been set.
func (o *SendingOptions) HasLabelFormat() bool {
	if o != nil && !IsNil(o.LabelFormat) {
		return true
	}

	return false
}

// SetLabelFormat gets a reference to the given string and assigns it to the LabelFormat field.
func (o *SendingOptions) SetLabelFormat(v string) {
	o.LabelFormat = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *SendingOptions) GetFromAddress() FromAddressV2 {
	if o == nil || IsNil(o.FromAddress) {
		var ret FromAddressV2
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendingOptions) GetFromAddressOk() (*FromAddressV2, bool) {
	if o == nil || IsNil(o.FromAddress) {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *SendingOptions) HasFromAddress() bool {
	if o != nil && !IsNil(o.FromAddress) {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given FromAddressV2 and assigns it to the FromAddress field.
func (o *SendingOptions) SetFromAddress(v FromAddressV2) {
	o.FromAddress = &v
}

// GetParcel returns the Parcel field value if set, zero value otherwise.
func (o *SendingOptions) GetParcel() ParcelV2 {
	if o == nil || IsNil(o.Parcel) {
		var ret ParcelV2
		return ret
	}
	return *o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendingOptions) GetParcelOk() (*ParcelV2, bool) {
	if o == nil || IsNil(o.Parcel) {
		return nil, false
	}
	return o.Parcel, true
}

// HasParcel returns a boolean if a field has been set.
func (o *SendingOptions) HasParcel() bool {
	if o != nil && !IsNil(o.Parcel) {
		return true
	}

	return false
}

// SetParcel gets a reference to the given ParcelV2 and assigns it to the Parcel field.
func (o *SendingOptions) SetParcel(v ParcelV2) {
	o.Parcel = &v
}

func (o SendingOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendingOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.CarrierAccounts) {
		toSerialize["carrierAccounts"] = o.CarrierAccounts
	}
	if !IsNil(o.LabelSize) {
		toSerialize["labelSize"] = o.LabelSize
	}
	if !IsNil(o.LabelType) {
		toSerialize["labelType"] = o.LabelType
	}
	if !IsNil(o.LabelFormat) {
		toSerialize["labelFormat"] = o.LabelFormat
	}
	if !IsNil(o.FromAddress) {
		toSerialize["fromAddress"] = o.FromAddress
	}
	if !IsNil(o.Parcel) {
		toSerialize["parcel"] = o.Parcel
	}
	return toSerialize, nil
}

type NullableSendingOptions struct {
	value *SendingOptions
	isSet bool
}

func (v NullableSendingOptions) Get() *SendingOptions {
	return v.value
}

func (v *NullableSendingOptions) Set(val *SendingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSendingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSendingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendingOptions(val *SendingOptions) *NullableSendingOptions {
	return &NullableSendingOptions{value: val, isSet: true}
}

func (v NullableSendingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


