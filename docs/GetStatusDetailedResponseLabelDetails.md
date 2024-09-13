# GetStatusDetailedResponseLabelDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layout** | Pointer to [**BulkShipmentResponseERRLabelDetailsLayout**](BulkShipmentResponseERRLabelDetailsLayout.md) |  | [optional] 
**Results** | Pointer to [**[]GetStatusDetailedResponseLabelDetailsResultsInner**](GetStatusDetailedResponseLabelDetailsResultsInner.md) |  This indicates the results of the label generation. | [optional] 

## Methods

### NewGetStatusDetailedResponseLabelDetails

`func NewGetStatusDetailedResponseLabelDetails() *GetStatusDetailedResponseLabelDetails`

NewGetStatusDetailedResponseLabelDetails instantiates a new GetStatusDetailedResponseLabelDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatusDetailedResponseLabelDetailsWithDefaults

`func NewGetStatusDetailedResponseLabelDetailsWithDefaults() *GetStatusDetailedResponseLabelDetails`

NewGetStatusDetailedResponseLabelDetailsWithDefaults instantiates a new GetStatusDetailedResponseLabelDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayout

`func (o *GetStatusDetailedResponseLabelDetails) GetLayout() BulkShipmentResponseERRLabelDetailsLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *GetStatusDetailedResponseLabelDetails) GetLayoutOk() (*BulkShipmentResponseERRLabelDetailsLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *GetStatusDetailedResponseLabelDetails) SetLayout(v BulkShipmentResponseERRLabelDetailsLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *GetStatusDetailedResponseLabelDetails) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetResults

`func (o *GetStatusDetailedResponseLabelDetails) GetResults() []GetStatusDetailedResponseLabelDetailsResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetStatusDetailedResponseLabelDetails) GetResultsOk() (*[]GetStatusDetailedResponseLabelDetailsResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetStatusDetailedResponseLabelDetails) SetResults(v []GetStatusDetailedResponseLabelDetailsResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetStatusDetailedResponseLabelDetails) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


