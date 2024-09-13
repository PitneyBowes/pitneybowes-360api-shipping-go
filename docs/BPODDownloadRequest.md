# BPODDownloadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentIds** | Pointer to **[]string** | &gt;- This shows the list of Shipment IDs. If this is present, then user does not need to provide &#x60;Date Range&#x60; filter. Else &#x60;startDate&#x60; and &#x60;endDate&#x60; need to be passed in the Query Parameters. | [optional] 

## Methods

### NewBPODDownloadRequest

`func NewBPODDownloadRequest() *BPODDownloadRequest`

NewBPODDownloadRequest instantiates a new BPODDownloadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBPODDownloadRequestWithDefaults

`func NewBPODDownloadRequestWithDefaults() *BPODDownloadRequest`

NewBPODDownloadRequestWithDefaults instantiates a new BPODDownloadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipmentIds

`func (o *BPODDownloadRequest) GetShipmentIds() []string`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *BPODDownloadRequest) GetShipmentIdsOk() (*[]string, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *BPODDownloadRequest) SetShipmentIds(v []string)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *BPODDownloadRequest) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


