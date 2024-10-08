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

// checks if the SingleRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleRate{}

// SingleRate struct for SingleRate
type SingleRate struct {
	// This defines the date of the Shipment in the format YYYY:MM:DD, required for future shipment rating
	DateOfShipment *string `json:"dateOfShipment,omitempty"`
	FromAddress SingleRateFromAddress `json:"fromAddress"`
	Parcel SingleRateParcel `json:"parcel"`
	//  This provides a single carrier account Id in case of single rate request. It can be referred from response of `Get Carrier Accounts` API.
	CarrierAccounts []string `json:"carrierAccounts"`
	// Parcel Type its value can be referred from response of `Get Parcel Types` API.
	ParcelType string `json:"parcelType"`
	// >-Parcel Id is optional and required to be given in case of branded parcels which have brandedDimension and/or brandedWeight. If parcel has brandedDimension, in that case user need not to pass dimensionUnit and dimension details(length, width and height) in the parcel object. And if brandedWeight is also available for the parcel then in that case weightUnit and weight need not to be passed  in parcel object  
	ParcelId *string `json:"parcelId,omitempty"`
	// Service to be used for rating, it can be referred from response of `Get Services` API
	ServiceId string `json:"serviceId"`
	// Special services to be used for rating, it can be referred from response of `Get Special Services` API
	SpecialServices []SpecialService `json:"specialServices,omitempty"`
	ToAddress SingleRateToAddress `json:"toAddress"`
}

type _SingleRate SingleRate

// NewSingleRate instantiates a new SingleRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleRate(fromAddress SingleRateFromAddress, parcel SingleRateParcel, carrierAccounts []string, parcelType string, serviceId string, toAddress SingleRateToAddress) *SingleRate {
	this := SingleRate{}
	this.FromAddress = fromAddress
	this.Parcel = parcel
	this.CarrierAccounts = carrierAccounts
	this.ParcelType = parcelType
	this.ServiceId = serviceId
	this.ToAddress = toAddress
	return &this
}

// NewSingleRateWithDefaults instantiates a new SingleRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleRateWithDefaults() *SingleRate {
	this := SingleRate{}
	return &this
}

// GetDateOfShipment returns the DateOfShipment field value if set, zero value otherwise.
func (o *SingleRate) GetDateOfShipment() string {
	if o == nil || IsNil(o.DateOfShipment) {
		var ret string
		return ret
	}
	return *o.DateOfShipment
}

// GetDateOfShipmentOk returns a tuple with the DateOfShipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRate) GetDateOfShipmentOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfShipment) {
		return nil, false
	}
	return o.DateOfShipment, true
}

// HasDateOfShipment returns a boolean if a field has been set.
func (o *SingleRate) HasDateOfShipment() bool {
	if o != nil && !IsNil(o.DateOfShipment) {
		return true
	}

	return false
}

// SetDateOfShipment gets a reference to the given string and assigns it to the DateOfShipment field.
func (o *SingleRate) SetDateOfShipment(v string) {
	o.DateOfShipment = &v
}

// GetFromAddress returns the FromAddress field value
func (o *SingleRate) GetFromAddress() SingleRateFromAddress {
	if o == nil {
		var ret SingleRateFromAddress
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *SingleRate) GetFromAddressOk() (*SingleRateFromAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *SingleRate) SetFromAddress(v SingleRateFromAddress) {
	o.FromAddress = v
}

// GetParcel returns the Parcel field value
func (o *SingleRate) GetParcel() SingleRateParcel {
	if o == nil {
		var ret SingleRateParcel
		return ret
	}

	return o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value
// and a boolean to check if the value has been set.
func (o *SingleRate) GetParcelOk() (*SingleRateParcel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parcel, true
}

// SetParcel sets field value
func (o *SingleRate) SetParcel(v SingleRateParcel) {
	o.Parcel = v
}

// GetCarrierAccounts returns the CarrierAccounts field value
func (o *SingleRate) GetCarrierAccounts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CarrierAccounts
}

// GetCarrierAccountsOk returns a tuple with the CarrierAccounts field value
// and a boolean to check if the value has been set.
func (o *SingleRate) GetCarrierAccountsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CarrierAccounts, true
}

// SetCarrierAccounts sets field value
func (o *SingleRate) SetCarrierAccounts(v []string) {
	o.CarrierAccounts = v
}

// GetParcelType returns the ParcelType field value
func (o *SingleRate) GetParcelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value
// and a boolean to check if the value has been set.
func (o *SingleRate) GetParcelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParcelType, true
}

// SetParcelType sets field value
func (o *SingleRate) SetParcelType(v string) {
	o.ParcelType = v
}

// GetParcelId returns the ParcelId field value if set, zero value otherwise.
func (o *SingleRate) GetParcelId() string {
	if o == nil || IsNil(o.ParcelId) {
		var ret string
		return ret
	}
	return *o.ParcelId
}

// GetParcelIdOk returns a tuple with the ParcelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRate) GetParcelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelId) {
		return nil, false
	}
	return o.ParcelId, true
}

// HasParcelId returns a boolean if a field has been set.
func (o *SingleRate) HasParcelId() bool {
	if o != nil && !IsNil(o.ParcelId) {
		return true
	}

	return false
}

// SetParcelId gets a reference to the given string and assigns it to the ParcelId field.
func (o *SingleRate) SetParcelId(v string) {
	o.ParcelId = &v
}

// GetServiceId returns the ServiceId field value
func (o *SingleRate) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *SingleRate) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *SingleRate) SetServiceId(v string) {
	o.ServiceId = v
}

// GetSpecialServices returns the SpecialServices field value if set, zero value otherwise.
func (o *SingleRate) GetSpecialServices() []SpecialService {
	if o == nil || IsNil(o.SpecialServices) {
		var ret []SpecialService
		return ret
	}
	return o.SpecialServices
}

// GetSpecialServicesOk returns a tuple with the SpecialServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRate) GetSpecialServicesOk() ([]SpecialService, bool) {
	if o == nil || IsNil(o.SpecialServices) {
		return nil, false
	}
	return o.SpecialServices, true
}

// HasSpecialServices returns a boolean if a field has been set.
func (o *SingleRate) HasSpecialServices() bool {
	if o != nil && !IsNil(o.SpecialServices) {
		return true
	}

	return false
}

// SetSpecialServices gets a reference to the given []SpecialService and assigns it to the SpecialServices field.
func (o *SingleRate) SetSpecialServices(v []SpecialService) {
	o.SpecialServices = v
}

// GetToAddress returns the ToAddress field value
func (o *SingleRate) GetToAddress() SingleRateToAddress {
	if o == nil {
		var ret SingleRateToAddress
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *SingleRate) GetToAddressOk() (*SingleRateToAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *SingleRate) SetToAddress(v SingleRateToAddress) {
	o.ToAddress = v
}

func (o SingleRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateOfShipment) {
		toSerialize["dateOfShipment"] = o.DateOfShipment
	}
	toSerialize["fromAddress"] = o.FromAddress
	toSerialize["parcel"] = o.Parcel
	toSerialize["carrierAccounts"] = o.CarrierAccounts
	toSerialize["parcelType"] = o.ParcelType
	if !IsNil(o.ParcelId) {
		toSerialize["parcelId"] = o.ParcelId
	}
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.SpecialServices) {
		toSerialize["specialServices"] = o.SpecialServices
	}
	toSerialize["toAddress"] = o.ToAddress
	return toSerialize, nil
}

func (o *SingleRate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fromAddress",
		"parcel",
		"carrierAccounts",
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

	varSingleRate := _SingleRate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSingleRate)

	if err != nil {
		return err
	}

	*o = SingleRate(varSingleRate)

	return err
}

type NullableSingleRate struct {
	value *SingleRate
	isSet bool
}

func (v NullableSingleRate) Get() *SingleRate {
	return v.value
}

func (v *NullableSingleRate) Set(val *SingleRate) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleRate) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleRate(val *SingleRate) *NullableSingleRate {
	return &NullableSingleRate{value: val, isSet: true}
}

func (v NullableSingleRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


