# FromAddressV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Sender to which this address points.&lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | 
**AddressLine1** | **string** | The addressLine1 contains the Flat number, Building or Apartment Name/number (if any) or company name (if not residential) of the Sender. &lt;br /&gt; &#x60;Max length &#x3D; 35&#x60;. | 
**CityTown** | **string** | The name of the city or town the Sender&#39;s address belongs to. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | 
**StateProvince** | **string** | The name of the State or Province the Sender belongs to. It is the &#x60;2-letter&#x60; State or Province Code for US or Canadian address(es). &lt;br /&gt; Below is the hyperlink for CA country that will navigate to its Province/State Codes page. Similarly, respective country users can check for their country- State/Province codes. &lt;br /&gt; Please switch to the &#x60;Search&#x60; tab, select &#x60;Country codes&#x60; radio button, enter the required country name or country code, and then click &#x60;SEARCH&#x60; button. &lt;br /&gt; &#x60;Max length &#x3D; 2&#x60; | 
**PostalCode** | **string** | The Postal Code or ZIP Code of the address. &lt;br /&gt; For CA addresses, use a &#x60;six-character&#x60; alphanumeric string Postal Code in this format: &#39;A1A 1A1&#39;. &lt;br /&gt; While for US addresses, use either the &#x60;5-digit&#x60; or &#x60;9-digit&#x60; ZIP Code in one of the following formats: &#39;12345&#39; or &#39;12345-6789&#39;. &lt;br /&gt; &#x60;Max length &#x3D; 10&#x60; | 
**CountryCode** | **string** | The country in which the sender&#39;s address is located. The value will be the two-character ISO Code of the country from the ISO country list. &lt;br /&gt; Use ISO 3166-1 Alpha-2 standard values. For best results this should be included, especially if the country name does not appear in any of the unparsedAddressLines. &lt;br /&gt; Below is the hyperlink, please select &#x60;Country codes&#x60; and then click &#x60;SEARCH&#x60; button. &lt;br /&gt; &#x60;Max length &#x3D; 2&#x60; | 
**Company** | Pointer to **string** | The name of the company, in case if the sender address is not residential. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 
**Phone** | **string** | This is sender&#39;s phone number. Enter the digits with or without spaces or hyphens. &lt;br /&gt; &#x60;Max length &#x3D; 15&#x60;. | 
**Email** | Pointer to **string** | This must be sender&#39;s valid email. Email is required if the customer is using GoFor Carrier. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;  | [optional] 
**Residential** | Pointer to **bool** | The specified address can be Residential or Official. In case if the address is Residential, the boolean value will be &#39;true&#39;, else it will take &#39;false&#39;. | [optional] 

## Methods

### NewFromAddressV2

`func NewFromAddressV2(name string, addressLine1 string, cityTown string, stateProvince string, postalCode string, countryCode string, phone string, ) *FromAddressV2`

NewFromAddressV2 instantiates a new FromAddressV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFromAddressV2WithDefaults

`func NewFromAddressV2WithDefaults() *FromAddressV2`

NewFromAddressV2WithDefaults instantiates a new FromAddressV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FromAddressV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FromAddressV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FromAddressV2) SetName(v string)`

SetName sets Name field to given value.


### GetAddressLine1

`func (o *FromAddressV2) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *FromAddressV2) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *FromAddressV2) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetCityTown

`func (o *FromAddressV2) GetCityTown() string`

GetCityTown returns the CityTown field if non-nil, zero value otherwise.

### GetCityTownOk

`func (o *FromAddressV2) GetCityTownOk() (*string, bool)`

GetCityTownOk returns a tuple with the CityTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTown

`func (o *FromAddressV2) SetCityTown(v string)`

SetCityTown sets CityTown field to given value.


### GetStateProvince

`func (o *FromAddressV2) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *FromAddressV2) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *FromAddressV2) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.


### GetPostalCode

`func (o *FromAddressV2) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *FromAddressV2) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *FromAddressV2) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCountryCode

`func (o *FromAddressV2) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *FromAddressV2) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *FromAddressV2) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCompany

`func (o *FromAddressV2) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *FromAddressV2) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *FromAddressV2) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *FromAddressV2) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetPhone

`func (o *FromAddressV2) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *FromAddressV2) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *FromAddressV2) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetEmail

`func (o *FromAddressV2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FromAddressV2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FromAddressV2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FromAddressV2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetResidential

`func (o *FromAddressV2) GetResidential() bool`

GetResidential returns the Residential field if non-nil, zero value otherwise.

### GetResidentialOk

`func (o *FromAddressV2) GetResidentialOk() (*bool, bool)`

GetResidentialOk returns a tuple with the Residential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidential

`func (o *FromAddressV2) SetResidential(v bool)`

SetResidential sets Residential field to given value.

### HasResidential

`func (o *FromAddressV2) HasResidential() bool`

HasResidential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


