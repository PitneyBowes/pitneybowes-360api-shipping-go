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

// checks if the ReturnLabelToAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnLabelToAddress{}

// ReturnLabelToAddress struct for ReturnLabelToAddress
type ReturnLabelToAddress struct {
	// The addressLine1 can contain the Flat number, Building or Apartment Name/number (if any) or company name (if not residential).
	AddressLine1 string `json:"addressLine1"`
	// Email id
	Email *string `json:"email,omitempty"`
	// The name of the city or town to where the address belongs.
	CityTown *string `json:"cityTown,omitempty"`
	// This indicates the two-character ISO code of the destination country from the ISO country list.
	CountryCode *string `json:"countryCode,omitempty"`
	// Name of the recipient to which this address points.
	Name *string `json:"name,omitempty"`
	// This indicates the name of the company, in case if the sender address is not residential.
	Company *string `json:"company,omitempty"`
	// This is recipient's phone number. Enter the digits with or without spaces or hyphens. The maximum characters for Phone number is 10 digits. 
	Phone *string `json:"phone,omitempty"`
	// The postal code or ZIP code of the address. For US addresses, use either the 5-digit or 9-digit ZIP code in one of the following formats: '12345' or '12345-6789'. If you use a different format, such as 12345- or 123451234, will receive an error.
	PostalCode *string `json:"postalCode,omitempty"`
	// The State or Province of the address. For a US or Canadian address, it is the 2-letter state or province code. 
	StateProvince *string `json:"stateProvince,omitempty"`
}

type _ReturnLabelToAddress ReturnLabelToAddress

// NewReturnLabelToAddress instantiates a new ReturnLabelToAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLabelToAddress(addressLine1 string) *ReturnLabelToAddress {
	this := ReturnLabelToAddress{}
	this.AddressLine1 = addressLine1
	return &this
}

// NewReturnLabelToAddressWithDefaults instantiates a new ReturnLabelToAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLabelToAddressWithDefaults() *ReturnLabelToAddress {
	this := ReturnLabelToAddress{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *ReturnLabelToAddress) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *ReturnLabelToAddress) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *ReturnLabelToAddress) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ReturnLabelToAddress) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelToAddress) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ReturnLabelToAddress) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ReturnLabelToAddress) SetEmail(v string) {
	o.Email = &v
}

// GetCityTown returns the CityTown field value if set, zero value otherwise.
func (o *ReturnLabelToAddress) GetCityTown() string {
	if o == nil || IsNil(o.CityTown) {
		var ret string
		return ret
	}
	return *o.CityTown
}

// GetCityTownOk returns a tuple with the CityTown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelToAddress) GetCityTownOk() (*string, bool) {
	if o == nil || IsNil(o.CityTown) {
		return nil, false
	}
	return o.CityTown, true
}

// HasCityTown returns a boolean if a field has been set.
func (o *ReturnLabelToAddress) HasCityTown() bool {
	if o != nil && !IsNil(o.CityTown) {
		return true
	}

	return false
}

// SetCityTown gets a reference to the given string and assigns it to the CityTown field.
func (o *ReturnLabelToAddress) SetCityTown(v string) {
	o.CityTown = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *ReturnLabelToAddress) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelToAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *ReturnLabelToAddress) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *ReturnLabelToAddress) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReturnLabelToAddress) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelToAddress) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReturnLabelToAddress) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReturnLabelToAddress) SetName(v string) {
	o.Name = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ReturnLabelToAddress) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelToAddress) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ReturnLabelToAddress) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *ReturnLabelToAddress) SetCompany(v string) {
	o.Company = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ReturnLabelToAddress) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelToAddress) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ReturnLabelToAddress) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ReturnLabelToAddress) SetPhone(v string) {
	o.Phone = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *ReturnLabelToAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelToAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *ReturnLabelToAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *ReturnLabelToAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStateProvince returns the StateProvince field value if set, zero value otherwise.
func (o *ReturnLabelToAddress) GetStateProvince() string {
	if o == nil || IsNil(o.StateProvince) {
		var ret string
		return ret
	}
	return *o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabelToAddress) GetStateProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.StateProvince) {
		return nil, false
	}
	return o.StateProvince, true
}

// HasStateProvince returns a boolean if a field has been set.
func (o *ReturnLabelToAddress) HasStateProvince() bool {
	if o != nil && !IsNil(o.StateProvince) {
		return true
	}

	return false
}

// SetStateProvince gets a reference to the given string and assigns it to the StateProvince field.
func (o *ReturnLabelToAddress) SetStateProvince(v string) {
	o.StateProvince = &v
}

func (o ReturnLabelToAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnLabelToAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addressLine1"] = o.AddressLine1
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.CityTown) {
		toSerialize["cityTown"] = o.CityTown
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.StateProvince) {
		toSerialize["stateProvince"] = o.StateProvince
	}
	return toSerialize, nil
}

func (o *ReturnLabelToAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addressLine1",
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

	varReturnLabelToAddress := _ReturnLabelToAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReturnLabelToAddress)

	if err != nil {
		return err
	}

	*o = ReturnLabelToAddress(varReturnLabelToAddress)

	return err
}

type NullableReturnLabelToAddress struct {
	value *ReturnLabelToAddress
	isSet bool
}

func (v NullableReturnLabelToAddress) Get() *ReturnLabelToAddress {
	return v.value
}

func (v *NullableReturnLabelToAddress) Set(val *ReturnLabelToAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLabelToAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLabelToAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLabelToAddress(val *ReturnLabelToAddress) *NullableReturnLabelToAddress {
	return &NullableReturnLabelToAddress{value: val, isSet: true}
}

func (v NullableReturnLabelToAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLabelToAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


