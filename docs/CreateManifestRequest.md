# CreateManifestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierAccountId** | **string** | A unique identifier associated with the specific carrier account, which is used in the Manifest process. | 
**FromAddress** | [**CreateManifestRequestFromAddress**](CreateManifestRequestFromAddress.md) |  | 
**SubmissionDate** | Pointer to **string** | The date the shipments are to be tendered to the carrier, entered as YYYY-MM-DD. | [optional] 

## Methods

### NewCreateManifestRequest

`func NewCreateManifestRequest(carrierAccountId string, fromAddress CreateManifestRequestFromAddress, ) *CreateManifestRequest`

NewCreateManifestRequest instantiates a new CreateManifestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateManifestRequestWithDefaults

`func NewCreateManifestRequestWithDefaults() *CreateManifestRequest`

NewCreateManifestRequestWithDefaults instantiates a new CreateManifestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierAccountId

`func (o *CreateManifestRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *CreateManifestRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *CreateManifestRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetFromAddress

`func (o *CreateManifestRequest) GetFromAddress() CreateManifestRequestFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *CreateManifestRequest) GetFromAddressOk() (*CreateManifestRequestFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *CreateManifestRequest) SetFromAddress(v CreateManifestRequestFromAddress)`

SetFromAddress sets FromAddress field to given value.


### GetSubmissionDate

`func (o *CreateManifestRequest) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *CreateManifestRequest) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *CreateManifestRequest) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *CreateManifestRequest) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


