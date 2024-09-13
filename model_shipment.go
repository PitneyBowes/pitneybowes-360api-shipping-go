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

// checks if the Shipment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Shipment{}

// Shipment struct for Shipment
type Shipment struct {
	// The external ID of the shipment. User can provide any custom value to it for its own reference
	ExternalId *string `json:"externalId,omitempty"`
	FromAddress MultipieceDomesticShipmentRequestFromAddress `json:"fromAddress"`
	Parcel ShipmentDomesticParcel `json:"parcel"`
	// A unique identifier associated with the Carrier account used by client users during shipment process.
	CarrierAccountId string `json:"carrierAccountId"`
	// >-Packaging type specific to the carrier, e.g., FRPKG, LGENV, TUBE,PKG.
	ParcelType string `json:"parcelType"`
	// >-The abbreviated name of the carrier-specific service. Required for creating a shipment. Optional for rating a parcel.
	ServiceId string `json:"serviceId"`
	// Current Shipment date
	DateOfShipment *string `json:"dateOfShipment,omitempty"`
	SpecialServices []SpecialService `json:"specialServices,omitempty"`
	ShipmentOptions *ShipmentOptions `json:"shipmentOptions,omitempty"`
	// Additional metadata that needs to be stored for this shipment can be added here. For now, `costAccountName` is supported.
	Metadata []ShipmentMetadataInner `json:"metadata,omitempty"`
	ToAddress ShipmentToAddress `json:"toAddress"`
}

type _Shipment Shipment

// NewShipment instantiates a new Shipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipment(fromAddress MultipieceDomesticShipmentRequestFromAddress, parcel ShipmentDomesticParcel, carrierAccountId string, parcelType string, serviceId string, toAddress ShipmentToAddress) *Shipment {
	this := Shipment{}
	this.FromAddress = fromAddress
	this.Parcel = parcel
	this.CarrierAccountId = carrierAccountId
	this.ParcelType = parcelType
	this.ServiceId = serviceId
	this.ToAddress = toAddress
	return &this
}

// NewShipmentWithDefaults instantiates a new Shipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentWithDefaults() *Shipment {
	this := Shipment{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Shipment) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Shipment) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Shipment) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetFromAddress returns the FromAddress field value
func (o *Shipment) GetFromAddress() MultipieceDomesticShipmentRequestFromAddress {
	if o == nil {
		var ret MultipieceDomesticShipmentRequestFromAddress
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetFromAddressOk() (*MultipieceDomesticShipmentRequestFromAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *Shipment) SetFromAddress(v MultipieceDomesticShipmentRequestFromAddress) {
	o.FromAddress = v
}

// GetParcel returns the Parcel field value
func (o *Shipment) GetParcel() ShipmentDomesticParcel {
	if o == nil {
		var ret ShipmentDomesticParcel
		return ret
	}

	return o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetParcelOk() (*ShipmentDomesticParcel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parcel, true
}

// SetParcel sets field value
func (o *Shipment) SetParcel(v ShipmentDomesticParcel) {
	o.Parcel = v
}

// GetCarrierAccountId returns the CarrierAccountId field value
func (o *Shipment) GetCarrierAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierAccountId, true
}

// SetCarrierAccountId sets field value
func (o *Shipment) SetCarrierAccountId(v string) {
	o.CarrierAccountId = v
}

// GetParcelType returns the ParcelType field value
func (o *Shipment) GetParcelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetParcelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParcelType, true
}

// SetParcelType sets field value
func (o *Shipment) SetParcelType(v string) {
	o.ParcelType = v
}

// GetServiceId returns the ServiceId field value
func (o *Shipment) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *Shipment) SetServiceId(v string) {
	o.ServiceId = v
}

// GetDateOfShipment returns the DateOfShipment field value if set, zero value otherwise.
func (o *Shipment) GetDateOfShipment() string {
	if o == nil || IsNil(o.DateOfShipment) {
		var ret string
		return ret
	}
	return *o.DateOfShipment
}

// GetDateOfShipmentOk returns a tuple with the DateOfShipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetDateOfShipmentOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfShipment) {
		return nil, false
	}
	return o.DateOfShipment, true
}

// HasDateOfShipment returns a boolean if a field has been set.
func (o *Shipment) HasDateOfShipment() bool {
	if o != nil && !IsNil(o.DateOfShipment) {
		return true
	}

	return false
}

// SetDateOfShipment gets a reference to the given string and assigns it to the DateOfShipment field.
func (o *Shipment) SetDateOfShipment(v string) {
	o.DateOfShipment = &v
}

// GetSpecialServices returns the SpecialServices field value if set, zero value otherwise.
func (o *Shipment) GetSpecialServices() []SpecialService {
	if o == nil || IsNil(o.SpecialServices) {
		var ret []SpecialService
		return ret
	}
	return o.SpecialServices
}

// GetSpecialServicesOk returns a tuple with the SpecialServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetSpecialServicesOk() ([]SpecialService, bool) {
	if o == nil || IsNil(o.SpecialServices) {
		return nil, false
	}
	return o.SpecialServices, true
}

// HasSpecialServices returns a boolean if a field has been set.
func (o *Shipment) HasSpecialServices() bool {
	if o != nil && !IsNil(o.SpecialServices) {
		return true
	}

	return false
}

// SetSpecialServices gets a reference to the given []SpecialService and assigns it to the SpecialServices field.
func (o *Shipment) SetSpecialServices(v []SpecialService) {
	o.SpecialServices = v
}

// GetShipmentOptions returns the ShipmentOptions field value if set, zero value otherwise.
func (o *Shipment) GetShipmentOptions() ShipmentOptions {
	if o == nil || IsNil(o.ShipmentOptions) {
		var ret ShipmentOptions
		return ret
	}
	return *o.ShipmentOptions
}

// GetShipmentOptionsOk returns a tuple with the ShipmentOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetShipmentOptionsOk() (*ShipmentOptions, bool) {
	if o == nil || IsNil(o.ShipmentOptions) {
		return nil, false
	}
	return o.ShipmentOptions, true
}

// HasShipmentOptions returns a boolean if a field has been set.
func (o *Shipment) HasShipmentOptions() bool {
	if o != nil && !IsNil(o.ShipmentOptions) {
		return true
	}

	return false
}

// SetShipmentOptions gets a reference to the given ShipmentOptions and assigns it to the ShipmentOptions field.
func (o *Shipment) SetShipmentOptions(v ShipmentOptions) {
	o.ShipmentOptions = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Shipment) GetMetadata() []ShipmentMetadataInner {
	if o == nil || IsNil(o.Metadata) {
		var ret []ShipmentMetadataInner
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetMetadataOk() ([]ShipmentMetadataInner, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Shipment) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []ShipmentMetadataInner and assigns it to the Metadata field.
func (o *Shipment) SetMetadata(v []ShipmentMetadataInner) {
	o.Metadata = v
}

// GetToAddress returns the ToAddress field value
func (o *Shipment) GetToAddress() ShipmentToAddress {
	if o == nil {
		var ret ShipmentToAddress
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetToAddressOk() (*ShipmentToAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *Shipment) SetToAddress(v ShipmentToAddress) {
	o.ToAddress = v
}

func (o Shipment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Shipment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	toSerialize["fromAddress"] = o.FromAddress
	toSerialize["parcel"] = o.Parcel
	toSerialize["carrierAccountId"] = o.CarrierAccountId
	toSerialize["parcelType"] = o.ParcelType
	toSerialize["serviceId"] = o.ServiceId
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
	toSerialize["toAddress"] = o.ToAddress
	return toSerialize, nil
}

func (o *Shipment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fromAddress",
		"parcel",
		"carrierAccountId",
		"parcelType",
		"serviceId",
		"toAddress",
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

	varShipment := _Shipment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipment)

	if err != nil {
		return err
	}

	*o = Shipment(varShipment)

	return err
}

type NullableShipment struct {
	value *Shipment
	isSet bool
}

func (v NullableShipment) Get() *Shipment {
	return v.value
}

func (v *NullableShipment) Set(val *Shipment) {
	v.value = val
	v.isSet = true
}

func (v NullableShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipment(val *Shipment) *NullableShipment {
	return &NullableShipment{value: val, isSet: true}
}

func (v NullableShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


