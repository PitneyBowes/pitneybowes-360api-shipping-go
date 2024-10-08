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

// checks if the ToAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToAddress{}

// ToAddress The complete address of the Recipient or Department (in case if the address is not pointed to any individual recipient).
type ToAddress struct {
	// The name of the Recipient. <br /> `Max length = 40`
	Name string `json:"name"`
	// The addressLine1 contains the Flat number, Building or Apartment Name/number (if any) or company name (if not residential) of the Recipient. <br /> `Max length = 40`
	AddressLine1 string `json:"addressLine1"`
	// The addressLine2 contains the Area or Street Name. This is an optional field. <br /> `Max length = 40`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// The addressLine3 contains other details for easy reach, e.g. Landmark. This is an optional field. <br /> `Max length = 40`
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// The name of the city or town the Recipient belongs to.
	CityTown string `json:"cityTown"`
	// The name of the State or Province the Sender belongs to. It is the `2-letter` State or Province Code for US or Canadian address(es). <br /> Below is the hyperlink for CA country that will navigate to its Province/State Codes page. Similarly, respective country users can check for their country- State/Province codes. <br /> Please switch to the `Search` tab, select `Country codes` radio button, enter the required country name or country code, and then click `SEARCH` button . <br /> `Max length = 2`
	StateProvince string `json:"stateProvince"`
	// The Postal Code or ZIP Code of the address.<br /> For US addresses, use either the `5-digit` or `9-digit` ZIP Code in one of the following formats: '12345' or '12345-6789'. <br /> While for CA addresses, use a `six-character` alphanumeric string Postal Code in this format: 'A1A 1A1'. ERR supports only US addresses.<br /> *NOTE: USPS supports only US location.*
	PostalCode string `json:"postalCode"`
	// The country in which the recipient's address is located. The value will be the two-character ISO Code of the country from the ISO country list. <br /> Use ISO 3166-1 Alpha-2 standard values. For best results this should be included, especially if the country name does not appear in any of the unparsedAddressLines. <br /> Below is the hyperlink, please select `Country codes` and then click `SEARCH` button.  <br /> *NOTE: USPS supports only US location.*
	CountryCode string `json:"countryCode"`
	// The name of the company, in case if the Recipient address is not residential.
	Company *string `json:"company,omitempty"`
	// This is recipient's phone number. Enter the digits with or without spaces or hyphens. <br /> `Max length = 15`.
	Phone *string `json:"phone,omitempty"`
	// This must be recipients's valid email. <br /> `Max length = 50` 
	Email *string `json:"email,omitempty"`
	// The specified address can be Residential or Official. In case if the address is Residential, the boolean value will be 'true', else it will take 'false'.
	Residential *bool `json:"residential,omitempty"`
}

type _ToAddress ToAddress

// NewToAddress instantiates a new ToAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToAddress(name string, addressLine1 string, cityTown string, stateProvince string, postalCode string, countryCode string) *ToAddress {
	this := ToAddress{}
	this.Name = name
	this.AddressLine1 = addressLine1
	this.CityTown = cityTown
	this.StateProvince = stateProvince
	this.PostalCode = postalCode
	this.CountryCode = countryCode
	return &this
}

// NewToAddressWithDefaults instantiates a new ToAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToAddressWithDefaults() *ToAddress {
	this := ToAddress{}
	return &this
}

// GetName returns the Name field value
func (o *ToAddress) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ToAddress) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ToAddress) SetName(v string) {
	o.Name = v
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *ToAddress) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *ToAddress) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *ToAddress) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *ToAddress) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToAddress) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *ToAddress) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *ToAddress) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *ToAddress) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToAddress) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *ToAddress) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *ToAddress) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetCityTown returns the CityTown field value
func (o *ToAddress) GetCityTown() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CityTown
}

// GetCityTownOk returns a tuple with the CityTown field value
// and a boolean to check if the value has been set.
func (o *ToAddress) GetCityTownOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CityTown, true
}

// SetCityTown sets field value
func (o *ToAddress) SetCityTown(v string) {
	o.CityTown = v
}

// GetStateProvince returns the StateProvince field value
func (o *ToAddress) GetStateProvince() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value
// and a boolean to check if the value has been set.
func (o *ToAddress) GetStateProvinceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateProvince, true
}

// SetStateProvince sets field value
func (o *ToAddress) SetStateProvince(v string) {
	o.StateProvince = v
}

// GetPostalCode returns the PostalCode field value
func (o *ToAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *ToAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *ToAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetCountryCode returns the CountryCode field value
func (o *ToAddress) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *ToAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *ToAddress) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ToAddress) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToAddress) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ToAddress) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *ToAddress) SetCompany(v string) {
	o.Company = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ToAddress) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToAddress) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ToAddress) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ToAddress) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ToAddress) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToAddress) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ToAddress) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ToAddress) SetEmail(v string) {
	o.Email = &v
}

// GetResidential returns the Residential field value if set, zero value otherwise.
func (o *ToAddress) GetResidential() bool {
	if o == nil || IsNil(o.Residential) {
		var ret bool
		return ret
	}
	return *o.Residential
}

// GetResidentialOk returns a tuple with the Residential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToAddress) GetResidentialOk() (*bool, bool) {
	if o == nil || IsNil(o.Residential) {
		return nil, false
	}
	return o.Residential, true
}

// HasResidential returns a boolean if a field has been set.
func (o *ToAddress) HasResidential() bool {
	if o != nil && !IsNil(o.Residential) {
		return true
	}

	return false
}

// SetResidential gets a reference to the given bool and assigns it to the Residential field.
func (o *ToAddress) SetResidential(v bool) {
	o.Residential = &v
}

func (o ToAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["addressLine1"] = o.AddressLine1
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if !IsNil(o.AddressLine3) {
		toSerialize["addressLine3"] = o.AddressLine3
	}
	toSerialize["cityTown"] = o.CityTown
	toSerialize["stateProvince"] = o.StateProvince
	toSerialize["postalCode"] = o.PostalCode
	toSerialize["countryCode"] = o.CountryCode
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Residential) {
		toSerialize["residential"] = o.Residential
	}
	return toSerialize, nil
}

func (o *ToAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"addressLine1",
		"cityTown",
		"stateProvince",
		"postalCode",
		"countryCode",
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

	varToAddress := _ToAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToAddress)

	if err != nil {
		return err
	}

	*o = ToAddress(varToAddress)

	return err
}

type NullableToAddress struct {
	value *ToAddress
	isSet bool
}

func (v NullableToAddress) Get() *ToAddress {
	return v.value
}

func (v *NullableToAddress) Set(val *ToAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableToAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableToAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToAddress(val *ToAddress) *NullableToAddress {
	return &NullableToAddress{value: val, isSet: true}
}

func (v NullableToAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


