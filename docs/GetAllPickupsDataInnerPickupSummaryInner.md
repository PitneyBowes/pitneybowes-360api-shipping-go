# GetAllPickupsDataInnerPickupSummaryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** | The service id | [optional] 
**PackageCount** | Pointer to **float32** | The total number of packages | [optional] 
**TotalWeight** | Pointer to **float32** | The total weight of the packages | [optional] 
**WeightUnit** | Pointer to **string** | Weight Unit, supported values are &#x60;OZ&#x60; and &#x60;GM&#x60; | [optional] 
**ParcelType** | Pointer to **string** | It indicates the parcel type, applicable values are- &#x60;PKG&#x60;,&#x60;LTR&#x60;,&#x60;TUBE&#x60;,&#x60;PACK&#x60;,&#x60;BOX&#x60;,&#x60;25KG&#x60;,&#x60;10KG&#x60;,&#x60;SMALL_EXP_BOX&#x60;,&#x60;MED_EXP_BOX&#x60;,&#x60;LG_EXP_BOX&#x60; | [optional] 
**ToAddressCountryCode** | Pointer to **string** | It indicates the 2 characters- ISO country code of recipient of the shipment. | [optional] 
**CurrencyCode** | Pointer to **string** | Currency code | [optional] 

## Methods

### NewGetAllPickupsDataInnerPickupSummaryInner

`func NewGetAllPickupsDataInnerPickupSummaryInner() *GetAllPickupsDataInnerPickupSummaryInner`

NewGetAllPickupsDataInnerPickupSummaryInner instantiates a new GetAllPickupsDataInnerPickupSummaryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllPickupsDataInnerPickupSummaryInnerWithDefaults

`func NewGetAllPickupsDataInnerPickupSummaryInnerWithDefaults() *GetAllPickupsDataInnerPickupSummaryInner`

NewGetAllPickupsDataInnerPickupSummaryInnerWithDefaults instantiates a new GetAllPickupsDataInnerPickupSummaryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GetAllPickupsDataInnerPickupSummaryInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *GetAllPickupsDataInnerPickupSummaryInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetPackageCount

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetPackageCount() float32`

GetPackageCount returns the PackageCount field if non-nil, zero value otherwise.

### GetPackageCountOk

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetPackageCountOk() (*float32, bool)`

GetPackageCountOk returns a tuple with the PackageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCount

`func (o *GetAllPickupsDataInnerPickupSummaryInner) SetPackageCount(v float32)`

SetPackageCount sets PackageCount field to given value.

### HasPackageCount

`func (o *GetAllPickupsDataInnerPickupSummaryInner) HasPackageCount() bool`

HasPackageCount returns a boolean if a field has been set.

### GetTotalWeight

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetTotalWeight() float32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetTotalWeightOk() (*float32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *GetAllPickupsDataInnerPickupSummaryInner) SetTotalWeight(v float32)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *GetAllPickupsDataInnerPickupSummaryInner) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *GetAllPickupsDataInnerPickupSummaryInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *GetAllPickupsDataInnerPickupSummaryInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetParcelType

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *GetAllPickupsDataInnerPickupSummaryInner) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *GetAllPickupsDataInnerPickupSummaryInner) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetToAddressCountryCode

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetToAddressCountryCode() string`

GetToAddressCountryCode returns the ToAddressCountryCode field if non-nil, zero value otherwise.

### GetToAddressCountryCodeOk

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetToAddressCountryCodeOk() (*string, bool)`

GetToAddressCountryCodeOk returns a tuple with the ToAddressCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddressCountryCode

`func (o *GetAllPickupsDataInnerPickupSummaryInner) SetToAddressCountryCode(v string)`

SetToAddressCountryCode sets ToAddressCountryCode field to given value.

### HasToAddressCountryCode

`func (o *GetAllPickupsDataInnerPickupSummaryInner) HasToAddressCountryCode() bool`

HasToAddressCountryCode returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GetAllPickupsDataInnerPickupSummaryInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GetAllPickupsDataInnerPickupSummaryInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GetAllPickupsDataInnerPickupSummaryInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


