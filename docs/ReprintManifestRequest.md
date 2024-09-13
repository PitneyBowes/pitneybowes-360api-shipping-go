# ReprintManifestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierAccountId** | **string** | The unique identifier associated with the Carrier account which is used while Manifest process. | 
**ManifestID** | **string** | A unique identifier associated to the created Manifest for a particular Carrier. | 

## Methods

### NewReprintManifestRequest

`func NewReprintManifestRequest(carrierAccountId string, manifestID string, ) *ReprintManifestRequest`

NewReprintManifestRequest instantiates a new ReprintManifestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReprintManifestRequestWithDefaults

`func NewReprintManifestRequestWithDefaults() *ReprintManifestRequest`

NewReprintManifestRequestWithDefaults instantiates a new ReprintManifestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierAccountId

`func (o *ReprintManifestRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *ReprintManifestRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *ReprintManifestRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetManifestID

`func (o *ReprintManifestRequest) GetManifestID() string`

GetManifestID returns the ManifestID field if non-nil, zero value otherwise.

### GetManifestIDOk

`func (o *ReprintManifestRequest) GetManifestIDOk() (*string, bool)`

GetManifestIDOk returns a tuple with the ManifestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestID

`func (o *ReprintManifestRequest) SetManifestID(v string)`

SetManifestID sets ManifestID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


