# CreateManifest200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierAccountId** | Pointer to **string** | A unique identifier associated with the Carrier account which is used while creating Manifest. | [optional] 
**CarrierName** | Pointer to **string** | Name of the Carrier. | [optional] 
**ManifestDocuments** | Pointer to [**[]ManifestCompactResponseManifestDocumentsInner**](ManifestCompactResponseManifestDocumentsInner.md) |  | [optional] 
**ManifestID** | Pointer to **string** |  | [optional] 
**ManifestTrackingNumber** | Pointer to **string** |  | [optional] 
**FromAddress** | Pointer to [**ManifestDetailedResponseFromAddress**](ManifestDetailedResponseFromAddress.md) |  | [optional] 
**ParcelTrackingNumbers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SubmissionDate** | Pointer to **string** | The date the shipments are to be tendered to the carrier, entered as YYYY-MM-DD. | [optional] 

## Methods

### NewCreateManifest200Response

`func NewCreateManifest200Response() *CreateManifest200Response`

NewCreateManifest200Response instantiates a new CreateManifest200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateManifest200ResponseWithDefaults

`func NewCreateManifest200ResponseWithDefaults() *CreateManifest200Response`

NewCreateManifest200ResponseWithDefaults instantiates a new CreateManifest200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierAccountId

`func (o *CreateManifest200Response) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *CreateManifest200Response) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *CreateManifest200Response) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *CreateManifest200Response) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetCarrierName

`func (o *CreateManifest200Response) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *CreateManifest200Response) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *CreateManifest200Response) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *CreateManifest200Response) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetManifestDocuments

`func (o *CreateManifest200Response) GetManifestDocuments() []ManifestCompactResponseManifestDocumentsInner`

GetManifestDocuments returns the ManifestDocuments field if non-nil, zero value otherwise.

### GetManifestDocumentsOk

`func (o *CreateManifest200Response) GetManifestDocumentsOk() (*[]ManifestCompactResponseManifestDocumentsInner, bool)`

GetManifestDocumentsOk returns a tuple with the ManifestDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestDocuments

`func (o *CreateManifest200Response) SetManifestDocuments(v []ManifestCompactResponseManifestDocumentsInner)`

SetManifestDocuments sets ManifestDocuments field to given value.

### HasManifestDocuments

`func (o *CreateManifest200Response) HasManifestDocuments() bool`

HasManifestDocuments returns a boolean if a field has been set.

### GetManifestID

`func (o *CreateManifest200Response) GetManifestID() string`

GetManifestID returns the ManifestID field if non-nil, zero value otherwise.

### GetManifestIDOk

`func (o *CreateManifest200Response) GetManifestIDOk() (*string, bool)`

GetManifestIDOk returns a tuple with the ManifestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestID

`func (o *CreateManifest200Response) SetManifestID(v string)`

SetManifestID sets ManifestID field to given value.

### HasManifestID

`func (o *CreateManifest200Response) HasManifestID() bool`

HasManifestID returns a boolean if a field has been set.

### GetManifestTrackingNumber

`func (o *CreateManifest200Response) GetManifestTrackingNumber() string`

GetManifestTrackingNumber returns the ManifestTrackingNumber field if non-nil, zero value otherwise.

### GetManifestTrackingNumberOk

`func (o *CreateManifest200Response) GetManifestTrackingNumberOk() (*string, bool)`

GetManifestTrackingNumberOk returns a tuple with the ManifestTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestTrackingNumber

`func (o *CreateManifest200Response) SetManifestTrackingNumber(v string)`

SetManifestTrackingNumber sets ManifestTrackingNumber field to given value.

### HasManifestTrackingNumber

`func (o *CreateManifest200Response) HasManifestTrackingNumber() bool`

HasManifestTrackingNumber returns a boolean if a field has been set.

### GetFromAddress

`func (o *CreateManifest200Response) GetFromAddress() ManifestDetailedResponseFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *CreateManifest200Response) GetFromAddressOk() (*ManifestDetailedResponseFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *CreateManifest200Response) SetFromAddress(v ManifestDetailedResponseFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *CreateManifest200Response) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcelTrackingNumbers

`func (o *CreateManifest200Response) GetParcelTrackingNumbers() []map[string]interface{}`

GetParcelTrackingNumbers returns the ParcelTrackingNumbers field if non-nil, zero value otherwise.

### GetParcelTrackingNumbersOk

`func (o *CreateManifest200Response) GetParcelTrackingNumbersOk() (*[]map[string]interface{}, bool)`

GetParcelTrackingNumbersOk returns a tuple with the ParcelTrackingNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumbers

`func (o *CreateManifest200Response) SetParcelTrackingNumbers(v []map[string]interface{})`

SetParcelTrackingNumbers sets ParcelTrackingNumbers field to given value.

### HasParcelTrackingNumbers

`func (o *CreateManifest200Response) HasParcelTrackingNumbers() bool`

HasParcelTrackingNumbers returns a boolean if a field has been set.

### GetSubmissionDate

`func (o *CreateManifest200Response) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *CreateManifest200Response) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *CreateManifest200Response) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *CreateManifest200Response) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


