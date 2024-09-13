# ParcelV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | Pointer to **float32** | Length is always the greatest of the three dimensions. The other two dimensions are used in the calculation of the girth. | [optional] 
**Width** | Pointer to **float32** | There is no strict rule as to which element is the width or the height, but the width is the second greatest dimension of a parcel by convention. | [optional] 
**Height** | Pointer to **float32** | By convention, the height is the smallest dimension of the parcel. | [optional] 
**DimUnit** | Pointer to **string** | DimUnit is a standard for measuring the physical quantities of specified dimension parameters. | [optional] 
**WeightUnit** | **string** | WeightUnit is a standard for measuring the physical quantities of specified weight. &lt;br /&gt; For USPS shipments, this is OZ. | 
**Weight** | Pointer to **float32** | Weight measures the heaviness of an object (how heavy an object is) . | [optional] 
**PackageValue** | Pointer to **float32** | Indicates value of the package. | [optional] 

## Methods

### NewParcelV2Response

`func NewParcelV2Response(weightUnit string, ) *ParcelV2Response`

NewParcelV2Response instantiates a new ParcelV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelV2ResponseWithDefaults

`func NewParcelV2ResponseWithDefaults() *ParcelV2Response`

NewParcelV2ResponseWithDefaults instantiates a new ParcelV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *ParcelV2Response) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ParcelV2Response) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ParcelV2Response) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ParcelV2Response) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWidth

`func (o *ParcelV2Response) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ParcelV2Response) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ParcelV2Response) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ParcelV2Response) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ParcelV2Response) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ParcelV2Response) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ParcelV2Response) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ParcelV2Response) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetDimUnit

`func (o *ParcelV2Response) GetDimUnit() string`

GetDimUnit returns the DimUnit field if non-nil, zero value otherwise.

### GetDimUnitOk

`func (o *ParcelV2Response) GetDimUnitOk() (*string, bool)`

GetDimUnitOk returns a tuple with the DimUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimUnit

`func (o *ParcelV2Response) SetDimUnit(v string)`

SetDimUnit sets DimUnit field to given value.

### HasDimUnit

`func (o *ParcelV2Response) HasDimUnit() bool`

HasDimUnit returns a boolean if a field has been set.

### GetWeightUnit

`func (o *ParcelV2Response) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ParcelV2Response) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ParcelV2Response) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.


### GetWeight

`func (o *ParcelV2Response) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ParcelV2Response) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ParcelV2Response) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ParcelV2Response) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPackageValue

`func (o *ParcelV2Response) GetPackageValue() float32`

GetPackageValue returns the PackageValue field if non-nil, zero value otherwise.

### GetPackageValueOk

`func (o *ParcelV2Response) GetPackageValueOk() (*float32, bool)`

GetPackageValueOk returns a tuple with the PackageValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageValue

`func (o *ParcelV2Response) SetPackageValue(v float32)`

SetPackageValue sets PackageValue field to given value.

### HasPackageValue

`func (o *ParcelV2Response) HasPackageValue() bool`

HasPackageValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


