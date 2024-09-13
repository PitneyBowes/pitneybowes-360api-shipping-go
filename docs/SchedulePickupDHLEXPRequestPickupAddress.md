# SchedulePickupDHLEXPRequestPickupAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the person, It should not contain special character or numeric value | 
**AddressLine1** | **string** | Address line1 of the pickup address | 
**CityTown** | **string** | City of the pickup address | 
**StateProvince** | **string** | State province of the pickup address | 
**PostalCode** | **string** | Postal Code of the pickup address | 
**CountryCode** | **string** | ISO-2 characters country code | 
**Phone** | **string** | Phone number | 
**Company** | Pointer to **string** | Company name of pickup address | [optional] 
**Email** | Pointer to **string** | Email. A Confirmation email will be send to this email id when pickup gets scheduled successfully. | [optional] 

## Methods

### NewSchedulePickupDHLEXPRequestPickupAddress

`func NewSchedulePickupDHLEXPRequestPickupAddress(name string, addressLine1 string, cityTown string, stateProvince string, postalCode string, countryCode string, phone string, ) *SchedulePickupDHLEXPRequestPickupAddress`

NewSchedulePickupDHLEXPRequestPickupAddress instantiates a new SchedulePickupDHLEXPRequestPickupAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupDHLEXPRequestPickupAddressWithDefaults

`func NewSchedulePickupDHLEXPRequestPickupAddressWithDefaults() *SchedulePickupDHLEXPRequestPickupAddress`

NewSchedulePickupDHLEXPRequestPickupAddressWithDefaults instantiates a new SchedulePickupDHLEXPRequestPickupAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchedulePickupDHLEXPRequestPickupAddress) SetName(v string)`

SetName sets Name field to given value.


### GetAddressLine1

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *SchedulePickupDHLEXPRequestPickupAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetCityTown

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetCityTown() string`

GetCityTown returns the CityTown field if non-nil, zero value otherwise.

### GetCityTownOk

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetCityTownOk() (*string, bool)`

GetCityTownOk returns a tuple with the CityTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTown

`func (o *SchedulePickupDHLEXPRequestPickupAddress) SetCityTown(v string)`

SetCityTown sets CityTown field to given value.


### GetStateProvince

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *SchedulePickupDHLEXPRequestPickupAddress) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.


### GetPostalCode

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *SchedulePickupDHLEXPRequestPickupAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCountryCode

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SchedulePickupDHLEXPRequestPickupAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPhone

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SchedulePickupDHLEXPRequestPickupAddress) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetCompany

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SchedulePickupDHLEXPRequestPickupAddress) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *SchedulePickupDHLEXPRequestPickupAddress) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEmail

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SchedulePickupDHLEXPRequestPickupAddress) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SchedulePickupDHLEXPRequestPickupAddress) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SchedulePickupDHLEXPRequestPickupAddress) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


