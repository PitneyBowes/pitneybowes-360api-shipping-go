# SpecialServicesServicesInnerParcelTypeRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandedName** | Pointer to **string** | The branded name of Parcel Type | [optional] 
**ParcelType** | Pointer to **string** | Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection. Here, it reflects as per the defined ParcelType Rules. ParcelType have categories like Package, Envelopes, Paks, Boxes, Tube, etc.  | [optional] 
**Trackable** | Pointer to **string** | Whether this parcel type is trackable. Valid Values are: TRACKABLE, NON_TRACKABLE, REQUIRES_TRACKABLE_SPECIAL_SERVICE | [optional] 
**SuggestedTrackableSpecialService** | Pointer to **string** | If trackable is set to REQUIRES_TRACKABLE_SPECIAL_SERVICE, this is a free or low-cost special service that allows the shipper to track the shipment. | [optional] 
**SpecialServiceRules** | Pointer to [**[]SpecialServicesServicesInnerParcelTypeRulesInnerSpecialServiceRulesInner**](SpecialServicesServicesInnerParcelTypeRulesInnerSpecialServiceRulesInner.md) | It displays all the available special services, their details and prerequisites and/or incompatible details with other special services | [optional] 

## Methods

### NewSpecialServicesServicesInnerParcelTypeRulesInner

`func NewSpecialServicesServicesInnerParcelTypeRulesInner() *SpecialServicesServicesInnerParcelTypeRulesInner`

NewSpecialServicesServicesInnerParcelTypeRulesInner instantiates a new SpecialServicesServicesInnerParcelTypeRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecialServicesServicesInnerParcelTypeRulesInnerWithDefaults

`func NewSpecialServicesServicesInnerParcelTypeRulesInnerWithDefaults() *SpecialServicesServicesInnerParcelTypeRulesInner`

NewSpecialServicesServicesInnerParcelTypeRulesInnerWithDefaults instantiates a new SpecialServicesServicesInnerParcelTypeRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandedName

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) GetBrandedName() string`

GetBrandedName returns the BrandedName field if non-nil, zero value otherwise.

### GetBrandedNameOk

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) GetBrandedNameOk() (*string, bool)`

GetBrandedNameOk returns a tuple with the BrandedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedName

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) SetBrandedName(v string)`

SetBrandedName sets BrandedName field to given value.

### HasBrandedName

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) HasBrandedName() bool`

HasBrandedName returns a boolean if a field has been set.

### GetParcelType

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetTrackable

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) GetTrackable() string`

GetTrackable returns the Trackable field if non-nil, zero value otherwise.

### GetTrackableOk

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) GetTrackableOk() (*string, bool)`

GetTrackableOk returns a tuple with the Trackable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackable

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) SetTrackable(v string)`

SetTrackable sets Trackable field to given value.

### HasTrackable

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) HasTrackable() bool`

HasTrackable returns a boolean if a field has been set.

### GetSuggestedTrackableSpecialService

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) GetSuggestedTrackableSpecialService() string`

GetSuggestedTrackableSpecialService returns the SuggestedTrackableSpecialService field if non-nil, zero value otherwise.

### GetSuggestedTrackableSpecialServiceOk

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) GetSuggestedTrackableSpecialServiceOk() (*string, bool)`

GetSuggestedTrackableSpecialServiceOk returns a tuple with the SuggestedTrackableSpecialService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedTrackableSpecialService

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) SetSuggestedTrackableSpecialService(v string)`

SetSuggestedTrackableSpecialService sets SuggestedTrackableSpecialService field to given value.

### HasSuggestedTrackableSpecialService

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) HasSuggestedTrackableSpecialService() bool`

HasSuggestedTrackableSpecialService returns a boolean if a field has been set.

### GetSpecialServiceRules

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) GetSpecialServiceRules() []SpecialServicesServicesInnerParcelTypeRulesInnerSpecialServiceRulesInner`

GetSpecialServiceRules returns the SpecialServiceRules field if non-nil, zero value otherwise.

### GetSpecialServiceRulesOk

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) GetSpecialServiceRulesOk() (*[]SpecialServicesServicesInnerParcelTypeRulesInnerSpecialServiceRulesInner, bool)`

GetSpecialServiceRulesOk returns a tuple with the SpecialServiceRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServiceRules

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) SetSpecialServiceRules(v []SpecialServicesServicesInnerParcelTypeRulesInnerSpecialServiceRulesInner)`

SetSpecialServiceRules sets SpecialServiceRules field to given value.

### HasSpecialServiceRules

`func (o *SpecialServicesServicesInnerParcelTypeRulesInner) HasSpecialServiceRules() bool`

HasSpecialServiceRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


