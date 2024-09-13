# GetAllPickups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**GetAllPickupsPageInfo**](GetAllPickupsPageInfo.md) |  | [optional] 
**Data** | Pointer to [**[]GetAllPickupsDataInner**](GetAllPickupsDataInner.md) |  | [optional] 

## Methods

### NewGetAllPickups

`func NewGetAllPickups() *GetAllPickups`

NewGetAllPickups instantiates a new GetAllPickups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllPickupsWithDefaults

`func NewGetAllPickupsWithDefaults() *GetAllPickups`

NewGetAllPickupsWithDefaults instantiates a new GetAllPickups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *GetAllPickups) GetPageInfo() GetAllPickupsPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *GetAllPickups) GetPageInfoOk() (*GetAllPickupsPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *GetAllPickups) SetPageInfo(v GetAllPickupsPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *GetAllPickups) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetData

`func (o *GetAllPickups) GetData() []GetAllPickupsDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllPickups) GetDataOk() (*[]GetAllPickupsDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllPickups) SetData(v []GetAllPickupsDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllPickups) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


