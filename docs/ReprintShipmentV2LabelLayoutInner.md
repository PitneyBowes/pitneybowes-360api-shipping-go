# ReprintShipmentV2LabelLayoutInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | This is used to encode binary data as printable text.&lt;br /&gt; ContentType of the document is URL if the fileformat is PDF, while it will be BASE64 if the fileFormat is ZPL2. | [optional] 
**Contents** | Pointer to **string** | Content/Identifier of document e.g. DOCUMENT_REFERECE_ID.&lt;br /&gt; Actual document name e.g. abc.pdf. [IN]. | [optional] 
**FileFormat** | Pointer to **string** | Defines the format of the document file the print takes. | [optional] 
**Size** | Pointer to **string** | Defines the label size of the Shipment, that is, the Shipping Label is available in different Doc Size. | [optional] 
**Type** | Pointer to **string** | Defines the type of the Shipment. | [optional] 

## Methods

### NewReprintShipmentV2LabelLayoutInner

`func NewReprintShipmentV2LabelLayoutInner() *ReprintShipmentV2LabelLayoutInner`

NewReprintShipmentV2LabelLayoutInner instantiates a new ReprintShipmentV2LabelLayoutInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReprintShipmentV2LabelLayoutInnerWithDefaults

`func NewReprintShipmentV2LabelLayoutInnerWithDefaults() *ReprintShipmentV2LabelLayoutInner`

NewReprintShipmentV2LabelLayoutInnerWithDefaults instantiates a new ReprintShipmentV2LabelLayoutInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *ReprintShipmentV2LabelLayoutInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ReprintShipmentV2LabelLayoutInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ReprintShipmentV2LabelLayoutInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ReprintShipmentV2LabelLayoutInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetContents

`func (o *ReprintShipmentV2LabelLayoutInner) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ReprintShipmentV2LabelLayoutInner) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ReprintShipmentV2LabelLayoutInner) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ReprintShipmentV2LabelLayoutInner) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetFileFormat

`func (o *ReprintShipmentV2LabelLayoutInner) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *ReprintShipmentV2LabelLayoutInner) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *ReprintShipmentV2LabelLayoutInner) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *ReprintShipmentV2LabelLayoutInner) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### GetSize

`func (o *ReprintShipmentV2LabelLayoutInner) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ReprintShipmentV2LabelLayoutInner) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ReprintShipmentV2LabelLayoutInner) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ReprintShipmentV2LabelLayoutInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *ReprintShipmentV2LabelLayoutInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReprintShipmentV2LabelLayoutInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReprintShipmentV2LabelLayoutInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReprintShipmentV2LabelLayoutInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


