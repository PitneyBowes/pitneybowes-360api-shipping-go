# SpecialServicesServicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** | The unique identifier given to the carrier specific service. | [optional] 
**BrandedName** | Pointer to **string** | The branded name of the service | [optional] 
**ParcelTypeRules** | Pointer to [**[]SpecialServicesServicesInnerParcelTypeRulesInner**](SpecialServicesServicesInnerParcelTypeRulesInner.md) | It displays special services for specific parcel type | [optional] 

## Methods

### NewSpecialServicesServicesInner

`func NewSpecialServicesServicesInner() *SpecialServicesServicesInner`

NewSpecialServicesServicesInner instantiates a new SpecialServicesServicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecialServicesServicesInnerWithDefaults

`func NewSpecialServicesServicesInnerWithDefaults() *SpecialServicesServicesInner`

NewSpecialServicesServicesInnerWithDefaults instantiates a new SpecialServicesServicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *SpecialServicesServicesInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SpecialServicesServicesInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SpecialServicesServicesInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SpecialServicesServicesInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetBrandedName

`func (o *SpecialServicesServicesInner) GetBrandedName() string`

GetBrandedName returns the BrandedName field if non-nil, zero value otherwise.

### GetBrandedNameOk

`func (o *SpecialServicesServicesInner) GetBrandedNameOk() (*string, bool)`

GetBrandedNameOk returns a tuple with the BrandedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedName

`func (o *SpecialServicesServicesInner) SetBrandedName(v string)`

SetBrandedName sets BrandedName field to given value.

### HasBrandedName

`func (o *SpecialServicesServicesInner) HasBrandedName() bool`

HasBrandedName returns a boolean if a field has been set.

### GetParcelTypeRules

`func (o *SpecialServicesServicesInner) GetParcelTypeRules() []SpecialServicesServicesInnerParcelTypeRulesInner`

GetParcelTypeRules returns the ParcelTypeRules field if non-nil, zero value otherwise.

### GetParcelTypeRulesOk

`func (o *SpecialServicesServicesInner) GetParcelTypeRulesOk() (*[]SpecialServicesServicesInnerParcelTypeRulesInner, bool)`

GetParcelTypeRulesOk returns a tuple with the ParcelTypeRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTypeRules

`func (o *SpecialServicesServicesInner) SetParcelTypeRules(v []SpecialServicesServicesInnerParcelTypeRulesInner)`

SetParcelTypeRules sets ParcelTypeRules field to given value.

### HasParcelTypeRules

`func (o *SpecialServicesServicesInner) HasParcelTypeRules() bool`

HasParcelTypeRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


