# AddressSuggestResponseSuggestionsAddressesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** | The addressLine1 can contain the Flat number, Building or Apartment Name/number (if any) or company name (if not residential). | [optional] 
**AddressLine2** | Pointer to **string** | The addressLine2 contains Street address or Landmark (if any). | [optional] 
**AddressLine3** | Pointer to **string** | The addressLine3 contains P.O.Box (if any) near the address. | [optional] 
**CityTown** | Pointer to **string** | The name of the city or town to where the address belongs. | [optional] 
**Name** | Pointer to **string** | Name of the person to which this address points. | [optional] 
**PostalCode** | Pointer to **string** | The postal code or ZIP code of the address. For US addresses, use either the 5-digit or 9-digit ZIP code in one of the following formats: &#39;12345&#39; or &#39;12345-6789&#39;. If you use a different format, such as 12345- or 123451234, will receive an error. | [optional] 
**Residential** | Pointer to **bool** | The specified adress can be Residential or Official. In case if the adress is Residential, the boolean value will be &#39;true&#39;, else it will take &#39;false&#39;. | [optional] 
**StateProvince** | Pointer to **string** | The State or Province of the address. For a US or Canadian address, it is the 2-letter state or province code.  | [optional] 

## Methods

### NewAddressSuggestResponseSuggestionsAddressesInner

`func NewAddressSuggestResponseSuggestionsAddressesInner() *AddressSuggestResponseSuggestionsAddressesInner`

NewAddressSuggestResponseSuggestionsAddressesInner instantiates a new AddressSuggestResponseSuggestionsAddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressSuggestResponseSuggestionsAddressesInnerWithDefaults

`func NewAddressSuggestResponseSuggestionsAddressesInnerWithDefaults() *AddressSuggestResponseSuggestionsAddressesInner`

NewAddressSuggestResponseSuggestionsAddressesInnerWithDefaults instantiates a new AddressSuggestResponseSuggestionsAddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *AddressSuggestResponseSuggestionsAddressesInner) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *AddressSuggestResponseSuggestionsAddressesInner) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AddressSuggestResponseSuggestionsAddressesInner) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *AddressSuggestResponseSuggestionsAddressesInner) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *AddressSuggestResponseSuggestionsAddressesInner) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *AddressSuggestResponseSuggestionsAddressesInner) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCityTown

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetCityTown() string`

GetCityTown returns the CityTown field if non-nil, zero value otherwise.

### GetCityTownOk

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetCityTownOk() (*string, bool)`

GetCityTownOk returns a tuple with the CityTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTown

`func (o *AddressSuggestResponseSuggestionsAddressesInner) SetCityTown(v string)`

SetCityTown sets CityTown field to given value.

### HasCityTown

`func (o *AddressSuggestResponseSuggestionsAddressesInner) HasCityTown() bool`

HasCityTown returns a boolean if a field has been set.

### GetName

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressSuggestResponseSuggestionsAddressesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddressSuggestResponseSuggestionsAddressesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPostalCode

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressSuggestResponseSuggestionsAddressesInner) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AddressSuggestResponseSuggestionsAddressesInner) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetResidential

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetResidential() bool`

GetResidential returns the Residential field if non-nil, zero value otherwise.

### GetResidentialOk

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetResidentialOk() (*bool, bool)`

GetResidentialOk returns a tuple with the Residential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidential

`func (o *AddressSuggestResponseSuggestionsAddressesInner) SetResidential(v bool)`

SetResidential sets Residential field to given value.

### HasResidential

`func (o *AddressSuggestResponseSuggestionsAddressesInner) HasResidential() bool`

HasResidential returns a boolean if a field has been set.

### GetStateProvince

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *AddressSuggestResponseSuggestionsAddressesInner) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *AddressSuggestResponseSuggestionsAddressesInner) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *AddressSuggestResponseSuggestionsAddressesInner) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


