# ManifestDetailedResponseManifestDocumentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | Determines type of the document as a URL. &lt;br&gt;- URL: The response returns a link to the label or manifest. | [optional] 
**Contents** | Pointer to **string** | When contentType is URL, this is the URL to access the label or manifest. Note: The document is available for 24 hours after it is created. | [optional] 
**Type** | Pointer to **string** | Specifies the type of manifest. | [optional] 

## Methods

### NewManifestDetailedResponseManifestDocumentsInner

`func NewManifestDetailedResponseManifestDocumentsInner() *ManifestDetailedResponseManifestDocumentsInner`

NewManifestDetailedResponseManifestDocumentsInner instantiates a new ManifestDetailedResponseManifestDocumentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManifestDetailedResponseManifestDocumentsInnerWithDefaults

`func NewManifestDetailedResponseManifestDocumentsInnerWithDefaults() *ManifestDetailedResponseManifestDocumentsInner`

NewManifestDetailedResponseManifestDocumentsInnerWithDefaults instantiates a new ManifestDetailedResponseManifestDocumentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *ManifestDetailedResponseManifestDocumentsInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ManifestDetailedResponseManifestDocumentsInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ManifestDetailedResponseManifestDocumentsInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ManifestDetailedResponseManifestDocumentsInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetContents

`func (o *ManifestDetailedResponseManifestDocumentsInner) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ManifestDetailedResponseManifestDocumentsInner) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ManifestDetailedResponseManifestDocumentsInner) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ManifestDetailedResponseManifestDocumentsInner) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetType

`func (o *ManifestDetailedResponseManifestDocumentsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManifestDetailedResponseManifestDocumentsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManifestDetailedResponseManifestDocumentsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManifestDetailedResponseManifestDocumentsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


