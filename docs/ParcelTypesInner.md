# ParcelTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandedName** | Pointer to **string** | The branded name of parcel type | [optional] 
**Carrier** | Pointer to **string** | A unique identifier associated with the specific carrier. | [optional] 
**DimensionRules** | Pointer to [**[]ParcelTypesInnerDimensionRulesInner**](ParcelTypesInnerDimensionRulesInner.md) | This defines the maximum and minimum dimension supported for given parcel type. | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**BrandedDimensions** | Pointer to [**ParcelTypesInnerBrandedDimensions**](ParcelTypesInnerBrandedDimensions.md) |  | [optional] 
**ParcelId** | Pointer to **string** | A unique identifier associated with the Parcel type. | [optional] 
**IsBranded** | Pointer to **bool** | If the Parcel is Branded. If yees, it will take &#x60;true&#x60;, else will take &#x60;false&#x60;. | [optional] 
**ParcelType** | Pointer to **string** | This defines the type of Parcel. | [optional] 
**ServiceId** | Pointer to **string** | A unique identifier associated with the carrier based service. | [optional] 
**ServiceName** | Pointer to **string** | Name of the Carrier Service. | [optional] 
**SuggestedTrackableSpecialserviceId** | Pointer to **string** | This defines the parcel has feature to track special serviceID. | [optional] 
**WeightRules** | Pointer to [**[]ParcelTypesInnerWeightRulesInner**](ParcelTypesInnerWeightRulesInner.md) | This defines the maximum and minimum weight supported for given parcel type. | [optional] 

## Methods

### NewParcelTypesInner

`func NewParcelTypesInner() *ParcelTypesInner`

NewParcelTypesInner instantiates a new ParcelTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelTypesInnerWithDefaults

`func NewParcelTypesInnerWithDefaults() *ParcelTypesInner`

NewParcelTypesInnerWithDefaults instantiates a new ParcelTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandedName

`func (o *ParcelTypesInner) GetBrandedName() string`

GetBrandedName returns the BrandedName field if non-nil, zero value otherwise.

### GetBrandedNameOk

`func (o *ParcelTypesInner) GetBrandedNameOk() (*string, bool)`

GetBrandedNameOk returns a tuple with the BrandedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedName

`func (o *ParcelTypesInner) SetBrandedName(v string)`

SetBrandedName sets BrandedName field to given value.

### HasBrandedName

`func (o *ParcelTypesInner) HasBrandedName() bool`

HasBrandedName returns a boolean if a field has been set.

### GetCarrier

`func (o *ParcelTypesInner) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *ParcelTypesInner) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *ParcelTypesInner) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *ParcelTypesInner) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetDimensionRules

`func (o *ParcelTypesInner) GetDimensionRules() []ParcelTypesInnerDimensionRulesInner`

GetDimensionRules returns the DimensionRules field if non-nil, zero value otherwise.

### GetDimensionRulesOk

`func (o *ParcelTypesInner) GetDimensionRulesOk() (*[]ParcelTypesInnerDimensionRulesInner, bool)`

GetDimensionRulesOk returns a tuple with the DimensionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionRules

`func (o *ParcelTypesInner) SetDimensionRules(v []ParcelTypesInnerDimensionRulesInner)`

SetDimensionRules sets DimensionRules field to given value.

### HasDimensionRules

`func (o *ParcelTypesInner) HasDimensionRules() bool`

HasDimensionRules returns a boolean if a field has been set.

### GetGroupName

`func (o *ParcelTypesInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ParcelTypesInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ParcelTypesInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ParcelTypesInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetBrandedDimensions

`func (o *ParcelTypesInner) GetBrandedDimensions() ParcelTypesInnerBrandedDimensions`

GetBrandedDimensions returns the BrandedDimensions field if non-nil, zero value otherwise.

### GetBrandedDimensionsOk

`func (o *ParcelTypesInner) GetBrandedDimensionsOk() (*ParcelTypesInnerBrandedDimensions, bool)`

GetBrandedDimensionsOk returns a tuple with the BrandedDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedDimensions

`func (o *ParcelTypesInner) SetBrandedDimensions(v ParcelTypesInnerBrandedDimensions)`

SetBrandedDimensions sets BrandedDimensions field to given value.

### HasBrandedDimensions

`func (o *ParcelTypesInner) HasBrandedDimensions() bool`

HasBrandedDimensions returns a boolean if a field has been set.

### GetParcelId

`func (o *ParcelTypesInner) GetParcelId() string`

GetParcelId returns the ParcelId field if non-nil, zero value otherwise.

### GetParcelIdOk

`func (o *ParcelTypesInner) GetParcelIdOk() (*string, bool)`

GetParcelIdOk returns a tuple with the ParcelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelId

`func (o *ParcelTypesInner) SetParcelId(v string)`

SetParcelId sets ParcelId field to given value.

### HasParcelId

`func (o *ParcelTypesInner) HasParcelId() bool`

HasParcelId returns a boolean if a field has been set.

### GetIsBranded

`func (o *ParcelTypesInner) GetIsBranded() bool`

GetIsBranded returns the IsBranded field if non-nil, zero value otherwise.

### GetIsBrandedOk

`func (o *ParcelTypesInner) GetIsBrandedOk() (*bool, bool)`

GetIsBrandedOk returns a tuple with the IsBranded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBranded

`func (o *ParcelTypesInner) SetIsBranded(v bool)`

SetIsBranded sets IsBranded field to given value.

### HasIsBranded

`func (o *ParcelTypesInner) HasIsBranded() bool`

HasIsBranded returns a boolean if a field has been set.

### GetParcelType

`func (o *ParcelTypesInner) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *ParcelTypesInner) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *ParcelTypesInner) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *ParcelTypesInner) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetServiceId

`func (o *ParcelTypesInner) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ParcelTypesInner) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ParcelTypesInner) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ParcelTypesInner) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *ParcelTypesInner) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ParcelTypesInner) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ParcelTypesInner) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ParcelTypesInner) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSuggestedTrackableSpecialserviceId

`func (o *ParcelTypesInner) GetSuggestedTrackableSpecialserviceId() string`

GetSuggestedTrackableSpecialserviceId returns the SuggestedTrackableSpecialserviceId field if non-nil, zero value otherwise.

### GetSuggestedTrackableSpecialserviceIdOk

`func (o *ParcelTypesInner) GetSuggestedTrackableSpecialserviceIdOk() (*string, bool)`

GetSuggestedTrackableSpecialserviceIdOk returns a tuple with the SuggestedTrackableSpecialserviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedTrackableSpecialserviceId

`func (o *ParcelTypesInner) SetSuggestedTrackableSpecialserviceId(v string)`

SetSuggestedTrackableSpecialserviceId sets SuggestedTrackableSpecialserviceId field to given value.

### HasSuggestedTrackableSpecialserviceId

`func (o *ParcelTypesInner) HasSuggestedTrackableSpecialserviceId() bool`

HasSuggestedTrackableSpecialserviceId returns a boolean if a field has been set.

### GetWeightRules

`func (o *ParcelTypesInner) GetWeightRules() []ParcelTypesInnerWeightRulesInner`

GetWeightRules returns the WeightRules field if non-nil, zero value otherwise.

### GetWeightRulesOk

`func (o *ParcelTypesInner) GetWeightRulesOk() (*[]ParcelTypesInnerWeightRulesInner, bool)`

GetWeightRulesOk returns a tuple with the WeightRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightRules

`func (o *ParcelTypesInner) SetWeightRules(v []ParcelTypesInnerWeightRulesInner)`

SetWeightRules sets WeightRules field to given value.

### HasWeightRules

`func (o *ParcelTypesInner) HasWeightRules() bool`

HasWeightRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


