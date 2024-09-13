# ReferenceV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference1** | Pointer to **string** | Reference 1 can have one of the above-indicated values/information, which is printed on Label, e.g. Cost Account No. (if any) or Invoice Number. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 
**Reference2** | Pointer to **string** | Reference 2 can have other details as indicated in the example values above. This is also printed on Label, e.g. Package Description . &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 
**Reference3** | Pointer to **string** | Reference 3 can have the information which were not fulfilled in Ref1 and Ref2, e.g. Order No. or Purchase Order ID. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 
**Reference4** | Pointer to **string** | Reference 4 can have more information which were not provided in Ref1, Ref2, or Ref3 e.g. Carrier Note. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 
**PoNumber** | Pointer to **string** | The Postal Office Number. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 
**Department** | Pointer to **string** | The department of the Recipient. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 
**AdditionalReference1** | Pointer to **string** | Additional Reference is hardly used, but sender can mention anything as per requirement, just for Recipient&#39;s information. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 
**AdditionalReference2** | Pointer to **string** | Any tags or information that to be shown to Recipient, can be mentioned by Sender, which is not indicated on AdditionalReference1 field, e.g. PO No, Order No. etc.&lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;. | [optional] 

## Methods

### NewReferenceV2

`func NewReferenceV2() *ReferenceV2`

NewReferenceV2 instantiates a new ReferenceV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceV2WithDefaults

`func NewReferenceV2WithDefaults() *ReferenceV2`

NewReferenceV2WithDefaults instantiates a new ReferenceV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference1

`func (o *ReferenceV2) GetReference1() string`

GetReference1 returns the Reference1 field if non-nil, zero value otherwise.

### GetReference1Ok

`func (o *ReferenceV2) GetReference1Ok() (*string, bool)`

GetReference1Ok returns a tuple with the Reference1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference1

`func (o *ReferenceV2) SetReference1(v string)`

SetReference1 sets Reference1 field to given value.

### HasReference1

`func (o *ReferenceV2) HasReference1() bool`

HasReference1 returns a boolean if a field has been set.

### GetReference2

`func (o *ReferenceV2) GetReference2() string`

GetReference2 returns the Reference2 field if non-nil, zero value otherwise.

### GetReference2Ok

`func (o *ReferenceV2) GetReference2Ok() (*string, bool)`

GetReference2Ok returns a tuple with the Reference2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference2

`func (o *ReferenceV2) SetReference2(v string)`

SetReference2 sets Reference2 field to given value.

### HasReference2

`func (o *ReferenceV2) HasReference2() bool`

HasReference2 returns a boolean if a field has been set.

### GetReference3

`func (o *ReferenceV2) GetReference3() string`

GetReference3 returns the Reference3 field if non-nil, zero value otherwise.

### GetReference3Ok

`func (o *ReferenceV2) GetReference3Ok() (*string, bool)`

GetReference3Ok returns a tuple with the Reference3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference3

`func (o *ReferenceV2) SetReference3(v string)`

SetReference3 sets Reference3 field to given value.

### HasReference3

`func (o *ReferenceV2) HasReference3() bool`

HasReference3 returns a boolean if a field has been set.

### GetReference4

`func (o *ReferenceV2) GetReference4() string`

GetReference4 returns the Reference4 field if non-nil, zero value otherwise.

### GetReference4Ok

`func (o *ReferenceV2) GetReference4Ok() (*string, bool)`

GetReference4Ok returns a tuple with the Reference4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference4

`func (o *ReferenceV2) SetReference4(v string)`

SetReference4 sets Reference4 field to given value.

### HasReference4

`func (o *ReferenceV2) HasReference4() bool`

HasReference4 returns a boolean if a field has been set.

### GetPoNumber

`func (o *ReferenceV2) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *ReferenceV2) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *ReferenceV2) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *ReferenceV2) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetDepartment

`func (o *ReferenceV2) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ReferenceV2) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ReferenceV2) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ReferenceV2) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetAdditionalReference1

`func (o *ReferenceV2) GetAdditionalReference1() string`

GetAdditionalReference1 returns the AdditionalReference1 field if non-nil, zero value otherwise.

### GetAdditionalReference1Ok

`func (o *ReferenceV2) GetAdditionalReference1Ok() (*string, bool)`

GetAdditionalReference1Ok returns a tuple with the AdditionalReference1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalReference1

`func (o *ReferenceV2) SetAdditionalReference1(v string)`

SetAdditionalReference1 sets AdditionalReference1 field to given value.

### HasAdditionalReference1

`func (o *ReferenceV2) HasAdditionalReference1() bool`

HasAdditionalReference1 returns a boolean if a field has been set.

### GetAdditionalReference2

`func (o *ReferenceV2) GetAdditionalReference2() string`

GetAdditionalReference2 returns the AdditionalReference2 field if non-nil, zero value otherwise.

### GetAdditionalReference2Ok

`func (o *ReferenceV2) GetAdditionalReference2Ok() (*string, bool)`

GetAdditionalReference2Ok returns a tuple with the AdditionalReference2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalReference2

`func (o *ReferenceV2) SetAdditionalReference2(v string)`

SetAdditionalReference2 sets AdditionalReference2 field to given value.

### HasAdditionalReference2

`func (o *ReferenceV2) HasAdditionalReference2() bool`

HasAdditionalReference2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


