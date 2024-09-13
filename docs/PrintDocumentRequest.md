# PrintDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrinterAliasName** | **string** | Name of the Printer connected (directly or via network) to a Computer. &#x60;Max length &#x3D; 30&#x60; | 
**Data** | **string** | Content/Identifier of document e.g. DOCUMENT_REFERECE_ID. Actual document name e.g. abc.pdf. [IN] i.e base64 string, URL, file path | 
**DataType** | **string** | Data Type of the document e.g. DOCUMENT_REFERENCE. [IN/OUT] | 
**DocumentType** | **string** | The format of the document file the print takes. | 
**FormName** | **string** | The name of the Document Form. | 
**Orientation** | Pointer to **string** | The orientation of the document layout: Portrait or Landscape. | [optional] 
**Reference** | Pointer to [**PrintDocumentRequestReference**](PrintDocumentRequestReference.md) |  | [optional] 

## Methods

### NewPrintDocumentRequest

`func NewPrintDocumentRequest(printerAliasName string, data string, dataType string, documentType string, formName string, ) *PrintDocumentRequest`

NewPrintDocumentRequest instantiates a new PrintDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintDocumentRequestWithDefaults

`func NewPrintDocumentRequestWithDefaults() *PrintDocumentRequest`

NewPrintDocumentRequestWithDefaults instantiates a new PrintDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrinterAliasName

`func (o *PrintDocumentRequest) GetPrinterAliasName() string`

GetPrinterAliasName returns the PrinterAliasName field if non-nil, zero value otherwise.

### GetPrinterAliasNameOk

`func (o *PrintDocumentRequest) GetPrinterAliasNameOk() (*string, bool)`

GetPrinterAliasNameOk returns a tuple with the PrinterAliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinterAliasName

`func (o *PrintDocumentRequest) SetPrinterAliasName(v string)`

SetPrinterAliasName sets PrinterAliasName field to given value.


### GetData

`func (o *PrintDocumentRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PrintDocumentRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PrintDocumentRequest) SetData(v string)`

SetData sets Data field to given value.


### GetDataType

`func (o *PrintDocumentRequest) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PrintDocumentRequest) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PrintDocumentRequest) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetDocumentType

`func (o *PrintDocumentRequest) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *PrintDocumentRequest) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *PrintDocumentRequest) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### GetFormName

`func (o *PrintDocumentRequest) GetFormName() string`

GetFormName returns the FormName field if non-nil, zero value otherwise.

### GetFormNameOk

`func (o *PrintDocumentRequest) GetFormNameOk() (*string, bool)`

GetFormNameOk returns a tuple with the FormName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormName

`func (o *PrintDocumentRequest) SetFormName(v string)`

SetFormName sets FormName field to given value.


### GetOrientation

`func (o *PrintDocumentRequest) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *PrintDocumentRequest) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *PrintDocumentRequest) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *PrintDocumentRequest) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetReference

`func (o *PrintDocumentRequest) GetReference() PrintDocumentRequestReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PrintDocumentRequest) GetReferenceOk() (*PrintDocumentRequestReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PrintDocumentRequest) SetReference(v PrintDocumentRequestReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PrintDocumentRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


