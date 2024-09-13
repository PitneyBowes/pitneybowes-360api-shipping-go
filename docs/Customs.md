# Customs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomsInfo** | [**CustomsInfo**](CustomsInfo.md) |  | 
**CustomsItems** | [**[]CustomsItem**](CustomsItem.md) | This operation asks the information for each commodity for Customs clearance. | 

## Methods

### NewCustoms

`func NewCustoms(customsInfo CustomsInfo, customsItems []CustomsItem, ) *Customs`

NewCustoms instantiates a new Customs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomsWithDefaults

`func NewCustomsWithDefaults() *Customs`

NewCustomsWithDefaults instantiates a new Customs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomsInfo

`func (o *Customs) GetCustomsInfo() CustomsInfo`

GetCustomsInfo returns the CustomsInfo field if non-nil, zero value otherwise.

### GetCustomsInfoOk

`func (o *Customs) GetCustomsInfoOk() (*CustomsInfo, bool)`

GetCustomsInfoOk returns a tuple with the CustomsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsInfo

`func (o *Customs) SetCustomsInfo(v CustomsInfo)`

SetCustomsInfo sets CustomsInfo field to given value.


### GetCustomsItems

`func (o *Customs) GetCustomsItems() []CustomsItem`

GetCustomsItems returns the CustomsItems field if non-nil, zero value otherwise.

### GetCustomsItemsOk

`func (o *Customs) GetCustomsItemsOk() (*[]CustomsItem, bool)`

GetCustomsItemsOk returns a tuple with the CustomsItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsItems

`func (o *Customs) SetCustomsItems(v []CustomsItem)`

SetCustomsItems sets CustomsItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


