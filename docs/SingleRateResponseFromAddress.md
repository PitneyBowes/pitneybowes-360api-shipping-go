# SingleRateResponseFromAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** | The addressLine1 can contain the Flat number, Building or Apartment Name/number (if any) or company name (if not residential). | [optional] 
**AddressLine2** | Pointer to **string** | The addressLine2 contains Street address or Landmark (if any). | [optional] 
**AddressLine3** | Pointer to **string** | The addressLine3 contains P.O.Box (if any) near the address. | [optional] 
**CityTown** | Pointer to **string** | The name of the city or town to where the sender&#39;s address belongs. | [optional] 
**Company** | Pointer to **string** | This indicates the name of the company, in case if the sender address is not residential. | [optional] 
**CountryCode** | Pointer to **string** | This indicates the two-character ISO code of the source country from the ISO country list. | [optional] 
**Name** | Pointer to **string** | Name of the sender to which this address points. | [optional] 
**PostalCode** | Pointer to **string** | The postal code or ZIP code of the address. For US addresses, use either the 5-digit or 9-digit ZIP code in one of the following formats: &#39;12345&#39; or &#39;12345-6789&#39;. If you use a different format, such as 12345- or 123451234, will receive an error. | [optional] 
**StateProvince** | Pointer to **string** | The State or Province of the address. For a US or Canadian address, it is the 2-letter state or province code.  | [optional] 

## Methods

### NewSingleRateResponseFromAddress

`func NewSingleRateResponseFromAddress() *SingleRateResponseFromAddress`

NewSingleRateResponseFromAddress instantiates a new SingleRateResponseFromAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleRateResponseFromAddressWithDefaults

`func NewSingleRateResponseFromAddressWithDefaults() *SingleRateResponseFromAddress`

NewSingleRateResponseFromAddressWithDefaults instantiates a new SingleRateResponseFromAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *SingleRateResponseFromAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *SingleRateResponseFromAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *SingleRateResponseFromAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *SingleRateResponseFromAddress) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *SingleRateResponseFromAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *SingleRateResponseFromAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *SingleRateResponseFromAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *SingleRateResponseFromAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *SingleRateResponseFromAddress) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *SingleRateResponseFromAddress) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *SingleRateResponseFromAddress) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *SingleRateResponseFromAddress) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCityTown

`func (o *SingleRateResponseFromAddress) GetCityTown() string`

GetCityTown returns the CityTown field if non-nil, zero value otherwise.

### GetCityTownOk

`func (o *SingleRateResponseFromAddress) GetCityTownOk() (*string, bool)`

GetCityTownOk returns a tuple with the CityTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTown

`func (o *SingleRateResponseFromAddress) SetCityTown(v string)`

SetCityTown sets CityTown field to given value.

### HasCityTown

`func (o *SingleRateResponseFromAddress) HasCityTown() bool`

HasCityTown returns a boolean if a field has been set.

### GetCompany

`func (o *SingleRateResponseFromAddress) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SingleRateResponseFromAddress) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SingleRateResponseFromAddress) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *SingleRateResponseFromAddress) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountryCode

`func (o *SingleRateResponseFromAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SingleRateResponseFromAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SingleRateResponseFromAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SingleRateResponseFromAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetName

`func (o *SingleRateResponseFromAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SingleRateResponseFromAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SingleRateResponseFromAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SingleRateResponseFromAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPostalCode

`func (o *SingleRateResponseFromAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *SingleRateResponseFromAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *SingleRateResponseFromAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *SingleRateResponseFromAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStateProvince

`func (o *SingleRateResponseFromAddress) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *SingleRateResponseFromAddress) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *SingleRateResponseFromAddress) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *SingleRateResponseFromAddress) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


