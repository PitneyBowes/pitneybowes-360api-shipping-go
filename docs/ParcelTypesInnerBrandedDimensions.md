# ParcelTypesInnerBrandedDimensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Girth** | Pointer to **float32** | This is the measure around parcel by area or circumference. | [optional] 
**Height** | Pointer to **float32** | Height is a part of Dimension objet where it helps determine a parcel’s girth. | [optional] 
**Length** | Pointer to **float32** | Length is a part of Dimension objet having highest numeric value out of three required parameters (length, width, and height) of Dimension. It helps determine a parcel’s girth. | [optional] 
**UnitOfMeasurement** | Pointer to **string** | UnitofMesurements is a standard for measuring the physical quantities of specified dimension parameters. | [optional] 
**Width** | Pointer to **float32** | Width is a part of Dimension object having lowest numeric value out of three required parameters of dimension (length, width, and height). This helps determine a parcel’s girth. | [optional] 

## Methods

### NewParcelTypesInnerBrandedDimensions

`func NewParcelTypesInnerBrandedDimensions() *ParcelTypesInnerBrandedDimensions`

NewParcelTypesInnerBrandedDimensions instantiates a new ParcelTypesInnerBrandedDimensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelTypesInnerBrandedDimensionsWithDefaults

`func NewParcelTypesInnerBrandedDimensionsWithDefaults() *ParcelTypesInnerBrandedDimensions`

NewParcelTypesInnerBrandedDimensionsWithDefaults instantiates a new ParcelTypesInnerBrandedDimensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGirth

`func (o *ParcelTypesInnerBrandedDimensions) GetGirth() float32`

GetGirth returns the Girth field if non-nil, zero value otherwise.

### GetGirthOk

`func (o *ParcelTypesInnerBrandedDimensions) GetGirthOk() (*float32, bool)`

GetGirthOk returns a tuple with the Girth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGirth

`func (o *ParcelTypesInnerBrandedDimensions) SetGirth(v float32)`

SetGirth sets Girth field to given value.

### HasGirth

`func (o *ParcelTypesInnerBrandedDimensions) HasGirth() bool`

HasGirth returns a boolean if a field has been set.

### GetHeight

`func (o *ParcelTypesInnerBrandedDimensions) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ParcelTypesInnerBrandedDimensions) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ParcelTypesInnerBrandedDimensions) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ParcelTypesInnerBrandedDimensions) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetLength

`func (o *ParcelTypesInnerBrandedDimensions) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ParcelTypesInnerBrandedDimensions) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ParcelTypesInnerBrandedDimensions) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ParcelTypesInnerBrandedDimensions) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetUnitOfMeasurement

`func (o *ParcelTypesInnerBrandedDimensions) GetUnitOfMeasurement() string`

GetUnitOfMeasurement returns the UnitOfMeasurement field if non-nil, zero value otherwise.

### GetUnitOfMeasurementOk

`func (o *ParcelTypesInnerBrandedDimensions) GetUnitOfMeasurementOk() (*string, bool)`

GetUnitOfMeasurementOk returns a tuple with the UnitOfMeasurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasurement

`func (o *ParcelTypesInnerBrandedDimensions) SetUnitOfMeasurement(v string)`

SetUnitOfMeasurement sets UnitOfMeasurement field to given value.

### HasUnitOfMeasurement

`func (o *ParcelTypesInnerBrandedDimensions) HasUnitOfMeasurement() bool`

HasUnitOfMeasurement returns a boolean if a field has been set.

### GetWidth

`func (o *ParcelTypesInnerBrandedDimensions) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ParcelTypesInnerBrandedDimensions) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ParcelTypesInnerBrandedDimensions) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ParcelTypesInnerBrandedDimensions) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


