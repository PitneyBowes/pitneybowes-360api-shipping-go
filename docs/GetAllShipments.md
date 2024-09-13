# GetAllShipments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**GetAllPickupsPageInfo**](GetAllPickupsPageInfo.md) |  | [optional] 
**Data** | Pointer to [**[]GetAllShipmentsDataInner**](GetAllShipmentsDataInner.md) |  | [optional] 

## Methods

### NewGetAllShipments

`func NewGetAllShipments() *GetAllShipments`

NewGetAllShipments instantiates a new GetAllShipments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllShipmentsWithDefaults

`func NewGetAllShipmentsWithDefaults() *GetAllShipments`

NewGetAllShipmentsWithDefaults instantiates a new GetAllShipments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *GetAllShipments) GetPageInfo() GetAllPickupsPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *GetAllShipments) GetPageInfoOk() (*GetAllPickupsPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *GetAllShipments) SetPageInfo(v GetAllPickupsPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *GetAllShipments) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetData

`func (o *GetAllShipments) GetData() []GetAllShipmentsDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllShipments) GetDataOk() (*[]GetAllShipmentsDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllShipments) SetData(v []GetAllShipmentsDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllShipments) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


