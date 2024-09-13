# SchedulePickupDHLEXPResponsePickupAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the person, It should not contain special character or numeric value | [optional] 
**AddressLine1** | Pointer to **string** | Address line1 of the pickup address | [optional] 
**CityTown** | Pointer to **string** | City of the pickup address | [optional] 
**StateProvince** | Pointer to **string** | State province of the pickup address | [optional] 
**PostalCode** | Pointer to **string** | Postal Code of the pickup address | [optional] 
**Residential** | Pointer to **bool** | The specified adress can be Residential or Official. In case if the adress is Residential, the boolean value will be &#39;true&#39;, else it will take &#39;false&#39;. | [optional] 
**CountryCode** | Pointer to **string** | ISO-2 characters country code | [optional] 
**Phone** | Pointer to **string** | Phone number | [optional] 
**Company** | Pointer to **string** | Company name of pickup address | [optional] 
**Email** | Pointer to **string** | Email. A Confirmation email will be send to this email id when pickup gets scheduled successfully. | [optional] 

## Methods

### NewSchedulePickupDHLEXPResponsePickupAddress

`func NewSchedulePickupDHLEXPResponsePickupAddress() *SchedulePickupDHLEXPResponsePickupAddress`

NewSchedulePickupDHLEXPResponsePickupAddress instantiates a new SchedulePickupDHLEXPResponsePickupAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupDHLEXPResponsePickupAddressWithDefaults

`func NewSchedulePickupDHLEXPResponsePickupAddressWithDefaults() *SchedulePickupDHLEXPResponsePickupAddress`

NewSchedulePickupDHLEXPResponsePickupAddressWithDefaults instantiates a new SchedulePickupDHLEXPResponsePickupAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchedulePickupDHLEXPResponsePickupAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchedulePickupDHLEXPResponsePickupAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *SchedulePickupDHLEXPResponsePickupAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *SchedulePickupDHLEXPResponsePickupAddress) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetCityTown

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetCityTown() string`

GetCityTown returns the CityTown field if non-nil, zero value otherwise.

### GetCityTownOk

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetCityTownOk() (*string, bool)`

GetCityTownOk returns a tuple with the CityTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTown

`func (o *SchedulePickupDHLEXPResponsePickupAddress) SetCityTown(v string)`

SetCityTown sets CityTown field to given value.

### HasCityTown

`func (o *SchedulePickupDHLEXPResponsePickupAddress) HasCityTown() bool`

HasCityTown returns a boolean if a field has been set.

### GetStateProvince

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *SchedulePickupDHLEXPResponsePickupAddress) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *SchedulePickupDHLEXPResponsePickupAddress) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.

### GetPostalCode

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *SchedulePickupDHLEXPResponsePickupAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *SchedulePickupDHLEXPResponsePickupAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetResidential

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetResidential() bool`

GetResidential returns the Residential field if non-nil, zero value otherwise.

### GetResidentialOk

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetResidentialOk() (*bool, bool)`

GetResidentialOk returns a tuple with the Residential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidential

`func (o *SchedulePickupDHLEXPResponsePickupAddress) SetResidential(v bool)`

SetResidential sets Residential field to given value.

### HasResidential

`func (o *SchedulePickupDHLEXPResponsePickupAddress) HasResidential() bool`

HasResidential returns a boolean if a field has been set.

### GetCountryCode

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SchedulePickupDHLEXPResponsePickupAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SchedulePickupDHLEXPResponsePickupAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhone

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SchedulePickupDHLEXPResponsePickupAddress) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SchedulePickupDHLEXPResponsePickupAddress) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCompany

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SchedulePickupDHLEXPResponsePickupAddress) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *SchedulePickupDHLEXPResponsePickupAddress) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEmail

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SchedulePickupDHLEXPResponsePickupAddress) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SchedulePickupDHLEXPResponsePickupAddress) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SchedulePickupDHLEXPResponsePickupAddress) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


