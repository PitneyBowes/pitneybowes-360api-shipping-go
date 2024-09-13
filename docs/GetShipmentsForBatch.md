# GetShipmentsForBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetShipmentsForBatchDataInner**](GetShipmentsForBatchDataInner.md) | It displays all the shipment details based on the paramter selected in the request body. | [optional] 
**PageInfo** | Pointer to [**GetShipmentsForBatchPageInfo**](GetShipmentsForBatchPageInfo.md) |  | [optional] 

## Methods

### NewGetShipmentsForBatch

`func NewGetShipmentsForBatch() *GetShipmentsForBatch`

NewGetShipmentsForBatch instantiates a new GetShipmentsForBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetShipmentsForBatchWithDefaults

`func NewGetShipmentsForBatchWithDefaults() *GetShipmentsForBatch`

NewGetShipmentsForBatchWithDefaults instantiates a new GetShipmentsForBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetShipmentsForBatch) GetData() []GetShipmentsForBatchDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetShipmentsForBatch) GetDataOk() (*[]GetShipmentsForBatchDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetShipmentsForBatch) SetData(v []GetShipmentsForBatchDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetShipmentsForBatch) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPageInfo

`func (o *GetShipmentsForBatch) GetPageInfo() GetShipmentsForBatchPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *GetShipmentsForBatch) GetPageInfoOk() (*GetShipmentsForBatchPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *GetShipmentsForBatch) SetPageInfo(v GetShipmentsForBatchPageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *GetShipmentsForBatch) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


