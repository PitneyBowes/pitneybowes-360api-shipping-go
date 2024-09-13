# ShipmentOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddToManifest** | Pointer to **bool** | This option asks if the shipment is to be added for Manifest, so that the shipment will reflect in the Manifest Form while compilation. The value can be &#39;true&#39; or &#39;false&#39;. Applicable for USPS | [optional] 
**PrintCustomMessage** | Pointer to **string** | This prints a custom message on shipping label | [optional] 
**ReceiptOption** | Pointer to **string** | It provides options to print receipt with shipping label. Only applicable for USPS, can take values- RECEIPT_ONLY or RECEIPT_WITH_INSTRUCTIONS or NO_OPTIONS | [optional] 
**PrintDepartment** | Pointer to **string** | This prints department on shipping label, applicable for FedEx | [optional] 
**PrintInvoiceNumber** | Pointer to **string** | This prints invoice number on shipping label, applicable for FedEx | [optional] 
**PrintPONumber** | Pointer to **string** | This prints po number on shipping label, applicable for FedEx | [optional] 
**PackageDescription** | Pointer to **string** | This is required for DHL Express shipment. It is used to provide description of package | [optional] 

## Methods

### NewShipmentOptions

`func NewShipmentOptions() *ShipmentOptions`

NewShipmentOptions instantiates a new ShipmentOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentOptionsWithDefaults

`func NewShipmentOptionsWithDefaults() *ShipmentOptions`

NewShipmentOptionsWithDefaults instantiates a new ShipmentOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddToManifest

`func (o *ShipmentOptions) GetAddToManifest() bool`

GetAddToManifest returns the AddToManifest field if non-nil, zero value otherwise.

### GetAddToManifestOk

`func (o *ShipmentOptions) GetAddToManifestOk() (*bool, bool)`

GetAddToManifestOk returns a tuple with the AddToManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToManifest

`func (o *ShipmentOptions) SetAddToManifest(v bool)`

SetAddToManifest sets AddToManifest field to given value.

### HasAddToManifest

`func (o *ShipmentOptions) HasAddToManifest() bool`

HasAddToManifest returns a boolean if a field has been set.

### GetPrintCustomMessage

`func (o *ShipmentOptions) GetPrintCustomMessage() string`

GetPrintCustomMessage returns the PrintCustomMessage field if non-nil, zero value otherwise.

### GetPrintCustomMessageOk

`func (o *ShipmentOptions) GetPrintCustomMessageOk() (*string, bool)`

GetPrintCustomMessageOk returns a tuple with the PrintCustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintCustomMessage

`func (o *ShipmentOptions) SetPrintCustomMessage(v string)`

SetPrintCustomMessage sets PrintCustomMessage field to given value.

### HasPrintCustomMessage

`func (o *ShipmentOptions) HasPrintCustomMessage() bool`

HasPrintCustomMessage returns a boolean if a field has been set.

### GetReceiptOption

`func (o *ShipmentOptions) GetReceiptOption() string`

GetReceiptOption returns the ReceiptOption field if non-nil, zero value otherwise.

### GetReceiptOptionOk

`func (o *ShipmentOptions) GetReceiptOptionOk() (*string, bool)`

GetReceiptOptionOk returns a tuple with the ReceiptOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptOption

`func (o *ShipmentOptions) SetReceiptOption(v string)`

SetReceiptOption sets ReceiptOption field to given value.

### HasReceiptOption

`func (o *ShipmentOptions) HasReceiptOption() bool`

HasReceiptOption returns a boolean if a field has been set.

### GetPrintDepartment

`func (o *ShipmentOptions) GetPrintDepartment() string`

GetPrintDepartment returns the PrintDepartment field if non-nil, zero value otherwise.

### GetPrintDepartmentOk

`func (o *ShipmentOptions) GetPrintDepartmentOk() (*string, bool)`

GetPrintDepartmentOk returns a tuple with the PrintDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintDepartment

`func (o *ShipmentOptions) SetPrintDepartment(v string)`

SetPrintDepartment sets PrintDepartment field to given value.

### HasPrintDepartment

`func (o *ShipmentOptions) HasPrintDepartment() bool`

HasPrintDepartment returns a boolean if a field has been set.

### GetPrintInvoiceNumber

`func (o *ShipmentOptions) GetPrintInvoiceNumber() string`

GetPrintInvoiceNumber returns the PrintInvoiceNumber field if non-nil, zero value otherwise.

### GetPrintInvoiceNumberOk

`func (o *ShipmentOptions) GetPrintInvoiceNumberOk() (*string, bool)`

GetPrintInvoiceNumberOk returns a tuple with the PrintInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintInvoiceNumber

`func (o *ShipmentOptions) SetPrintInvoiceNumber(v string)`

SetPrintInvoiceNumber sets PrintInvoiceNumber field to given value.

### HasPrintInvoiceNumber

`func (o *ShipmentOptions) HasPrintInvoiceNumber() bool`

HasPrintInvoiceNumber returns a boolean if a field has been set.

### GetPrintPONumber

`func (o *ShipmentOptions) GetPrintPONumber() string`

GetPrintPONumber returns the PrintPONumber field if non-nil, zero value otherwise.

### GetPrintPONumberOk

`func (o *ShipmentOptions) GetPrintPONumberOk() (*string, bool)`

GetPrintPONumberOk returns a tuple with the PrintPONumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintPONumber

`func (o *ShipmentOptions) SetPrintPONumber(v string)`

SetPrintPONumber sets PrintPONumber field to given value.

### HasPrintPONumber

`func (o *ShipmentOptions) HasPrintPONumber() bool`

HasPrintPONumber returns a boolean if a field has been set.

### GetPackageDescription

`func (o *ShipmentOptions) GetPackageDescription() string`

GetPackageDescription returns the PackageDescription field if non-nil, zero value otherwise.

### GetPackageDescriptionOk

`func (o *ShipmentOptions) GetPackageDescriptionOk() (*string, bool)`

GetPackageDescriptionOk returns a tuple with the PackageDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDescription

`func (o *ShipmentOptions) SetPackageDescription(v string)`

SetPackageDescription sets PackageDescription field to given value.

### HasPackageDescription

`func (o *ShipmentOptions) HasPackageDescription() bool`

HasPackageDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


