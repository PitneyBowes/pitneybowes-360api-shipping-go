# ServicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandedName** | Pointer to **string** | The branded name of service. | [optional] 
**ServiceId** | Pointer to **string** | The unique identifier given to the carrier specific service. | [optional] 

## Methods

### NewServicesInner

`func NewServicesInner() *ServicesInner`

NewServicesInner instantiates a new ServicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesInnerWithDefaults

`func NewServicesInnerWithDefaults() *ServicesInner`

NewServicesInnerWithDefaults instantiates a new ServicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandedName

`func (o *ServicesInner) GetBrandedName() string`

GetBrandedName returns the BrandedName field if non-nil, zero value otherwise.

### GetBrandedNameOk

`func (o *ServicesInner) GetBrandedNameOk() (*string, bool)`

GetBrandedNameOk returns a tuple with the BrandedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedName

`func (o *ServicesInner) SetBrandedName(v string)`

SetBrandedName sets BrandedName field to given value.

### HasBrandedName

`func (o *ServicesInner) HasBrandedName() bool`

HasBrandedName returns a boolean if a field has been set.

### GetServiceId

`func (o *ServicesInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServicesInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServicesInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServicesInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


