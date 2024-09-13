# DocTabItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | This is a mandatory field. It will be displayed on the label | [optional] 
**Value** | Pointer to **string** | If the field is part of a request or response, the value will be picked up from there. In the case of custom fields, the user-provided value will be printed. | [optional] 
**Row** | Pointer to **int32** | Row Position of the Item. The min value is 1. | [optional] 
**Column** | Pointer to **int32** | Column Position of the Item. The min value is 1. | [optional] 

## Methods

### NewDocTabItem

`func NewDocTabItem() *DocTabItem`

NewDocTabItem instantiates a new DocTabItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocTabItemWithDefaults

`func NewDocTabItemWithDefaults() *DocTabItem`

NewDocTabItemWithDefaults instantiates a new DocTabItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *DocTabItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DocTabItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DocTabItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DocTabItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetValue

`func (o *DocTabItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DocTabItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DocTabItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DocTabItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRow

`func (o *DocTabItem) GetRow() int32`

GetRow returns the Row field if non-nil, zero value otherwise.

### GetRowOk

`func (o *DocTabItem) GetRowOk() (*int32, bool)`

GetRowOk returns a tuple with the Row field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRow

`func (o *DocTabItem) SetRow(v int32)`

SetRow sets Row field to given value.

### HasRow

`func (o *DocTabItem) HasRow() bool`

HasRow returns a boolean if a field has been set.

### GetColumn

`func (o *DocTabItem) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *DocTabItem) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *DocTabItem) SetColumn(v int32)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *DocTabItem) HasColumn() bool`

HasColumn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


