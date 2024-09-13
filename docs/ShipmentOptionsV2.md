# ShipmentOptionsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddToManifest** | Pointer to **bool** | This option asks if the shipment is to be added for Manifest, so that the shipment will reflect in the Manifest Form while compilation. The value can be &#39;true&#39; or &#39;false&#39;. Applicable for USPS | [optional] 
**PrintCustomMessage** | Pointer to **string** | This prints a custom message on shipping label | [optional] 
**ReceiptOption** | Pointer to **string** | It provides options to print receipt with shipping label. Only applicable for USPS, can take values- RECEIPT_ONLY or RECEIPT_WITH_INSTRUCTIONS or NO_OPTIONS | [optional] 
**PrintDepartment** | Pointer to **string** | This prints department on shipping label, applicable for FedEx | [optional] 
**PrintInvoiceNumber** | Pointer to **string** | This prints invoice number on shipping label, applicable for FedEx | [optional] 
**PrintPONumber** | Pointer to **string** | This prints po number on shipping label, applicable for FedEx | [optional] 
**MinimalAddressValidation** | Pointer to **bool** | If true, then only City, State, and PostalCode (zip) are validated. This option is specific for US domestic addresses only. | [optional] 

## Methods

### NewShipmentOptionsV2

`func NewShipmentOptionsV2() *ShipmentOptionsV2`

NewShipmentOptionsV2 instantiates a new ShipmentOptionsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentOptionsV2WithDefaults

`func NewShipmentOptionsV2WithDefaults() *ShipmentOptionsV2`

NewShipmentOptionsV2WithDefaults instantiates a new ShipmentOptionsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddToManifest

`func (o *ShipmentOptionsV2) GetAddToManifest() bool`

GetAddToManifest returns the AddToManifest field if non-nil, zero value otherwise.

### GetAddToManifestOk

`func (o *ShipmentOptionsV2) GetAddToManifestOk() (*bool, bool)`

GetAddToManifestOk returns a tuple with the AddToManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToManifest

`func (o *ShipmentOptionsV2) SetAddToManifest(v bool)`

SetAddToManifest sets AddToManifest field to given value.

### HasAddToManifest

`func (o *ShipmentOptionsV2) HasAddToManifest() bool`

HasAddToManifest returns a boolean if a field has been set.

### GetPrintCustomMessage

`func (o *ShipmentOptionsV2) GetPrintCustomMessage() string`

GetPrintCustomMessage returns the PrintCustomMessage field if non-nil, zero value otherwise.

### GetPrintCustomMessageOk

`func (o *ShipmentOptionsV2) GetPrintCustomMessageOk() (*string, bool)`

GetPrintCustomMessageOk returns a tuple with the PrintCustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintCustomMessage

`func (o *ShipmentOptionsV2) SetPrintCustomMessage(v string)`

SetPrintCustomMessage sets PrintCustomMessage field to given value.

### HasPrintCustomMessage

`func (o *ShipmentOptionsV2) HasPrintCustomMessage() bool`

HasPrintCustomMessage returns a boolean if a field has been set.

### GetReceiptOption

`func (o *ShipmentOptionsV2) GetReceiptOption() string`

GetReceiptOption returns the ReceiptOption field if non-nil, zero value otherwise.

### GetReceiptOptionOk

`func (o *ShipmentOptionsV2) GetReceiptOptionOk() (*string, bool)`

GetReceiptOptionOk returns a tuple with the ReceiptOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptOption

`func (o *ShipmentOptionsV2) SetReceiptOption(v string)`

SetReceiptOption sets ReceiptOption field to given value.

### HasReceiptOption

`func (o *ShipmentOptionsV2) HasReceiptOption() bool`

HasReceiptOption returns a boolean if a field has been set.

### GetPrintDepartment

`func (o *ShipmentOptionsV2) GetPrintDepartment() string`

GetPrintDepartment returns the PrintDepartment field if non-nil, zero value otherwise.

### GetPrintDepartmentOk

`func (o *ShipmentOptionsV2) GetPrintDepartmentOk() (*string, bool)`

GetPrintDepartmentOk returns a tuple with the PrintDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintDepartment

`func (o *ShipmentOptionsV2) SetPrintDepartment(v string)`

SetPrintDepartment sets PrintDepartment field to given value.

### HasPrintDepartment

`func (o *ShipmentOptionsV2) HasPrintDepartment() bool`

HasPrintDepartment returns a boolean if a field has been set.

### GetPrintInvoiceNumber

`func (o *ShipmentOptionsV2) GetPrintInvoiceNumber() string`

GetPrintInvoiceNumber returns the PrintInvoiceNumber field if non-nil, zero value otherwise.

### GetPrintInvoiceNumberOk

`func (o *ShipmentOptionsV2) GetPrintInvoiceNumberOk() (*string, bool)`

GetPrintInvoiceNumberOk returns a tuple with the PrintInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintInvoiceNumber

`func (o *ShipmentOptionsV2) SetPrintInvoiceNumber(v string)`

SetPrintInvoiceNumber sets PrintInvoiceNumber field to given value.

### HasPrintInvoiceNumber

`func (o *ShipmentOptionsV2) HasPrintInvoiceNumber() bool`

HasPrintInvoiceNumber returns a boolean if a field has been set.

### GetPrintPONumber

`func (o *ShipmentOptionsV2) GetPrintPONumber() string`

GetPrintPONumber returns the PrintPONumber field if non-nil, zero value otherwise.

### GetPrintPONumberOk

`func (o *ShipmentOptionsV2) GetPrintPONumberOk() (*string, bool)`

GetPrintPONumberOk returns a tuple with the PrintPONumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintPONumber

`func (o *ShipmentOptionsV2) SetPrintPONumber(v string)`

SetPrintPONumber sets PrintPONumber field to given value.

### HasPrintPONumber

`func (o *ShipmentOptionsV2) HasPrintPONumber() bool`

HasPrintPONumber returns a boolean if a field has been set.

### GetMinimalAddressValidation

`func (o *ShipmentOptionsV2) GetMinimalAddressValidation() bool`

GetMinimalAddressValidation returns the MinimalAddressValidation field if non-nil, zero value otherwise.

### GetMinimalAddressValidationOk

`func (o *ShipmentOptionsV2) GetMinimalAddressValidationOk() (*bool, bool)`

GetMinimalAddressValidationOk returns a tuple with the MinimalAddressValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalAddressValidation

`func (o *ShipmentOptionsV2) SetMinimalAddressValidation(v bool)`

SetMinimalAddressValidation sets MinimalAddressValidation field to given value.

### HasMinimalAddressValidation

`func (o *ShipmentOptionsV2) HasMinimalAddressValidation() bool`

HasMinimalAddressValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


