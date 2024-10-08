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

// checks if the ShipmentERRCoversheet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentERRCoversheet{}

// ShipmentERRCoversheet struct for ShipmentERRCoversheet
type ShipmentERRCoversheet struct {
	// This is a user-defined value provided by users just for their reference. This is for mapping purpose against each shipment.
	ExternalId *string `json:"externalId,omitempty"`
	FromAddress FromAddress `json:"fromAddress"`
	ToAddress ToAddress `json:"toAddress"`
	Parcel Parcel `json:"parcel"`
	// A unique identifier associated with the user's registered USPS account which is used by client users while shipment process.
	CarrierAccountId *string `json:"carrierAccountId,omitempty"`
	// >-Packaging type varies as per USPS selected services, e.g., LTR, LGENV.
	ParcelType *string `json:"parcelType,omitempty"`
	// >-A unique identifier given to the carrier-specific service. ERR supports two services: First Class Mail (FCM) and Priority Mail (PM).
	ServiceId *string `json:"serviceId,omitempty"`
	// The date when shipment gets created.
	DateOfShipment *string `json:"dateOfShipment,omitempty"`
	SpecialServices []SpecialServiceERRInner `json:"specialServices,omitempty"`
	ShipmentOptions *ShipmentOptionsERR `json:"shipmentOptions,omitempty"`
	// Additional metadata that needs to be stored for this shipment, can be added here. For now, `costAccountName` is supported.
	Metadata []ShipmentERRCoversheetMetadataInner `json:"metadata,omitempty"`
}

type _ShipmentERRCoversheet ShipmentERRCoversheet

// NewShipmentERRCoversheet instantiates a new ShipmentERRCoversheet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentERRCoversheet(fromAddress FromAddress, toAddress ToAddress, parcel Parcel) *ShipmentERRCoversheet {
	this := ShipmentERRCoversheet{}
	this.FromAddress = fromAddress
	this.ToAddress = toAddress
	this.Parcel = parcel
	return &this
}

// NewShipmentERRCoversheetWithDefaults instantiates a new ShipmentERRCoversheet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentERRCoversheetWithDefaults() *ShipmentERRCoversheet {
	this := ShipmentERRCoversheet{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ShipmentERRCoversheet) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERRCoversheet) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ShipmentERRCoversheet) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ShipmentERRCoversheet) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetFromAddress returns the FromAddress field value
func (o *ShipmentERRCoversheet) GetFromAddress() FromAddress {
	if o == nil {
		var ret FromAddress
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *ShipmentERRCoversheet) GetFromAddressOk() (*FromAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *ShipmentERRCoversheet) SetFromAddress(v FromAddress) {
	o.FromAddress = v
}

// GetToAddress returns the ToAddress field value
func (o *ShipmentERRCoversheet) GetToAddress() ToAddress {
	if o == nil {
		var ret ToAddress
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *ShipmentERRCoversheet) GetToAddressOk() (*ToAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *ShipmentERRCoversheet) SetToAddress(v ToAddress) {
	o.ToAddress = v
}

// GetParcel returns the Parcel field value
func (o *ShipmentERRCoversheet) GetParcel() Parcel {
	if o == nil {
		var ret Parcel
		return ret
	}

	return o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value
// and a boolean to check if the value has been set.
func (o *ShipmentERRCoversheet) GetParcelOk() (*Parcel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parcel, true
}

// SetParcel sets field value
func (o *ShipmentERRCoversheet) SetParcel(v Parcel) {
	o.Parcel = v
}

// GetCarrierAccountId returns the CarrierAccountId field value if set, zero value otherwise.
func (o *ShipmentERRCoversheet) GetCarrierAccountId() string {
	if o == nil || IsNil(o.CarrierAccountId) {
		var ret string
		return ret
	}
	return *o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERRCoversheet) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierAccountId) {
		return nil, false
	}
	return o.CarrierAccountId, true
}

// HasCarrierAccountId returns a boolean if a field has been set.
func (o *ShipmentERRCoversheet) HasCarrierAccountId() bool {
	if o != nil && !IsNil(o.CarrierAccountId) {
		return true
	}

	return false
}

// SetCarrierAccountId gets a reference to the given string and assigns it to the CarrierAccountId field.
func (o *ShipmentERRCoversheet) SetCarrierAccountId(v string) {
	o.CarrierAccountId = &v
}

// GetParcelType returns the ParcelType field value if set, zero value otherwise.
func (o *ShipmentERRCoversheet) GetParcelType() string {
	if o == nil || IsNil(o.ParcelType) {
		var ret string
		return ret
	}
	return *o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERRCoversheet) GetParcelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelType) {
		return nil, false
	}
	return o.ParcelType, true
}

// HasParcelType returns a boolean if a field has been set.
func (o *ShipmentERRCoversheet) HasParcelType() bool {
	if o != nil && !IsNil(o.ParcelType) {
		return true
	}

	return false
}

// SetParcelType gets a reference to the given string and assigns it to the ParcelType field.
func (o *ShipmentERRCoversheet) SetParcelType(v string) {
	o.ParcelType = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ShipmentERRCoversheet) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERRCoversheet) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ShipmentERRCoversheet) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *ShipmentERRCoversheet) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetDateOfShipment returns the DateOfShipment field value if set, zero value otherwise.
func (o *ShipmentERRCoversheet) GetDateOfShipment() string {
	if o == nil || IsNil(o.DateOfShipment) {
		var ret string
		return ret
	}
	return *o.DateOfShipment
}

// GetDateOfShipmentOk returns a tuple with the DateOfShipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERRCoversheet) GetDateOfShipmentOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfShipment) {
		return nil, false
	}
	return o.DateOfShipment, true
}

// HasDateOfShipment returns a boolean if a field has been set.
func (o *ShipmentERRCoversheet) HasDateOfShipment() bool {
	if o != nil && !IsNil(o.DateOfShipment) {
		return true
	}

	return false
}

// SetDateOfShipment gets a reference to the given string and assigns it to the DateOfShipment field.
func (o *ShipmentERRCoversheet) SetDateOfShipment(v string) {
	o.DateOfShipment = &v
}

// GetSpecialServices returns the SpecialServices field value if set, zero value otherwise.
func (o *ShipmentERRCoversheet) GetSpecialServices() []SpecialServiceERRInner {
	if o == nil || IsNil(o.SpecialServices) {
		var ret []SpecialServiceERRInner
		return ret
	}
	return o.SpecialServices
}

// GetSpecialServicesOk returns a tuple with the SpecialServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERRCoversheet) GetSpecialServicesOk() ([]SpecialServiceERRInner, bool) {
	if o == nil || IsNil(o.SpecialServices) {
		return nil, false
	}
	return o.SpecialServices, true
}

// HasSpecialServices returns a boolean if a field has been set.
func (o *ShipmentERRCoversheet) HasSpecialServices() bool {
	if o != nil && !IsNil(o.SpecialServices) {
		return true
	}

	return false
}

// SetSpecialServices gets a reference to the given []SpecialServiceERRInner and assigns it to the SpecialServices field.
func (o *ShipmentERRCoversheet) SetSpecialServices(v []SpecialServiceERRInner) {
	o.SpecialServices = v
}

// GetShipmentOptions returns the ShipmentOptions field value if set, zero value otherwise.
func (o *ShipmentERRCoversheet) GetShipmentOptions() ShipmentOptionsERR {
	if o == nil || IsNil(o.ShipmentOptions) {
		var ret ShipmentOptionsERR
		return ret
	}
	return *o.ShipmentOptions
}

// GetShipmentOptionsOk returns a tuple with the ShipmentOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERRCoversheet) GetShipmentOptionsOk() (*ShipmentOptionsERR, bool) {
	if o == nil || IsNil(o.ShipmentOptions) {
		return nil, false
	}
	return o.ShipmentOptions, true
}

// HasShipmentOptions returns a boolean if a field has been set.
func (o *ShipmentERRCoversheet) HasShipmentOptions() bool {
	if o != nil && !IsNil(o.ShipmentOptions) {
		return true
	}

	return false
}

// SetShipmentOptions gets a reference to the given ShipmentOptionsERR and assigns it to the ShipmentOptions field.
func (o *ShipmentERRCoversheet) SetShipmentOptions(v ShipmentOptionsERR) {
	o.ShipmentOptions = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ShipmentERRCoversheet) GetMetadata() []ShipmentERRCoversheetMetadataInner {
	if o == nil || IsNil(o.Metadata) {
		var ret []ShipmentERRCoversheetMetadataInner
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentERRCoversheet) GetMetadataOk() ([]ShipmentERRCoversheetMetadataInner, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ShipmentERRCoversheet) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []ShipmentERRCoversheetMetadataInner and assigns it to the Metadata field.
func (o *ShipmentERRCoversheet) SetMetadata(v []ShipmentERRCoversheetMetadataInner) {
	o.Metadata = v
}

func (o ShipmentERRCoversheet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentERRCoversheet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	toSerialize["fromAddress"] = o.FromAddress
	toSerialize["toAddress"] = o.ToAddress
	toSerialize["parcel"] = o.Parcel
	if !IsNil(o.CarrierAccountId) {
		toSerialize["carrierAccountId"] = o.CarrierAccountId
	}
	if !IsNil(o.ParcelType) {
		toSerialize["parcelType"] = o.ParcelType
	}
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !IsNil(o.DateOfShipment) {
		toSerialize["dateOfShipment"] = o.DateOfShipment
	}
	if !IsNil(o.SpecialServices) {
		toSerialize["specialServices"] = o.SpecialServices
	}
	if !IsNil(o.ShipmentOptions) {
		toSerialize["shipmentOptions"] = o.ShipmentOptions
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *ShipmentERRCoversheet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fromAddress",
		"toAddress",
		"parcel",
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

	varShipmentERRCoversheet := _ShipmentERRCoversheet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipmentERRCoversheet)

	if err != nil {
		return err
	}

	*o = ShipmentERRCoversheet(varShipmentERRCoversheet)

	return err
}

type NullableShipmentERRCoversheet struct {
	value *ShipmentERRCoversheet
	isSet bool
}

func (v NullableShipmentERRCoversheet) Get() *ShipmentERRCoversheet {
	return v.value
}

func (v *NullableShipmentERRCoversheet) Set(val *ShipmentERRCoversheet) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentERRCoversheet) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentERRCoversheet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentERRCoversheet(val *ShipmentERRCoversheet) *NullableShipmentERRCoversheet {
	return &NullableShipmentERRCoversheet{value: val, isSet: true}
}

func (v NullableShipmentERRCoversheet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentERRCoversheet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


