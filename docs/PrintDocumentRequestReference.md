# PrintDocumentRequestReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference1** | Pointer to **string** | Reference 1 can have one of the above-indicated values/information, which is printed on Label, e.g. Cost Account No. (if any) or Invoice Number. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 
**Reference2** | Pointer to **string** | Reference 2 can have other details as indicated in the list above. This is also printed on Label, e.g. Package Description . &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 
**Reference3** | Pointer to **string** | Reference 3 can have the information which were not fulfilled in Ref1 and Ref2, e.g. Order No. or Purchase Order ID. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 
**Reference4** | Pointer to **string** | Reference 4 can have more information which were not provided in Ref1, Ref2, or Ref3 e.g. Carrier Note. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 

## Methods

### NewPrintDocumentRequestReference

`func NewPrintDocumentRequestReference() *PrintDocumentRequestReference`

NewPrintDocumentRequestReference instantiates a new PrintDocumentRequestReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintDocumentRequestReferenceWithDefaults

`func NewPrintDocumentRequestReferenceWithDefaults() *PrintDocumentRequestReference`

NewPrintDocumentRequestReferenceWithDefaults instantiates a new PrintDocumentRequestReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference1

`func (o *PrintDocumentRequestReference) GetReference1() string`

GetReference1 returns the Reference1 field if non-nil, zero value otherwise.

### GetReference1Ok

`func (o *PrintDocumentRequestReference) GetReference1Ok() (*string, bool)`

GetReference1Ok returns a tuple with the Reference1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference1

`func (o *PrintDocumentRequestReference) SetReference1(v string)`

SetReference1 sets Reference1 field to given value.

### HasReference1

`func (o *PrintDocumentRequestReference) HasReference1() bool`

HasReference1 returns a boolean if a field has been set.

### GetReference2

`func (o *PrintDocumentRequestReference) GetReference2() string`

GetReference2 returns the Reference2 field if non-nil, zero value otherwise.

### GetReference2Ok

`func (o *PrintDocumentRequestReference) GetReference2Ok() (*string, bool)`

GetReference2Ok returns a tuple with the Reference2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference2

`func (o *PrintDocumentRequestReference) SetReference2(v string)`

SetReference2 sets Reference2 field to given value.

### HasReference2

`func (o *PrintDocumentRequestReference) HasReference2() bool`

HasReference2 returns a boolean if a field has been set.

### GetReference3

`func (o *PrintDocumentRequestReference) GetReference3() string`

GetReference3 returns the Reference3 field if non-nil, zero value otherwise.

### GetReference3Ok

`func (o *PrintDocumentRequestReference) GetReference3Ok() (*string, bool)`

GetReference3Ok returns a tuple with the Reference3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference3

`func (o *PrintDocumentRequestReference) SetReference3(v string)`

SetReference3 sets Reference3 field to given value.

### HasReference3

`func (o *PrintDocumentRequestReference) HasReference3() bool`

HasReference3 returns a boolean if a field has been set.

### GetReference4

`func (o *PrintDocumentRequestReference) GetReference4() string`

GetReference4 returns the Reference4 field if non-nil, zero value otherwise.

### GetReference4Ok

`func (o *PrintDocumentRequestReference) GetReference4Ok() (*string, bool)`

GetReference4Ok returns a tuple with the Reference4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference4

`func (o *PrintDocumentRequestReference) SetReference4(v string)`

SetReference4 sets Reference4 field to given value.

### HasReference4

`func (o *PrintDocumentRequestReference) HasReference4() bool`

HasReference4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


