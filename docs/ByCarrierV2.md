# ByCarrierV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierAccountId** | **string** | This is a unique identifier associated with the specific sub-carrier account, which must be valid.&lt;br /&gt; This is used in the shipment creation (if this value is defined, Carrier properties will be skipped). | 
**Carrier** | **string** | A unique identifier associated with the specific carrier, i.e. CarrierID, which must be valid. | 
**Service** | **string** | Indicates a unique identifier associated with the carrier specific service, which is ServiceID, which must be valid. | 

## Methods

### NewByCarrierV2

`func NewByCarrierV2(carrierAccountId string, carrier string, service string, ) *ByCarrierV2`

NewByCarrierV2 instantiates a new ByCarrierV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByCarrierV2WithDefaults

`func NewByCarrierV2WithDefaults() *ByCarrierV2`

NewByCarrierV2WithDefaults instantiates a new ByCarrierV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierAccountId

`func (o *ByCarrierV2) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *ByCarrierV2) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *ByCarrierV2) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetCarrier

`func (o *ByCarrierV2) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *ByCarrierV2) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *ByCarrierV2) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.


### GetService

`func (o *ByCarrierV2) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ByCarrierV2) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ByCarrierV2) SetService(v string)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


