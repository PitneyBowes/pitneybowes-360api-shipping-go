# ShipmentOptionsERR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddToManifest** | Pointer to **bool** | This option asks if the shipment is to be added for Manifest, so that the shipment can be reflected in the Manifest Form while compilation. The value can be &#39;true&#39; or &#39;false&#39;. Applicable for USPS. | [optional] 
**PrintCustomMessage** | Pointer to **string** | This prints a custom message on the Shipping Label. | [optional] 
**ReceiptOption** | Pointer to **string** | It provides options to print receipt with Shipping Label. This is only applicable for USPS, and takes values: &#x60;RECEIPT_ONLY&#x60;, &#x60;RECEIPT_WITH_INSTRUCTIONS&#x60;, or &#x60;NO_OPTIONS&#x60;. | [optional] 
**PrintDepartment** | Pointer to **string** | It displays/prints the Department on the Shipping Label. | [optional] 
**PrintInvoiceNumber** | Pointer to **string** | It displays/prints Invoice Number on the Shipping Label. | [optional] 
**PrintPONumber** | Pointer to **string** | It displays/prints PO Number on the Shipping Label. | [optional] 

## Methods

### NewShipmentOptionsERR

`func NewShipmentOptionsERR() *ShipmentOptionsERR`

NewShipmentOptionsERR instantiates a new ShipmentOptionsERR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentOptionsERRWithDefaults

`func NewShipmentOptionsERRWithDefaults() *ShipmentOptionsERR`

NewShipmentOptionsERRWithDefaults instantiates a new ShipmentOptionsERR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddToManifest

`func (o *ShipmentOptionsERR) GetAddToManifest() bool`

GetAddToManifest returns the AddToManifest field if non-nil, zero value otherwise.

### GetAddToManifestOk

`func (o *ShipmentOptionsERR) GetAddToManifestOk() (*bool, bool)`

GetAddToManifestOk returns a tuple with the AddToManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToManifest

`func (o *ShipmentOptionsERR) SetAddToManifest(v bool)`

SetAddToManifest sets AddToManifest field to given value.

### HasAddToManifest

`func (o *ShipmentOptionsERR) HasAddToManifest() bool`

HasAddToManifest returns a boolean if a field has been set.

### GetPrintCustomMessage

`func (o *ShipmentOptionsERR) GetPrintCustomMessage() string`

GetPrintCustomMessage returns the PrintCustomMessage field if non-nil, zero value otherwise.

### GetPrintCustomMessageOk

`func (o *ShipmentOptionsERR) GetPrintCustomMessageOk() (*string, bool)`

GetPrintCustomMessageOk returns a tuple with the PrintCustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintCustomMessage

`func (o *ShipmentOptionsERR) SetPrintCustomMessage(v string)`

SetPrintCustomMessage sets PrintCustomMessage field to given value.

### HasPrintCustomMessage

`func (o *ShipmentOptionsERR) HasPrintCustomMessage() bool`

HasPrintCustomMessage returns a boolean if a field has been set.

### GetReceiptOption

`func (o *ShipmentOptionsERR) GetReceiptOption() string`

GetReceiptOption returns the ReceiptOption field if non-nil, zero value otherwise.

### GetReceiptOptionOk

`func (o *ShipmentOptionsERR) GetReceiptOptionOk() (*string, bool)`

GetReceiptOptionOk returns a tuple with the ReceiptOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptOption

`func (o *ShipmentOptionsERR) SetReceiptOption(v string)`

SetReceiptOption sets ReceiptOption field to given value.

### HasReceiptOption

`func (o *ShipmentOptionsERR) HasReceiptOption() bool`

HasReceiptOption returns a boolean if a field has been set.

### GetPrintDepartment

`func (o *ShipmentOptionsERR) GetPrintDepartment() string`

GetPrintDepartment returns the PrintDepartment field if non-nil, zero value otherwise.

### GetPrintDepartmentOk

`func (o *ShipmentOptionsERR) GetPrintDepartmentOk() (*string, bool)`

GetPrintDepartmentOk returns a tuple with the PrintDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintDepartment

`func (o *ShipmentOptionsERR) SetPrintDepartment(v string)`

SetPrintDepartment sets PrintDepartment field to given value.

### HasPrintDepartment

`func (o *ShipmentOptionsERR) HasPrintDepartment() bool`

HasPrintDepartment returns a boolean if a field has been set.

### GetPrintInvoiceNumber

`func (o *ShipmentOptionsERR) GetPrintInvoiceNumber() string`

GetPrintInvoiceNumber returns the PrintInvoiceNumber field if non-nil, zero value otherwise.

### GetPrintInvoiceNumberOk

`func (o *ShipmentOptionsERR) GetPrintInvoiceNumberOk() (*string, bool)`

GetPrintInvoiceNumberOk returns a tuple with the PrintInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintInvoiceNumber

`func (o *ShipmentOptionsERR) SetPrintInvoiceNumber(v string)`

SetPrintInvoiceNumber sets PrintInvoiceNumber field to given value.

### HasPrintInvoiceNumber

`func (o *ShipmentOptionsERR) HasPrintInvoiceNumber() bool`

HasPrintInvoiceNumber returns a boolean if a field has been set.

### GetPrintPONumber

`func (o *ShipmentOptionsERR) GetPrintPONumber() string`

GetPrintPONumber returns the PrintPONumber field if non-nil, zero value otherwise.

### GetPrintPONumberOk

`func (o *ShipmentOptionsERR) GetPrintPONumberOk() (*string, bool)`

GetPrintPONumberOk returns a tuple with the PrintPONumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintPONumber

`func (o *ShipmentOptionsERR) SetPrintPONumber(v string)`

SetPrintPONumber sets PrintPONumber field to given value.

### HasPrintPONumber

`func (o *ShipmentOptionsERR) HasPrintPONumber() bool`

HasPrintPONumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


