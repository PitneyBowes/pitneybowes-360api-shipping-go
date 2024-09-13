# SignatureFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentId** | Pointer to **string** | Indicates the shipment identifier, a unique identifier for an individual shipment. | [optional] 
**SignatureFileURL** | Pointer to **string** | A link or URL that consists of Proof of Delivery (POD) - Recipient&#39;s Signature image file. Clicking the link will download the POD Signature. | [optional] 

## Methods

### NewSignatureFileResponse

`func NewSignatureFileResponse() *SignatureFileResponse`

NewSignatureFileResponse instantiates a new SignatureFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureFileResponseWithDefaults

`func NewSignatureFileResponseWithDefaults() *SignatureFileResponse`

NewSignatureFileResponseWithDefaults instantiates a new SignatureFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipmentId

`func (o *SignatureFileResponse) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *SignatureFileResponse) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *SignatureFileResponse) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *SignatureFileResponse) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetSignatureFileURL

`func (o *SignatureFileResponse) GetSignatureFileURL() string`

GetSignatureFileURL returns the SignatureFileURL field if non-nil, zero value otherwise.

### GetSignatureFileURLOk

`func (o *SignatureFileResponse) GetSignatureFileURLOk() (*string, bool)`

GetSignatureFileURLOk returns a tuple with the SignatureFileURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureFileURL

`func (o *SignatureFileResponse) SetSignatureFileURL(v string)`

SetSignatureFileURL sets SignatureFileURL field to given value.

### HasSignatureFileURL

`func (o *SignatureFileResponse) HasSignatureFileURL() bool`

HasSignatureFileURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


