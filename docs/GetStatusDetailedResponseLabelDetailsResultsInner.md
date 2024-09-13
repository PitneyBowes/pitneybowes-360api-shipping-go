# GetStatusDetailedResponseLabelDetailsResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  This indicates URL of the shipping label in PDF format. | [optional] 
**ShipmentIdentifiers** | Pointer to [**[]GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner**](GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner.md) |  Indicates the following identifiers related to Shipment. | [optional] 

## Methods

### NewGetStatusDetailedResponseLabelDetailsResultsInner

`func NewGetStatusDetailedResponseLabelDetailsResultsInner() *GetStatusDetailedResponseLabelDetailsResultsInner`

NewGetStatusDetailedResponseLabelDetailsResultsInner instantiates a new GetStatusDetailedResponseLabelDetailsResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatusDetailedResponseLabelDetailsResultsInnerWithDefaults

`func NewGetStatusDetailedResponseLabelDetailsResultsInnerWithDefaults() *GetStatusDetailedResponseLabelDetailsResultsInner`

NewGetStatusDetailedResponseLabelDetailsResultsInnerWithDefaults instantiates a new GetStatusDetailedResponseLabelDetailsResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GetStatusDetailedResponseLabelDetailsResultsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetStatusDetailedResponseLabelDetailsResultsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetStatusDetailedResponseLabelDetailsResultsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetStatusDetailedResponseLabelDetailsResultsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetShipmentIdentifiers

`func (o *GetStatusDetailedResponseLabelDetailsResultsInner) GetShipmentIdentifiers() []GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner`

GetShipmentIdentifiers returns the ShipmentIdentifiers field if non-nil, zero value otherwise.

### GetShipmentIdentifiersOk

`func (o *GetStatusDetailedResponseLabelDetailsResultsInner) GetShipmentIdentifiersOk() (*[]GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner, bool)`

GetShipmentIdentifiersOk returns a tuple with the ShipmentIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIdentifiers

`func (o *GetStatusDetailedResponseLabelDetailsResultsInner) SetShipmentIdentifiers(v []GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner)`

SetShipmentIdentifiers sets ShipmentIdentifiers field to given value.

### HasShipmentIdentifiers

`func (o *GetStatusDetailedResponseLabelDetailsResultsInner) HasShipmentIdentifiers() bool`

HasShipmentIdentifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


