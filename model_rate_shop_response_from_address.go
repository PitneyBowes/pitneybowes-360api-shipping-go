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

// checks if the RateShopResponseFromAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateShopResponseFromAddress{}

// RateShopResponseFromAddress Sender address details
type RateShopResponseFromAddress struct {
	// The addressLine1 can contain the Flat number, Building or Apartment Name/number (if any) or company name (if not residential).
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// The addressLine2 contains Street address or Landmark (if any).
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// The addressLine3 contains P.O.Box (if any) near the address.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// The name of the city or town to where the sender's address belongs.
	CityTown *string `json:"cityTown,omitempty"`
	// This indicates the name of the company, in case if the sender address is not residential.
	Company *string `json:"company,omitempty"`
	// This indicates the two-character ISO code of the source country from the ISO country list.
	CountryCode *string `json:"countryCode,omitempty"`
	// The email address of the sender. It can be person's email address or company email address (for non-residential).
	Email *string `json:"email,omitempty"`
	// Name of the sender to which this address points.
	Name *string `json:"name,omitempty"`
	// This is sender's phone number. Enter the digits with or without spaces or hyphens. The maximum characters for Phone number is 10 digits. 
	Phone *string `json:"phone,omitempty"`
	// The postal code or ZIP code of the address. For US addresses, use either the 5-digit or 9-digit ZIP code in one of the following formats: '12345' or '12345-6789'. If you use a different format, such as 12345- or 123451234, will receive an error.
	PostalCode *string `json:"postalCode,omitempty"`
	// The specified adress can be Residential or Official. In case if the adress is Residential, the boolean value will be 'true', else it will take 'false'.
	Residential *bool `json:"residential,omitempty"`
	// The State or Province of the address. For a US or Canadian address, it is the 2-letter state or province code. 
	StateProvince *string `json:"stateProvince,omitempty"`
}

// NewRateShopResponseFromAddress instantiates a new RateShopResponseFromAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateShopResponseFromAddress() *RateShopResponseFromAddress {
	this := RateShopResponseFromAddress{}
	return &this
}

// NewRateShopResponseFromAddressWithDefaults instantiates a new RateShopResponseFromAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateShopResponseFromAddressWithDefaults() *RateShopResponseFromAddress {
	this := RateShopResponseFromAddress{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *RateShopResponseFromAddress) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseFromAddress) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *RateShopResponseFromAddress) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *RateShopResponseFromAddress) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *RateShopResponseFromAddress) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseFromAddress) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *RateShopResponseFromAddress) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *RateShopResponseFromAddress) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *RateShopResponseFromAddress) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseFromAddress) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *RateShopResponseFromAddress) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *RateShopResponseFromAddress) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetCityTown returns the CityTown field value if set, zero value otherwise.
func (o *RateShopResponseFromAddress) GetCityTown() string {
	if o == nil || IsNil(o.CityTown) {
		var ret string
		return ret
	}
	return *o.CityTown
}

// GetCityTownOk returns a tuple with the CityTown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseFromAddress) GetCityTownOk() (*string, bool) {
	if o == nil || IsNil(o.CityTown) {
		return nil, false
	}
	return o.CityTown, true
}

// HasCityTown returns a boolean if a field has been set.
func (o *RateShopResponseFromAddress) HasCityTown() bool {
	if o != nil && !IsNil(o.CityTown) {
		return true
	}

	return false
}

// SetCityTown gets a reference to the given string and assigns it to the CityTown field.
func (o *RateShopResponseFromAddress) SetCityTown(v string) {
	o.CityTown = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *RateShopResponseFromAddress) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseFromAddress) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *RateShopResponseFromAddress) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *RateShopResponseFromAddress) SetCompany(v string) {
	o.Company = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *RateShopResponseFromAddress) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseFromAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *RateShopResponseFromAddress) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *RateShopResponseFromAddress) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RateShopResponseFromAddress) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseFromAddress) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *RateShopResponseFromAddress) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RateShopResponseFromAddress) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RateShopResponseFromAddress) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseFromAddress) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RateShopResponseFromAddress) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RateShopResponseFromAddress) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *RateShopResponseFromAddress) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseFromAddress) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *RateShopResponseFromAddress) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *RateShopResponseFromAddress) SetPhone(v string) {
	o.Phone = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *RateShopResponseFromAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseFromAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *RateShopResponseFromAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *RateShopResponseFromAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetResidential returns the Residential field value if set, zero value otherwise.
func (o *RateShopResponseFromAddress) GetResidential() bool {
	if o == nil || IsNil(o.Residential) {
		var ret bool
		return ret
	}
	return *o.Residential
}

// GetResidentialOk returns a tuple with the Residential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseFromAddress) GetResidentialOk() (*bool, bool) {
	if o == nil || IsNil(o.Residential) {
		return nil, false
	}
	return o.Residential, true
}

// HasResidential returns a boolean if a field has been set.
func (o *RateShopResponseFromAddress) HasResidential() bool {
	if o != nil && !IsNil(o.Residential) {
		return true
	}

	return false
}

// SetResidential gets a reference to the given bool and assigns it to the Residential field.
func (o *RateShopResponseFromAddress) SetResidential(v bool) {
	o.Residential = &v
}

// GetStateProvince returns the StateProvince field value if set, zero value otherwise.
func (o *RateShopResponseFromAddress) GetStateProvince() string {
	if o == nil || IsNil(o.StateProvince) {
		var ret string
		return ret
	}
	return *o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponseFromAddress) GetStateProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.StateProvince) {
		return nil, false
	}
	return o.StateProvince, true
}

// HasStateProvince returns a boolean if a field has been set.
func (o *RateShopResponseFromAddress) HasStateProvince() bool {
	if o != nil && !IsNil(o.StateProvince) {
		return true
	}

	return false
}

// SetStateProvince gets a reference to the given string and assigns it to the StateProvince field.
func (o *RateShopResponseFromAddress) SetStateProvince(v string) {
	o.StateProvince = &v
}

func (o RateShopResponseFromAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateShopResponseFromAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressLine1) {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if !IsNil(o.AddressLine3) {
		toSerialize["addressLine3"] = o.AddressLine3
	}
	if !IsNil(o.CityTown) {
		toSerialize["cityTown"] = o.CityTown
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.Residential) {
		toSerialize["residential"] = o.Residential
	}
	if !IsNil(o.StateProvince) {
		toSerialize["stateProvince"] = o.StateProvince
	}
	return toSerialize, nil
}

type NullableRateShopResponseFromAddress struct {
	value *RateShopResponseFromAddress
	isSet bool
}

func (v NullableRateShopResponseFromAddress) Get() *RateShopResponseFromAddress {
	return v.value
}

func (v *NullableRateShopResponseFromAddress) Set(val *RateShopResponseFromAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableRateShopResponseFromAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableRateShopResponseFromAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateShopResponseFromAddress(val *RateShopResponseFromAddress) *NullableRateShopResponseFromAddress {
	return &NullableRateShopResponseFromAddress{value: val, isSet: true}
}

func (v NullableRateShopResponseFromAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateShopResponseFromAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


