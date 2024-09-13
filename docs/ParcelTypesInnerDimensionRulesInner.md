# ParcelTypesInnerDimensionRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxLengthPlusGirth** | Pointer to **float32** | This measures the parcel&#39;s maximum length and determine parcel’s girth. | [optional] 
**MaxParcelDimensions** | Pointer to [**ParcelTypesInnerDimensionRulesInnerMaxParcelDimensions**](ParcelTypesInnerDimensionRulesInnerMaxParcelDimensions.md) |  | [optional] 
**MinLengthPlusGirth** | Pointer to **float32** | This measures the parcel&#39;s minimum length and determine parcel’s girth. | [optional] 
**MinParcelDimensions** | Pointer to [**ParcelTypesInnerDimensionRulesInnerMinParcelDimensions**](ParcelTypesInnerDimensionRulesInnerMinParcelDimensions.md) |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**UnitOfMeasurement** | Pointer to **string** | UnitofMesurement is a standard for measuring the physical quantities of specified dimension parameters. | [optional] 

## Methods

### NewParcelTypesInnerDimensionRulesInner

`func NewParcelTypesInnerDimensionRulesInner() *ParcelTypesInnerDimensionRulesInner`

NewParcelTypesInnerDimensionRulesInner instantiates a new ParcelTypesInnerDimensionRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelTypesInnerDimensionRulesInnerWithDefaults

`func NewParcelTypesInnerDimensionRulesInnerWithDefaults() *ParcelTypesInnerDimensionRulesInner`

NewParcelTypesInnerDimensionRulesInnerWithDefaults instantiates a new ParcelTypesInnerDimensionRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxLengthPlusGirth

`func (o *ParcelTypesInnerDimensionRulesInner) GetMaxLengthPlusGirth() float32`

GetMaxLengthPlusGirth returns the MaxLengthPlusGirth field if non-nil, zero value otherwise.

### GetMaxLengthPlusGirthOk

`func (o *ParcelTypesInnerDimensionRulesInner) GetMaxLengthPlusGirthOk() (*float32, bool)`

GetMaxLengthPlusGirthOk returns a tuple with the MaxLengthPlusGirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLengthPlusGirth

`func (o *ParcelTypesInnerDimensionRulesInner) SetMaxLengthPlusGirth(v float32)`

SetMaxLengthPlusGirth sets MaxLengthPlusGirth field to given value.

### HasMaxLengthPlusGirth

`func (o *ParcelTypesInnerDimensionRulesInner) HasMaxLengthPlusGirth() bool`

HasMaxLengthPlusGirth returns a boolean if a field has been set.

### GetMaxParcelDimensions

`func (o *ParcelTypesInnerDimensionRulesInner) GetMaxParcelDimensions() ParcelTypesInnerDimensionRulesInnerMaxParcelDimensions`

GetMaxParcelDimensions returns the MaxParcelDimensions field if non-nil, zero value otherwise.

### GetMaxParcelDimensionsOk

`func (o *ParcelTypesInnerDimensionRulesInner) GetMaxParcelDimensionsOk() (*ParcelTypesInnerDimensionRulesInnerMaxParcelDimensions, bool)`

GetMaxParcelDimensionsOk returns a tuple with the MaxParcelDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParcelDimensions

`func (o *ParcelTypesInnerDimensionRulesInner) SetMaxParcelDimensions(v ParcelTypesInnerDimensionRulesInnerMaxParcelDimensions)`

SetMaxParcelDimensions sets MaxParcelDimensions field to given value.

### HasMaxParcelDimensions

`func (o *ParcelTypesInnerDimensionRulesInner) HasMaxParcelDimensions() bool`

HasMaxParcelDimensions returns a boolean if a field has been set.

### GetMinLengthPlusGirth

`func (o *ParcelTypesInnerDimensionRulesInner) GetMinLengthPlusGirth() float32`

GetMinLengthPlusGirth returns the MinLengthPlusGirth field if non-nil, zero value otherwise.

### GetMinLengthPlusGirthOk

`func (o *ParcelTypesInnerDimensionRulesInner) GetMinLengthPlusGirthOk() (*float32, bool)`

GetMinLengthPlusGirthOk returns a tuple with the MinLengthPlusGirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLengthPlusGirth

`func (o *ParcelTypesInnerDimensionRulesInner) SetMinLengthPlusGirth(v float32)`

SetMinLengthPlusGirth sets MinLengthPlusGirth field to given value.

### HasMinLengthPlusGirth

`func (o *ParcelTypesInnerDimensionRulesInner) HasMinLengthPlusGirth() bool`

HasMinLengthPlusGirth returns a boolean if a field has been set.

### GetMinParcelDimensions

`func (o *ParcelTypesInnerDimensionRulesInner) GetMinParcelDimensions() ParcelTypesInnerDimensionRulesInnerMinParcelDimensions`

GetMinParcelDimensions returns the MinParcelDimensions field if non-nil, zero value otherwise.

### GetMinParcelDimensionsOk

`func (o *ParcelTypesInnerDimensionRulesInner) GetMinParcelDimensionsOk() (*ParcelTypesInnerDimensionRulesInnerMinParcelDimensions, bool)`

GetMinParcelDimensionsOk returns a tuple with the MinParcelDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinParcelDimensions

`func (o *ParcelTypesInnerDimensionRulesInner) SetMinParcelDimensions(v ParcelTypesInnerDimensionRulesInnerMinParcelDimensions)`

SetMinParcelDimensions sets MinParcelDimensions field to given value.

### HasMinParcelDimensions

`func (o *ParcelTypesInnerDimensionRulesInner) HasMinParcelDimensions() bool`

HasMinParcelDimensions returns a boolean if a field has been set.

### GetRequired

`func (o *ParcelTypesInnerDimensionRulesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ParcelTypesInnerDimensionRulesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ParcelTypesInnerDimensionRulesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ParcelTypesInnerDimensionRulesInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetUnitOfMeasurement

`func (o *ParcelTypesInnerDimensionRulesInner) GetUnitOfMeasurement() string`

GetUnitOfMeasurement returns the UnitOfMeasurement field if non-nil, zero value otherwise.

### GetUnitOfMeasurementOk

`func (o *ParcelTypesInnerDimensionRulesInner) GetUnitOfMeasurementOk() (*string, bool)`

GetUnitOfMeasurementOk returns a tuple with the UnitOfMeasurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasurement

`func (o *ParcelTypesInnerDimensionRulesInner) SetUnitOfMeasurement(v string)`

SetUnitOfMeasurement sets UnitOfMeasurement field to given value.

### HasUnitOfMeasurement

`func (o *ParcelTypesInnerDimensionRulesInner) HasUnitOfMeasurement() bool`

HasUnitOfMeasurement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


