# ManifestDetailedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierAccountId** | Pointer to **string** | A unique identifier associated with the Carrier account which is used while creating Manifest. | [optional] 
**CarrierName** | Pointer to **string** | Name of the Carrier. | [optional] 
**ManifestDocuments** | Pointer to [**[]ManifestDetailedResponseManifestDocumentsInner**](ManifestDetailedResponseManifestDocumentsInner.md) | The electronically generated document that has manifest (end-of-day) records of all shipments of the day. | [optional] 
**ManifestID** | Pointer to **string** | The unique manifest ID. This field is not returned for APAC Services. | [optional] 
**ManifestTrackingNumber** | Pointer to **string** | The manifest tracking number. This is returned only if the carrier has a pre-defined valid value, e.g., UPS, FedEX, or USPS. | [optional] 
**FromAddress** | Pointer to [**ManifestDetailedResponseFromAddress**](ManifestDetailedResponseFromAddress.md) |  | [optional] 
**ParcelTrackingNumbers** | Pointer to **[]interface{}** |  | [optional] 
**SubmissionDate** | Pointer to **string** | The date the shipments are to be tendered to the carrier, entered as YYYY-MM-DD. | [optional] 

## Methods

### NewManifestDetailedResponse

`func NewManifestDetailedResponse() *ManifestDetailedResponse`

NewManifestDetailedResponse instantiates a new ManifestDetailedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManifestDetailedResponseWithDefaults

`func NewManifestDetailedResponseWithDefaults() *ManifestDetailedResponse`

NewManifestDetailedResponseWithDefaults instantiates a new ManifestDetailedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierAccountId

`func (o *ManifestDetailedResponse) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *ManifestDetailedResponse) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *ManifestDetailedResponse) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.

### HasCarrierAccountId

`func (o *ManifestDetailedResponse) HasCarrierAccountId() bool`

HasCarrierAccountId returns a boolean if a field has been set.

### GetCarrierName

`func (o *ManifestDetailedResponse) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *ManifestDetailedResponse) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *ManifestDetailedResponse) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *ManifestDetailedResponse) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetManifestDocuments

`func (o *ManifestDetailedResponse) GetManifestDocuments() []ManifestDetailedResponseManifestDocumentsInner`

GetManifestDocuments returns the ManifestDocuments field if non-nil, zero value otherwise.

### GetManifestDocumentsOk

`func (o *ManifestDetailedResponse) GetManifestDocumentsOk() (*[]ManifestDetailedResponseManifestDocumentsInner, bool)`

GetManifestDocumentsOk returns a tuple with the ManifestDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestDocuments

`func (o *ManifestDetailedResponse) SetManifestDocuments(v []ManifestDetailedResponseManifestDocumentsInner)`

SetManifestDocuments sets ManifestDocuments field to given value.

### HasManifestDocuments

`func (o *ManifestDetailedResponse) HasManifestDocuments() bool`

HasManifestDocuments returns a boolean if a field has been set.

### GetManifestID

`func (o *ManifestDetailedResponse) GetManifestID() string`

GetManifestID returns the ManifestID field if non-nil, zero value otherwise.

### GetManifestIDOk

`func (o *ManifestDetailedResponse) GetManifestIDOk() (*string, bool)`

GetManifestIDOk returns a tuple with the ManifestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestID

`func (o *ManifestDetailedResponse) SetManifestID(v string)`

SetManifestID sets ManifestID field to given value.

### HasManifestID

`func (o *ManifestDetailedResponse) HasManifestID() bool`

HasManifestID returns a boolean if a field has been set.

### GetManifestTrackingNumber

`func (o *ManifestDetailedResponse) GetManifestTrackingNumber() string`

GetManifestTrackingNumber returns the ManifestTrackingNumber field if non-nil, zero value otherwise.

### GetManifestTrackingNumberOk

`func (o *ManifestDetailedResponse) GetManifestTrackingNumberOk() (*string, bool)`

GetManifestTrackingNumberOk returns a tuple with the ManifestTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestTrackingNumber

`func (o *ManifestDetailedResponse) SetManifestTrackingNumber(v string)`

SetManifestTrackingNumber sets ManifestTrackingNumber field to given value.

### HasManifestTrackingNumber

`func (o *ManifestDetailedResponse) HasManifestTrackingNumber() bool`

HasManifestTrackingNumber returns a boolean if a field has been set.

### GetFromAddress

`func (o *ManifestDetailedResponse) GetFromAddress() ManifestDetailedResponseFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ManifestDetailedResponse) GetFromAddressOk() (*ManifestDetailedResponseFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ManifestDetailedResponse) SetFromAddress(v ManifestDetailedResponseFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *ManifestDetailedResponse) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcelTrackingNumbers

`func (o *ManifestDetailedResponse) GetParcelTrackingNumbers() []interface{}`

GetParcelTrackingNumbers returns the ParcelTrackingNumbers field if non-nil, zero value otherwise.

### GetParcelTrackingNumbersOk

`func (o *ManifestDetailedResponse) GetParcelTrackingNumbersOk() (*[]interface{}, bool)`

GetParcelTrackingNumbersOk returns a tuple with the ParcelTrackingNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumbers

`func (o *ManifestDetailedResponse) SetParcelTrackingNumbers(v []interface{})`

SetParcelTrackingNumbers sets ParcelTrackingNumbers field to given value.

### HasParcelTrackingNumbers

`func (o *ManifestDetailedResponse) HasParcelTrackingNumbers() bool`

HasParcelTrackingNumbers returns a boolean if a field has been set.

### GetSubmissionDate

`func (o *ManifestDetailedResponse) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *ManifestDetailedResponse) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *ManifestDetailedResponse) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *ManifestDetailedResponse) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


