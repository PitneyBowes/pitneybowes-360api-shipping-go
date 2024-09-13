# DomesticShipmentResponseFromAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** | The addressLine1 can contain the Flat number, Building or Apartment Name/number (if any) or company name (if not residential). | [optional] 
**AddressLine2** | Pointer to **string** | The addressLine2 contains Street address or Landmark (if any). | [optional] 
**AddressLine3** | Pointer to **string** | The addressLine3 contains P.O.Box (if any) near the address. | [optional] 
**CityTown** | Pointer to **string** | The name of the city or town to where the address belongs. | [optional] 
**Company** | Pointer to **string** | This indicates the name of the company, in case if the sender address is not residential. | [optional] 
**CountryCode** | Pointer to **string** | This indicates the two-character ISO code of the source country from the ISO country list. | [optional] 
**Email** | Pointer to **string** | The email address of the sender. It can be person&#39;s email address or company email address (for non-residential). | [optional] 
**Name** | Pointer to **string** | Name of the sender to which this address points. | [optional] 
**Phone** | Pointer to **string** | This is sender&#39;s phone number. Enter the digits with or without spaces or hyphens. The maximum characters for Phone number is 10 digits.  | [optional] 
**PostalCode** | Pointer to **string** | The postal code or ZIP code of the address. For US addresses, use either the 5-digit or 9-digit ZIP code in one of the following formats: &#39;12345&#39; or &#39;12345-6789&#39;. If you use a different format, such as 12345- or 123451234, will receive an error. | [optional] 
**Residential** | Pointer to **bool** | The specified adress can be Residential or Official. In case if the adress is Residential, the boolean value will be &#39;true&#39;, else it will take &#39;false&#39;. | [optional] 
**StateProvince** | Pointer to **string** | The State or Province of the address. For a US or Canadian address, it is the 2-letter state or province code.  | [optional] 
**InductionPostalCode** | Pointer to **string** | The postal code where the shipment is tendered to the carrier. If an induction postal code is specified in the \&quot;fromAddress\&quot;, it will be used for rate calculations and determining manifest eligibility instead of the standard postal code. If not specified, the postal code from the \&quot;fromAddress\&quot; will be used. | [optional] 

## Methods

### NewDomesticShipmentResponseFromAddress

`func NewDomesticShipmentResponseFromAddress() *DomesticShipmentResponseFromAddress`

NewDomesticShipmentResponseFromAddress instantiates a new DomesticShipmentResponseFromAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomesticShipmentResponseFromAddressWithDefaults

`func NewDomesticShipmentResponseFromAddressWithDefaults() *DomesticShipmentResponseFromAddress`

NewDomesticShipmentResponseFromAddressWithDefaults instantiates a new DomesticShipmentResponseFromAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *DomesticShipmentResponseFromAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *DomesticShipmentResponseFromAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *DomesticShipmentResponseFromAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *DomesticShipmentResponseFromAddress) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *DomesticShipmentResponseFromAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *DomesticShipmentResponseFromAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *DomesticShipmentResponseFromAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *DomesticShipmentResponseFromAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *DomesticShipmentResponseFromAddress) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *DomesticShipmentResponseFromAddress) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *DomesticShipmentResponseFromAddress) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *DomesticShipmentResponseFromAddress) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCityTown

`func (o *DomesticShipmentResponseFromAddress) GetCityTown() string`

GetCityTown returns the CityTown field if non-nil, zero value otherwise.

### GetCityTownOk

`func (o *DomesticShipmentResponseFromAddress) GetCityTownOk() (*string, bool)`

GetCityTownOk returns a tuple with the CityTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTown

`func (o *DomesticShipmentResponseFromAddress) SetCityTown(v string)`

SetCityTown sets CityTown field to given value.

### HasCityTown

`func (o *DomesticShipmentResponseFromAddress) HasCityTown() bool`

HasCityTown returns a boolean if a field has been set.

### GetCompany

`func (o *DomesticShipmentResponseFromAddress) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *DomesticShipmentResponseFromAddress) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *DomesticShipmentResponseFromAddress) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *DomesticShipmentResponseFromAddress) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountryCode

`func (o *DomesticShipmentResponseFromAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *DomesticShipmentResponseFromAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *DomesticShipmentResponseFromAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *DomesticShipmentResponseFromAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEmail

`func (o *DomesticShipmentResponseFromAddress) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DomesticShipmentResponseFromAddress) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DomesticShipmentResponseFromAddress) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DomesticShipmentResponseFromAddress) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *DomesticShipmentResponseFromAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomesticShipmentResponseFromAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomesticShipmentResponseFromAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomesticShipmentResponseFromAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *DomesticShipmentResponseFromAddress) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *DomesticShipmentResponseFromAddress) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *DomesticShipmentResponseFromAddress) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *DomesticShipmentResponseFromAddress) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *DomesticShipmentResponseFromAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *DomesticShipmentResponseFromAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *DomesticShipmentResponseFromAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *DomesticShipmentResponseFromAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetResidential

`func (o *DomesticShipmentResponseFromAddress) GetResidential() bool`

GetResidential returns the Residential field if non-nil, zero value otherwise.

### GetResidentialOk

`func (o *DomesticShipmentResponseFromAddress) GetResidentialOk() (*bool, bool)`

GetResidentialOk returns a tuple with the Residential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidential

`func (o *DomesticShipmentResponseFromAddress) SetResidential(v bool)`

SetResidential sets Residential field to given value.

### HasResidential

`func (o *DomesticShipmentResponseFromAddress) HasResidential() bool`

HasResidential returns a boolean if a field has been set.

### GetStateProvince

`func (o *DomesticShipmentResponseFromAddress) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *DomesticShipmentResponseFromAddress) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *DomesticShipmentResponseFromAddress) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *DomesticShipmentResponseFromAddress) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.

### GetInductionPostalCode

`func (o *DomesticShipmentResponseFromAddress) GetInductionPostalCode() string`

GetInductionPostalCode returns the InductionPostalCode field if non-nil, zero value otherwise.

### GetInductionPostalCodeOk

`func (o *DomesticShipmentResponseFromAddress) GetInductionPostalCodeOk() (*string, bool)`

GetInductionPostalCodeOk returns a tuple with the InductionPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInductionPostalCode

`func (o *DomesticShipmentResponseFromAddress) SetInductionPostalCode(v string)`

SetInductionPostalCode sets InductionPostalCode field to given value.

### HasInductionPostalCode

`func (o *DomesticShipmentResponseFromAddress) HasInductionPostalCode() bool`

HasInductionPostalCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


